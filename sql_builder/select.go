package sql_builder

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/lib/pq"
)

type Case struct {
	Whens []When `json:"whens"`
	Else  Select `json:"else,omitempty"`
}

type When struct {
	When []Where `json:"when"`
	Then Select  `json:"then"`
}

type Select struct {
	Table  string `json:"table,omitempty"`
	Column string `json:"column,omitempty"`
	As     string `json:"as,omitempty"`
	Value  string `json:"value,omitempty"`
	Fns    []Fn   `json:"fns,omitempty"`
	Type   string `json:"type,omitempty"`
	Case   *Case  `json:"case,omitempty"`
}

type SelectOutput struct {
	Sql string
}

type QueryParameters struct {
	RecursionInQuery   int
	RowLimit           int
	NumColumnsSelected int
	Values             []string
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

func SelectBuildCase(
	inputSelect Select,
	outputSql string,
) (string, error) {

	return outputSql, nil
}

func SelectBuildFns(
	inputSelect Select,
	outputSql string,
) (string, error) {

	return outputSql, nil
}

// Build the sql string
func SelectBuildSqlString(
	inputSelect Select,
	tableAndColumnDefined bool,
	valueDefined bool,
	caseDefined bool,
) (string, error) {

	outputSql := ""
	err := error(nil)

	if tableAndColumnDefined {
		outputSql = fmt.Sprintf(`%s.%s`,
			pq.QuoteIdentifier(inputSelect.Table),
			pq.QuoteIdentifier(inputSelect.Column),
		)

	} else if valueDefined {
		outputSql = pq.QuoteLiteral(inputSelect.Value)

	} else if caseDefined {

		outputSql, err = SelectBuildCase(inputSelect, outputSql)
		if err != nil {
			return "", err
		}

	}

	if len(inputSelect.Fns) > 0 {
		outputSql, err = SelectBuildFns(inputSelect, outputSql)
		if err != nil {
			return "", err
		}
	}

	if inputSelect.Type != "" {
		outputSql = fmt.Sprintf("%s::%s", outputSql, inputSelect.Type)
	}

	if inputSelect.As != "" {
		outputSql = fmt.Sprintf("%s as %s", outputSql, pq.QuoteIdentifier(inputSelect.As))
	}

	log.Print("outputSql after SelectBuildSqlString: ", outputSql)

	return outputSql, nil
}

// replace "where fy_year > 2018" with "where fy_year > $1"
//
// store the original values inside of queryParameters.Values,
//
// len(queryParameters.Values)+1 slice gives database safe dollar value
//
// ($1, $2, $3 etc) which is stored back in inputSelect.Values
func ReplaceValuesQueryParameter(
	inputSelect Select,
	queryParameters QueryParameters,
) (
	Select,
	QueryParameters,
) {

	if inputSelect.Value != "" {

		// add the valuesItem to queryParameters, and replace with query parameter
		queryParameters.Values = append(queryParameters.Values, inputSelect.Value)

		// substitute $1, $2 etc for the original value, using the length of the `queryParameters.Values` array
		newSubstitutionValue := fmt.Sprintf("$%d", len(queryParameters.Values))
		inputSelect.Value = newSubstitutionValue

	}

	return inputSelect, queryParameters
}

func SelectValidateAndBuild(
	inputJson map[string]any,
	queryParameters QueryParameters,
	maxQueryParameters QueryParameters,
	permissions Permissions,
	allTables []string,
	allTablesColumns map[string][]string,
	// allDbs []string,
	// allSchemas []string,
) (string, error) {

	inputBytes, err := json.Marshal(inputJson)
	if err != nil {
		errorString := fmt.Sprintf("error creating Select from inputData: %+v", inputJson)
		log.Print(errorString)
		return "", fmt.Errorf(errorString)
	}
	inputSelect := Select{}
	json.Unmarshal(inputBytes, &inputSelect)

	log.Printf("The value of inputSelect: %+v", inputSelect)

	tableAndColumnDefined,
		valueDefined,
		caseDefined, err := SelectValidate(inputSelect, permissions)
	if err != nil {
		return "", err
	}

	// substitute inputSelect.Values with [$1,$2,$3 etc] and store in queryParameters.ValueList
	inputSelect, queryParameters = ReplaceValuesQueryParameter(
		inputSelect,
		queryParameters,
	)

	log.Print(inputSelect)
	log.Print(queryParameters)
	//===========================
	sqlText, err := SelectBuildSqlString(
		inputSelect,
		tableAndColumnDefined,
		valueDefined,
		caseDefined,
	)
	if err != nil {
		return "", err
	}

	log.Printf("sqlText: %s", sqlText)

	// return sqlText, nil

	return "", nil
}
