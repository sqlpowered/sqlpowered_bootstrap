package sql_builder

// Note we differentiate between left, right, full via the name in the JSON, the underlying data type is the same
//
//	input field name, not the type
type Join struct {
	Column1  string `json:"column1"`
	Table1   string `json:"table1,omitempty"`
	Operator string `json:"operator,omitempty"`
	Column2  string `json:"column2"`
	Table2   string `json:"table2,omitempty"`
}

type JoinOutput struct {
	Sql string
}
