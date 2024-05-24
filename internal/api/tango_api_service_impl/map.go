package tango_api_service_impl

import (
	"reflect"

	"gitlab.com/zigal0-group/nica/tango-api/internal/domain"
	pb "gitlab.com/zigal0-group/nica/tango-api/internal/generated/api/tango_api_service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func paramTypeToPB(domainType reflect.Kind) pb.GetTangoParamsV1Response_ParamType {
	switch domainType {
	case reflect.Float32:
		return pb.GetTangoParamsV1Response_FLOAT32
	case reflect.Float64:
		return pb.GetTangoParamsV1Response_FLOAT64
	case reflect.Bool:
		return pb.GetTangoParamsV1Response_BOOL
	case reflect.String:
		return pb.GetTangoParamsV1Response_STRING
	case reflect.Int8:
		return pb.GetTangoParamsV1Response_INT8
	case reflect.Int16:
		return pb.GetTangoParamsV1Response_INT16
	case reflect.Int32:
		return pb.GetTangoParamsV1Response_INT32
	case reflect.Int64:
		return pb.GetTangoParamsV1Response_UINT64
	case reflect.Uint8:
		return pb.GetTangoParamsV1Response_UINT8
	case reflect.Uint16:
		return pb.GetTangoParamsV1Response_UINT16
	case reflect.Uint32:
		return pb.GetTangoParamsV1Response_UINT32
	case reflect.Uint64:
		return pb.GetTangoParamsV1Response_UINT64
	default:
		return pb.GetTangoParamsV1Response_INVALID
	}
}

func scalarParamToPB(param domain.ScalarParam) *pb.GetTangoParamsV1Response_ScalarParam {
	return &pb.GetTangoParamsV1Response_ScalarParam{
		RawValueR: param.ValueR,
		RawValueW: param.ValueW,
		DataTime:  timestamppb.New(param.DataTime),
	}
}

func arrayParamToPB(param domain.ArrayParam) *pb.GetTangoParamsV1Response_ArrayParam {
	return &pb.GetTangoParamsV1Response_ArrayParam{
		RawValuesR: param.ValuesR,
		RawValuesW: param.ValuesW,
		DataTime:   timestamppb.New(param.DataTime),
	}
}
