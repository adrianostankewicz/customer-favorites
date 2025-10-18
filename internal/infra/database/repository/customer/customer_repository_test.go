package repository

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	customer "github.com/adrianostankewicz/customer-favorites/internal/customer/entity"
	database "github.com/adrianostankewicz/customer-favorites/internal/infra/database"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

type CustomerRepositoryTestSuite struct {
	suite.Suite
	db         *sql.DB
	repository *CustomerRepositoryPostgres
}

func (suite *CustomerRepositoryTestSuite) SetupSuite() {
	err := godotenv.Load("../../../../../.env")
	if err != nil {
		fmt.Println("Erro:", err)
	}

	dbConfig := database.Config{
		Host:     getEnv("POSTGRES_HOST", "postgres"),
		Port:     5433,
		User:     getEnv("POSTGRES_USER", "postgres"),
		Password: getEnv("POSTGRES_PASSWORD", "postgres"),
		DBName:   getEnv("POSTGRES_DB", "postgres"),
		SSLMode:  getEnv("POSTGRES_SSL_MODE", "disable"),
	}

	suite.db, err = database.NewConnection(dbConfig)
	if err != nil {
		suite.T().Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	_, err = suite.db.Exec("DELETE FROM CUSTOMER")
	if err != nil {
		suite.T().Fatalf("Erro ao limpar a tabela de clientes: %v", err)
	}

	suite.repository = NewCustomerRepositoryPostgres(suite.db)
}

func (suite *CustomerRepositoryTestSuite) TearDownSuite() {
	_, err := suite.db.Exec("DELETE FROM CUSTOMER")
	if err != nil {
		suite.T().Fatalf("Erro ao limpar a tabela de clientes: %v", err)
	}

	if suite.db != nil {
		suite.db.Close()
	}
}

func (suite *CustomerRepositoryTestSuite) TestCreate() {
	customer, err := customer.NewCustomer("Fulano da Silva", "fulano@customer_favorites.com")
	suite.Nil(err)
	err = suite.repository.Create(customer)
	suite.NoError(err)

	var count int
	err = suite.db.QueryRow("SELECT COUNT(*) FROM customer WHERE id = $1", customer.ID).Scan(&count)
	suite.NoError(err)
	suite.Equal(1, count)
}

func (suite *CustomerRepositoryTestSuite) TestFindById() {
	customer, _ := customer.NewCustomer("Fulano da Silva", "fulano@customer_favoritess.com")
	suite.repository.Create(customer)

	aCustomer, err := suite.repository.FindById(customer.ID)
	suite.Nil(err)
	suite.Equal(customer.ID, aCustomer.ID)
	suite.Equal(customer.Name, aCustomer.Name)
	suite.Equal(customer.Email, aCustomer.Email)
	suite.NotNil(aCustomer.CreatedAt)
	suite.NotNil(aCustomer.UpdatedAt)
	suite.Nil(aCustomer.DeletedAt)
}

func (suite *CustomerRepositoryTestSuite) TestFindByEmail() {
	customer, _ := customer.NewCustomer("Fulano da Silva", "fulano@customer_favoritess.com")
	suite.repository.Create(customer)

	aCustomer, err := suite.repository.FindByEmail("fulano@customer_favoritess.com")
	suite.Nil(err)
	suite.Equal(customer.ID, aCustomer.ID)
	suite.Equal(customer.Name, aCustomer.Name)
	suite.Equal(customer.Email, aCustomer.Email)
	suite.NotNil(aCustomer.CreatedAt)
	suite.NotNil(aCustomer.UpdatedAt)
	suite.Nil(aCustomer.DeletedAt)
}

func (suite *CustomerRepositoryTestSuite) TestUpdate() {
	customer, _ := customer.NewCustomer("Fulano da Silva", "fulanos@customer_favorites.com")
	err := suite.repository.Create(customer)
	suite.NoError(err)

	customer.Name = "Fulano da Silva Sauro"
	err = suite.repository.Update(customer)
	suite.NoError(err)

	aCustomer, err := suite.repository.FindByEmail("fulanos@customer_favorites.com")
	suite.Nil(err)
	suite.Equal(customer.ID, aCustomer.ID)
	suite.Equal(customer.Name, aCustomer.Name)
	suite.Equal(customer.Email, aCustomer.Email)
	suite.NotNil(aCustomer.CreatedAt)
	suite.NotNil(aCustomer.UpdatedAt)
	suite.Nil(aCustomer.DeletedAt)
}

func (suite *CustomerRepositoryTestSuite) TestDelete() {
	customer, _ := customer.NewCustomer("Ciclano da Silva", "ciclano@customer_favorites.com")
	err := suite.repository.Create(customer)
	suite.NoError(err)

	err = suite.repository.Delete(customer.ID)
	suite.NoError(err)

	var count int
	err = suite.db.QueryRow("SELECT COUNT(*) FROM customer WHERE id = $1", customer.ID).Scan(&count)
	suite.NoError(err)
	suite.Equal(0, count)
}

func TestCustomerRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerRepositoryTestSuite))
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
