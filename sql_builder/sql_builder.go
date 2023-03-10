package sql_builder

// A clear separation between input and output data
type SqlInput struct {
	Select    []Select
	From      []From
	InnerJoin []Join
	LeftJoin  []Join
	RightJoin []Join
	OuterJoin []Join
	Where     []Where
	GroupBy   []GroupBy
	Having    []Where
}

// type SqlOutput struct {
// 	Select []SelectOutput
// 	From   []FromOutput
// 	Join   []JoinOutput
// 	Where  []WhereOutput
// }
