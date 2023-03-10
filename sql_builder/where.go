package sql_builder

type Where struct {
	BooleanOperator string   `json:"boolean_operator,omitempty"`
	Left            WhereArg `json:"left"`
	Operator        string   `json:"operator,omitempty"`
	Right           WhereArg `json:"right"`
}

type WhereArg struct {
	Column string   `json:"column,omitempty"`
	Table  string   `json:"table,omitempty"`
	Values []string `json:"values,omitempty"`
}

type WhereOutput struct {
	Sql string
}
