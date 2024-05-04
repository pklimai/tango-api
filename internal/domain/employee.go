package domain

import "time"

type Employee struct {
	ID        int64
	Name      string
	Age       int32
	CreatedAt time.Time
}
