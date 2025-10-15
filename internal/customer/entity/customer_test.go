package customer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCustomer(t *testing.T) {
	customer, err := NewCustomer("Fulano da Silva", "fulano@customerfavorites.com.br")
	assert.Nil(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "Fulano da Silva", customer.Name)
	assert.Equal(t, "fulano@customerfavorites.com.br", customer.Email)
}
