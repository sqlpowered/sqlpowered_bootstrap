package sql_builder

type From struct {
	Tables []string `json:"tables,omitempty"`
}

type FromOutput struct {
	Sql string
}
