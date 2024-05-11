package employee_manager_test

import (
	"math"
	"testing"
	"time"

	employee_manager "gitlab.com/zigal0-group/nica/tango-api/internal/business/manager/employee"
	"gitlab.com/zigal0-group/nica/tango-api/internal/domain"
	"syreclabs.com/go/faker"
)

func Test_GetEmployeeByID(t *testing.T) {
	t.Parallel()

	var (
		id = faker.RandomInt64(1, math.MaxInt64)
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		t.Run("full flow", func(t *testing.T) {
			t.Parallel()

			// arrange
			f := setUp(t)

			employee := domain.Employee{
				ID:        id,
				Name:      faker.RandomString(10),
				Age:       int32(faker.RandomInt64(18, 100)),
				CreatedAt: time.Now(),
			}

			f.emplyeeRepo.GetByIDMock.
				Expect(f.ctx, id).
				Return(employee, nil)

			// act
			res, err := f.testEntity.GetEmployeeByID(f.ctx, id)

			// assert
			f.NoError(err)
			f.Equal(employee, res)
		})
	})

	t.Run("fail", func(t *testing.T) {
		t.Parallel()

		t.Run("negative id", func(t *testing.T) {
			t.Parallel()

			// arrange
			f := setUp(t)

			// act
			res, err := f.testEntity.GetEmployeeByID(f.ctx, -1)

			// assert
			f.ErrorIs(err, employee_manager.ErrNotPositiveEmployeeID)
			f.Empty(res)
		})

		t.Run("full flow", func(t *testing.T) {
			t.Parallel()

			// arrange
			f := setUp(t)

			f.emplyeeRepo.GetByIDMock.
				Expect(f.ctx, id).
				Return(domain.Employee{}, errForTest)

			// act
			res, err := f.testEntity.GetEmployeeByID(f.ctx, id)

			// assert
			f.ErrorIs(err, errForTest)
			f.Empty(res)
		})
	})
}
