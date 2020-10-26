package mocks

import (
	"github.com/Hajime3778/api-creator-backend/pkg/domain"

	"github.com/stretchr/testify/mock"
)

// ModelUsecase is mock
type ModelUsecase struct {
	mock.Mock
}

// GetAll is mock function
func (_m *ModelUsecase) GetAll() ([]domain.Model, error) {
	ret := _m.Called()
	return ret.Get(0).([]domain.Model), ret.Error(1)
}

// GetByID is mock function
func (_m *ModelUsecase) GetByID(id string) (domain.Model, error) {
	ret := _m.Called(id)
	return ret.Get(0).(domain.Model), ret.Error(1)
}

// Create is mock function
func (_m *ModelUsecase) Create(model domain.Model) (string, error) {
	ret := _m.Called(model)
	return model.ID, ret.Error(0)
}

// Update is mock function
func (_m *ModelUsecase) Update(model domain.Model) error {
	ret := _m.Called(model)
	return ret.Error(0)
}

// Delete is mock function
func (_m *ModelUsecase) Delete(id string) error {
	ret := _m.Called(id)
	return ret.Error(0)
}

// ModelRepository is mock
type ModelRepository struct {
	mock.Mock
}

// GetAll is mock function
func (_m *ModelRepository) GetAll() ([]domain.Model, error) {
	ret := _m.Called()
	return ret.Get(0).([]domain.Model), ret.Error(1)
}

// GetByID is mock function
func (_m *ModelRepository) GetByID(id string) (domain.Model, error) {
	ret := _m.Called(id)
	return ret.Get(0).(domain.Model), ret.Error(1)
}

// Create is mock function
func (_m *ModelRepository) Create(model domain.Model) (string, error) {
	ret := _m.Called(model)
	return model.ID, ret.Error(0)
}

// Update is mock function
func (_m *ModelRepository) Update(model domain.Model) error {
	ret := _m.Called(model)
	return ret.Error(0)
}

// Delete is mock function
func (_m *ModelRepository) Delete(id string) error {
	ret := _m.Called(id)
	return ret.Error(0)
}
