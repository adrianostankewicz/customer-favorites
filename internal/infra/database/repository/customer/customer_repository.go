package repository

import (
	"database/sql"

	customer "github.com/adrianostankewicz/customer-favorites/internal/customer/entity"
	customerRepository "github.com/adrianostankewicz/customer-favorites/internal/customer/repository"
)

type CustomerRepositoryPostgres struct {
	DB *sql.DB
}

func NewCustomerRepositoryPostgres(db *sql.DB) *CustomerRepositoryPostgres {
	return &CustomerRepositoryPostgres{
		DB: db,
	}
}

func (r *CustomerRepositoryPostgres) Create(customer *customer.Customer) error {
	stmt, err := r.DB.Prepare("INSERT INTO customer (id, name, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(customer.ID, customer.Name, customer.Email, customer.CreatedAt, customer.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *CustomerRepositoryPostgres) FindById(id string) (*customer.Customer, error) {
	customer := &customer.Customer{}
	stmt, err := r.DB.Prepare("SELECT id, name, email, created_at, updated_at, deleted_at FROM customer WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	if err := row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.CreatedAt, &customer.UpdatedAt, &customer.DeletedAt); err != nil {
		return nil, err
	}
	return customer, nil
}

func (r *CustomerRepositoryPostgres) FindByEmail(email string) (*customer.Customer, error) {
	customer := &customer.Customer{}
	stmt, err := r.DB.Prepare("SELECT id, name, email, created_at, updated_at, deleted_at FROM customer WHERE email = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(email)
	if err := row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.CreatedAt, &customer.UpdatedAt, &customer.DeletedAt); err != nil {
		return nil, err
	}
	return customer, nil
}

func (r *CustomerRepositoryPostgres) Update(customer *customer.Customer) error {
	stmt, err := r.DB.Prepare("UPDATE customer SET name = $2, email = $3 WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(customer.ID, customer.Name, customer.Email)
	if err != nil {
		return err
	}
	return nil
}

func (r *CustomerRepositoryPostgres) Delete(id string) error {
	stmt, err := r.DB.Prepare("DELETE FROM customer WHERE id = $1;")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

var _ customerRepository.CustomerRepository = (*CustomerRepositoryPostgres)(nil)
