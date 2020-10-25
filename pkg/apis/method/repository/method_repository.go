package repository

import (
	"github.com/Hajime3778/api-creator-backend/pkg/domain"
	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/database"

	"github.com/jinzhu/gorm"
)

// MethodRepository Interface
type MethodRepository interface {
	GetAll() ([]domain.Method, error)
	GetByID(id string) (domain.Method, error)
	Create(method domain.Method) (string, error)
	Update(method domain.Method) error
	Delete(id string) error
}

type methodRepository struct {
	db *gorm.DB
}

// NewMethodRepository MethodRepositoryインターフェイスを表すオブジェクトを作成します
func NewMethodRepository(db *database.DB) MethodRepository {
	return &methodRepository{
		db: db.Connection,
	}
}

// GetAll すべてのMethodを取得します
func (r *methodRepository) GetAll() ([]domain.Method, error) {
	methods := []domain.Method{}
	err := r.db.Find(&methods).Error

	return methods, err
}

// GetByID Methodを1件取得します
func (r *methodRepository) GetByID(id string) (domain.Method, error) {
	method := domain.Method{}
	err := r.db.Where("id = ?", id).First(&method).Error

	return method, err
}

// Create Methodを追加します
func (r *methodRepository) Create(method domain.Method) (string, error) {
	err := r.db.Create(&method).Error
	id := method.ID
	return id, err
}

// Update Methodを更新します
func (r *methodRepository) Update(method domain.Method) error {
	targetMethod := domain.Method{}

	err := r.db.Where("id = ?", method.ID).First(&targetMethod).Error
	if err != nil {
		return err
	}

	return r.db.Save(&method).Error
}

// Delete Methodを削除します
func (r *methodRepository) Delete(id string) error {
	method := domain.Method{}

	method.ID = id
	result := r.db.Delete(&method)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}
