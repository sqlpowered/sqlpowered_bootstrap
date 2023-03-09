package sql_builder

// Note we differentiate between left, right, full via the
//
//	input field name, not the type
type Join struct {
	Column1Name string `json:"column1_name"`
	Table1Name  string `json:"table1_name,omitempty"`
	Operator    string `json:"operator,omitempty"`
	Column2Name string `json:"column2_name"`
	Table2Name  string `json:"table2_name,omitempty"`
}

type JoinOutput struct {
	Sql string
}
