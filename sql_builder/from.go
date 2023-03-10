package sql_builder

type From struct {
	Tables []string
}

type FromOutput struct {
	Sql string
}
