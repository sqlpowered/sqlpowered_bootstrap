package sql_builder

type Cte struct {
	Sql SqlQuery `json:"sql"`
	As  string   `json:"as"`
}

type SqlCte struct {
	Select    []Select  `json:"select,omitempty"`
	From      []string  `json:"from,omitempty"`
	InnerJoin []Join    `json:"inner_join,omitempty"`
	LeftJoin  []Join    `json:"left_join,omitempty"`
	RightJoin []Join    `json:"right_join,omitempty"`
	OuterJoin []Join    `json:"outer_join,omitempty"`
	Where     []Where   `json:"where,omitempty"`
	GroupBy   []GroupBy `json:"group_by,omitempty"`
	Having    []Where   `json:"having,omitempty"`
	OrderBy   []OrderBy `json:"order_by,omitempty"`
}
