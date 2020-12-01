package mocks

import (
	"net/http"

	"github.com/Hajime3778/api-creator-backend/pkg/domain"

	"github.com/stretchr/testify/mock"
)

// APIUsecase is mock
type APIUsecase struct {
	mock.Mock
}

// GetAll is mock function
func (_m *APIUsecase) GetAll() ([]domain.API, error) {
	ret := _m.Called()
	return ret.Get(0).([]domain.API), ret.Error(1)
}

// GetByID is mock function
func (_m *APIUsecase) GetByID(id string) (domain.API, error) {
	ret := _m.Called(id)
	return ret.Get(0).(domain.API), ret.Error(1)
}

// Create is mock function
func (_m *APIUsecase) Create(api domain.API) (int, string, error) {
	ret := _m.Called(api)
	return http.StatusCreated, api.ID, ret.Error(0)
}

// Update is mock function
func (_m *APIUsecase) Update(api domain.API) (int, error) {
	ret := _m.Called(api)
	return http.StatusOK, ret.Error(0)
}

// Delete is mock function
func (_m *APIUsecase) Delete(id string) error {
	ret := _m.Called(id)
	return ret.Error(0)
}

// APIRepository is mock
type APIRepository struct {
	mock.Mock
}

// GetAll is mock function
func (_m *APIRepository) GetAll() ([]domain.API, error) {
	ret := _m.Called()
	return ret.Get(0).([]domain.API), ret.Error(1)
}

// GetByID is mock function
func (_m *APIRepository) GetByID(id string) (domain.API, error) {
	ret := _m.Called(id)
	return ret.Get(0).(domain.API), ret.Error(1)
}

// GetByURL is mock function
func (_m *APIRepository) GetByURL(url string) (domain.API, error) {
	ret := _m.Called(url)
	return ret.Get(0).(domain.API), ret.Error(1)
}

// Create is mock function
func (_m *APIRepository) Create(api domain.API) (string, error) {
	ret := _m.Called(api)
	return api.ID, ret.Error(0)
}

// Update is mock function
func (_m *APIRepository) Update(api domain.API) error {
	ret := _m.Called(api)
	return ret.Error(0)
}

// Delete is mock function
func (_m *APIRepository) Delete(id string) error {
	ret := _m.Called(id)
	return ret.Error(0)
}
