package repository

import (
	"github.com/Hajime3778/api-creator-backend/pkg/domain"

	"github.com/jinzhu/gorm"
)

// ModelRepository Interface
type ModelRepository interface {
	GetAll() ([]domain.Model, error)
	GetByID(id string) (domain.Model, error)
	GetByAPIID(apiID string) (domain.Model, error)
	Create(model domain.Model) (string, error)
	Update(model domain.Model) error
	Delete(id string) error
}

type modelRepository struct {
	db *gorm.DB
}

// NewModelRepository ModelRepositoryインターフェイスを表すオブジェクトを作成します
func NewModelRepository(db *gorm.DB) ModelRepository {
	return &modelRepository{
		db: db,
	}
}

// GetAll すべてのModelを取得します
func (r *modelRepository) GetAll() ([]domain.Model, error) {
	models := []domain.Model{}
	err := r.db.Find(&models).Error

	return models, err
}

// GetByID Modelを1件取得します
func (r *modelRepository) GetByID(id string) (domain.Model, error) {
	model := domain.Model{}
	err := r.db.Where("id = ?", id).First(&model).Error

	return model, err
}

// GetByID Modelを1件取得します
func (r *modelRepository) GetByAPIID(apiID string) (domain.Model, error) {
	model := domain.Model{}
	err := r.db.Where("api_id = ?", apiID).First(&model).Error

	return model, err
}

// Create Modelを追加します
func (r *modelRepository) Create(model domain.Model) (string, error) {
	err := r.db.Create(&model).Error
	id := model.ID
	return id, err
}

// Update Modelを更新します
func (r *modelRepository) Update(model domain.Model) error {
	targetModel := domain.Model{}

	err := r.db.Where("id = ?", model.ID).First(&targetModel).Error
	if err != nil {
		return err
	}

	return r.db.Save(&model).Error
}

// Delete Modelを削除します
func (r *modelRepository) Delete(id string) error {
	model := domain.Model{}

	model.ID = id
	result := r.db.Delete(&model)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}
