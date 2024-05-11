package employee_manager_test

import (
	"context"
	"errors"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	employee_manager "gitlab.com/zigal0-group/nica/tango-api/internal/business/manager/employee"
	"gitlab.com/zigal0-group/nica/tango-api/internal/business/manager/employee/mocks"
)

var (
	errForTest = errors.New("error for test")
)

type fixture struct {
	ctx context.Context
	*assert.Assertions

	emplyeeRepo *mocks.EmployeeRepoMock

	testEntity *employee_manager.Manager
}

func setUp(t *testing.T) (f *fixture) {
	t.Helper()

	ctrl := minimock.NewController(t)

	emplyeeRepo := mocks.NewEmployeeRepoMock(ctrl)

	testEntity := employee_manager.New(emplyeeRepo)

	return &fixture{
		ctx:        context.Background(),
		Assertions: assert.New(ctrl),

		emplyeeRepo: emplyeeRepo,

		testEntity: testEntity,
	}
}
