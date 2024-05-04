package employee_manager

import (
	"context"

	"gitlab.com/zigal0-group/nica/tango-api/internal/domain"
)

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i * -o ./mocks -s "_mock.go" -g

type employeeRepo interface {
	GetByID(ctx context.Context, id int64) (domain.Employee, error)
	Insert(ctx context.Context, employee domain.Employee) (int64, error)
}
