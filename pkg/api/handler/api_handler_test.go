package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Hajime3778/api-creator-backend/pkg/api/handler"
	"github.com/Hajime3778/api-creator-backend/pkg/domain"
	"github.com/Hajime3778/api-creator-backend/test/mocks"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

func newMockRouter() (*gin.Engine, *gin.RouterGroup) {
	router := gin.Default()
	apiV1 := router.Group("api/v1")

	return router, apiV1
}

func TestGetAll(t *testing.T) {
	apiId, _ := uuid.NewRandom()

	mockAPIs := make([]domain.API, 0)
	mockAPI := domain.API{}
	mockAPI.ID = apiId.String()
	mockAPI.Name = "mockapi"
	mockAPI.URL = "url"
	mockAPI.Description = "mock@mock.com"
	mockAPI.CreatedAt = time.Now()
	mockAPI.UpdatedAt = time.Now()
	mockAPIs = append(mockAPIs, mockAPI)

	gin.SetMode(gin.TestMode)

	mockAPIUsecase := new(mocks.APIUsecase)
	mockAPIUsecase.On("GetAll").Return(mockAPIs, nil).Once()

	router, rg := newMockRouter()
	handler.NewAPIHandler(rg, mockAPIUsecase)

	getAllRes := httptest.NewRecorder()
	getAllReq, _ := http.NewRequest("GET", "/api/v1/apis", nil)
	router.ServeHTTP(getAllRes, getAllReq)

	assert.Equal(t, 200, getAllRes.Code)
}

func TestGetByID(t *testing.T) {
	apiId, _ := uuid.NewRandom()

	mockAPI := domain.API{}
	mockAPI.ID = apiId.String()
	mockAPI.Name = "mockapi"
	mockAPI.URL = "url"
	mockAPI.Description = "mock@mock.com"
	mockAPI.CreatedAt = time.Now()
	mockAPI.UpdatedAt = time.Now()

	gin.SetMode(gin.TestMode)

	mockAPIUsecase := new(mocks.APIUsecase)
	mockAPIUsecase.On("GetByID", mockAPI.ID).Return(mockAPI, nil).Once()

	router, rg := newMockRouter()
	handler.NewAPIHandler(rg, mockAPIUsecase)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/apis/"+apiId.String(), nil)
	router.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
}

func TestCreate(t *testing.T) {
	apiId, _ := uuid.NewRandom()

	mockAPI := domain.API{}
	mockAPI.ID = apiId.String()
	mockAPI.Name = "mockapi"
	mockAPI.URL = "url"
	mockAPI.Description = "mock@mock.com"

	gin.SetMode(gin.TestMode)

	mockAPIUsecase := new(mocks.APIUsecase)
	mockAPIUsecase.On("Create", mockAPI).Return(nil).Once()

	router, rg := newMockRouter()
	handler.NewAPIHandler(rg, mockAPIUsecase)

	api_json, _ := json.Marshal(mockAPI)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/apis", bytes.NewReader(api_json))
	router.ServeHTTP(res, req)

	assert.Equal(t, 201, res.Code)
}

func TestUpdate(t *testing.T) {
	apiId, _ := uuid.NewRandom()

	mockAPI := domain.API{}
	mockAPI.ID = apiId.String()
	mockAPI.Name = "mockapi"
	mockAPI.URL = "url"
	mockAPI.Description = "mock@mock.com"

	gin.SetMode(gin.TestMode)

	mockAPIUsecase := new(mocks.APIUsecase)
	mockAPIUsecase.On("Update", mockAPI).Return(nil).Once()

	router, rg := newMockRouter()
	handler.NewAPIHandler(rg, mockAPIUsecase)

	api_json, _ := json.Marshal(mockAPI)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/v1/apis", bytes.NewReader(api_json))
	router.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
}

func TestDelete(t *testing.T) {
	apiId, _ := uuid.NewRandom()

	mockAPI := domain.API{}
	mockAPI.ID = apiId.String()
	mockAPI.Name = "mockapi"
	mockAPI.URL = "url"
	mockAPI.Description = "mock@mock.com"
	mockAPI.CreatedAt = time.Now()
	mockAPI.UpdatedAt = time.Now()

	gin.SetMode(gin.TestMode)

	mockAPIUsecase := new(mocks.APIUsecase)
	mockAPIUsecase.On("Delete", mockAPI.ID).Return(nil).Once()

	router, rg := newMockRouter()
	handler.NewAPIHandler(rg, mockAPIUsecase)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/apis/"+apiId.String(), nil)
	router.ServeHTTP(res, req)

	assert.Equal(t, 204, res.Code)
}
