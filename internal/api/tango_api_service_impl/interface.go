package tango_api_service_impl

import (
	"context"

	"gitlab.com/zigal0-group/nica/tango-api/internal/domain"
)

type employeeManager interface {
	GetEmployeeByID(ctx context.Context, id int64) (domain.Employee, error)
}
