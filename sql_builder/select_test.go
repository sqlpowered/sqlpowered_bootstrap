package sql_builder

import (
	"log"
	"sqlpowered_bootstrap/api_config_management"
	"sqlpowered_bootstrap/lookup"
	"testing"

	"golang.org/x/exp/slices"
)

func TestSelectBuild(t *testing.T) {

	inputData := map[string]any{
		"column": "name",
		"table":  "product",
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

func TestReplaceValuesQueryParameter(t *testing.T) {

	// expectedQueryParametersOutput := []string{"drop tables;", "--", ";"}
	// selectOutput := []string{"$1", "$2", "$3"}

	// TODO: add more items to this select as support expands
	expectedQueryParametersOutput := []string{"drop tables;"}
	selectOutput := "$1"

	inputSelect := Select{Value: "drop tables;"}
	queryParameters := QueryParameters{}

	log.Printf("%+v", inputSelect)
	log.Printf("%+v\n\n", queryParameters)
	inputSelect, queryParameters = ReplaceValuesQueryParameter(
		inputSelect,
		queryParameters,
	)
	log.Printf("%+v", inputSelect)
	log.Printf("%+v", queryParameters)

	if !slices.Equal(queryParameters.Values, expectedQueryParametersOutput) {
		log.Fatalf("unable to produce expected output: %v in queryParameters.Values: %v",
			expectedQueryParametersOutput,
			queryParameters.Values,
		)
	}

	if inputSelect.Value != selectOutput {
		log.Fatalf(`unable to produce selectOutput: %v in "values": %v`,
			selectOutput,
			inputSelect.Value,
		)
	}

}
