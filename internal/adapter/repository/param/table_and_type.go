package param_repository

const (
	TableNameAttArrayDevdouble = "att_array_devdouble"
	TableNameAttArrayBoolean   = "att_array_devboolean"

	TableNameAttScalarDevdouble = "att_scalar_devdouble"
	TableNameAttScalarBoolean   = "att_scalar_devboolean"
)

var scalarTables = []string{
	TableNameAttScalarDevdouble,
	TableNameAttScalarBoolean,
}

var arrayTables = []string{
	TableNameAttArrayDevdouble,
	TableNameAttArrayBoolean,
}

var allTables = append(scalarTables, arrayTables...)
