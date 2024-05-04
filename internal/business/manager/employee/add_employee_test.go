package employee_manager_test

import (
	"fmt"
	"math"
	"testing"

	employee_manager "gitlab.com/zigal0-group/nica/tango-api/internal/business/manager/employee"
	"gitlab.com/zigal0-group/nica/tango-api/internal/domain"
	"syreclabs.com/go/faker"
)

func Test_AddEmployee(t *testing.T) {
	t.Parallel()

	var (
		name = faker.RandomString(10)
		age  = int32(faker.RandomInt64(18, 100))
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		t.Run("full flow", func(t *testing.T) {
			t.Parallel()

			// arrange
			f := setUp(t)

			var (
				id = faker.RandomInt64(1, math.MaxInt64)
			)

			f.emplyeeRepo.InsertMock.
				Expect(f.ctx, domain.Employee{Name: name, Age: age}).
				Return(id, nil)

			// act
			res, err := f.testEntity.AddEmployee(f.ctx, name, age)

			// assert
			f.NoError(err)
			f.Equal(id, res)
		})
	})

	t.Run("fail", func(t *testing.T) {
		t.Parallel()

		t.Run("empty name", func(t *testing.T) {
			t.Parallel()

			// arrange
			f := setUp(t)

			// act
			res, err := f.testEntity.AddEmployee(f.ctx, "", age)

			// assert
			f.ErrorIs(err, employee_manager.ErrEmptyName)
			f.Empty(res)
		})

		t.Run("age out of range", func(t *testing.T) {
			t.Parallel()

			// arrange
			f := setUp(t)

			expectedErr := fmt.Errorf(employee_manager.FormatErrAgeOutOfRange, 18, 100)

			// act
			res, err := f.testEntity.AddEmployee(f.ctx, name, 0)

			// assert
			f.EqualError(err, expectedErr.Error())
			f.Empty(res)
		})

		t.Run("employeeRepo.Insert", func(t *testing.T) {
			t.Parallel()

			// arrange
			f := setUp(t)

			f.emplyeeRepo.InsertMock.
				Expect(f.ctx, domain.Employee{Name: name, Age: age}).
				Return(0, errForTest)

			// act
			res, err := f.testEntity.AddEmployee(f.ctx, name, age)

			// assert
			f.ErrorIs(err, errForTest)
			f.Empty(res)
		})
	})
}
