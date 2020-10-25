package usecase

import (
	"github.com/Hajime3778/api-creator-backend/pkg/apis/method/repository"
	"github.com/Hajime3778/api-creator-backend/pkg/domain"
)

// MethodUsecase Interface
type MethodUsecase interface {
	GetAll() ([]domain.Method, error)
	GetByID(id string) (domain.Method, error)
	Create(method domain.Method) (string, error)
	Update(method domain.Method) error
	Delete(id string) error
}

type methodUsecase struct {
	repo repository.MethodRepository
}

// NewMethodUsecase MethodUsecaseインターフェイスを表すオブジェクトを作成します
func NewMethodUsecase(repo repository.MethodRepository) MethodUsecase {
	return &methodUsecase{
		repo: repo,
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

// Create Methodを作成します
func (u *methodUsecase) Create(method domain.Method) (string, error) {
	return u.repo.Create(method)
}

// Update Methodを更新します。
func (u *methodUsecase) Update(method domain.Method) error {
	return u.repo.Update(method)
}

// Delete Methodを削除します
func (u *methodUsecase) Delete(id string) error {
	return u.repo.Delete(id)
}
