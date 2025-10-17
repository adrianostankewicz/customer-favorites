package customer

import (
	"testing"

	customer "github.com/adrianostankewicz/customer-favorites/internal/customer/entity"
	"github.com/adrianostankewicz/customer-favorites/internal/customer/service/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCustomer(t *testing.T) {
	m := new(mocks.CustomerRepositoryMock)
	m.On("Create", mock.Anything).Return(nil)

	s := NewCustomerService(m)
	err := s.Create(CreateCustomerInputDTO{
		Name:  "Fulano da Silva",
		Email: "fulano@customer_repositories.com",
	})

	assert.Nil(t, err)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Create", 1)
}

func TestFindCustomerById(t *testing.T) {
	m := new(mocks.CustomerRepositoryMock)
	id := uuid.New().String()
	email := "fulano@customer_favorites.com"
	customer := &customer.Customer{
		ID:    id,
		Name:  "Fulano da Silva",
		Email: email,
	}
	m.On("FindById", id).Return(customer, nil)

	s := NewCustomerService(m)
	output, err := s.FindById(FindCustomerByIdInputDTO{
		ID: id,
	})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, customer.ID, output.ID)
	assert.Equal(t, customer.Name, output.Name)
	assert.Equal(t, customer.Email, output.Email)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "FindById", 1)
}

func TestFindCustomerByEmail(t *testing.T) {
	m := new(mocks.CustomerRepositoryMock)
	id := uuid.New().String()
	email := "fulano@customer_favorites.com"
	customer := &customer.Customer{
		ID:    id,
		Name:  "Fulano da Silva",
		Email: email,
	}
	m.On("FindByEmail", email).Return(customer, nil)

	s := NewCustomerService(m)
	output, err := s.FindByEmail(FindCustomerByEmailInputDTO{
		Email: email,
	})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, customer.ID, output.ID)
	assert.Equal(t, customer.Name, output.Name)
	assert.Equal(t, customer.Email, output.Email)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "FindByEmail", 1)
}

func TestUpdateCustomer(t *testing.T) {
	m := new(mocks.CustomerRepositoryMock)

	id := uuid.New().String()
	customer := &customer.Customer{
		ID:    id,
		Name:  "Fulano da Silva",
		Email: "fulano@customer_favorites.com",
	}
	m.On("FindById", id).Return(customer, nil)

	m.On("Update", mock.Anything).Return(nil)

	s := NewCustomerService(m)
	err := s.Update(UpdateCustomerInputDTO{
		ID:    id,
		Name:  "Ciclano da Silva",
		Email: "ciclano@customer_favorites.com",
	})

	assert.Nil(t, err)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "FindById", 1)
	m.AssertNumberOfCalls(t, "Update", 1)
}

func TestDeleteCustomer(t *testing.T) {
	m := new(mocks.CustomerRepositoryMock)

	id := uuid.New().String()
	customer := &customer.Customer{
		ID:    id,
		Name:  "Fulano da Silva",
		Email: "fulano@customer_favorites.com",
	}
	m.On("FindById", id).Return(customer, nil)

	m.On("Delete", id).Return(nil)

	s := NewCustomerService(m)
	err := s.Delete(DeleteCustomerInputDTO{
		ID: id,
	})

	assert.Nil(t, err)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "FindById", 1)
	m.AssertNumberOfCalls(t, "Delete", 1)
}
