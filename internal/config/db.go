package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const (
	DriverName string = "postgres"
	Db_host    string = "DATABASE_HOST"
	Db_port    string = "DATABASE_PORT"
	Db_user    string = "DATABASE_USER"
	Db_pass    string = "DATABASE_PASSWORD"
	Db_name    string = "DATABASE_NAME"
)

func OpenDB() (*sql.DB, error) {
	dsn, err := buildDSN()
	if err != nil {
		return nil, fmt.Errorf("error to build DSN: %w", err)
	}

	db, err := sql.Open(DriverName, dsn)
	if err != nil {
		return nil, fmt.Errorf("error to connect db: %w", err)
	}

	setConnPool(db)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db ping error: %w", err)
	}

	log.Println("Success to connect DB")
	return db, nil

}

func buildDSN() (string, error) {
	requiredVars := []string{Db_host, Db_port, Db_user, Db_pass, Db_name}
	env := make(map[string]string)

	for _, key := range requiredVars {
		val, exists := os.LookupEnv(key)
		if val == "" || !exists {
			return "", fmt.Errorf("enviroment variable %s is not set", key)
		}
		env[key] = val
	}

	dsnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env[Db_host], env[Db_port], env[Db_user], env[Db_pass], env[Db_name])

	return dsnStr, nil
}

func setConnPool(db *sql.DB) {
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(50)
	db.SetConnMaxLifetime(2 * time.Minute)
	db.SetConnMaxIdleTime(30 * time.Second)
}
