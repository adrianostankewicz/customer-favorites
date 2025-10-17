package customer

import (
	customer "github.com/adrianostankewicz/customer-favorites/internal/customer/entity"
	repository "github.com/adrianostankewicz/customer-favorites/internal/customer/repository"
)

type CreateCustomerInputDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type FindCustomerByIdInputDTO struct {
	ID string `json:"id"`
}

type FindCustomerByIdOutputDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type FindCustomerByEmailInputDTO struct {
	Email string `json:"email"`
}

type FindCustomerByEmailOutputDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateCustomerInputDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type DeleteCustomerInputDTO struct {
	ID string `json:"id"`
}

type CustomerService struct {
	CustomerRepository repository.CustomerRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository) *CustomerService {
	return &CustomerService{
		CustomerRepository: customerRepository,
	}
}

func (cs *CustomerService) Create(input CreateCustomerInputDTO) error {
	customer, err := customer.NewCustomer(input.Name, input.Email)
	if err != nil {
		return err
	}

	err = cs.CustomerRepository.Create(customer)
	if err != nil {
		return err
	}
	return nil
}

func (cs *CustomerService) FindById(input FindCustomerByIdInputDTO) (*FindCustomerByIdOutputDTO, error) {
	customer, err := cs.CustomerRepository.FindById(input.ID)
	if err != nil {
		return nil, err
	}

	return &FindCustomerByIdOutputDTO{
		ID:    customer.ID,
		Name:  customer.Name,
		Email: customer.Email,
	}, nil
}

func (cs *CustomerService) FindByEmail(input FindCustomerByEmailInputDTO) (*FindCustomerByEmailOutputDTO, error) {
	customer, err := cs.CustomerRepository.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	return &FindCustomerByEmailOutputDTO{
		ID:    customer.ID,
		Name:  customer.Name,
		Email: customer.Email,
	}, nil
}

func (cs *CustomerService) Update(input UpdateCustomerInputDTO) error {
	customer, err := cs.CustomerRepository.FindById(input.ID)
	if err != nil {
		return err
	}

	customer.Name = input.Name
	customer.Email = input.Email
	err = cs.CustomerRepository.Update(customer)
	if err != nil {
		return err
	}
	return nil
}

func (cs *CustomerService) Delete(input DeleteCustomerInputDTO) error {
	customer, err := cs.CustomerRepository.FindById(input.ID)
	if err != nil {
		return err
	}

	err = cs.CustomerRepository.Delete(customer.ID)
	if err != nil {
		return err
	}
	return nil
}
