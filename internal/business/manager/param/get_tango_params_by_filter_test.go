package param_manager_test

import (
	"reflect"
	"testing"
	"time"

	"gitlab.com/zigal0-group/nica/tango-api/internal/domain"
	"syreclabs.com/go/faker"
)

func Test_GetTangoParamByFilter(t *testing.T) {
	t.Parallel()

	filter := domain.ParamFilter{
		Domain:   faker.RandomString(10),
		Name:     faker.RandomString(10),
		Member:   faker.RandomString(10),
		TimeFrom: time.Now(),
		TimeTo:   time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		// arrange
		f := setUp(t)

		params := domain.Params{
			Type: reflect.Bool,
			ScalarParams: []domain.ScalarParam{
				{
					ValueR:   "true",
					ValueW:   "true",
					DataTime: time.Now(),
				},
				{
					ValueR:   "false",
					ValueW:   "false",
					DataTime: time.Now(),
				},
			},
		}

		f.paramRepo.GetParamByFilterMock.
			Expect(f.ctx, filter).
			Return(params, nil)

		// act
		res, err := f.testEntity.GetTangoParamByFilter(f.ctx, filter)

		// assert
		f.NoError(err)
		f.Equal(params, res)
	})

	t.Run("fail", func(t *testing.T) {
		t.Parallel()

		// arrange
		f := setUp(t)

		f.paramRepo.GetParamByFilterMock.
			Expect(f.ctx, filter).
			Return(domain.Params{}, errForTest)

		// act
		res, err := f.testEntity.GetTangoParamByFilter(f.ctx, filter)

		// assert
		f.ErrorIs(err, errForTest)
		f.Empty(res)
	})
}
