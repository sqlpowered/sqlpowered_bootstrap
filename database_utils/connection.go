package database_utils

import (
	"database/sql"
	"fmt"
	"log"
	"sqlpowered_bootstrap/api_config_management"
)

// Establish database connection.
//
// The required api_config.json settings:
//
// databaseSslmode
//
// databaseName
//
// databaseHost
//
// databasePort
//
// databaseUsername
//
// databasePassword
func Connect(
	apiConfig api_config_management.ApiConfig,
) (*sql.DB, error) {

	if apiConfig.DatabaseSslmode == "" {
		log.Fatalf(`required api_config.json setting: "databaseSslmode" is not defined, exiting`)
	}
	if apiConfig.DatabaseName == "" {
		log.Fatalf(`required api_config.json setting: "databaseName" is not defined, exiting`)
	}
	if apiConfig.DatabaseHost == "" {
		log.Fatalf(`required api_config.json setting: "databaseHost" is not defined, exiting`)
	}
	if apiConfig.DatabasePort == "" {
		log.Fatalf(`required api_config.json setting: "databasePort" is not defined, exiting`)
	}
	if apiConfig.DatabaseUsername == "" {
		log.Fatalf(`required api_config.json setting: "databaseUsername" is not defined, exiting`)
	}
	if apiConfig.DatabasePassword == "" {
		log.Fatalf(`required api_config.json setting: "databasePassword" is not defined, exiting`)
	}

	connectionString := fmt.Sprintf("sslmode=%s dbname=%s host=%s port=%s user=%s password=%s ",
		apiConfig.DatabaseSslmode,
		apiConfig.DatabaseName,
		apiConfig.DatabaseHost,
		apiConfig.DatabasePort,
		apiConfig.DatabaseUsername,
		apiConfig.DatabasePassword,
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		errorString := fmt.Sprintf(`database connection string does not conform to standard format.
		See docs: https://pkg.go.dev/github.com/lib/pq#hdr-Connection_String_Parameters.
		Error: %s`, err)
		log.Print(errorString)
		return &sql.DB{}, fmt.Errorf(errorString)
	}

	err = db.Ping()
	if err != nil {
		errorString := fmt.Sprintf("Database connection Error: %s", err)
		log.Print(errorString)
		return &sql.DB{}, fmt.Errorf(errorString)
	}

	return db, nil
}
