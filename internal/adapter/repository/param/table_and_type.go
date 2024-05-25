package param_repository

import "reflect"

const (
	TableNameAttArrayDevushort  = "att_array_devushort"
	TableNameAttArrayDevulong64 = "att_array_devulong64"
	TableNameAttArrayDevulong   = "att_array_devulong"
	TableNameAttArrayDevuchar   = "att_array_devuchar"
	TableNameAttArrayDevstring  = "att_array_devstring"
	TableNameAttArrayDevstate   = "att_array_devstate"
	TableNameAttArrayDevshort   = "att_array_devshort"
	TableNameAttArrayDevlong64  = "att_array_devlong64"
	TableNameAttArrayDevlong    = "att_array_devlong"
	TableNameAttArrayDevfloat   = "att_array_devfloat"
	TableNameAttArrayDevenum    = "att_array_devenum"    // additional fields with labels
	TableNameAttArrayDevencoded = "att_array_devencoded" // bytes, can be painfull to convert to string
	TableNameAttArrayDevdouble  = "att_array_devdouble"
	TableNameAttArrayDevboolean = "att_array_devboolean"

	TableNameAttScalarDevushort  = "att_scalar_devushort"
	TableNameAttScalarDevulong64 = "att_scalar_devulong64"
	TableNameAttScalarDevulong   = "att_scalar_devulong"
	TableNameAttScalarDevuchar   = "att_scalar_devuchar"
	TableNameAttScalarDevfloat   = "att_scalar_devfloat"
	TableNameAttScalarDevenum    = "att_scalar_devenum"    // additional fields with labels
	TableNameAttScalarDevencoded = "att_scalar_devencoded" // bytes, can be painfull to convert to string
	TableNameAttScalarDevdouble  = "att_scalar_devdouble"
	TableNameAttScalarDevboolean = "att_scalar_devboolean"
	TableNameAttScalarDevstate   = "att_scalar_devstate"
	TableNameAttScalarDevstring  = "att_scalar_devstring"
	TableNameAttScalarDevshort   = "att_scalar_devshort"
	TableNameAttScalarDevlong64  = "att_scalar_devlong64"
	TableNameAttScalarDevlong    = "att_scalar_devlong"
)

var scalarTables = []string{
	TableNameAttScalarDevushort,
	TableNameAttScalarDevulong64,
	TableNameAttScalarDevulong,
	TableNameAttScalarDevuchar,
	TableNameAttScalarDevfloat,
	// TableNameAttScalarDevenum,
	// TableNameAttScalarDevencoded,
	TableNameAttScalarDevdouble,
	TableNameAttScalarDevboolean,
	TableNameAttScalarDevstate,
	TableNameAttScalarDevstring,
	TableNameAttScalarDevshort,
	TableNameAttScalarDevlong64,
	TableNameAttScalarDevlong,
}

var arrayTables = []string{
	TableNameAttArrayDevushort,
	TableNameAttArrayDevulong64,
	TableNameAttArrayDevulong,
	TableNameAttArrayDevuchar,
	TableNameAttArrayDevstring,
	TableNameAttArrayDevstate,
	TableNameAttArrayDevshort,
	TableNameAttArrayDevlong64,
	TableNameAttArrayDevlong,
	TableNameAttArrayDevfloat,
	// TableNameAttArrayDevenum,
	// TableNameAttArrayDevencoded,
	TableNameAttArrayDevdouble,
	TableNameAttArrayDevboolean,
}

var allTables = append(scalarTables, arrayTables...)

// nolint: gomnd
func paramTypeToKind(attConfTypeID int16) reflect.Kind {
	switch attConfTypeID {
	case 1:
		return reflect.Bool
	case 2:
		return reflect.Int16
	case 3:
		return reflect.Int32
	case 4:
		return reflect.Float32
	case 5:
		return reflect.Float64
	case 6:
		return reflect.Uint16
	case 7:
		return reflect.Uint32
	case 8:
		return reflect.String
	case 9:
		return reflect.Int32
	case 10:
		return reflect.Uint8
	case 11:
		return reflect.Int64
	case 12:
		return reflect.Uint64
	case 13: // If bytes willbe required - need to refuse from reflect types or find workaround...
		return reflect.Invalid
	case 14:
		return reflect.Int16
	default:
		return reflect.Invalid
	}
}
