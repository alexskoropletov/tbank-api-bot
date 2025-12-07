package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

// InitDB инициализирует подключение к базе данных PostgreSQL
func InitDB() error {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "moneymoneymoney")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println("Successfully connected to PostgreSQL database")
	return nil
}

// CloseDB закрывает подключение к базе данных
func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

// getEnv получает переменную окружения или возвращает значение по умолчанию
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// CreateTables создает необходимые таблицы в базе данных
func CreateTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS statements (
		id SERIAL PRIMARY KEY,
		account_number VARCHAR(50) NOT NULL,
		transaction_date TIMESTAMP,
		amount DECIMAL(15, 2),
		description TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}

	fmt.Println("Tables created successfully")
	return nil
}

