package lookup

import (
	"log"
	"sqlpowered_bootstrap/api_config_management"
	"sqlpowered_bootstrap/testing_utils"
	"testing"

	"golang.org/x/exp/slices"
)

func TestListAllTables(t *testing.T) {

	apiConfigFileName := "../api_config.json"
	apiConfig, err := api_config_management.Load(apiConfigFileName)
	if err != nil {
		log.Fatalf("failed to load: %s, error: %v", apiConfigFileName, err)
	}

	apiConfig.ApiAccessibleTables = []string{
		"city",
		"district",
	}

	allTables, err := ListAllTables(
		apiConfig,
	)
	if err != nil {
		log.Fatalf("failed to ListAllTables, error: %v", err)
	}

	expectedAllTables := []string{
		"city",
		"district",
	}

	testing_utils.SlicesEqualOrderIndependent(
		allTables,
		expectedAllTables,
	)

}

func TestListAllTablesColumns(t *testing.T) {
	tablesList := []string{
		"city",
		"district",
		"retail_shop",
		"product",
		"retail_shop_product_stock",
	}

	apiConfigFileName := "../api_config.json"
	apiConfig, err := api_config_management.Load(apiConfigFileName)
	if err != nil {
		log.Fatalf("failed to load api_config.json")
	}

	tablesColumns, err := ListAllTablesColumns(
		apiConfig,
		tablesList,
	)
	if err != nil {
		log.Fatalf("unable to populate columns for provided tables")
	}
	log.Print(tablesColumns)

	expectedTablesColumns := map[string][]string{
		"city": {
			"id",
			"name",
		},
		"district": {
			"id",
			"name",
			"city",
		},
		"product": {
			"id",
			"name",
			"price",
		},
		"retail_shop": {
			"id",
			"name",
			"district_id",
		},
		"retail_shop_product_stock": {
			"id",
			"product_id",
			"retail_shop_id",
			"number_stocked",
		},
	}

	// order agnostic equals comparisons for slice
	for table, columnNameSlice := range expectedTablesColumns {
		if tablesColumns[table] == nil {
			log.Fatalf("unable to produce output for table: %s", table)
		}
		for _, columnName := range columnNameSlice {
			if !slices.Contains(tablesColumns[table], columnName) {
				log.Fatalf("missing column %s for table %s ", columnName, table)
			}
		}
	}

}

func TestListAllDb(t *testing.T) {
	apiConfig, err := api_config_management.Load("../api_config.json")
	if err != nil {
		log.Fatal(err)
	}

	excludedDbs := []string{"postgres"}
	excludedSchemas := []string{
		"information_schema",
		"pg_catalog",
		"pg_toast",
	}

	allDbs, err := ListAllDb(
		apiConfig,
		excludedDbs,
		excludedSchemas,
	)
	if err != nil {
		log.Fatalf("unable to ListAllDb")
	}
	log.Print(allDbs)
	if !slices.Equal(allDbs, []string{"sqlpowered_bootstrap"}) {
		log.Fatalf("unable to find standard db")
	}
}

func TestListAllSchemas(t *testing.T) {

	apiConfig, err := api_config_management.Load("../api_config.json")
	if err != nil {
		log.Fatal(err)
	}

	excludedSchemas := []string{
		"information_schema",
		"pg_catalog",
		"pg_toast",
	}

	allSchemas, err := ListAllSchemas(
		apiConfig,
		excludedSchemas,
	)
	if err != nil {
		log.Fatalf("unable to ListAllSchemas")
	}

	log.Print(allSchemas)

	// exits if the condition is false
	testing_utils.SlicesEqualOrderIndependent(allSchemas, []string{"public"})
}
