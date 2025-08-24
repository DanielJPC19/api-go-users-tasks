package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")

	// Cargar archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("No se pudo cargar el archivo .env: %v", err)
	}

	// Conectar a la base de datos
	db, err := connectDB()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

}

func connectDB() (*sql.DB, error) {

	connStr := os.Getenv("MYSQL_URI")

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")
	return db, nil
}
