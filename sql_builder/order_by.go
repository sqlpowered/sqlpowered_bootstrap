package sql_builder

type OrderBy struct {
	Table  string `json:"table"`
	Column string `json:"column"`
	Order  string `json:"order"`
}
