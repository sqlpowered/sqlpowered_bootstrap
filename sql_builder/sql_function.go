package sql_builder

type Case struct {
	CaseArgs       []CaseArgs `json:"case_args"`
	ElseExpression Select     `json:"else_expression"`
}

type CaseArgs struct {
	Conditions []Where `json:"conditions"`
	Expression Select  `json:"expression"`
}

type SqlFunction struct {
	Function     string         `json:"function"`
	FunctionArgs []FunctionArgs `json:"function_args"`
	TypeCast     string         `json:"type_cast"`
}

// TODO: think about how we handle JSON data accessors
type FunctionArgs struct {
	ColumnName   string        `json:"column_name"`
	TableName    string        `json:"table_name"`
	ValuesList   []string      `json:"values_list"`
	FunctionList []SqlFunction `json:"function_list"`
	TypeCast     string        `json:"type_cast"`
}

type SqlFunctionOutput struct {
	Sql string
}
