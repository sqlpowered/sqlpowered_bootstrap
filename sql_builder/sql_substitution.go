package sql_builder

import (
	"fmt"
)

// replace "where fy_year > 2018" with "where fy_year > $1"
//
// store the original values inside of queryParameters.ValuesList,
//
// len(queryParameters.ValuesList)+1 slice gives database safe dollar value
//
// ($1, $2, $3 etc) which is stored back in selectInstance.ValuesList
func ReplaceValuesListQueryParameter(
	selectInstance Select,
	queryParameters QueryParameters,
) (
	Select,
	QueryParameters,
) {

	if len(selectInstance.ValuesList) > 0 {
		for index, valuesItem := range selectInstance.ValuesList {

			// length is 1 indexed, so only need to add one
			newSubstitutionValue := fmt.Sprintf("$%d", len(queryParameters.ValuesList)+1)

			// add the valuesItem to queryParameters, and replace with query parameter
			queryParameters.ValuesList = append(queryParameters.ValuesList, valuesItem)
			selectInstance.ValuesList[index] = newSubstitutionValue
		}
	}

	return selectInstance, queryParameters
}
