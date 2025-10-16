package customer

type CustomerRepository interface {
	Create(customer *Customer) error
	Update(customer *Customer) error
	FindByEmail(email string) (*Customer, error)
	Delete(id string) error
}
