package mocks

import (
	"github.com/Hajime3778/api-creator-backend/pkg/domain"

	"github.com/stretchr/testify/mock"
)

// MethodUsecase is mock
type MethodUsecase struct {
	mock.Mock
}

// GetAll is mock function
func (_m *MethodUsecase) GetAll() ([]domain.Method, error) {
	ret := _m.Called()
	return ret.Get(0).([]domain.Method), ret.Error(1)
}

// GetByID is mock function
func (_m *MethodUsecase) GetByID(id string) (domain.Method, error) {
	ret := _m.Called(id)
	return ret.Get(0).(domain.Method), ret.Error(1)
}

// Create is mock function
func (_m *MethodUsecase) Create(method domain.Method) (string, error) {
	ret := _m.Called(method)
	return method.ID, ret.Error(0)
}

// Update is mock function
func (_m *MethodUsecase) Update(method domain.Method) error {
	ret := _m.Called(method)
	return ret.Error(0)
}

// Delete is mock function
func (_m *MethodUsecase) Delete(id string) error {
	ret := _m.Called(id)
	return ret.Error(0)
}

// MethodRepository is mock
type MethodRepository struct {
	mock.Mock
}

// GetAll is mock function
func (_m *MethodRepository) GetAll() ([]domain.Method, error) {
	ret := _m.Called()
	return ret.Get(0).([]domain.Method), ret.Error(1)
}

// GetByID is mock function
func (_m *MethodRepository) GetByID(id string) (domain.Method, error) {
	ret := _m.Called(id)
	return ret.Get(0).(domain.Method), ret.Error(1)
}

// Create is mock function
func (_m *MethodRepository) Create(method domain.Method) (string, error) {
	ret := _m.Called(method)
	return method.ID, ret.Error(0)
}

// Update is mock function
func (_m *MethodRepository) Update(method domain.Method) error {
	ret := _m.Called(method)
	return ret.Error(0)
}

// Delete is mock function
func (_m *MethodRepository) Delete(id string) error {
	ret := _m.Called(id)
	return ret.Error(0)
}
