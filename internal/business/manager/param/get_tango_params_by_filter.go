package param_manager

import (
	"context"
	"fmt"

	"gitlab.com/zigal0-group/nica/tango-api/internal/domain"
)

func (m *Manager) GetTangoParamByFilter(
	ctx context.Context,
	filter domain.ParamFilter,
) (domain.Params, error) {
	res, err := m.paramRepo.GetParamByFilter(ctx, filter)
	if err != nil {
		return domain.Params{}, fmt.Errorf("paramRepo.GetParamByFilter: %w", err)
	}

	return res, nil
}
