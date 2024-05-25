package tango_api_service_impl

import (
	"context"
	"errors"
	"fmt"
	"time"

	tool_collection "gitlab.com/zigal0-group/nica/tango-api/internal/business/tool/collection"
	"gitlab.com/zigal0-group/nica/tango-api/internal/domain"
	pb "gitlab.com/zigal0-group/nica/tango-api/internal/generated/api/tango_api_service"
	"gitlab.com/zigal0/architect/pkg/business_error"
)

func (s *Service) GetTangoParamsV1(
	ctx context.Context,
	req *pb.GetTangoParamsV1Request,
) (*pb.GetTangoParamsV1Response, error) {
	timeFrom, err := time.Parse(time.DateOnly, req.GetStartTime())
	if err != nil {
		return nil, business_error.New(
			err,
			fmt.Sprintf("invalid start_time: %s", req.GetStartTime()),
			business_error.InvalidArgument,
		)
	}

	var timeTo time.Time

	if req.GetEndTime() != "" {
		timeTo, err = time.Parse(time.DateOnly, req.GetEndTime())
		if err != nil {
			return nil, business_error.New(
				err,
				fmt.Sprintf("invalid end_time: %s", req.GetStartTime()),
				business_error.InvalidArgument,
			)
		}

		if timeTo.Before(timeFrom) {
			err = errors.New("strart_time > end_time")

			return nil, business_error.New(
				err,
				err.Error(),
				business_error.InvalidArgument,
			)
		}
	}

	params, err := s.paramManager.GetTangoParamByFilter(ctx, domain.ParamFilter{
		Domain:   req.GetSystemName(),
		Name:     req.GetParameterName(),
		Member:   req.GetMemberName(),
		TimeFrom: timeFrom,
		TimeTo:   timeTo,
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
