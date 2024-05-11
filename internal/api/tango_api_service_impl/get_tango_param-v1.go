package tango_api_service_impl

import (
	"context"

	pb "gitlab.com/zigal0-group/nica/tango-api/internal/generated/api/tango_api_service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) GetTangoParamV1(
	_ context.Context,
	_ *pb.GetTangoParamsV1Request,
) (*pb.GetTangoParamsV1Response, error) {
	return &pb.GetTangoParamsV1Response{
		ParamType: pb.GetTangoParamsV1Response_FLOAT64,
		ScalarParams: []*pb.GetTangoParamsV1Response_ScalarParam{
			{
				RawValueR: "1.1",
				RawValueW: "2.2",
				DataTime:  timestamppb.Now(),
			},
			{
				RawValueR: "3.3",
				RawValueW: "4.4",
				DataTime:  timestamppb.Now(),
			},
		},
		ArrayParmas: []*pb.GetTangoParamsV1Response_ArrayParam{
			{
				RawValuesR: []string{"1.1", "1.1"},
				RawValuesW: []string{"2.2", "2.2"},
				DataTime:   timestamppb.Now(),
			},
			{
				RawValuesR: []string{"3.3", "3.3"},
				RawValuesW: []string{"4.4", "4.4"},
				DataTime:   timestamppb.Now(),
			},
		},
	}, nil
}
