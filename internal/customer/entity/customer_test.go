package customer

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewCustomer(t *testing.T) {
	customer, err := NewCustomer("Fulano da Silva", "fulano@customerfavorites.com.br")
	assert.Nil(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "Fulano da Silva", customer.Name)
	assert.Equal(t, "fulano@customerfavorites.com.br", customer.Email)
}

func TestCustomer_Validate(t *testing.T) {
	customer := &Customer{uuid.New().String(), "", "fulano@customerfavorites.com.br", time.Now(), time.Now(), nil}

	err := customer.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "customer name is required", err.Error())

	now := time.Now()
	customer.Name = "Fulano da Silva"
	customer.Email = ""
	customer.CreatedAt = now
	customer.UpdatedAt = now
	customer.DeletedAt = nil

	err = customer.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "customer email is required", err.Error())
}
