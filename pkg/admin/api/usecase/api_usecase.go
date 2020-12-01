package usecase

import (
	"net/http"

	"github.com/Hajime3778/api-creator-backend/pkg/admin/api/repository"
	"github.com/Hajime3778/api-creator-backend/pkg/domain"
	"github.com/google/uuid"
)

// APIUsecase Interface
type APIUsecase interface {
	GetAll() ([]domain.API, error)
	GetByID(id string) (domain.API, error)
	Create(api domain.API) (int, string, error)
	Update(api domain.API) (int, error)
	Delete(id string) error
}

type apiUsecase struct {
	repo repository.APIRepository
}

// NewAPIUsecase APIUsecaseインターフェイスを表すオブジェクトを作成します
func NewAPIUsecase(repo repository.APIRepository) APIUsecase {
	return &apiUsecase{
		repo: repo,
	}
}

// GetAll 複数のAPIを取得します
func (u *apiUsecase) GetAll() ([]domain.API, error) {
	return u.repo.GetAll()
}

// GetByID APIを1件取得します
func (u *apiUsecase) GetByID(id string) (domain.API, error) {
	return u.repo.GetByID(id)
}

// Create APIを作成します
func (u *apiUsecase) Create(api domain.API) (int, string, error) {
	if api.ID == "" {
		id, _ := uuid.NewRandom()
		api.ID = id.String()
	}
	id, err := u.repo.Create(api)
	if err != nil {
		return http.StatusInternalServerError, "", nil
	}
	return http.StatusCreated, id, nil
}

// Update APIを更新します
func (u *apiUsecase) Update(api domain.API) (int, error) {
	err := u.repo.Update(api)
	if err != nil {
		return http.StatusInternalServerError, nil
	}
	return http.StatusOK, nil
}

// Delete APIを削除します
func (u *apiUsecase) Delete(id string) error {
	return u.repo.Delete(id)
}
