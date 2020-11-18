package usecase

import (
	"errors"
	"log"
	"net/http"
	"regexp"
	"strings"

	_apiRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/api/repository"
	_methodRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/method/repository"
	_modelRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/model/repository"
	_apiserverRepository "github.com/Hajime3778/api-creator-backend/pkg/apiserver/repository"
	"github.com/Hajime3778/api-creator-backend/pkg/domain"
	"github.com/xeipuuv/gojsonschema"
)

var (
	// ErrAPINotFound "api not found"
	ErrAPINotFound = errors.New("api not found")
	// ErrInvalidRequest "invalid request"
	ErrInvalidRequest = errors.New("invalid request")
)

// APIServerUsecase Interface
type APIServerUsecase interface {
	RequestDocumentServer(httpMethod string, url string, body []byte) (string, int, error)
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
func (u *apiServerUsecase) RequestDocumentServer(httpMethod string, url string, body []byte) (string, int, error) {
	api, err := u.apiRepo.GetByURL(url)
	if err != nil {
		// APIが見つかりません
		return "", http.StatusNotFound, err
	}

	// 対象のメソッドを取得
	method, err := u.getRequestedMethod(httpMethod, url, api)

	if err != nil {
		// APIが見つかりません
		return "", http.StatusNotFound, ErrAPINotFound
	}

	if method.Type != httpMethod {
		// APIが見つかりません
		return "", http.StatusNotFound, ErrAPINotFound
	}

	// リクエストされたパラメータを取得
	param := getRequestedURLParameter(url, api.URL, method.URL)

	model, err := u.modelRepo.GetByAPIID(api.ID)

	schemaLoader := gojsonschema.NewStringLoader(model.Schema)
	documentLoader := gojsonschema.NewStringLoader(string(body))

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		// リクエストされたJsonの形式が違います
		return "", http.StatusBadRequest, ErrInvalidRequest
	}

	if !result.Valid() {
		// モデルの形式が違います
		return "", http.StatusBadRequest, ErrInvalidRequest

		// fmt.Printf("The document is not valid. see errors :\n")
		// for _, desc := range result.Errors() {
		// 	fmt.Printf("- %s\n", desc)
		// }
		// panic(result.Errors())
	}

	switch method.Type {
	case "GET":
		log.Println("GET")
	case "POST":
		u.apiserverRepo.Create(body)
	case "PUT":
		log.Println("PUT")
	case "DELETE":
		log.Println("DELETE")
	}

	return param, http.StatusOK, err
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

// getRequestedURLParameter リクエストされたURLパラメータを取得します
func getRequestedURLParameter(requestedURL string, apiURL string, methodURL string) string {
	if methodURL != "" {
		// /{}で囲まれた箇所を削除
		re := regexp.MustCompile(`/\{.+?\}`)
		methodURL = re.ReplaceAllString(methodURL, "")
	}

	// リクエストされたパラメータを取得
	param := strings.Replace(requestedURL, apiURL+methodURL, "", 1)

	// 最初の/を削除する
	if param != "" {
		param = param[1:]
	}

	return param
}
