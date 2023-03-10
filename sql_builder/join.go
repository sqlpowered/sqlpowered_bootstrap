package sql_builder

// Note we differentiate between left, right, full
// via the name in the JSON, the underlying data type is the same

type Join struct {
	Arg1 JoinArg `json:"arg1"`
	Op   string  `json:"op,omitempty"`
	Arg2 JoinArg `json:"arg2"`
}

type JoinArg struct {
	Column string `json:"column"`
	Table  string `json:"table"`
}

type JoinOutput struct {
	Sql string
}
