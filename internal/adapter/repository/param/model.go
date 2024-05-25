package param_repository

import (
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
