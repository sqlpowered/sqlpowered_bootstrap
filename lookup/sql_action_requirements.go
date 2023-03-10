package lookup

/**

This file aims to define the specifications of what is allowed for each of the 4 main statement types:
insert, update, delete, select

**/

type SqlActionRequirements struct {
	ValidFields   []string
	MinimumFields []string
	FieldsSpec    []FieldSpec
}

type FieldSpec struct {
	FieldName         string
	MinLength         int
	MaxLength         int
	MaxRecursionDepth int
	ComponentSpec     *FieldSpec
}

func DeleteRequirements(maxNumWhereConditions int) SqlActionRequirements {

	// TODO: Nested specifications
	deleteFromSpec := FieldSpec{}
	whereSpec := FieldSpec{}

	requirements := SqlActionRequirements{
		ValidFields:   ValidDeleteKeys(),
		MinimumFields: MinimumDeleteKeys(),
		FieldsSpec: []FieldSpec{
			{
				FieldName:         "delete_from",
				MinLength:         1,
				MaxLength:         1,
				MaxRecursionDepth: 1,
				ComponentSpec:     &deleteFromSpec,
			},
			{
				FieldName:         "where",
				MinLength:         1,
				MaxLength:         maxNumWhereConditions,
				MaxRecursionDepth: 1,
				ComponentSpec:     &whereSpec,
			},
		},
	}

	return requirements
}

func UpdateRequirements(
	maxNumColumns int,
	maxNumWhereConditions int,
) SqlActionRequirements {

	// TODO: Nested specifications
	updateSpec := FieldSpec{}
	setSpec := FieldSpec{}
	whereSpec := FieldSpec{}

	requirements := SqlActionRequirements{
		ValidFields:   ValidUpdateKeys(),
		MinimumFields: MinimumUpdateKeys(),
		FieldsSpec: []FieldSpec{
			{
				FieldName:         "update",
				MinLength:         1,
				MaxLength:         1,
				MaxRecursionDepth: 1,
				ComponentSpec:     &updateSpec,
			},
			{
				FieldName:         "set",
				MinLength:         1,
				MaxLength:         maxNumColumns,
				MaxRecursionDepth: 1,
				ComponentSpec:     &setSpec,
			},
			{
				FieldName:         "where",
				MinLength:         1,
				MaxLength:         maxNumWhereConditions,
				MaxRecursionDepth: 1,
				ComponentSpec:     &whereSpec,
			},
		},
	}

	return requirements
}

func InsertRequirements(
	maxNumColumns int,
	maxNumInsertRows int,
) SqlActionRequirements {

	// TODO: Nested specifications
	insertIntoSpec := FieldSpec{}
	valuesSpec := FieldSpec{}

	requirements := SqlActionRequirements{
		ValidFields:   ValidInsertKeys(),
		MinimumFields: MinimumInsertKeys(),
		FieldsSpec: []FieldSpec{
			{
				FieldName:         "insert_into",
				MinLength:         1,
				MaxLength:         1,
				MaxRecursionDepth: 1,
				ComponentSpec:     &insertIntoSpec,
			},
			{
				FieldName:         "values",
				MinLength:         1,
				MaxLength:         maxNumInsertRows,
				MaxRecursionDepth: 1,
				ComponentSpec:     &valuesSpec,
			},
		},
	}

	return requirements
}

func SelectRequirements(
	maxNumColumns int,
	maxRecursionDepth int,
	maxNumFromTables int,
	maxNumJoinTables int,
	maxNumWhereConditions int,
	maxNumGroupBy int,
) SqlActionRequirements {

	// TODO: Nested specifications
	selectSpec := FieldSpec{}
	fromSpec := FieldSpec{}
	joinSpec := FieldSpec{}
	whereSpec := FieldSpec{}
	groupBySpec := FieldSpec{}
	havingSpec := FieldSpec{}

	requirements := SqlActionRequirements{
		ValidFields:   ValidSelectKeys(),
		MinimumFields: MinimumSelectKeys(),
		FieldsSpec: []FieldSpec{
			{
				FieldName:         "select",
				MinLength:         1,
				MaxLength:         1,
				MaxRecursionDepth: maxRecursionDepth,
				ComponentSpec:     &selectSpec,
			},
			{
				FieldName:         "from",
				MinLength:         1,
				MaxLength:         maxNumFromTables,
				MaxRecursionDepth: 1,
				ComponentSpec:     &fromSpec,
			},
			{
				FieldName:         "join",
				MinLength:         1,
				MaxLength:         maxNumJoinTables,
				MaxRecursionDepth: 1,
				ComponentSpec:     &joinSpec,
			},
			{
				FieldName:         "where",
				MinLength:         1,
				MaxLength:         maxNumWhereConditions,
				MaxRecursionDepth: 1,
				ComponentSpec:     &whereSpec,
			},
			{
				FieldName:         "group_by",
				MinLength:         1,
				MaxLength:         maxNumGroupBy,
				MaxRecursionDepth: 1,
				ComponentSpec:     &groupBySpec,
			},
			{
				FieldName:         "having",
				MinLength:         1,
				MaxLength:         maxNumWhereConditions,
				MaxRecursionDepth: 1,
				ComponentSpec:     &havingSpec,
			},
		},
	}

	return requirements
}
