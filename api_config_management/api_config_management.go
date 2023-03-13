package api_config_management

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Load(fileName string) (map[string]string, error) {

	apiConfigBytes, err := os.ReadFile(fileName)
	if err != nil {
		errorString := fmt.Sprintf("unable to load data from: api_config.json error: %+v", err)
		log.Print(errorString)
		return nil, fmt.Errorf(errorString)
	}

	apiConfig := map[string]string{}
	err = json.Unmarshal(apiConfigBytes, &apiConfig)
	if err != nil {
		errorString := fmt.Sprintf("unable to process data loaded from: api_config.json error: %+v", err)
		log.Print(errorString)
		return nil, fmt.Errorf(errorString)
	}

	return apiConfig, nil
}
