package usecase

import (
	"go-mysql-api/pkg/domain"
	"go-mysql-api/pkg/user/repository"
)

// UserUsecase usecase
type UserUsecase interface {
	GetAll() ([]domain.User, error)
	GetByID(id int) (domain.User, error)
	Create(user domain.User) (int, error)
	Update(user domain.User) error
	Delete(id int) error
}

// userUsecase usecase
type userUsecase struct {
	repo repository.UserRepository
}

// NewUserUsecase is init for UserUsecase
func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

// GetAll 複数のUserを取得します
func (u *userUsecase) GetAll() ([]domain.User, error) {
	return u.repo.GetAll()
}

// GetByID 1件のUserを取得します
func (u *userUsecase) GetByID(id int) (domain.User, error) {
	return u.repo.GetByID(id)
}

// Create Userを作成します
func (u *userUsecase) Create(user domain.User) (int, error) {
	return u.repo.Create(user)
}

// Update Userを更新します。
func (u *userUsecase) Update(user domain.User) error {
	return u.repo.Update(user)
}

// Delete Userを削除します
func (u *userUsecase) Delete(id int) error {
	return u.repo.Delete(id)
}
