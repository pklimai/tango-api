package domain

import "time"

type ParamFilter struct {
	Domain   string
	Name     string
	Member   string
	TimeFrom time.Time
	TimeTo   time.Time
}
