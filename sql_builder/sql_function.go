package sql_builder

type Fn struct {
	Fn   string   `json:"fn"`
	Args []FnArgs `json:"args"`
	Type string   `json:"type"`
}

// TODO: think about how we handle JSON data accessors
type FnArgs struct {
	Table  string   `json:"table"`
	Column string   `json:"column"`
	Values []string `json:"values"`
	Fns    []Fn     `json:"fns"`
	Type   string   `json:"type"`
	Case   Case     `json:"case"`
}

type FnOutput struct {
	Sql string
}
