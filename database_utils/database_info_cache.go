package database_utils

import (
	"sync"
	"time"
)

type AllSchemasCache struct {
	AllSchemas        []string
	WriteLock         sync.Mutex
	TimeCacheAcquired time.Time
}

type AllDatabasesCache struct {
	AllDatabases      []string
	WriteLock         sync.Mutex
	TimeCacheAcquired time.Time
}

type AllTablesCache struct {
	AllTables         []string
	WriteLock         sync.Mutex
	TimeCacheAcquired time.Time
}

type AllTablesColumnsCache struct {
	AllTablesColumns  map[string][]string
	WriteLock         sync.Mutex
	TimeCacheAcquired time.Time
}

// TODO: add query cache
//
// TODO: add tiered caching, akin to CPU cache structure, but inverted,
// with repeatedly accessed queries being promoted to higher priority caches
type QueryCache struct {
	QueryHash         map[string]string
	WriteLock         sync.Mutex
	TimeCacheAcquired time.Time
}

// TODO: finish
func UpdateAllDatabaseCaches(
	allTablesCache *AllTablesCache,
	allTablesColumnsCache *AllTablesColumnsCache,
	apiConfigFilename string,
) {

	return
}

// allTables, err := lookup.ListAllTables(
// 	[]string{},
// 	apiConfig["databaseSchemaName"],
// 	apiConfig["databaseName"],
// 	apiConfigFilename,
// )
// if err != nil {
// 	log.Fatalf("%v", err)
// }

// allTablesColumns, err := lookup.ListAllTablesColumns(
// 	allTables,
// 	apiConfig["databaseSchemaName"],
// 	apiConfig["databaseName"],
// 	apiConfigFilename,
// )
// if err != nil {
// 	log.Fatalf("%v", err)
// }
