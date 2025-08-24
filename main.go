package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type User struct {
	ID    int
	Name  string
	Email string
}

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

	createTable(db)
	insertUser(db, User{Name: "John Doe", Email: "john@example.com"})
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

// Example of how to create a table
func createTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	_, err := db.Exec(query)
	return err
}

func insertUser(db *sql.DB, user User) int {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"

	result, err := db.Exec(query, user.Name, user.Email)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return 0
	}

	userID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID: %v", err)
		return 0
	}

	return int(userID)
}
