package sql_builder

type Where struct {
	Cond  string   `json:"cond,omitempty"`
	Left  WhereArg `json:"left"`
	Op    string   `json:"operator,omitempty"`
	Right WhereArg `json:"right"`
}

type WhereArg struct {
	Column string   `json:"column,omitempty"`
	Table  string   `json:"table,omitempty"`
	Values []string `json:"values,omitempty"`
	Fns    []Fn     `json:"fns,omitempty"`
	Case   Case     `json:"case,omitempty"`
}

type WhereOutput struct {
	Sql string
}
