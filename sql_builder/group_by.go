package sql_builder

type GroupBy struct {
	ColumnName string `json:"column_name"`
	TableName  string `json:"table_name,omitempty"`
}

type GroupByOutput struct {
	Sql string
}
