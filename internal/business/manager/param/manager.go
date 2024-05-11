package param_manager

import (
	"context"

	"gitlab.com/zigal0-group/nica/tango-api/internal/domain"
)

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i * -o ./mocks -s "_mock.go" -g

type Manager struct {
	paramRepo paramRepo
}

func New(
	paramRepo paramRepo,
) *Manager {
	return &Manager{
		paramRepo: paramRepo,
	}
}

type paramRepo interface {
	GetParamByFilter(
		ctx context.Context,
		filter domain.ParamFilter,
	) (domain.Params, error)
}
