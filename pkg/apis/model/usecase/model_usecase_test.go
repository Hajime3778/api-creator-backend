package usecase_test

import (
	"testing"
	"time"

	"github.com/Hajime3778/api-creator-backend/pkg/apis/model/usecase"
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
	mockModel.Name = "test"
	mockModel.Description = "test"
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
	mockModel.Name = "test"
	mockModel.Description = "test"
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
	mockModel.Name = "test"
	mockModel.Description = "test"

	mockModelRepo := new(mocks.ModelRepository)

	t.Run("test1", func(t *testing.T) {
		mockModelRepo.On("Create", mockModel).Return(nil).Once()
		usecase := usecase.NewModelUsecase(mockModelRepo)

		_, err := usecase.Create(mockModel)

		assert.NoError(t, err)

		mockModelRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	modelId, _ := uuid.NewRandom()

	mockModel := domain.Model{}
	mockModel.ID = modelId.String()
	mockModel.Name = "test"
	mockModel.Description = "test"

	mockModelRepo := new(mocks.ModelRepository)

	t.Run("test1", func(t *testing.T) {
		mockModelRepo.On("Update", mockModel).Return(nil).Once()
		usecase := usecase.NewModelUsecase(mockModelRepo)

		err := usecase.Update(mockModel)

		assert.NoError(t, err)

		mockModelRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	modelId, _ := uuid.NewRandom()

	mockModel := domain.Model{}
	mockModel.ID = modelId.String()
	mockModel.Name = "test"
	mockModel.Description = "test"

	mockModelRepo := new(mocks.ModelRepository)

	t.Run("test1", func(t *testing.T) {
		mockModelRepo.On("Delete", mockModel.ID).Return(nil).Once()
		usecase := usecase.NewModelUsecase(mockModelRepo)

		err := usecase.Delete(mockModel.ID)

		assert.NoError(t, err)

		mockModelRepo.AssertExpectations(t)
	})
}
