package param_manager_test

import (
	"context"
	"errors"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	param_manager "gitlab.com/zigal0-group/nica/tango-api/internal/business/manager/param"
	"gitlab.com/zigal0-group/nica/tango-api/internal/business/manager/param/mocks"
)

var (
	errForTest = errors.New("error for test")
)

type fixture struct {
	ctx context.Context
	*assert.Assertions

	paramRepo *mocks.ParamRepoMock

	testEntity *param_manager.Manager
}

func setUp(t *testing.T) (f *fixture) {
	t.Helper()

	ctrl := minimock.NewController(t)

	paramRepo := mocks.NewParamRepoMock(ctrl)

	testEntity := param_manager.New(paramRepo)

	return &fixture{
		ctx:        context.Background(),
		Assertions: assert.New(ctrl),

		paramRepo: paramRepo,

		testEntity: testEntity,
	}
}
