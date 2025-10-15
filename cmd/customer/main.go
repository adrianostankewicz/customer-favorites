package main

import (
	"database/sql"
	"fmt"
	"os"
)

func main() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:5433/%s?sslmode=false",
		user, password, host, dbname,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		message := "Erro ao conectar com o banco de dados: " + err.Error()
		fmt.Println(message)
	}
}
