package sql_builder

import (
	"log"
	"testing"

	"golang.org/x/exp/slices"
)

func TestReplaceValuesListQueryParameter(t *testing.T) {

	expectedQueryParametersOutput := []string{"drop tables;", "--", ";"}
	expectedSelectInstanceOutput := []string{"$1", "$2", "$3"}

	selectInstance := Select{Values: slices.Clone(expectedQueryParametersOutput)}
	queryParameters := QueryParameters{}

	log.Printf("%+v", selectInstance)
	log.Printf("%+v\n\n", queryParameters)
	selectInstance, queryParameters = ReplaceValuesListQueryParameter(
		selectInstance,
		queryParameters,
	)
	log.Printf("%+v", selectInstance)
	log.Printf("%+v", queryParameters)

	if !slices.Equal(queryParameters.Values, expectedQueryParametersOutput) {
		log.Fatalf("unable to produce expected output: %v in queryParameters.ValuesList: %v",
			expectedQueryParametersOutput,
			queryParameters.Values,
		)
	}

	if !slices.Equal(selectInstance.Values, expectedSelectInstanceOutput) {
		log.Fatalf("unable to produce expected output: %v in queryParameters.ValuesList: %v",
			expectedSelectInstanceOutput,
			selectInstance.Values,
		)
	}

}
