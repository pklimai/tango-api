package employee_repository

import (
	"time"

	"gitlab.com/zigal0-group/nica/tango-api/internal/domain"
)

type Employee struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Age       int32     `db:"age"`
	CreatedAt time.Time `db:"created_at"`
}

func toDomain(e Employee) domain.Employee {
	return domain.Employee{
		ID:        e.ID,
		Name:      e.Name,
		Age:       e.Age,
		CreatedAt: e.CreatedAt,
	}
}
