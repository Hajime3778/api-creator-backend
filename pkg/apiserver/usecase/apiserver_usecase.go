package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	_apiRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/api/repository"
	_methodRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/method/repository"
	_modelRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/model/repository"
	_apiserverRepository "github.com/Hajime3778/api-creator-backend/pkg/apiserver/repository"
	"github.com/Hajime3778/api-creator-backend/pkg/domain"
	"github.com/xeipuuv/gojsonschema"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	// ErrAPINotFound "api not found"
	ErrAPINotFound = errors.New("api not found")
	// ErrInvalidRequest "invalid request"
	ErrInvalidRequest = errors.New("invalid request")
	// ErrModelNotDeclare "model not declare"
	ErrModelNotDeclare = errors.New("model not declare")
)

// APIServerUsecase Interface
type APIServerUsecase interface {
	RequestDocumentServer(httpMethod string, url string, body []byte) (interface{}, int, error)
}

type apiServerUsecase struct {
	apiRepo       _apiRepository.APIRepository
	methodRepo    _methodRepository.MethodRepository
	modelRepo     _modelRepository.ModelRepository
	apiserverRepo _apiserverRepository.APIServerRepository
}

// NewAPIServerUsecase APIServerUsecaseインターフェイスを表すオブジェクトを作成します
func NewAPIServerUsecase(apiRepo _apiRepository.APIRepository, methodRepo _methodRepository.MethodRepository, modelRepo _modelRepository.ModelRepository, apiserverRepo _apiserverRepository.APIServerRepository) APIServerUsecase {
	return &apiServerUsecase{
		apiRepo:       apiRepo,
		methodRepo:    methodRepo,
		modelRepo:     modelRepo,
		apiserverRepo: apiserverRepo,
	}
}

// RequestDocumentServer リクエスト情報からMethodを特定し、ドキュメントに対してCRUDします
func (u *apiServerUsecase) RequestDocumentServer(httpMethod string, url string, body []byte) (interface{}, int, error) {
	api, err := u.apiRepo.GetByURL(url)
	if err != nil {
		// APIが見つかりません
		return "", http.StatusNotFound, err
	}
	model, err := u.modelRepo.GetByAPIID(api.ID)
	if err != nil {
		return "", http.StatusBadRequest, ErrModelNotDeclare
	}

	// コレクションを削除する、システム規定のメソッド
	pathParam := strings.Replace(url, api.URL+"/", "", 1)
	if pathParam == "remove-target-collection" && httpMethod == "DELETE" {
		return u.apiserverRepo.RemoveCollection(model.Name)
	}

	// 対象のメソッドを取得
	method, err := u.getRequestedMethod(httpMethod, url, api)

	if err != nil {
		return "", http.StatusNotFound, ErrAPINotFound
	}

	if method.Type != httpMethod {
		return "", http.StatusNotFound, ErrAPINotFound
	}

	// リクエストされたパラメータを取得
	paramKey, paramValue, err := getRequestedURLParameter(url, api.URL, method.URL, model.Schema)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	switch method.Type {
	case "GET":
		return u.get(method.IsArray, model.Name, paramKey, paramValue)

	case "POST":
		return u.create(model, body)

	case "PUT":
		return u.update(model, body)

	case "DELETE":
		return u.delete(model.Name, paramKey, paramValue)

	default:
		return "", http.StatusInternalServerError, errors.New("incorrect http method")
	}
}

func (u *apiServerUsecase) get(isArray bool, modelName string, key string, value interface{}) (interface{}, int, error) {
	if isArray {
		return u.apiserverRepo.GetList(modelName, key, value)
	}
	return u.apiserverRepo.Get(modelName, key, value)
}

func (u *apiServerUsecase) create(model domain.Model, body []byte) (interface{}, int, error) {
	err := getRequestedSchemaValidate(model.Schema, body)
	if err != nil {
		return "", http.StatusBadRequest, err
	}
	keys, err := model.GetKeyNames()
	if err != nil {
		return "", http.StatusBadRequest, err
	}

	value, err := getPropertyValue(keys[0], body)
	if err != nil {
		return "", http.StatusBadRequest, err
	}

	if _, status, _ := u.apiserverRepo.Get(model.Name, keys[0], value); status != http.StatusNotFound {
		return "", http.StatusBadRequest, errors.New("record is exists")
	}

	return u.apiserverRepo.Create(model.Name, keys[0], body)
}

func (u *apiServerUsecase) update(model domain.Model, body []byte) (interface{}, int, error) {
	err := getRequestedSchemaValidate(model.Schema, body)
	if err != nil {
		return "", http.StatusBadRequest, err
	}
	keys, err := model.GetKeyNames()
	if err != nil {
		return "", http.StatusBadRequest, err
	}

	value, err := getPropertyValue(keys[0], body)

	if _, status, _ := u.apiserverRepo.Get(model.Name, keys[0], value); status == http.StatusNotFound {
		return "", http.StatusBadRequest, errors.New("record is not found")
	}

	return u.apiserverRepo.Update(model.Name, keys[0], body)
}

func (u *apiServerUsecase) delete(modelName string, key string, value interface{}) (interface{}, int, error) {
	return u.apiserverRepo.Delete(modelName, key, value)
}

// getRequestedMethod リクエストされたHTTPメソッド、URL、APIから、対象のMethodを返却します
func (u *apiServerUsecase) getRequestedMethod(httpMethod string, requestedURL string, api domain.API) (domain.Method, error) {

	var returnMethod domain.Method

	methods, err := u.methodRepo.GetListByAPIID(api.ID)
	if err != nil {
		return returnMethod, err
	}

	// MethodのURL部分を抽出
	requestedMethodURL := strings.Replace(requestedURL, api.URL, "", 1)

	// /{}で囲まれた箇所を削除
	re := regexp.MustCompile(`/\{.+?\}`)

	// ※パラメータが2つ以上やクエリストリングは現在の仕様にないので今は考えない！
	for _, method := range methods {
		if method.Type != httpMethod {
			continue
		}

		// methodにURLがないパターン
		if requestedMethodURL == "" && method.URL == "" {
			returnMethod = method
			break
		}

		// メソッド名が指定されているパターン(例："/foo/{bar}")
		// methodのURLからパラメータ部分を削除
		methodURL := re.ReplaceAllString(method.URL, "")
		if methodURL != "" {
			methodURL = methodURL + "/"
		}
		// パラメータ以外のmethodのURL(メソッド名)が合致していれば返却
		if methodURL != "" && strings.Contains(requestedMethodURL, methodURL) {
			returnMethod = method
			break
		} else {
			// パラメータのみ指定されているパターン(例："/{bar}")
			params := re.FindAllStringSubmatch(method.URL, -1)
			paramsCount := len(params)
			methodURLSlushCount := strings.Count(method.URL, "/")
			requestedSlushCount := strings.Count(requestedMethodURL, "/")

			// MethodURL,リクエスト内の"/"の数と、Methodのパラメータの数が同じであれば返却
			if methodURLSlushCount == requestedSlushCount && requestedSlushCount == paramsCount {
				returnMethod = method
				break
			}
		}
	}

	if returnMethod.ID == "" {
		return returnMethod, errors.New("対象のメソッドが見つかりません")
	}

	return returnMethod, err
}

// getRequestedURLParameter リクエストされたURLパラメータとKeyを取得します
func getRequestedURLParameter(requestedURL string, apiURL string, methodURL string, modelSchema string) (string, interface{}, error) {
	var key string
	var value string

	if methodURL != "" {
		// {}で囲まれた中身のみ取得
		re := regexp.MustCompile(`/\{.+?\}`)
		key = re.FindString(methodURL)
		key = strings.Replace(key, "/{", "", -1)
		key = strings.Replace(key, "}", "", -1)

		// /{}で囲まれた箇所を削除
		methodURL = re.ReplaceAllString(methodURL, "")
	}

	if key == "" {
		return "", "", nil
	}

	// リクエストされたパラメータを取得
	value = strings.Replace(requestedURL, apiURL+methodURL, "", 1)

	// 最初の/を削除する
	if value != "" {
		value = value[1:]
	}

    var schemaMap map[string]interface{}
    err := json.Unmarshal([]byte(modelSchema), &schemaMap)
    if err != nil {
        fmt.Println(err)
        return "", "", errors.New("model schema Unmarshal failed")
	}

	property := schemaMap["properties"].(map[string]interface{})[key].(map[string]interface{})
	keyType := property["type"].(string)

	switch keyType {
	case "number":
		intValue, err := strconv.Atoi(value)
		if err != nil {
			errMsg := fmt.Sprintf("invalid convert number%s", value)
			return "", "", errors.New(errMsg)
		}
		return key, intValue, nil
	case "integer":
		intValue, err := strconv.Atoi(value)
		if err != nil {
			errMsg := fmt.Sprintf("invalid convert number%s", value)
			return "", "", errors.New(errMsg)
		}
		return key, intValue, nil
    case "string":
        return key, value, nil
	default:
		errMsg := fmt.Sprintf("not an allowed id type%s", keyType)
        return "", "", errors.New(errMsg)
    }
}

// getRequestedSchemaValidate リクエストBodyがSchemaに則っているか検証します
func getRequestedSchemaValidate(modelSchema string, requestBody []byte) error {

	schemaLoader := gojsonschema.NewStringLoader(modelSchema)
	documentLoader := gojsonschema.NewStringLoader(string(requestBody))

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return ErrInvalidRequest
	}

	if !result.Valid() {
		errMsg := ""
		for _, desc := range result.Errors() {
			errMsg = errMsg + desc.String()
		}
		return errors.New(errMsg)
	}

	return nil
}

// getPropertyValue リクエストBody内を、keyから項目の値を取得します
func getPropertyValue(key string, body []byte) (interface{}, error) {
	var bsonBody bson.M

	err := bson.UnmarshalExtJSON(body, false, &bsonBody)
	if err != nil {
		return "", err
	}

	if bsonBody[key] == nil {
		return "", errors.New("target property is not found")
	}

	return bsonBody[key], nil
}
