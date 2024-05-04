package employee_manager

import (
	"context"
	"fmt"

	"gitlab.com/zigal0-group/nica/tango-api/internal/domain"
	"gitlab.com/zigal0/architect/pkg/business_error"
)

const (
	lowerAgeLimit = 18
	upperAgeLimit = 100
)

func (m *Manager) AddEmployee(ctx context.Context, name string, age int32) (int64, error) {
	if len(name) == 0 {
		return 0, business_error.New(
			ErrEmptyName,
			"имя сотрудника не может быть пустым",
			business_error.InvalidArgument,
		)
	}

	if age < lowerAgeLimit || upperAgeLimit > 100 {
		return 0, business_error.New(
			fmt.Errorf(FormatErrAgeOutOfRange, lowerAgeLimit, upperAgeLimit),
			fmt.Sprintf(FormatMsgAgeOutOfRange, lowerAgeLimit, upperAgeLimit),
			business_error.InvalidArgument,
		)
	}

	id, err := m.employeeRepo.Insert(ctx, domain.Employee{Name: name, Age: age})
	if err != nil {
		return 0, fmt.Errorf("employeeRepo.Insert: %w", err)
	}

	return id, nil
}
