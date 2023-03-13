package lookup

// Valid operators for join statement, for joining two tables
func ValidJoinOperators() []string {
	return []string{
		"eq",
		"neq",
		"gt",
		"lt",
	}
}

// Valid operators for where statement, for comparing values
func ValidWhereOperators() []string {
	return []string{
		"exists",
		"not exists",
		"eq",
		"neq",
		"in",
		"not in",
		"gt",
		"lt",
		"gte",
		"lte",
	}
}

// Valid functions we can use in the select statment
func ValidFunctions() []string {
	return []string{
		"count",
		"distinct",
		"avg",
		"min",
		"max",

		"mult",
		"div",
		"sub",
		"add",

		//================================
		// TODO: these need reviewing
		// Had a stab at making names for these operators
		// look for standard ones if possible
		// Ref: http://www.silota.com/docs/recipes/sql-postgres-json-data-types.html
		// Ref: https://www.postgresql.org/docs/15/functions-json.html
		// Ref: https://www.buckenhofer.com/2019/01/json-and-iso-sql-standard/
		// Lack of standardisation in different dialects for working with SQL: https://blog.jooq.org/standard-sql-json-the-sobering-parts/
		// "jsonb_del",          // -
		// "jsonb_add",          // ||
		// "jsonb_contains_obj", // @>
		// "jsonb_into",         // ->
		// "jsonb_return_str",   // ->>
		// //================================
		// "jsonb_each",         // jsonb_each
		// "jsonb_object_keys",  // jsonb_object_keys
		// "jsonb_extract_path", // jsonb_extract_path
		// "jsonb_pretty",       // jsonb_pretty
	}
}

func ValidTypes() []string {
	return []string{
		"varchar",
		"text",
		"float",
		"integer",
		"numeric",
		"timestamp",
		"interval",
		"date",
		"money",
	}
}
