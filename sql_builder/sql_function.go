package sql_builder

type Fn struct {
	Fn string `json:"fn"`
	// Args []FnArgs `json:"args"`
	Args []Select `json:"args"` // "as" is not valid in this Select
	Type string   `json:"type"`
}

type FnOutput struct {
	Sql string
}
