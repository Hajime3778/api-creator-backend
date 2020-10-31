package usecase

import (
	"github.com/Hajime3778/api-creator-backend/pkg/apis/api/repository"
	"github.com/Hajime3778/api-creator-backend/pkg/domain"
	"github.com/google/uuid"
)

// APIUsecase Interface
type APIUsecase interface {
	GetAll() ([]domain.API, error)
	GetByID(id string) (domain.API, error)
	Create(api domain.API) (string, error)
	Update(api domain.API) error
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
func (u *apiUsecase) Create(api domain.API) (string, error) {
	if api.ID == "" {
		id, _ := uuid.NewRandom()
		api.ID = id.String()
	}
	return u.repo.Create(api)
}

// Update APIを更新します
func (u *apiUsecase) Update(api domain.API) error {
	return u.repo.Update(api)
}

// Delete APIを削除します
func (u *apiUsecase) Delete(id string) error {
	return u.repo.Delete(id)
}
