package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Hajime3778/api-creator-backend/pkg/apis/method/handler"
	"github.com/Hajime3778/api-creator-backend/pkg/domain"
	"github.com/Hajime3778/api-creator-backend/test/mocks"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

func newMockRouter() (*gin.Engine, *gin.RouterGroup) {
	router := gin.Default()
	methodV1 := router.Group("method/v1")

	return router, methodV1
}

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

	gin.SetMode(gin.TestMode)

	mockMethodUsecase := new(mocks.MethodUsecase)
	mockMethodUsecase.On("GetAll").Return(mockMethods, nil).Once()

	router, rg := newMockRouter()
	handler.NewMethodHandler(rg, mockMethodUsecase)

	getAllRes := httptest.NewRecorder()
	getAllReq, _ := http.NewRequest("GET", "/method/v1/methods", nil)
	router.ServeHTTP(getAllRes, getAllReq)

	assert.Equal(t, 200, getAllRes.Code)
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

	gin.SetMode(gin.TestMode)

	mockMethodUsecase := new(mocks.MethodUsecase)
	mockMethodUsecase.On("GetByID", mockMethod.ID).Return(mockMethod, nil).Once()

	router, rg := newMockRouter()
	handler.NewMethodHandler(rg, mockMethodUsecase)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/method/v1/methods/"+methodId.String(), nil)
	router.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
}

func TestCreate(t *testing.T) {
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

	gin.SetMode(gin.TestMode)

	mockMethodUsecase := new(mocks.MethodUsecase)
	mockMethodUsecase.On("Create", mockMethod).Return(nil).Once()

	router, rg := newMockRouter()
	handler.NewMethodHandler(rg, mockMethodUsecase)

	method_json, _ := json.Marshal(mockMethod)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/method/v1/methods", bytes.NewReader(method_json))
	router.ServeHTTP(res, req)

	assert.Equal(t, 201, res.Code)
}

func TestUpdate(t *testing.T) {
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

	gin.SetMode(gin.TestMode)

	mockMethodUsecase := new(mocks.MethodUsecase)
	mockMethodUsecase.On("Update", mockMethod).Return(nil).Once()

	router, rg := newMockRouter()
	handler.NewMethodHandler(rg, mockMethodUsecase)

	method_json, _ := json.Marshal(mockMethod)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/method/v1/methods", bytes.NewReader(method_json))
	router.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
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
	mockMethod.CreatedAt = time.Now()
	mockMethod.UpdatedAt = time.Now()

	gin.SetMode(gin.TestMode)

	mockMethodUsecase := new(mocks.MethodUsecase)
	mockMethodUsecase.On("Delete", mockMethod.ID).Return(nil).Once()

	router, rg := newMockRouter()
	handler.NewMethodHandler(rg, mockMethodUsecase)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/method/v1/methods/"+methodId.String(), nil)
	router.ServeHTTP(res, req)

	assert.Equal(t, 204, res.Code)
}
