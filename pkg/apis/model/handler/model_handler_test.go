package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Hajime3778/api-creator-backend/pkg/apis/model/handler"
	"github.com/Hajime3778/api-creator-backend/pkg/domain"
	"github.com/Hajime3778/api-creator-backend/test/mocks"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

func newMockRouter() (*gin.Engine, *gin.RouterGroup) {
	router := gin.Default()
	modelV1 := router.Group("model/v1")

	return router, modelV1
}

func TestGetAll(t *testing.T) {
	modelId, _ := uuid.NewRandom()

	mockModels := make([]domain.Model, 0)
	mockModel := domain.Model{}
	mockModel.ID = modelId.String()
	mockModel.Name = "name"
	mockModel.Description = "description"
	mockModel.Scheme = "scheme"
	mockModel.CreatedAt = time.Now()
	mockModel.UpdatedAt = time.Now()
	mockModels = append(mockModels, mockModel)

	gin.SetMode(gin.TestMode)

	mockModelUsecase := new(mocks.ModelUsecase)
	mockModelUsecase.On("GetAll").Return(mockModels, nil).Once()

	router, rg := newMockRouter()
	handler.NewModelHandler(rg, mockModelUsecase)

	getAllRes := httptest.NewRecorder()
	getAllReq, _ := http.NewRequest("GET", "/model/v1/models", nil)
	router.ServeHTTP(getAllRes, getAllReq)

	assert.Equal(t, 200, getAllRes.Code)
}

func TestGetByID(t *testing.T) {
	modelId, _ := uuid.NewRandom()

	mockModel := domain.Model{}
	mockModel.ID = modelId.String()
	mockModel.Name = "name"
	mockModel.Description = "description"
	mockModel.Scheme = "scheme"
	mockModel.CreatedAt = time.Now()
	mockModel.UpdatedAt = time.Now()

	gin.SetMode(gin.TestMode)

	mockModelUsecase := new(mocks.ModelUsecase)
	mockModelUsecase.On("GetByID", mockModel.ID).Return(mockModel, nil).Once()

	router, rg := newMockRouter()
	handler.NewModelHandler(rg, mockModelUsecase)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/model/v1/models/"+modelId.String(), nil)
	router.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
}

func TestCreate(t *testing.T) {
	modelId, _ := uuid.NewRandom()

	mockModel := domain.Model{}
	mockModel.ID = modelId.String()
	mockModel.Name = "name"
	mockModel.Description = "description"
	mockModel.Scheme = "scheme"

	gin.SetMode(gin.TestMode)

	mockModelUsecase := new(mocks.ModelUsecase)
	mockModelUsecase.On("Create", mockModel).Return(nil).Once()

	router, rg := newMockRouter()
	handler.NewModelHandler(rg, mockModelUsecase)

	model_json, _ := json.Marshal(mockModel)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/model/v1/models", bytes.NewReader(model_json))
	router.ServeHTTP(res, req)

	assert.Equal(t, 201, res.Code)
}

func TestUpdate(t *testing.T) {
	modelId, _ := uuid.NewRandom()

	mockModel := domain.Model{}
	mockModel.ID = modelId.String()
	mockModel.Name = "name"
	mockModel.Description = "description"
	mockModel.Scheme = "scheme"

	gin.SetMode(gin.TestMode)

	mockModelUsecase := new(mocks.ModelUsecase)
	mockModelUsecase.On("Update", mockModel).Return(nil).Once()

	router, rg := newMockRouter()
	handler.NewModelHandler(rg, mockModelUsecase)

	model_json, _ := json.Marshal(mockModel)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/model/v1/models", bytes.NewReader(model_json))
	router.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
}

func TestDelete(t *testing.T) {
	modelId, _ := uuid.NewRandom()

	mockModel := domain.Model{}
	mockModel.ID = modelId.String()
	mockModel.Name = "name"
	mockModel.Description = "description"
	mockModel.Scheme = "scheme"

	gin.SetMode(gin.TestMode)

	mockModelUsecase := new(mocks.ModelUsecase)
	mockModelUsecase.On("Delete", mockModel.ID).Return(nil).Once()

	router, rg := newMockRouter()
	handler.NewModelHandler(rg, mockModelUsecase)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/model/v1/models/"+modelId.String(), nil)
	router.ServeHTTP(res, req)

	assert.Equal(t, 204, res.Code)
}
