package usecase

import (
	"github.com/Hajime3778/api-creator-backend/pkg/admin/model/repository"
	"github.com/Hajime3778/api-creator-backend/pkg/domain"
	"github.com/google/uuid"
)

// ModelUsecase Interface
type ModelUsecase interface {
	GetAll() ([]domain.Model, error)
	GetByID(id string) (domain.Model, error)
	GetByAPIID(apiID string) (domain.Model, error)
	Create(model domain.Model) (string, error)
	Update(model domain.Model) error
	Delete(id string) error
}

type modelUsecase struct {
	repo repository.ModelRepository
}

// NewModelUsecase ModelUsecaseインターフェイスを表すオブジェクトを作成します
func NewModelUsecase(repo repository.ModelRepository) ModelUsecase {
	return &modelUsecase{
		repo: repo,
	}
}

// GetAll 複数のModelを取得します
func (u *modelUsecase) GetAll() ([]domain.Model, error) {
	return u.repo.GetAll()
}

// GetByID 1件のModelを取得します
func (u *modelUsecase) GetByID(id string) (domain.Model, error) {
	return u.repo.GetByID(id)
}

// GetByAPIID APIIDから1件のModelを取得します
func (u *modelUsecase) GetByAPIID(apiID string) (domain.Model, error) {
	return u.repo.GetByAPIID(apiID)
}

// Create Modelを作成します
func (u *modelUsecase) Create(model domain.Model) (string, error) {
	if model.ID == "" {
		id, _ := uuid.NewRandom()
		model.ID = id.String()
	}
	// JsonSchemaが正しい形式か検証
	err := model.ValidateSchema()
	if err != nil {
		return "", err
	}

	return u.repo.Create(model)
}

// Update Modelを更新します。
func (u *modelUsecase) Update(model domain.Model) error {
	// JsonSchemaが正しい形式か検証
	err := model.ValidateSchema()
	if err != nil {
		return err
	}

	return u.repo.Update(model)
}

// Delete Modelを削除します
func (u *modelUsecase) Delete(id string) error {
	return u.repo.Delete(id)
}
