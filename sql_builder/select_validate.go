package sql_builder

import (
	"fmt"
	"log"
	"sqlpowered_bootstrap/lookup"

	"golang.org/x/exp/slices"
)

// Check the role is allowed to do their current actions
func SelectCheckPermissions(
	inputSelect Select,
	permissions Permissions,
) error {
	return nil
}

// Validate the Select category is one of:
//
// 1.) table and column
//
// 2.) values
//
// 3.) case
func SelectValidateCategory(
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
	if inputSelect.Case != nil && len(inputSelect.Case.Whens) > 0 {
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

	// Validate select is in a valid category
	if !tableAndColumnDefined && !valueDefined && !caseDefined {

		logString := fmt.Sprintf(
			`select not in known category: both "table"(%v) and "column"(%v) are not defined, "value"(%v) is not defined, "case"(%v) is not defined`,
			tableDefined,
			columnDefined,
			valueDefined,
			caseDefined,
		)

		log.Print(logString)
		return false, false, false, fmt.Errorf(logString)
	}

	return tableAndColumnDefined, valueDefined, caseDefined, nil
}

func SelectValidateType(
	inputSelect Select,
) error {

	if inputSelect.Type != "" {

		// validate the "type" is valid
		if !slices.Contains(lookup.ValidTypes(), inputSelect.Type) {

			logString := fmt.Sprintf(`invalid "type": %v, valid values: %v`,
				inputSelect.Type,
				lookup.ValidTypes(),
			)
			log.Print(logString)
			return fmt.Errorf(logString)
		}

	}
	return nil

}

func SelectValidateFns(
	inputSelect Select,
	permissions Permissions,
) error {
	// when a "type" is defined, the first item in "Fns" cannot also have a type
	if inputSelect.Type != "" && len(inputSelect.Fns) > 0 {
		if inputSelect.Fns[0].Type != "" {

			logString := `when a top level "type" is defined, the first "fn" in "fns" cannot also have a "type"`
			log.Print(logString)
			return fmt.Errorf(logString)
		}
	}

	for _, fnItem := range inputSelect.Fns {

		// validate the function names are valid
		if !slices.Contains(lookup.ValidFunctions(), fnItem.Fn) {

			logString := fmt.Sprintf(`invalid "fn": %v, valid values: %v`,
				fnItem.Fn,
				lookup.ValidFunctions(),
			)
			log.Print(logString)
			return fmt.Errorf(logString)
		}

		// validate the function names are valid
		if !slices.Contains(lookup.ValidTypes(), fnItem.Type) {

			logString := fmt.Sprintf(`invalid "Type" on "fn": %v, valid values: %v`,
				fnItem.Type,
				lookup.ValidTypes(),
			)
			log.Print(logString)
			return fmt.Errorf(logString)
		}

		// validate inputSelect.Args[] which are Select[]
		for _, arg := range fnItem.Args {
			_, _, _, err := SelectValidate(arg, permissions)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

// TODO: validate "case.whens[].when" (which is of type "where")`
func SelectValidateCase(
	inputSelect Select,
	permissions Permissions,
) error {

	if len(inputSelect.Case.Whens) == 0 {
		logString := `if "case" is defined, "case.whens" cannot be empty, either define "whens" or remove from query`
		log.Print(logString)
		return fmt.Errorf(logString)
	}

	for _, whenItem := range inputSelect.Case.Whens {

		// TODO: validate when (which is of type "where")
		log.Print(`TODO: validate "case.whens[].when" (which is of type "where")`)

		// validate case.whens.then
		_, _, _, err := SelectValidate(whenItem.Then, permissions)
		if err != nil {
			return err
		}

	}

	// validate case.else
	_, _, _, err := SelectValidate(inputSelect.Case.Else, permissions)
	if err != nil {
		return err
	}

	return nil
}

// Check the Select is not exceeding the allowed maximums for query parameters
func SelectCheckParameters(
	inputSelect Select,
	queryParameters QueryParameters,
	maxQueryParameters QueryParameters,
) error {
	return nil
}

func SelectValidate(
	inputSelect Select,
	permissions Permissions,
) (bool, bool, bool, error) {

	tableAndColumnDefined,
		valueDefined,
		caseDefined,
		err := SelectValidateCategory(inputSelect)
	if err != nil {
		return false, false, false, err
	}

	log.Printf("tableAndColumnDefined(%v) valueDefined(%v) caseDefined(%v)",
		tableAndColumnDefined,
		valueDefined,
		caseDefined,
	)

	err = SelectValidateFns(inputSelect, permissions)
	if err != nil {
		return tableAndColumnDefined, valueDefined, caseDefined, err
	}

	err = SelectValidateType(inputSelect)
	if err != nil {
		return tableAndColumnDefined, valueDefined, caseDefined, err
	}

	if caseDefined {
		err = SelectValidateCase(inputSelect, permissions)
		if err != nil {
			return tableAndColumnDefined, valueDefined, caseDefined, err
		}
	}

	//==========================
	err = SelectCheckPermissions(
		inputSelect,
		permissions,
	)
	if err != nil {
		errorString := fmt.Sprintf("permissions error creating Select, error: %+v", err)
		log.Print(errorString)
		return tableAndColumnDefined, valueDefined, caseDefined, fmt.Errorf(errorString)
	}

	return tableAndColumnDefined, valueDefined, caseDefined, nil
}
