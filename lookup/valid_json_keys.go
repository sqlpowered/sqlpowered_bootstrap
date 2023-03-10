package lookup

/**

This file aims to define valid and minimum required statement components for each of the 4 main statement types:
insert, update, delete, select

**/

// =================================
func ValidSelectKeys() []string {
	return []string{
		"select",
		"from",
		"join",
		"where",
		"group_by",
		"having",
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
