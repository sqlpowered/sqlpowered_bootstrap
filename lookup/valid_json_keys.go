package lookup

// =================================
func ValidSelectKeys() []string {
	return []string{
		"select",
		"from",
		"join",
		"where",
		"group_by",
	}
}
func MinimumSelectKeys() []string {
	return []string{
		"select",
		"from",
	}
}

// =================================
func ValidInsertKeys() []string {
	return []string{
		"insert_into",
		"values",
	}
}
func MinimumInsertKeys() []string {
	return []string{
		"insert_into",
		"values",
	}
}

// =================================
func ValidUpdateKeys() []string {
	return []string{
		"update",
		"set",
		"where",
	}
}
func MinimumUpdateKeys() []string {
	return []string{
		"update",
		"set",
		"where",
	}
}

// =================================
func ValidDeleteKeys() []string {
	return []string{
		"delete_from",
		"where",
	}
}
func MinimumDeleteKeys() []string {
	return []string{
		"delete_from",
		"where",
	}
}
