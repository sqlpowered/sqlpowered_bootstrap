package sql_builder

type GroupBy struct {
	Column string `json:"column"`
	Table  string `json:"table"`
}

type GroupByOutput struct {
	Sql string
}
