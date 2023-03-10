package sql_builder

import (
	"encoding/json"
	"fmt"
	"log"
	"sqlpowered_bootstrap/lookup"

	"golang.org/x/exp/slices"
)

type Case struct {
	CaseWhens      []CaseWhens `json:"whens"`
	ElseExpression Select      `json:"else"`
}

type CaseWhens struct {
	When []Where `json:"when"`
	Then Select  `json:"then"`
}

type Select struct {
	Table  string   `json:"table"`
	Column string   `json:"column"`
	As     string   `json:"as"`
	Values []string `json:"values"`
	Fns    []Fn     `json:"fns"`
	Type   string   `json:"type"`
	Case   *Case    `json:"case"`
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
) (bool, bool, bool, bool, error) {

	tableDefined := false
	columnDefined := false
	valuesDefined := false
	caseDefined := false

	if inputSelect.Table != "" {
		tableDefined = true
	}
	if inputSelect.Column != "" {
		columnDefined = true
	}
	if len(inputSelect.Values) > 0 {
		valuesDefined = true
	}
	if len(inputSelect.Case.CaseWhens) > 0 {
		caseDefined = true
	}

	// we force table and column to both be defined together
	if tableDefined && !columnDefined {
		logString := fmt.Sprintf(
			`when "table"(%v) is defined, "column"(%v) must also be defined`,
			tableDefined,
			columnDefined,
		)

		log.Print(logString)
		return false, false, false, false, fmt.Errorf(logString)
	}
	if !tableDefined && columnDefined {
		logString := fmt.Sprintf(
			`when "column"(%v) is defined, "table"(%v) must also be defined`,
			columnDefined,
			tableDefined,
		)

		log.Print(logString)
		return false, false, false, false, fmt.Errorf(logString)
	}

	// validate the table and column only situation
	if tableDefined && columnDefined {
		if valuesDefined || caseDefined {
			logString := fmt.Sprintf(
				`when "table"(%v) and "column"(%v) are defined, cannot also define: "values"(%v) or "case"(%v)`,
				tableDefined,
				columnDefined,
				valuesDefined,
				caseDefined,
			)

			log.Print(logString)
			return false, false, false, false, fmt.Errorf(logString)
		}
	}

	// validate case only situation
	if caseDefined {
		if tableDefined || columnDefined || valuesDefined {
			logString := fmt.Sprintf(
				`when "case"(%v) is defined, cannot also define: "table"(%v), "column"(%v) or "values"(%v)`,
				caseDefined,
				tableDefined,
				columnDefined,
				valuesDefined,
			)

			log.Print(logString)
			return false, false, false, false, fmt.Errorf(logString)
		}
	}

	// validate values only situation
	if valuesDefined {
		if caseDefined || tableDefined || columnDefined {

			logString := fmt.Sprintf(
				`when "values"(%v) is defined, cannot also define: "table"(%v), "column"(%v) or "case"(%v)`,
				valuesDefined,
				tableDefined,
				columnDefined,
				caseDefined,
			)

			log.Print(logString)
			return false, false, false, false, fmt.Errorf(logString)
		}
	}

	return tableDefined, columnDefined, valuesDefined, caseDefined, nil
}

// Build the sql string
func SelectBuildSqlString(
	inputSelect Select,
) (string, error) {

	if len(inputSelect.Fns) > 0 {
		for _, fnItem := range inputSelect.Fns {
			if !slices.Contains(lookup.ValidFunctions(), fnItem.Fn) {

				logString := fmt.Sprintf(`invalid "fn" value: %v, valid values: %v`,
					fnItem.Fn,
					lookup.ValidFunctions(),
				)
				log.Print(logString)
				return "", fmt.Errorf(logString)
			}
		}
	}

	if inputSelect.Type != "" {
		if !slices.Contains(lookup.ValidTypeCasts(), inputSelect.Type) {

			logString := fmt.Sprintf(`invalid "type" value: %v, valid values: %v`,
				inputSelect.Type,
				lookup.ValidTypeCasts(),
			)
			log.Print(logString)
			return "", fmt.Errorf(logString)

		}
	}

	// when a "type" is defined, the first item in "Fns" cannot also have a type
	if inputSelect.Type != "" && len(inputSelect.Fns) > 0 {
		if inputSelect.Fns[0].Type != "" {

			logString := `when a "type" is defined, the first item in "Fns" cannot also have a type`
			log.Print(logString)
			return "", fmt.Errorf(logString)
		}
	}

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

	if len(inputSelect.Values) > 0 {
		for index, valuesItem := range inputSelect.Values {

			// add the valuesItem to queryParameters, and replace with query parameter
			queryParameters.Values = append(queryParameters.Values, valuesItem)

			// substitute $1, $2 etc for the original value, using the length of the `queryParameters.Values` array
			newSubstitutionValue := fmt.Sprintf("$%d", len(queryParameters.Values))
			inputSelect.Values[index] = newSubstitutionValue
		}
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

	// helper functions
	log.Print(lookup.ValidFunctions())
	log.Print(lookup.ValidTypeCasts())

	inputBytes, err := json.Marshal(inputJson)
	if err != nil {
		errorString := fmt.Sprintf("error creating Select from inputData: %+v", inputJson)
		log.Print(errorString)
		return "", fmt.Errorf(errorString)
	}
	inputSelect := Select{}
	json.Unmarshal(inputBytes, &inputSelect)

	log.Printf("The value of inputSelect: %+v", inputSelect)

	tableDefined,
		columnDefined,
		valuesDefined,
		caseDefined,
		err := ValidateSelectType(
		inputSelect,
	)
	if err != nil {
		return "", err
	}
	log.Printf("tableDefined(%v) columnDefined(%v) valuesDefined(%v) caseDefined(%v)",
		tableDefined,
		columnDefined,
		valuesDefined,
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
	// sqlText, err := SelectBuildSqlString(
	// 	inputSelect,
	// )
	// if err != nil {
	// 	return "", err
	// }

	// return sqlText, nil

	return "", nil
}
