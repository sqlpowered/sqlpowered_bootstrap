package database_utils

import (
	"database/sql"
	"fmt"
	"log"
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
func Connect(apiConfig map[string]string) (*sql.DB, error) {

	for _, envVar := range []string{
		"databaseSslmode",
		"databaseName",
		"databaseHost",
		"databasePort",
		"databaseUsername",
		"databasePassword",
	} {
		if apiConfig[envVar] == "" {
			log.Fatalf("required api_config.json setting: %s is not defined, exiting", envVar)
		}
	}

	connectionString := fmt.Sprintf("sslmode=%s dbname=%s host=%s port=%s user=%s password=%s ",
		apiConfig["databaseSslmode"],
		apiConfig["databaseName"],
		apiConfig["databaseHost"],
		apiConfig["databasePort"],
		apiConfig["databaseUsername"],
		apiConfig["databasePassword"],
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
