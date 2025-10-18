package main

import (
	"fmt"
	"os"

	service "github.com/adrianostankewicz/customer-favorites/internal/customer/service"
	"github.com/adrianostankewicz/customer-favorites/internal/infra/database"
	repository "github.com/adrianostankewicz/customer-favorites/internal/infra/database/repository/customer"
	"github.com/adrianostankewicz/customer-favorites/internal/infra/web"
)

func main() {
	dbConfig := database.Config{
		Host:     getEnv("DB_HOST", "postgres"),
		Port:     5432,
		User:     getEnv("POSTGRES_USER", "postgres"),
		Password: getEnv("POSTGRES_PASSWORD", "postgres"),
		DBName:   getEnv("POSTGRES_DB", "postgres"),
		SSLMode:  getEnv("POSTGRES_SSL_MODE", "disable"),
	}

	db, err := database.NewConnection(dbConfig)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		message := "Erro ao conectar com o banco de dados: " + err.Error()
		fmt.Println(message)
	}

	customerRepository := repository.NewCustomerRepositoryPostgres(db)
	customerService := service.NewCustomerService(customerRepository)

	webserver := web.NewWebServer(":3000")
	webserver.AddHandler(customerService)
	webserver.Start()
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
