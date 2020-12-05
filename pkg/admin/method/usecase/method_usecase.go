package usecase

import (
	"errors"
	"net/http"
	"regexp"
	"strings"

	_apiRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/api/repository"
	_methodRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/method/repository"
	_modelRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/model/repository"
	"github.com/Hajime3778/api-creator-backend/pkg/domain"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// MethodUsecase Interface
type MethodUsecase interface {
	GetAll() ([]domain.Method, error)
	GetByID(id string) (domain.Method, error)
	GetListByAPIID(apiID string) ([]domain.Method, error)
	Create(method domain.Method) (int, string, error)
	CreateDefaultMethods(apiID string) (int, []domain.Method, error)
	Update(method domain.Method) (int, error)
	Delete(id string) error
}

type methodUsecase struct {
	apiRepo    _apiRepository.APIRepository
	methodRepo _methodRepository.MethodRepository
	modelRepo  _modelRepository.ModelRepository
}

// NewMethodUsecase MethodUsecaseインターフェイスを表すオブジェクトを作成します
func NewMethodUsecase(apiRepo _apiRepository.APIRepository, methodRepo _methodRepository.MethodRepository, modelRepo _modelRepository.ModelRepository) MethodUsecase {
	return &methodUsecase{
		apiRepo:    apiRepo,
		modelRepo:  modelRepo,
		methodRepo: methodRepo,
	}
}

// GetAll 複数のMethodを取得します
func (u *methodUsecase) GetAll() ([]domain.Method, error) {
	return u.methodRepo.GetAll()
}

// GetByID 1件のMethodを取得します
func (u *methodUsecase) GetByID(id string) (domain.Method, error) {
	return u.methodRepo.GetByID(id)
}

// GetListByAPIID MethodをAPIIDで複数取得します
func (u *methodUsecase) GetListByAPIID(apiID string) ([]domain.Method, error) {
	return u.methodRepo.GetListByAPIID(apiID)
}

// Create Methodを作成します
func (u *methodUsecase) Create(method domain.Method) (int, string, error) {
	if method.ID == "" {
		id, _ := uuid.NewRandom()
		method.ID = id.String()
	}
	err := u.validateMethodURL(method)
	if err != nil {
		return http.StatusBadRequest, "", err
	}
	id, err := u.methodRepo.Create(method)
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	return http.StatusCreated, id, nil
}

// CreateDefaultMethods デフォルトのCRUDMethodを作成します
func (u *methodUsecase) CreateDefaultMethods(apiID string) (int, []domain.Method, error) {
	_, err := u.apiRepo.GetByID(apiID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return http.StatusNotFound, nil, err
		}
		return http.StatusInternalServerError, nil, err
	}

	model, err := u.modelRepo.GetByAPIID(apiID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return http.StatusNotFound, nil, err
		}
		return http.StatusInternalServerError, nil, err
	}

	keys, err := model.GetKeyNames()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	var methods []domain.Method
	for i := 0; i < 5; i++ {
		id, _ := uuid.NewRandom()
		initMethod := domain.Method{
			ID:      id.String(),
			APIID:   apiID,
			IsArray: false,
		}
		methods = append(methods, initMethod)
	}
	// 本来は、1SQLの実行でやるのがベスト
	// RollBackをgormを使用して行う必要あり
	// getAll メソッド作成
	methods[0].Type = "GET"
	methods[0].Description = "すべての" + model.Name + "を取得します。"
	methods[0].IsArray = true

	// getOne メソッド作成
	methods[1].Type = "GET"
	methods[1].Description = keys[0] + "から1件の" + model.Name + "を取得します。"
	methods[1].URL = "/{" + keys[0] + "}"

	// create メソッド作成
	methods[2].Type = "POST"
	methods[2].Description = model.Name + "を1件作成します。"

	// update メソッド作成
	methods[3].Type = "PUT"
	methods[3].Description = model.Name + "を1件更新します。"

	// delete メソッド作成
	methods[4].Type = "DELETE"
	methods[4].Description = model.Name + "を1件削除します。"
	methods[4].URL = "/{" + keys[0] + "}"

	for _, method := range methods {
		_, err := u.methodRepo.Create(method)
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}
	}

	createdMethods, err := u.methodRepo.GetListByAPIID(apiID)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, createdMethods, nil
}

// Update Methodを更新します。
func (u *methodUsecase) Update(method domain.Method) (int, error) {
	err := u.validateMethodURL(method)
	if err != nil {
		return http.StatusBadRequest, err
	}
	err = u.methodRepo.Update(method)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, err
}

// Delete Methodを削除します
func (u *methodUsecase) Delete(id string) error {
	return u.methodRepo.Delete(id)
}

// validateMethodURL メソッドURLを検証します
func (u *methodUsecase) validateMethodURL(method domain.Method) error {
	newMethod := method

	methods, err := u.methodRepo.GetListByAPIIDAndType(newMethod.APIID, newMethod.Type)
	if err != nil {
		return nil
	}

	if method.URL != "" {
		ret := regexp.MustCompile(`^/.+?[^/]$`)
		if !ret.MatchString(method.URL) {
			return errors.New("URLが正しい形式ではありません")
		}
	}

	// /{}で囲まれた箇所
	re := regexp.MustCompile(`/\{.+?\}`)

	paramNames := re.FindAllStringSubmatch(newMethod.URL, -1)

	// TODO:パラメータ複数には現時点では対応しない
	if len(paramNames) > 1 {
		return errors.New("パラメータを複数指定することはできません")
	}
	for _, paramName := range paramNames {
		if paramName[0] == "" || strings.Count(paramName[0], "/") > 1 {
			return errors.New("URLが正しい形式ではありません")
		}
	}
	newMethodParamCount := len(paramNames)
	newMethodSlushCount := strings.Count(newMethod.URL, "/")

	for _, method := range methods {
		// 同じメソッドの場合はスルー
		if method.ID == newMethod.ID {
			continue
		}
		methodParamCount := len(re.FindAllStringSubmatch(method.URL, -1))
		methodSlushCount := strings.Count(method.URL, "/")

		//スラッシュの数とパラメータの数が同じメソッドがすでに存在する場合
		if methodParamCount == newMethodParamCount &&
			methodSlushCount == newMethodSlushCount {
			return errors.New("同じHTTPメソッド、URLのメソッドがすでに存在しています。" + "\n：" + method.URL)
		}
	}
	return nil
}
