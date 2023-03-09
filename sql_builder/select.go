package sql_builder

import (
	"encoding/json"
	"fmt"
	"log"
	"sqlpowered_bootstrap/lookup"
)

type Select struct {
	ColumnName   string        `json:"column_name"`
	TableName    string        `json:"table_name"`
	ValuesList   []string      `json:"values_list"`
	FunctionList []SqlFunction `json:"function_list"`
	TypeCast     string        `json:"type_cast"`
	Case         *Case         `json:"case"`
}

type SelectOutput struct {
	Sql string
}

type QueryParameters struct {
	RecursionInQuery   int
	RowLimit           int
	NumColumnsSelected int
	ValuesList         []string
}

type Permissions struct {
	Select []TablePermissions
	Delete []TablePermissions
	Update []TablePermissions
	Insert []TablePermissions
}

type TablePermissions struct {
	ColumnList     []string
	ColumnsInWhere []string
}

// Check the role is allowed to do their current actions
func SelectCheckPermissions(
	inputSelect Select,
	permissions Permissions,
) error {
	return nil
}

// Build the sql string
func SelectBuildSqlString(
	inputSelect Select,
) string {
	return ""
}

// Check the Select is not exceeding the allowed maximums for query parameters
func SelectCheckParameters(
	inputSelect Select,
	queryParameters QueryParameters,
	maxQueryParameters QueryParameters,
) error {
	return nil
}

func SelectBuild(
	inputJson map[string]any,
	queryParameters QueryParameters,
	maxQueryParameters QueryParameters,
	permissions Permissions,
	allTables []string,
	allTablesColumns map[string][]string,
	// allDbs []string,
	// allSchemas []string,
) (string, error) {

	// helper functions
	log.Print(lookup.ValidFunctions())
	log.Print(lookup.ValidTypeCasts())

	inputBytes, err := json.Marshal(inputJson)
	if err != nil {
		errorString := fmt.Sprintf("error creating Select from inputData: %+v", inputJson)
		log.Print(errorString)
		return "", fmt.Errorf(errorString)
	}
	selectInstance := Select{}
	json.Unmarshal(inputBytes, &selectInstance)

	log.Printf("The value of selectInstance: %+v", selectInstance)

	//==========================
	err = SelectCheckPermissions(
		selectInstance,
		permissions,
	)
	if err != nil {
		errorString := fmt.Sprintf("permissions error creating Select, error: %+v", err)
		log.Print(errorString)
		return "", fmt.Errorf(errorString)
	}

	// substitute selectInstance.ValuesList with [$1,$2,$3 etc] and store in queryParameters.ValueList
	selectInstance, queryParameters = ReplaceValuesListQueryParameter(
		selectInstance,
		queryParameters,
	)

	//===========================
	sqlText := SelectBuildSqlString(
		selectInstance,
	)

	return sqlText, nil
}
