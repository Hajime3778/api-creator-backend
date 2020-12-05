package usecase_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/Hajime3778/api-creator-backend/pkg/admin/api/usecase"
	"github.com/Hajime3778/api-creator-backend/pkg/domain"
	"github.com/Hajime3778/api-creator-backend/test/mocks"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	apiId, _ := uuid.NewRandom()

	mockAPIs := make([]domain.API, 0)
	mockAPI := domain.API{}
	mockAPI.ID = apiId.String()
	mockAPI.Name = "name"
	mockAPI.URL = "url"
	mockAPI.Description = "test"
	mockAPI.CreatedAt = time.Now()
	mockAPI.UpdatedAt = time.Now()
	mockAPIs = append(mockAPIs, mockAPI)

	// モック
	mockAPIRepo := new(mocks.APIRepository)
	mockMethodRepo := new(mocks.MethodRepository)
	mockModelRepo := new(mocks.ModelRepository)

	t.Run("test1", func(t *testing.T) {
		mockAPIRepo.On("GetAll").Return(mockAPIs, nil).Once()
		usecase := usecase.NewAPIUsecase(mockAPIRepo, mockMethodRepo, mockModelRepo)

		apis, err := usecase.GetAll()

		assert.NoError(t, err)
		assert.NotNil(t, apis)

		mockAPIRepo.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	apiId, _ := uuid.NewRandom()

	mockAPI := domain.API{}
	mockAPI.ID = apiId.String()
	mockAPI.Name = "name"
	mockAPI.URL = "url"
	mockAPI.Description = "test"

	mockAPI.CreatedAt = time.Now()
	mockAPI.UpdatedAt = time.Now()

	// モック
	mockAPIRepo := new(mocks.APIRepository)
	mockMethodRepo := new(mocks.MethodRepository)
	mockModelRepo := new(mocks.ModelRepository)

	t.Run("test1", func(t *testing.T) {
		mockAPIRepo.On("GetByID", mockAPI.ID).Return(mockAPI, nil).Once()
		usecase := usecase.NewAPIUsecase(mockAPIRepo, mockMethodRepo, mockModelRepo)

		api, err := usecase.GetByID(mockAPI.ID)

		assert.NoError(t, err)
		assert.NotNil(t, api)

		mockAPIRepo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	apiId, _ := uuid.NewRandom()

	mockAPI := domain.API{}
	mockAPI.ID = apiId.String()
	mockAPI.Name = "name"
	mockAPI.URL = "url"
	mockAPI.Description = "test"

	// モック
	mockAPIRepo := new(mocks.APIRepository)
	mockMethodRepo := new(mocks.MethodRepository)
	mockModelRepo := new(mocks.ModelRepository)

	t.Run("test1", func(t *testing.T) {
		mockAPIRepo.On("Create", mockAPI).Return(nil).Once()
		usecase := usecase.NewAPIUsecase(mockAPIRepo, mockMethodRepo, mockModelRepo)

		status, _, err := usecase.Create(mockAPI)

		assert.NoError(t, err)
		assert.Equal(t, status, http.StatusCreated)

		mockAPIRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	apiId, _ := uuid.NewRandom()

	mockAPI := domain.API{}
	mockAPI.ID = apiId.String()
	mockAPI.Name = "name"
	mockAPI.URL = "url"
	mockAPI.Description = "test"

	// モック
	mockAPIRepo := new(mocks.APIRepository)
	mockMethodRepo := new(mocks.MethodRepository)
	mockModelRepo := new(mocks.ModelRepository)

	t.Run("test1", func(t *testing.T) {
		mockAPIRepo.On("Update", mockAPI).Return(nil).Once()
		usecase := usecase.NewAPIUsecase(mockAPIRepo, mockMethodRepo, mockModelRepo)

		status, err := usecase.Update(mockAPI)

		assert.NoError(t, err)
		assert.Equal(t, status, http.StatusOK)

		mockAPIRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	apiId, _ := uuid.NewRandom()

	mockAPI := domain.API{}
	mockAPI.ID = apiId.String()
	mockAPI.Name = "name"
	mockAPI.URL = "url"
	mockAPI.Description = "test"

	// モック
	mockAPIRepo := new(mocks.APIRepository)
	mockMethodRepo := new(mocks.MethodRepository)
	mockModelRepo := new(mocks.ModelRepository)

	t.Run("test1", func(t *testing.T) {
		mockAPIRepo.On("GetByID", apiId.String()).Return(mockAPI, nil).Once()
		mockMethodRepo.On("GetListByAPIID", apiId.String()).Return([]domain.Method{}, nil).Once()
		mockModelRepo.On("GetByAPIID", apiId.String()).Return(domain.Model{}, nil).Once()
		mockAPIRepo.On("Delete", apiId.String()).Return(nil).Once()
		usecase := usecase.NewAPIUsecase(mockAPIRepo, mockMethodRepo, mockModelRepo)

		status, err := usecase.Delete(apiId.String())

		assert.NoError(t, err)
		assert.Equal(t, status, http.StatusNoContent)

		mockAPIRepo.AssertExpectations(t)
	})
}
