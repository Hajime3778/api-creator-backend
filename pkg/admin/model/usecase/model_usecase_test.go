package usecase_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/Hajime3778/api-creator-backend/pkg/admin/model/usecase"
	"github.com/Hajime3778/api-creator-backend/pkg/domain"
	"github.com/Hajime3778/api-creator-backend/test/mocks"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	modelId, _ := uuid.NewRandom()

	mockModels := make([]domain.Model, 0)
	mockModel := domain.Model{}
	mockModel.ID = modelId.String()
	mockModel.Name = "name"
	mockModel.Description = "description"
	mockModel.Schema = "schema"
	mockModels = append(mockModels, mockModel)

	// モック
	mockModelRepo := new(mocks.ModelRepository)

	t.Run("test1", func(t *testing.T) {
		mockModelRepo.On("GetAll").Return(mockModels, nil).Once()
		usecase := usecase.NewModelUsecase(mockModelRepo)

		models, err := usecase.GetAll()

		assert.NoError(t, err)
		assert.NotNil(t, models)

		mockModelRepo.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	modelId, _ := uuid.NewRandom()

	mockModel := domain.Model{}
	mockModel.ID = modelId.String()
	mockModel.Name = "name"
	mockModel.Description = "description"
	mockModel.Schema = "schema"
	mockModel.CreatedAt = time.Now()
	mockModel.UpdatedAt = time.Now()

	mockModelRepo := new(mocks.ModelRepository)

	t.Run("test1", func(t *testing.T) {
		mockModelRepo.On("GetByID", mockModel.ID).Return(mockModel, nil).Once()
		usecase := usecase.NewModelUsecase(mockModelRepo)

		model, err := usecase.GetByID(mockModel.ID)

		assert.NoError(t, err)
		assert.NotNil(t, model)

		mockModelRepo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	modelId, _ := uuid.NewRandom()

	mockModel := domain.Model{}
	mockModel.ID = modelId.String()
	mockModel.Name = "name"
	mockModel.Description = "description"
	mockModel.Schema = "{\"type\": \"object\", \"keys\": [\"id\"], \"properties\": {\"id\": {\"type\":\"string\"}}}"

	mockModelRepo := new(mocks.ModelRepository)

	t.Run("test1", func(t *testing.T) {
		mockModelRepo.On("Create", mockModel).Return(nil).Once()
		usecase := usecase.NewModelUsecase(mockModelRepo)

		status, _, err := usecase.Create(mockModel)

		assert.NoError(t, err)
		assert.Equal(t, status, http.StatusCreated)

		mockModelRepo.AssertExpectations(t)
	})
	t.Run("jsonschema形式でない", func(t *testing.T) {
		mockModelRepo.On("Create", mockModel).Return(nil).Once()
		usecase := usecase.NewModelUsecase(mockModelRepo)

		mockModel.Schema = "test"

		status, _, err := usecase.Create(mockModel)

		assert.Error(t, err)
		assert.Equal(t, status, http.StatusBadRequest)
	})
	t.Run("keysがnil", func(t *testing.T) {
		mockModelRepo.On("Create", mockModel).Return(nil).Once()
		usecase := usecase.NewModelUsecase(mockModelRepo)

		mockModel.Schema = "{\"type\": \"object\", \"properties\": {\"id\": {\"type\":\"string\"}}}"

		status, _, err := usecase.Create(mockModel)

		assert.Error(t, err)
		assert.Equal(t, status, http.StatusBadRequest)
	})
	t.Run("keysにpropertyが指定されていない", func(t *testing.T) {
		mockModelRepo.On("Create", mockModel).Return(nil).Once()
		usecase := usecase.NewModelUsecase(mockModelRepo)

		mockModel.Schema = "{\"type\": \"object\", \"keys\": [], \"properties\": {\"id\": {\"type\":\"string\"}}}"

		status, _, err := usecase.Create(mockModel)

		assert.Error(t, err)
		assert.Equal(t, status, http.StatusBadRequest)
	})
	t.Run("存在しないプロパティをkeysで指定している", func(t *testing.T) {
		mockModel.Schema = "{\"type\": \"object\", \"keys\": [\"id\", \"foo\"], \"properties\": {\"id\": {\"type\":\"string\"}}}"
		mockModelRepo.On("Create", mockModel).Return(nil).Once()
		usecase := usecase.NewModelUsecase(mockModelRepo)

		status, _, err := usecase.Create(mockModel)

		assert.Error(t, err)
		assert.Equal(t, status, http.StatusBadRequest)
	})
}

func TestUpdate(t *testing.T) {
	modelId, _ := uuid.NewRandom()

	mockModel := domain.Model{}
	mockModel.ID = modelId.String()
	mockModel.Name = "name"
	mockModel.Description = "description"
	mockModel.Schema = "{\"type\": \"object\", \"keys\": [\"id\"], \"properties\": {\"id\": {\"type\":\"string\"}}}"

	mockModelRepo := new(mocks.ModelRepository)

	t.Run("test1", func(t *testing.T) {
		mockModelRepo.On("Update", mockModel).Return(nil).Once()
		usecase := usecase.NewModelUsecase(mockModelRepo)

		status, err := usecase.Update(mockModel)

		assert.NoError(t, err)
		assert.Equal(t, status, http.StatusOK)

		mockModelRepo.AssertExpectations(t)
	})
	t.Run("jsonschema形式でない", func(t *testing.T) {
		mockModelRepo.On("Create", mockModel).Return(nil).Once()
		usecase := usecase.NewModelUsecase(mockModelRepo)

		mockModel.Schema = "test"

		status, err := usecase.Update(mockModel)

		assert.Error(t, err)
		assert.Equal(t, status, http.StatusBadRequest)
	})
	t.Run("keysがnil", func(t *testing.T) {
		mockModelRepo.On("Create", mockModel).Return(nil).Once()
		usecase := usecase.NewModelUsecase(mockModelRepo)

		mockModel.Schema = "{\"type\": \"object\", \"properties\": {\"id\": {\"type\":\"string\"}}}"

		status, err := usecase.Update(mockModel)

		assert.Error(t, err)
		assert.Equal(t, status, http.StatusBadRequest)
	})
	t.Run("keysにpropertyが指定されていない", func(t *testing.T) {
		mockModelRepo.On("Create", mockModel).Return(nil).Once()
		usecase := usecase.NewModelUsecase(mockModelRepo)

		mockModel.Schema = "{\"type\": \"object\", \"keys\": [], \"properties\": {\"id\": {\"type\":\"string\"}}}"

		status, err := usecase.Update(mockModel)

		assert.Error(t, err)
		assert.Equal(t, status, http.StatusBadRequest)
	})
	t.Run("存在しないプロパティをkeysで指定している", func(t *testing.T) {
		mockModel.Schema = "{\"type\": \"object\", \"keys\": [\"id\", \"foo\"], \"properties\": {\"id\": {\"type\":\"string\"}}}"
		mockModelRepo.On("Create", mockModel).Return(nil).Once()
		usecase := usecase.NewModelUsecase(mockModelRepo)

		status, err := usecase.Update(mockModel)

		assert.Error(t, err)
		assert.Equal(t, status, http.StatusBadRequest)
	})
}

func TestDelete(t *testing.T) {
	modelId, _ := uuid.NewRandom()

	mockModel := domain.Model{}
	mockModel.ID = modelId.String()
	mockModel.Name = "name"
	mockModel.Description = "description"
	mockModel.Schema = "schema"

	mockModelRepo := new(mocks.ModelRepository)

	t.Run("test1", func(t *testing.T) {
		mockModelRepo.On("Delete", mockModel.ID).Return(nil).Once()
		usecase := usecase.NewModelUsecase(mockModelRepo)

		err := usecase.Delete(mockModel.ID)

		assert.NoError(t, err)

		mockModelRepo.AssertExpectations(t)
	})
}
