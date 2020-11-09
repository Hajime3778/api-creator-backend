package repository

import (
	"github.com/Hajime3778/api-creator-backend/pkg/domain"

	"github.com/jinzhu/gorm"
)

// APIRepository Interface
type APIRepository interface {
	GetAll() ([]domain.API, error)
	GetByID(id string) (domain.API, error)
	Create(api domain.API) (string, error)
	Update(api domain.API) error
	Delete(id string) error
}

type apiRepository struct {
	db *gorm.DB
}

// NewAPIRepository APIRepositoryインターフェイスを表すオブジェクトを作成します
func NewAPIRepository(db *gorm.DB) APIRepository {
	return &apiRepository{
		db: db,
	}
}

// GetAll すべてのAPIを取得します
func (r *apiRepository) GetAll() ([]domain.API, error) {
	apis := []domain.API{}
	err := r.db.Find(&apis).Error

	return apis, err
}

// GetByID APIを1件取得します
func (r *apiRepository) GetByID(id string) (domain.API, error) {
	api := domain.API{}
	err := r.db.Where("id = ?", id).First(&api).Error

	return api, err
}

// Create APIを作成します
func (r *apiRepository) Create(api domain.API) (string, error) {
	err := r.db.Create(&api).Error
	id := api.ID
	return id, err
}

// Update APIを更新します
func (r *apiRepository) Update(api domain.API) error {
	targetAPI := domain.API{}

	err := r.db.Where("id = ?", api.ID).First(&targetAPI).Error
	if err != nil {
		return err
	}

	return r.db.Save(&api).Error
}

// Delete APIを削除します
func (r *apiRepository) Delete(id string) error {
	api := domain.API{}

	api.ID = id
	result := r.db.Delete(&api)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}
