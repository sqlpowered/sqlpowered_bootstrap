package lookup

import (
	"fmt"
	"log"
	"sqlpowered_bootstrap/database_utils"
	"strings"

	"github.com/lib/pq"
)

func ListAllTables(
	apiConfig map[string]string,
	excludedTables []string,
) ([]string, error) {

	db, err := database_utils.Connect(apiConfig)
	if err != nil {
		return nil, err
	}

	sqlQuery := ""

	// TODO: use QueryIntoJson to run queries, and pass values to the db safely
	// Then can support values from an API as well as the current config file
	if len(excludedTables) > 0 {
		excludedTablesQuoted := []string{}
		for _, table := range excludedTables {
			excludedTablesQuoted = append(excludedTablesQuoted, pq.QuoteLiteral(table))
		}

		sqlQuery = fmt.Sprintf(`
		select
			table_name
		from
			information_schema.tables
		where
			table_schema = %s
			and table_catalog = %s
			and table_name not in (%s)
		order by 
			table_name;
		`,
			pq.QuoteLiteral(apiConfig["databaseSchemaName"]),
			pq.QuoteLiteral(apiConfig["databaseName"]),
			strings.Join(excludedTablesQuoted, ","),
		)

	} else if len(excludedTables) == 0 {

		sqlQuery = fmt.Sprintf(`
		select
			table_name
		from
			information_schema.tables
		where
			table_schema = %s
			and table_catalog = %s
		order by 
			table_name;
		`,
			pq.QuoteLiteral(apiConfig["databaseSchemaName"]),
			pq.QuoteLiteral(apiConfig["databaseName"]),
		)
	}

	log.Print(sqlQuery)

	sqlRows, err := db.Query(sqlQuery)

	if err != nil {
		errorString := fmt.Sprintf("Failed querying database, error: %s", err)
		log.Print(errorString)
		return nil, err
	}
	defer sqlRows.Close()

	tableSlice := []string{}
	table := ""

	for sqlRows.Next() {

		err := sqlRows.Scan(
			&table,
		)
		if err != nil {
			errorString := fmt.Sprintf("Failed querying database, error: %s", err)
			log.Print(errorString)
			return nil, err
		}
		tableSlice = append(tableSlice, table)
	}

	return tableSlice, nil
}

func ListAllTablesColumns(
	apiConfig map[string]string,
	tablesList []string,
) (map[string][]string, error) {

	db, err := database_utils.Connect(apiConfig)
	if err != nil {
		return nil, err
	}

	includedTablesQuoted := []string{}
	for _, table := range tablesList {
		includedTablesQuoted = append(includedTablesQuoted, pq.QuoteLiteral(table))
	}

	// this outputs easy to split strings
	sqlQuery := fmt.Sprintf(`
	select
		table_name,
		string_agg(column_name, ',') as column_name
	from
		information_schema.columns
	where
		table_name in (%s)
		and table_schema = %s
		and table_catalog = %s
	group by
		table_name
	order by 
		table_name, column_name;
	`,
		strings.Join(includedTablesQuoted, ","),
		pq.QuoteLiteral(apiConfig["databaseSchemaName"]),
		pq.QuoteLiteral(apiConfig["databaseName"]),
	)

	log.Print(sqlQuery)

	sqlRows, err := db.Query(sqlQuery)

	if err != nil {
		errorString := fmt.Sprintf("Failed querying database, error: %s", err)
		log.Print(errorString)
		return nil, fmt.Errorf(errorString)
	}
	defer sqlRows.Close()

	table := ""
	columnNamesString := ""

	tableColumnsMap := map[string][]string{}

	for sqlRows.Next() {

		err = sqlRows.Scan(
			&table,
			&columnNamesString,
		)
		if err != nil {
			errorString := fmt.Sprintf("unable to ListAllTablesColumns, error: %v", err)
			log.Print(errorString)
			return nil, fmt.Errorf(errorString)
		}
		// query outputs array as a string, which we split back into an array in the function here
		tableColumnsMap[table] = strings.Split(columnNamesString, ",")
	}
	return tableColumnsMap, nil

}

func ListAllDb(
	apiConfig map[string]string,
	excludedDbs []string,
	excludedSchemas []string,
) ([]string, error) {

	db, err := database_utils.Connect(apiConfig)
	if err != nil {
		return nil, err
	}

	excludedDbsQuoted := []string{}
	for _, dbName := range excludedDbs {
		excludedDbsQuoted = append(excludedDbsQuoted, pq.QuoteLiteral(dbName))
	}
	excludedSchemasQuoted := []string{}
	for _, schemaName := range excludedSchemas {
		excludedSchemasQuoted = append(excludedSchemasQuoted, pq.QuoteLiteral(schemaName))
	}

	// sqlQuery := fmt.Sprintf(`
	// select
	//     datname
	// from
	//     pg_database
	// where
	//     datistemplate = false
	// 	and datname not in (%s)
	// order by
	// 	datname;
	// `, strings.Join(excludedDbsQuoted, ","))

	sqlQuery := fmt.Sprintf(`
	select
		distinct table_catalog, table_schema
	from
		information_schema.columns
	where 
		table_schema not in (%s)
		and table_catalog not in (%s);
	`,
		strings.Join(excludedSchemasQuoted, ","),
		strings.Join(excludedDbsQuoted, ","),
	)

	log.Print(sqlQuery)

	sqlRows, err := db.Query(sqlQuery)

	if err != nil {
		errorString := fmt.Sprintf("Failed querying database, error: %s", err)
		log.Print(errorString)
		return nil, fmt.Errorf(errorString)
	}
	defer sqlRows.Close()

	dbName := ""
	schemaName := ""
	dbNameSlice := []string{}

	for sqlRows.Next() {

		err = sqlRows.Scan(
			&dbName,
			&schemaName,
		)
		if err != nil {
			errorString := fmt.Sprintf("unable to ListAllTablesColumns, error: %v", err)
			log.Print(errorString)
			return nil, fmt.Errorf(errorString)
		}

		dbNameSlice = append(dbNameSlice, dbName)
	}
	return dbNameSlice, nil
}

func ListAllSchemas(
	apiConfig map[string]string,
	excludedSchemas []string,
) ([]string, error) {

	db, err := database_utils.Connect(apiConfig)
	if err != nil {
		return nil, err
	}

	excludedSchemasQuoted := []string{}
	for _, schemaName := range excludedSchemas {
		excludedSchemasQuoted = append(excludedSchemasQuoted, pq.QuoteLiteral(schemaName))
	}

	sqlQuery := fmt.Sprintf(`
	select
		schema_name
	from 
		information_schema.schemata
	where
		schema_name not in (%s)
	order by
		schema_name;
    `, strings.Join(excludedSchemasQuoted, ","))

	log.Print(sqlQuery)

	sqlRows, err := db.Query(sqlQuery)

	if err != nil {
		errorString := fmt.Sprintf("Failed querying database, error: %s", err)
		log.Print(errorString)
		return nil, fmt.Errorf(errorString)
	}
	defer sqlRows.Close()

	schemaName := ""
	schemaNameSlice := []string{}

	for sqlRows.Next() {

		err = sqlRows.Scan(
			&schemaName,
		)
		if err != nil {
			errorString := fmt.Sprintf("unable to ListAllTablesColumns, error: %v", err)
			log.Print(errorString)
			return nil, fmt.Errorf(errorString)
		}

		schemaNameSlice = append(schemaNameSlice, schemaName)
	}
	return schemaNameSlice, nil
}
