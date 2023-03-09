package sql_builder

type Where struct {
	BooleanOperator string `json:"boolean_operator,omitempty"`
	Left            Select `json:"left"`
	Operator        string `json:"operator,omitempty"`
	Right           Select `json:"right"`
}

type WhereArgs struct {
	ColumnName string   `json:"column_name,omitempty"`
	TableName  string   `json:"table_name,omitempty"`
	ValuesList []string `json:"values_list,omitempty"`
}

type WhereOutput struct {
	Sql string
}
