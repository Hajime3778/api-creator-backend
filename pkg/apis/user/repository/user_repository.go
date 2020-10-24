package repository

import (
	"github.com/Hajime3778/api-creator-backend/pkg/domain"
	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/database"

	"github.com/jinzhu/gorm"
)

// UserRepository repository
type UserRepository interface {
	GetAll() ([]domain.User, error)
	GetByID(id int) (domain.User, error)
	Create(user domain.User) (int, error)
	Update(user domain.User) error
	Delete(id int) error
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository is init for UserController
func NewUserRepository(db *database.DB) UserRepository {
	return &userRepository{
		db: db.Connection,
	}
}

// GetAll Get all usersdata
func (r *userRepository) GetAll() ([]domain.User, error) {
	users := []domain.User{}
	err := r.db.Find(&users).Error

	return users, err
}

// GetByID Get single usersdata
func (r *userRepository) GetByID(id int) (domain.User, error) {
	user := domain.User{}
	err := r.db.First(&user, id).Error

	return user, err
}

// Create Add user
func (r *userRepository) Create(user domain.User) (int, error) {
	err := r.db.Create(&user).Error
	id := user.ID
	return id, err
}

// Update Update user
func (r *userRepository) Update(user domain.User) error {
	targetUser := domain.User{}
	if err := r.db.First(&targetUser, user.ID).Error; err != nil {
		return err
	}

	return r.db.Save(&user).Error
}

// Delete Delete userdata
func (r *userRepository) Delete(id int) error {
	user := domain.User{}

	if id <= 0 {
		return gorm.ErrRecordNotFound
	}

	user.ID = id
	result := r.db.Delete(&user)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}
