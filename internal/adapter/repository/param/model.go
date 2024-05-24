package param_repository

import (
	"reflect"
	"time"

	"github.com/lib/pq"
	"gitlab.com/zigal0-group/nica/tango-api/internal/domain"
)

type SearchOptions struct {
	AttConfID     int32  `db:"att_conf_id"`
	AttConfTypeID int16  `db:"att_conf_type_id"`
	TableName     string `db:"table_name"`
}

type ScalarParam struct {
	ValueR   string    `db:"value_r"`
	ValueW   string    `db:"value_w"`
	DataTime time.Time `db:"data_time"`
}

func scalarParamToDomain(param ScalarParam) domain.ScalarParam {
	return domain.ScalarParam{
		ValueR:   param.ValueR,
		ValueW:   param.ValueW,
		DataTime: param.DataTime,
	}
}

type ArrayParam struct {
	RawValuesR pq.StringArray `db:"value_r"`
	RawValuesW pq.StringArray `db:"value_w"`
	DataTime   time.Time      `db:"data_time"`
}

func arrayParamToDomain(param ArrayParam) domain.ArrayParam {
	return domain.ArrayParam{
		ValuesR:  param.RawValuesR,
		ValuesW:  param.RawValuesW,
		DataTime: param.DataTime,
	}
}

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
