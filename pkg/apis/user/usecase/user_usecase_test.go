package usecase_test

import (
	"testing"
	"time"

	"github.com/Hajime3778/api-creator-backend/pkg/apis/user/usecase"
	"github.com/Hajime3778/api-creator-backend/pkg/domain"
	"github.com/Hajime3778/api-creator-backend/test/mocks"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	mockUsers := make([]domain.User, 0)
	mockUser := domain.User{}
	mockUser.ID = 1
	mockUser.Name = "mockuser"
	mockUser.Email = "mock@mock.com"
	mockUser.CreatedAt = time.Now()
	mockUser.UpdatedAt = time.Now()
	mockUsers = append(mockUsers, mockUser)

	// モック
	mockUserRepo := new(mocks.UserRepository)

	t.Run("test1", func(t *testing.T) {
		mockUserRepo.On("GetAll").Return(mockUsers, nil).Once()
		usecase := usecase.NewUserUsecase(mockUserRepo)

		users, err := usecase.GetAll()

		assert.NoError(t, err)
		assert.NotNil(t, users)

		mockUserRepo.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	mockUser := domain.User{}
	mockUser.ID = 1
	mockUser.Name = "mockuser"
	mockUser.Email = "mock@mock.com"
	mockUser.CreatedAt = time.Now()
	mockUser.UpdatedAt = time.Now()

	mockUserRepo := new(mocks.UserRepository)

	t.Run("test1", func(t *testing.T) {
		mockUserRepo.On("GetByID", mockUser.ID).Return(mockUser, nil).Once()
		usecase := usecase.NewUserUsecase(mockUserRepo)

		user, err := usecase.GetByID(mockUser.ID)

		assert.NoError(t, err)
		assert.NotNil(t, user)

		mockUserRepo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	mockUser := domain.User{}
	mockUser.ID = 1
	mockUser.Name = "mockuser"
	mockUser.Email = "mock@mock.com"

	mockUserRepo := new(mocks.UserRepository)

	t.Run("test1", func(t *testing.T) {
		mockUserRepo.On("Create", mockUser).Return(nil).Once()
		usecase := usecase.NewUserUsecase(mockUserRepo)

		_, err := usecase.Create(mockUser)

		assert.NoError(t, err)

		mockUserRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockUser := domain.User{}
	mockUser.ID = 1
	mockUser.Name = "mockuser"
	mockUser.Email = "mock@mock.com"

	// モック
	mockUserRepo := new(mocks.UserRepository)

	t.Run("test1", func(t *testing.T) {
		mockUserRepo.On("Update", mockUser).Return(nil).Once()
		usecase := usecase.NewUserUsecase(mockUserRepo)

		err := usecase.Update(mockUser)

		assert.NoError(t, err)

		mockUserRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockUser := domain.User{}
	mockUser.ID = 1
	mockUser.Name = "mockuser"
	mockUser.Email = "mock@mock.com"
	mockUserRepo := new(mocks.UserRepository)

	t.Run("test1", func(t *testing.T) {
		mockUserRepo.On("Delete", mockUser.ID).Return(nil).Once()
		usecase := usecase.NewUserUsecase(mockUserRepo)

		err := usecase.Delete(mockUser.ID)

		assert.NoError(t, err)

		mockUserRepo.AssertExpectations(t)
	})
}
