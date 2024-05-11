package tango_api_service_impl

import (
	"context"
	"fmt"

	tool_collection "gitlab.com/zigal0-group/nica/tango-api/internal/business/tool/collection"
	"gitlab.com/zigal0-group/nica/tango-api/internal/domain"
	pb "gitlab.com/zigal0-group/nica/tango-api/internal/generated/api/tango_api_service"
)

func (s *Service) GetTangoParamsV1(
	ctx context.Context,
	req *pb.GetTangoParamsV1Request,
) (*pb.GetTangoParamsV1Response, error) {
	filterPB := req.GetFilter()

	params, err := s.paramManager.GetTangoParamByFilter(ctx, domain.ParamFilter{
		Domain:   filterPB.GetDomain(),
		Name:     filterPB.GetName(),
		Member:   filterPB.GetMember(),
		TimeFrom: filterPB.TimeFrom.AsTime(),
		TimeTo:   filterPB.TimeTo.AsTime(),
	})
	if err != nil {
		return nil, fmt.Errorf("paramManager.GetTangoParamByFilter: %w", err)
	}

	return &pb.GetTangoParamsV1Response{
		ParamType:    paramTypeToPB(params.Type),
		ScalarParams: tool_collection.Map(params.ScalarParams, scalarParamToPB),
		ArrayParmas:  tool_collection.Map(params.ArrayParams, arrayParamToPB),
	}, nil
}
