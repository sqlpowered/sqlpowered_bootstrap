package sql_builder

import (
	"log"
	"sqlpowered_bootstrap/api_config_management"
	"sqlpowered_bootstrap/lookup"
	"testing"
)

func TestSelectBuild(t *testing.T) {

	// ColumnName   string        `json:"column_name"`
	// TableName    string        `json:"table_name"`
	// ValuesList   []string      `json:"values_list"`
	// FunctionList []SqlFunction `json:"function_list"`
	// TypeCast     string        `json:"type_cast"`
	inputData := map[string]any{
		"column_name": "name",
		"table_name":  "product",
	}

	apiConfigFilename := "../api_config.json"
	apiConfig, err := api_config_management.Load(apiConfigFilename)
	if err != nil {
		log.Fatalf("%v", err)
	}
	allTables, err := lookup.ListAllTables(
		apiConfig,
		[]string{},
	)
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Printf("allTables: %+v", allTables)

	allTablesColumns, err := lookup.ListAllTablesColumns(
		apiConfig,
		allTables,
	)
	if err != nil {
		log.Fatalf("%v", err)
	}

	SelectBuild(
		inputData,
		QueryParameters{},
		QueryParameters{
			RecursionInQuery:   5,
			RowLimit:           1000,
			NumColumnsSelected: 100,
			Values:             []string{},
		},
		Permissions{},
		allTables,
		allTablesColumns,
	)

}
