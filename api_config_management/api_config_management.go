package api_config_management

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ApiConfig struct {
	DatabaseSslmode     string   `json:"databaseSslmode"`
	DatabaseSchemaName  string   `json:"databaseSchemaName"`
	DatabaseName        string   `json:"databaseName"`
	DatabaseHost        string   `json:"databaseHost"`
	DatabasePort        string   `json:"databasePort"`
	DatabaseUsername    string   `json:"databaseUsername"`
	DatabasePassword    string   `json:"databasePassword"`
	ApiAccessibleTables []string `json:"apiAccessibleTables"`
}

func Load(
	fileName string,
) (ApiConfig, error) {

	apiConfigBytes, err := os.ReadFile(fileName)
	if err != nil {
		errorString := fmt.Sprintf("unable to load data from: api_config.json error: %+v", err)
		log.Print(errorString)
		return ApiConfig{}, fmt.Errorf(errorString)
	}

	apiConfig := ApiConfig{}

	err = json.Unmarshal(apiConfigBytes, &apiConfig)
	if err != nil {
		errorString := fmt.Sprintf("unable to process data loaded from: api_config.json error: %+v", err)
		log.Print(errorString)
		return ApiConfig{}, fmt.Errorf(errorString)
	}

	return apiConfig, nil
}
