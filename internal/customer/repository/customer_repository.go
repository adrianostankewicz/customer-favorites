package customer

import customer "github.com/adrianostankewicz/customer-favorites/internal/customer/entity"

type CustomerRepository interface {
	Create(customer *customer.Customer) error
	Update(customer *customer.Customer) error
	FindByEmail(email string) (*customer.Customer, error)
	Delete(id string) error
}
