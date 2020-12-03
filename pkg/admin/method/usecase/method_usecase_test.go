package usecase_test

import (
	"testing"
	"time"

	"github.com/Hajime3778/api-creator-backend/pkg/admin/method/usecase"
	"github.com/Hajime3778/api-creator-backend/pkg/domain"
	"github.com/Hajime3778/api-creator-backend/test/mocks"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	methodId, _ := uuid.NewRandom()
	apiId, _ := uuid.NewRandom()

	mockMethods := make([]domain.Method, 0)
	mockMethod := domain.Method{}
	mockMethod.ID = methodId.String()
	mockMethod.APIID = apiId.String()
	mockMethod.Type = "GET"
	mockMethod.URL = "url"
	mockMethod.Description = "test"
	mockMethod.RequestParameter = ""
	mockMethod.RequestModelID = ""
	mockMethod.ResponseModelID = ""
	mockMethod.IsArray = false
	mockMethod.CreatedAt = time.Now()
	mockMethod.UpdatedAt = time.Now()
	mockMethods = append(mockMethods, mockMethod)

	// モック
	mockAPIRepo := new(mocks.APIRepository)
	mockModelRepo := new(mocks.ModelRepository)
	mockMethodRepo := new(mocks.MethodRepository)

	t.Run("test1", func(t *testing.T) {
		mockMethodRepo.On("GetAll").Return(mockMethods, nil).Once()
		usecase := usecase.NewMethodUsecase(mockAPIRepo, mockModelRepo, mockMethodRepo)

		methods, err := usecase.GetAll()

		assert.NoError(t, err)
		assert.NotNil(t, methods)

		mockMethodRepo.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	methodId, _ := uuid.NewRandom()
	apiId, _ := uuid.NewRandom()

	mockMethod := domain.Method{}
	mockMethod.ID = methodId.String()
	mockMethod.APIID = apiId.String()
	mockMethod.Type = "GET"
	mockMethod.URL = "url"
	mockMethod.Description = "test"
	mockMethod.RequestParameter = ""
	mockMethod.RequestModelID = ""
	mockMethod.ResponseModelID = ""
	mockMethod.IsArray = false
	mockMethod.CreatedAt = time.Now()
	mockMethod.UpdatedAt = time.Now()

	mockAPIRepo := new(mocks.APIRepository)
	mockModelRepo := new(mocks.ModelRepository)
	mockMethodRepo := new(mocks.MethodRepository)

	t.Run("test1", func(t *testing.T) {
		mockMethodRepo.On("GetByID", mockMethod.ID).Return(mockMethod, nil).Once()
		usecase := usecase.NewMethodUsecase(mockAPIRepo, mockModelRepo, mockMethodRepo)

		method, err := usecase.GetByID(mockMethod.ID)

		assert.NoError(t, err)
		assert.NotNil(t, method)

		mockMethodRepo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	methodId, _ := uuid.NewRandom()
	apiId, _ := uuid.NewRandom()

	mockMethod := domain.Method{}
	mockMethod.ID = methodId.String()
	mockMethod.APIID = apiId.String()
	mockMethod.Type = "GET"
	mockMethod.URL = "/url"
	mockMethod.Description = "test"
	mockMethod.RequestParameter = ""
	mockMethod.RequestModelID = ""
	mockMethod.ResponseModelID = ""
	mockMethod.IsArray = false

	mockAPIRepo := new(mocks.APIRepository)
	mockModelRepo := new(mocks.ModelRepository)
	mockMethodRepo := new(mocks.MethodRepository)

	t.Run("test1", func(t *testing.T) {
		mockMethodRepo.On("Create", mockMethod).Return(nil).Once()
		mockMethodRepo.On("GetListByAPIIDAndType", mockMethod.APIID, mockMethod.Type).Return([]domain.Method{}, nil).Once()
		usecase := usecase.NewMethodUsecase(mockAPIRepo, mockModelRepo, mockMethodRepo)

		_, _, err := usecase.Create(mockMethod)

		assert.NoError(t, err)

		mockMethodRepo.AssertExpectations(t)
	})

	t.Run("test2", func(t *testing.T) {
		mockMethod.URL = "url"
		mockMethodRepo.On("Create", mockMethod).Return(nil).Once()
		mockMethodRepo.On("GetListByAPIIDAndType", mockMethod.APIID, mockMethod.Type).Return([]domain.Method{}, nil).Once()
		usecase := usecase.NewMethodUsecase(mockAPIRepo, mockModelRepo, mockMethodRepo)

		_, err := usecase.Update(mockMethod)

		assert.Error(t, err)
	})
}

func TestUpdate(t *testing.T) {
	methodId, _ := uuid.NewRandom()
	apiId, _ := uuid.NewRandom()

	mockMethod := domain.Method{}
	mockMethod.ID = methodId.String()
	mockMethod.APIID = apiId.String()
	mockMethod.Type = "GET"
	mockMethod.URL = "/url"
	mockMethod.Description = "test"
	mockMethod.RequestParameter = ""
	mockMethod.RequestModelID = ""
	mockMethod.ResponseModelID = ""
	mockMethod.IsArray = false

	mockAPIRepo := new(mocks.APIRepository)
	mockModelRepo := new(mocks.ModelRepository)
	mockMethodRepo := new(mocks.MethodRepository)

	t.Run("test1", func(t *testing.T) {
		mockMethodRepo.On("Update", mockMethod).Return(nil).Once()
		mockMethodRepo.On("GetListByAPIIDAndType", mockMethod.APIID, mockMethod.Type).Return([]domain.Method{}, nil).Once()
		usecase := usecase.NewMethodUsecase(mockAPIRepo, mockModelRepo, mockMethodRepo)

		_, err := usecase.Update(mockMethod)

		assert.NoError(t, err)

		mockMethodRepo.AssertExpectations(t)
	})

	t.Run("test2", func(t *testing.T) {
		mockMethod.URL = "url"
		mockMethodRepo.On("Update", mockMethod).Return(nil).Once()
		mockMethodRepo.On("GetListByAPIIDAndType", mockMethod.APIID, mockMethod.Type).Return([]domain.Method{}, nil).Once()
		usecase := usecase.NewMethodUsecase(mockAPIRepo, mockModelRepo, mockMethodRepo)

		_, err := usecase.Update(mockMethod)

		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	methodId, _ := uuid.NewRandom()
	apiId, _ := uuid.NewRandom()

	mockMethod := domain.Method{}
	mockMethod.ID = methodId.String()
	mockMethod.APIID = apiId.String()
	mockMethod.Type = "GET"
	mockMethod.URL = "url"
	mockMethod.Description = "test"
	mockMethod.RequestParameter = ""
	mockMethod.RequestModelID = ""
	mockMethod.ResponseModelID = ""
	mockMethod.IsArray = false

	mockAPIRepo := new(mocks.APIRepository)
	mockModelRepo := new(mocks.ModelRepository)
	mockMethodRepo := new(mocks.MethodRepository)

	t.Run("test1", func(t *testing.T) {
		mockMethodRepo.On("Delete", mockMethod.ID).Return(nil).Once()
		usecase := usecase.NewMethodUsecase(mockAPIRepo, mockModelRepo, mockMethodRepo)

		err := usecase.Delete(mockMethod.ID)

		assert.NoError(t, err)

		mockMethodRepo.AssertExpectations(t)
	})
}
