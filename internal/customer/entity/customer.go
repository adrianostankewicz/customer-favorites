package customer

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt time.Time
}

func NewCustomer(name string, email string) (*Customer, error) {
	customer := &Customer{
		ID:    uuid.New().String(),
		Name:  name,
		Email: email,
	}
	return customer, nil
}
