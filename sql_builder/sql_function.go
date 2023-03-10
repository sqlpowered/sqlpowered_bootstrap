package sql_builder

type SqlFunction struct {
	Function string         `json:"function"`
	Args     []FunctionArgs `json:"args"`
	Type     string         `json:"type"`
}

// TODO: think about how we handle JSON data accessors
type FunctionArgs struct {
	Table     string        `json:"table"`
	Column    string        `json:"column"`
	Values    []string      `json:"values"`
	Functions []SqlFunction `json:"functions"`
	Type      string        `json:"type"`
	Case      Case          `json:"case"`
}

type SqlFunctionOutput struct {
	Sql string
}
