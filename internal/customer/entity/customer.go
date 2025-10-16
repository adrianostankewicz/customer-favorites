package customer

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt *time.Time
}

func NewCustomer(name string, email string) (*Customer, error) {
	customer := &Customer{
		ID:    uuid.New().String(),
		Name:  name,
		Email: email,
	}
	return customer, nil
}

func (c *Customer) Validate() error {
	if c.Name == "" {
		return errors.New("customer name is required")
	}

	if c.Email == "" {
		return errors.New("customer email is required")
	}

	return nil
}
