package employee_manager

import (
	"context"
	"fmt"

	"gitlab.com/zigal0-group/nica/tango-api/internal/domain"
	"gitlab.com/zigal0/architect/pkg/business_error"
)

func (m *Manager) GetEmployeeByID(ctx context.Context, id int64) (domain.Employee, error) {
	if id < 0 {
		return domain.Employee{}, business_error.New(
			ErrNotPositiveEmployeeID,
			"id работника может быть только положительным числом",
			business_error.InvalidArgument,
		)
	}

	employee, err := m.employeeRepo.GetByID(ctx, id)
	if err != nil {
		return domain.Employee{}, fmt.Errorf("employeeRepo.GetByID: %w", err)
	}

	return employee, nil
}
