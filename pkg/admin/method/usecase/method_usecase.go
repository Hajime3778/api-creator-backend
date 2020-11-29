package usecase

import (
	"errors"
	"log"
	"net/http"
	"regexp"
	"strings"

	_apiRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/api/repository"
	_methodRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/method/repository"
	"github.com/Hajime3778/api-creator-backend/pkg/domain"
	"github.com/google/uuid"
)

// MethodUsecase Interface
type MethodUsecase interface {
	GetAll() ([]domain.Method, error)
	GetByID(id string) (domain.Method, error)
	GetListByAPIID(apiID string) ([]domain.Method, error)
	Create(method domain.Method) (int, string, error)
	CreateDefaultMethods(apiID string) (int, error)
	Update(method domain.Method) (int, error)
	Delete(id string) error
}

type methodUsecase struct {
	apiRepo _apiRepository.APIRepository
	repo    _methodRepository.MethodRepository
}

// NewMethodUsecase MethodUsecaseインターフェイスを表すオブジェクトを作成します
func NewMethodUsecase(apiRepo _apiRepository.APIRepository, repo _methodRepository.MethodRepository) MethodUsecase {
	return &methodUsecase{
		apiRepo: apiRepo,
		repo:    repo,
	}
}

// GetAll 複数のMethodを取得します
func (u *methodUsecase) GetAll() ([]domain.Method, error) {
	return u.repo.GetAll()
}

// GetByID 1件のMethodを取得します
func (u *methodUsecase) GetByID(id string) (domain.Method, error) {
	return u.repo.GetByID(id)
}

// GetListByAPIID MethodをAPIIDで複数取得します
func (u *methodUsecase) GetListByAPIID(apiID string) ([]domain.Method, error) {
	return u.repo.GetListByAPIID(apiID)
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
	id, err := u.repo.Create(method)
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	return http.StatusCreated, id, nil
}

// Create Methodを作成します
func (u *methodUsecase) CreateDefaultMethods(apiID string) (int, error) {
	api, err := u.apiRepo.GetByID(apiID)
	log.Println(api)
	log.Println(err)
	return http.StatusNotImplemented, nil
}

// Update Methodを更新します。
func (u *methodUsecase) Update(method domain.Method) (int, error) {
	err := u.validateMethodURL(method)
	if err != nil {
		return http.StatusBadRequest, err
	}
	err = u.repo.Update(method)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, err
}

// Delete Methodを削除します
func (u *methodUsecase) Delete(id string) error {
	return u.repo.Delete(id)
}

// validateMethodURL メソッドURLを検証します
func (u *methodUsecase) validateMethodURL(method domain.Method) error {
	newMethod := method

	methods, err := u.repo.GetListByAPIIDAndType(newMethod.APIID, newMethod.Type)
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
