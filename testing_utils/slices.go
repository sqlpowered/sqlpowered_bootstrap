package testing_utils

import (
	"log"

	"golang.org/x/exp/slices"
)

func SlicesEqualOrderIndependent[inputType comparable](
	slice1 []inputType,
	slice2 []inputType,
) {

	for _, slice1Item := range slice1 {
		if !slices.Contains(slice2, slice1Item) {
			log.Fatalf("missing %v from %+v", slice1Item, slice2)
		}
	}

	for _, slice2Item := range slice2 {
		if !slices.Contains(slice1, slice2Item) {
			log.Printf("missing %v from %+v", slice2Item, slice1)
		}
	}
}
