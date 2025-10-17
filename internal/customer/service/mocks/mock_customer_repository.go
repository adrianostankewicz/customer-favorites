package mocks

import (
	customer "github.com/adrianostankewicz/customer-favorites/internal/customer/entity"
	"github.com/stretchr/testify/mock"
)

type CustomerRepositoryMock struct {
	mock.Mock
}

func (m *CustomerRepositoryMock) Create(customer *customer.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *CustomerRepositoryMock) Update(customer *customer.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *CustomerRepositoryMock) FindById(id string) (*customer.Customer, error) {
	args := m.Called(id)
	return args.Get(0).(*customer.Customer), args.Error(1)
}

func (m *CustomerRepositoryMock) FindByEmail(email string) (*customer.Customer, error) {
	args := m.Called(email)
	return args.Get(0).(*customer.Customer), args.Error(1)
}

func (m *CustomerRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
