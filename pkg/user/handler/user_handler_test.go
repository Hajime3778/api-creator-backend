package handler_test

import (
	"bytes"
	"encoding/json"
	"go-mysql-api/pkg/domain"
	"go-mysql-api/pkg/user/handler"
	"go-mysql-api/test/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

func newMockRouter() (*gin.Engine, *gin.RouterGroup) {
	router := gin.Default()
	apiV1 := router.Group("api/v1")

	return router, apiV1
}

func TestGetAll(t *testing.T) {

	mockUsers := make([]domain.User, 0)
	mockUser := domain.User{}
	mockUser.ID = 1
	mockUser.Name = "mockuser"
	mockUser.Email = "mock@mock.com"
	mockUser.CreatedAt = time.Now()
	mockUser.UpdatedAt = time.Now()
	mockUsers = append(mockUsers, mockUser)

	gin.SetMode(gin.TestMode)

	mockUserUsecase := new(mocks.UserUsecase)
	mockUserUsecase.On("GetAll").Return(mockUsers, nil).Once()

	router, rg := newMockRouter()
	handler.NewUserHandler(rg, mockUserUsecase)

	getAllRes := httptest.NewRecorder()
	getAllReq, _ := http.NewRequest("GET", "/api/v1/users", nil)
	router.ServeHTTP(getAllRes, getAllReq)

	assert.Equal(t, 200, getAllRes.Code)
}

func TestGetByID(t *testing.T) {
	mockUser := domain.User{}
	mockUser.ID = 1
	mockUser.Name = "mockuser"
	mockUser.Email = "mock@mock.com"
	mockUser.CreatedAt = time.Now()
	mockUser.UpdatedAt = time.Now()

	gin.SetMode(gin.TestMode)

	mockUserUsecase := new(mocks.UserUsecase)
	mockUserUsecase.On("GetByID", mockUser.ID).Return(mockUser, nil).Once()

	router, rg := newMockRouter()
	handler.NewUserHandler(rg, mockUserUsecase)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users/1", nil)
	router.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
}

func TestCreate(t *testing.T) {
	mockUser := domain.User{}
	mockUser.ID = 1
	mockUser.Name = "mockuser"
	mockUser.Email = "mock@mock.com"

	gin.SetMode(gin.TestMode)

	mockUserUsecase := new(mocks.UserUsecase)
	mockUserUsecase.On("Create", mockUser).Return(nil).Once()

	router, rg := newMockRouter()
	handler.NewUserHandler(rg, mockUserUsecase)

	user_json, _ := json.Marshal(mockUser)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewReader(user_json))
	router.ServeHTTP(res, req)

	assert.Equal(t, 201, res.Code)
}

func TestUpdate(t *testing.T) {
	mockUser := domain.User{}
	mockUser.ID = 1
	mockUser.Name = "mockuser"
	mockUser.Email = "mock@mock.com"

	gin.SetMode(gin.TestMode)

	mockUserUsecase := new(mocks.UserUsecase)
	mockUserUsecase.On("Update", mockUser).Return(nil).Once()

	router, rg := newMockRouter()
	handler.NewUserHandler(rg, mockUserUsecase)

	user_json, _ := json.Marshal(mockUser)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/v1/users", bytes.NewReader(user_json))
	router.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
}

func TestDelete(t *testing.T) {
	mockUser := domain.User{}
	mockUser.ID = 1
	mockUser.Name = "mockuser"
	mockUser.Email = "mock@mock.com"
	mockUser.CreatedAt = time.Now()
	mockUser.UpdatedAt = time.Now()

	gin.SetMode(gin.TestMode)

	mockUserUsecase := new(mocks.UserUsecase)
	mockUserUsecase.On("Delete", mockUser.ID).Return(nil).Once()

	router, rg := newMockRouter()
	handler.NewUserHandler(rg, mockUserUsecase)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/users/1", nil)
	router.ServeHTTP(res, req)

	assert.Equal(t, 204, res.Code)
}
