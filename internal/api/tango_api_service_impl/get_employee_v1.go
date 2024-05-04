package tango_api_service_impl

import (
	"context"
	"fmt"

	pb "gitlab.com/zigal0-group/nica/tango-api/internal/generated/api/tango_api_service"
)

func (s *Service) GetEmployeeV1(
	ctx context.Context,
	req *pb.GetEmployeeV1Request,
) (*pb.GetEmployeeV1Response, error) {
	employee, err := s.employeeManager.GetEmployeeByID(ctx, req.GetEmployeeId())
	if err != nil {
		return nil, fmt.Errorf("employeeManager.GetEmployeeByID: %w", err)
	}

	return &pb.GetEmployeeV1Response{
		Employee: &pb.Employee{
			Id:   employee.ID,
			Name: employee.Name,
		},
	}, nil
}
