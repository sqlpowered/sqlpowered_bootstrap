package sql_builder

// A clear separation between input and output data
type SqlQuery struct {
	Cte       []Cte     `json:"cte,omitempty"`
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

// type SqlOutput struct {
// 	Select []SelectOutput
// 	From   []FromOutput
// 	Join   []JoinOutput
// 	Where  []WhereOutput
// }
