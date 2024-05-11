package domain

import (
	"reflect"
	"time"
)

type Params struct {
	Type         reflect.Kind
	ArrayParams  []ArrayParam
	ScalarParams []ScalarParam
}

type ArrayParam struct {
	ValuesR  []string
	ValuesW  []string
	DataTime time.Time
}

type ScalarParam struct {
	ValueR   string
	ValueW   string
	DataTime time.Time
}
