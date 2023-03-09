package database_utils

import (
	"database/sql"
	"fmt"
	"log"
)

//============================================
// JSON key value pairs output

func processRowKeyValuePairs(
	sqlRows *sql.Rows,
	columnNamesSlice []string,
	queryOutputJson []map[string]any,
) ([]map[string]any, error) {

	resultValues := make([]interface{}, len(columnNamesSlice))
	scanIntoPointerSlice := make([]interface{}, len(columnNamesSlice))

	for i := range columnNamesSlice {
		scanIntoPointerSlice[i] = &resultValues[i]
	}

	err := sqlRows.Scan(scanIntoPointerSlice...)
	if err != nil {
		errorString := fmt.Sprintf("error Scanning data from sqlRows, error: %v", err)
		log.Print(errorString)
		return nil, fmt.Errorf(errorString)
	}

	rowDataKeyValuePair := map[string]any{}

	for index, columnName := range columnNamesSlice {
		rowDataKeyValuePair[columnName] = resultValues[index]
	}
	queryOutputJson = append(queryOutputJson, rowDataKeyValuePair)

	return queryOutputJson, nil
}

func QueryIntoJson(
	db *sql.DB,
	sqlString string,
	unsafeDataSlice []string,
) ([]map[string]any, error) {

	unsafeDataSliceAny := []any{}
	for _, item := range unsafeDataSlice {
		unsafeDataSliceAny = append(unsafeDataSliceAny, item)
	}

	sqlRows, err := db.Query(sqlString, unsafeDataSliceAny...)
	if err != nil {
		errorString := fmt.Sprintf("query failed, error: %v, query: %v", err, sqlString)
		log.Print(errorString)
		return nil, fmt.Errorf(errorString)
	}
	defer sqlRows.Close()

	columnNamesSlice, err := sqlRows.Columns()
	if err != nil {
		errorString := fmt.Sprintf("fetching column names for query failed, error: %v, query: %v", err, sqlString)
		log.Print(errorString)
		return nil, fmt.Errorf(errorString)
	}

	queryOutputJson := []map[string]any{}
	for sqlRows.Next() {

		queryOutputJson, err = processRowKeyValuePairs(
			sqlRows,
			columnNamesSlice,
			queryOutputJson,
		)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}
	return queryOutputJson, nil
}

//============================================
// JSON Array output

func processRowToArrays(
	sqlRows *sql.Rows,
	columnNamesSlice []string,
	queryOutputJson map[string][]any,
) (map[string][]any, error) {

	resultValues := make([]interface{}, len(columnNamesSlice))
	scanIntoPointerSlice := make([]interface{}, len(columnNamesSlice))

	for i := range columnNamesSlice {
		scanIntoPointerSlice[i] = &resultValues[i]
	}

	err := sqlRows.Scan(scanIntoPointerSlice...)
	if err != nil {
		errorString := fmt.Sprintf("error Scanning data from sqlRows, error: %v", err)
		log.Print(errorString)
		return nil, fmt.Errorf(errorString)
	}

	for index, columnName := range columnNamesSlice {
		queryOutputJson[columnName] = append(queryOutputJson[columnName], resultValues[index])
	}

	return queryOutputJson, nil
}

func QueryIntoJsonArrays(
	db *sql.DB,
	sqlString string,
	unsafeDataSlice []string,
) (map[string][]any, error) {

	unsafeDataSliceAny := []any{}
	for _, item := range unsafeDataSlice {
		unsafeDataSliceAny = append(unsafeDataSliceAny, item)
	}

	sqlRows, err := db.Query(sqlString, unsafeDataSliceAny...)
	if err != nil {
		errorString := fmt.Sprintf("query failed, error: %v, query: %v", err, sqlString)
		log.Print(errorString)
		return nil, fmt.Errorf(errorString)
	}
	defer sqlRows.Close()

	columnNamesSlice, err := sqlRows.Columns()
	if err != nil {
		errorString := fmt.Sprintf("fetching column names for query failed, error: %v, query: %v", err, sqlString)
		log.Print(errorString)
		return nil, fmt.Errorf(errorString)
	}

	queryOutputJson := map[string][]any{}
	for sqlRows.Next() {

		queryOutputJson, err = processRowToArrays(
			sqlRows,
			columnNamesSlice,
			queryOutputJson,
		)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}
	return queryOutputJson, nil
}

//============================================
// CSV output

func processRowCsv(
	sqlRows *sql.Rows,
	columnNamesSlice []string,
	// columnTypesSlice []*sql.ColumnType,
	queryOutputCsv [][]any,
) ([][]any, error) {

	resultValues := make([]interface{}, len(columnNamesSlice))
	scanIntoPointerSlice := make([]interface{}, len(columnNamesSlice))

	for i := range columnNamesSlice {
		scanIntoPointerSlice[i] = &resultValues[i]
	}

	err := sqlRows.Scan(scanIntoPointerSlice...)
	if err != nil {
		errorString := fmt.Sprintf("error Scanning data from sqlRows, error: %v", err)
		log.Print(errorString)
		return nil, fmt.Errorf(errorString)
	}

	queryOutputCsv = append(queryOutputCsv, resultValues)

	return queryOutputCsv, nil
}

func QueryIntoCsv(
	db *sql.DB,
	sqlString string,
	unsafeDataSlice []string,
) ([][]any, error) {

	unsafeDataSliceAny := []any{}
	for _, item := range unsafeDataSlice {
		unsafeDataSliceAny = append(unsafeDataSliceAny, item)
	}

	sqlRows, err := db.Query(sqlString, unsafeDataSliceAny...)
	if err != nil {
		errorString := fmt.Sprintf("query failed, error: %v, query: %v", err, sqlString)
		log.Print(errorString)
		return nil, fmt.Errorf(errorString)
	}
	defer sqlRows.Close()

	columnNamesSlice, err := sqlRows.Columns()
	if err != nil {
		errorString := fmt.Sprintf("fetching column names for query failed, error: %v, query: %v", err, sqlString)
		log.Print(errorString)
		return nil, fmt.Errorf(errorString)
	}

	// each row of the csv is a slice in the slice of slices
	queryOutputCsv := [][]any{}

	// add the row of column labels
	rowData := make([]any, len(columnNamesSlice))

	// add column header to CSV
	for index, data := range columnNamesSlice {
		rowData[index] = data
	}
	queryOutputCsv = append(queryOutputCsv, rowData)

	for sqlRows.Next() {

		queryOutputCsv, err = processRowCsv(
			sqlRows,
			columnNamesSlice,
			// columnTypesSlice,
			queryOutputCsv,
		)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}
	return queryOutputCsv, nil
}
