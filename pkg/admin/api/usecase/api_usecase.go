package usecase

import (
	"errors"
	"net/http"

	_apiRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/api/repository"
	_methodRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/method/repository"
	_modelRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/model/repository"
	"github.com/Hajime3778/api-creator-backend/pkg/domain"
	"github.com/Hajime3778/api-creator-backend/pkg/validation"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// APIUsecase Interface
type APIUsecase interface {
	GetAll() ([]domain.API, error)
	GetByID(id string) (domain.API, error)
	Create(api domain.API) (int, string, error)
	Update(api domain.API) (int, error)
	Delete(id string) (int, error)
}

type apiUsecase struct {
	apiRepo    _apiRepository.APIRepository
	methodRepo _methodRepository.MethodRepository
	modelRepo  _modelRepository.ModelRepository
}

// NewAPIUsecase APIUsecaseインターフェイスを表すオブジェクトを作成します
func NewAPIUsecase(apiRepo _apiRepository.APIRepository, methodRepo _methodRepository.MethodRepository, modelRepo _modelRepository.ModelRepository) APIUsecase {
	return &apiUsecase{
		apiRepo:    apiRepo,
		methodRepo: methodRepo,
		modelRepo:  modelRepo,
	}
}

// GetAll 複数のAPIを取得します
func (u *apiUsecase) GetAll() ([]domain.API, error) {
	return u.apiRepo.GetAll()
}

// GetByID APIを1件取得します
func (u *apiUsecase) GetByID(id string) (domain.API, error) {
	return u.apiRepo.GetByID(id)
}

// Create APIを作成します
func (u *apiUsecase) Create(api domain.API) (int, string, error) {
	if api.ID == "" {
		id, _ := uuid.NewRandom()
		api.ID = id.String()
	}
	if !validation.IsHalfWidthOnly(api.URL) {
		return http.StatusBadRequest, "", errors.New("url is halfwidth only")
	}
	id, err := u.apiRepo.Create(api)
	if err != nil {
		return http.StatusInternalServerError, "", nil
	}
	return http.StatusCreated, id, nil
}

// Update APIを更新します
func (u *apiUsecase) Update(api domain.API) (int, error) {
	err := u.apiRepo.Update(api)
	if !validation.IsHalfWidthOnly(api.URL) {
		return http.StatusBadRequest, errors.New("url is halfwidth only")
	}
	if err != nil {
		return http.StatusInternalServerError, nil
	}
	return http.StatusOK, nil
}

// Delete APIを削除します(関連するメソッド、モデルも含めて)
func (u *apiUsecase) Delete(id string) (int, error) {
	if _, err := u.apiRepo.GetByID(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return http.StatusNotFound, err
		}
		return http.StatusInternalServerError, err
	}

	methods, err := u.methodRepo.GetListByAPIID(id)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return http.StatusInternalServerError, err
		}
	}

	model, err := u.modelRepo.GetByAPIID(id)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return http.StatusInternalServerError, err
		}
	}

	err = u.apiRepo.Delete(id, methods, model)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusNoContent, nil
}
