package sql_builder

type From struct {
	TableNames []string `json:"tableNames,omitempty"`
}

type FromOutput struct {
	Sql string
}
