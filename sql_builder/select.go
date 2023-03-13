package sql_builder

import (
	"encoding/json"
	"fmt"
	"log"
	"sqlpowered_bootstrap/lookup"

	"github.com/lib/pq"
	"golang.org/x/exp/slices"
)

type Case struct {
	CaseWhens []Whens `json:"whens"`
	Else      Select  `json:"else,omitempty"`
}

type Whens struct {
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

// Check the role is allowed to do their current actions
func SelectCheckPermissions(
	inputSelect Select,
	permissions Permissions,
) error {
	return nil
}

// Validate the Select type is one of:
//
// 1.) table and column
//
// 2.) values
//
// 3.) case
func ValidateSelectType(
	inputSelect Select,
) (bool, bool, bool, error) {

	tableDefined := false
	columnDefined := false
	tableAndColumnDefined := false
	valueDefined := false
	caseDefined := false

	if inputSelect.Table != "" {
		tableDefined = true
	}
	if inputSelect.Column != "" {
		columnDefined = true
	}
	if inputSelect.Value != "" {
		valueDefined = true
	}
	if inputSelect.Case != nil && len(inputSelect.Case.CaseWhens) > 0 {
		caseDefined = true
	}

	// validate the table and column only situation
	if tableDefined && columnDefined {

		tableAndColumnDefined = true

		if valueDefined || caseDefined {
			logString := fmt.Sprintf(
				`when "table"(%v) and "column"(%v) are defined, cannot also define: "value"(%v) or "case"(%v)`,
				tableDefined,
				columnDefined,
				valueDefined,
				caseDefined,
			)

			log.Print(logString)
			return false, false, false, fmt.Errorf(logString)
		}

		// must define both table and column
	} else if tableDefined && !columnDefined {

		logString := fmt.Sprintf(
			`when "table"(%v) is defined, "column"(%v) must also be defined`,
			tableDefined,
			columnDefined,
		)

		log.Print(logString)
		return false, false, false, fmt.Errorf(logString)

		// must define both table and column
	} else if !tableDefined && columnDefined {

		logString := fmt.Sprintf(
			`when "column"(%v) is defined, "table"(%v) must also be defined`,
			columnDefined,
			tableDefined,
		)

		log.Print(logString)
		return false, false, false, fmt.Errorf(logString)
	}

	// validate case only situation
	if caseDefined {
		if tableDefined || columnDefined || valueDefined {
			logString := fmt.Sprintf(
				`when "case"(%v) is defined, cannot also define: "table"(%v), "column"(%v) or "value"(%v)`,
				caseDefined,
				tableDefined,
				columnDefined,
				valueDefined,
			)

			log.Print(logString)
			return false, false, false, fmt.Errorf(logString)
		}
	}

	// validate values only situation
	if valueDefined {
		if caseDefined || tableDefined || columnDefined {

			logString := fmt.Sprintf(
				`when "value"(%v) is defined, cannot also define: "table"(%v), "column"(%v) or "case"(%v)`,
				valueDefined,
				tableDefined,
				columnDefined,
				caseDefined,
			)

			log.Print(logString)
			return false, false, false, fmt.Errorf(logString)
		}
	}

	return tableAndColumnDefined, valueDefined, caseDefined, nil
}

func SelectBuildFns(
	inputSelect Select,
	outputSql string,
) (string, error) {

	// when a "type" is defined, the first item in "Fns" cannot also have a type
	if inputSelect.Type != "" && len(inputSelect.Fns) > 0 {
		if inputSelect.Fns[0].Type != "" {

			logString := `when a top level "type" is defined, the first "fn" in "fns" cannot also have a "type"`
			log.Print(logString)
			return "", fmt.Errorf(logString)
		}
	}

	for _, fnItem := range inputSelect.Fns {
		// validate the function names are valid
		if !slices.Contains(lookup.ValidFunctions(), fnItem.Fn) {

			logString := fmt.Sprintf(`invalid "fn" value: %v, valid values: %v`,
				fnItem.Fn,
				lookup.ValidFunctions(),
			)
			log.Print(logString)
			return "", fmt.Errorf(logString)
		}

		// evaluate Fns: [{"fn": "first"}, {"fn": "second"}, {"fn": "third"}]

	}

	return "", nil
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
	}

	if inputSelect.Type != "" {

		// validate the "type" is valid
		if !slices.Contains(lookup.ValidTypeCasts(), inputSelect.Type) {

			logString := fmt.Sprintf(`invalid "type" value: %v, valid values: %v`,
				inputSelect.Type,
				lookup.ValidTypeCasts(),
			)
			log.Print(logString)
			return "", fmt.Errorf(logString)
		}

		outputSql = fmt.Sprintf("%s::%s", outputSql, inputSelect.Type)

	}

	if len(inputSelect.Fns) > 0 {
		outputSql, err = SelectBuildFns(inputSelect, outputSql)
		if err != nil {
			return "", err
		}
	}

	log.Print("outputSql after SelectBuildFns: ", outputSql)

	return "", nil
}

// Check the Select is not exceeding the allowed maximums for query parameters
func SelectCheckParameters(
	inputSelect Select,
	queryParameters QueryParameters,
	maxQueryParameters QueryParameters,
) error {
	return nil
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
		caseDefined,
		err := ValidateSelectType(
		inputSelect,
	)
	if err != nil {
		return "", err
	}
	log.Printf("tableAndColumnDefined(%v) valueDefined(%v) caseDefined(%v)",
		tableAndColumnDefined,
		valueDefined,
		caseDefined,
	)
	//==========================
	err = SelectCheckPermissions(
		inputSelect,
		permissions,
	)
	if err != nil {
		errorString := fmt.Sprintf("permissions error creating Select, error: %+v", err)
		log.Print(errorString)
		return "", fmt.Errorf(errorString)
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
