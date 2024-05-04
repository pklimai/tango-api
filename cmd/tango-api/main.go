package main

import (
	"context"

	"gitlab.com/zigal0-group/nica/tango-api/config"
	employee_repository "gitlab.com/zigal0-group/nica/tango-api/internal/adapter/repository/employee"
	"gitlab.com/zigal0-group/nica/tango-api/internal/api/tango_api_service_impl"
	employee_manager "gitlab.com/zigal0-group/nica/tango-api/internal/business/manager/employee"
	"gitlab.com/zigal0-group/nica/tango-api/internal/database"
	"gitlab.com/zigal0-group/nica/tango-api/internal/generated/swagger"
	"gitlab.com/zigal0/architect"
	"gitlab.com/zigal0/architect/pkg/logger"
)

func main() {
	ctx := context.Background()

	err := config.InitConfig()
	if err != nil {
		logger.Fatalf("failed to init config: %v", err)
	}

	appSettings, err := createAppSettingsFromConfig()
	if err != nil {
		logger.Fatalf("failed to create app settings: %v", err)
	}

	a, err := architect.NewApp(appSettings)
	if err != nil {
		logger.Fatalf("failed to create app: %v", err)
	}

	pgDB, err := database.InitPostgreSQLConn(ctx)
	if err != nil {
		logger.Fatalf("failed to connect to pg: %v", err)
	}

	employeeRepo := employee_repository.New(pgDB)

	employeeManager := employee_manager.New(employeeRepo)

	err = a.Run(
		tango_api_service_impl.New(employeeManager),
	)
	if err != nil {
		logger.Fatalf("faile to run app: %v", err)
	}
}

func createAppSettingsFromConfig() (architect.AppSettings, error) {
	logLevel, err := config.GetValue(config.Key_LogLevel)
	if err != nil {
		return architect.AppSettings{}, err
	}

	host, err := config.GetValue(config.Key_Host)
	if err != nil {
		return architect.AppSettings{}, err
	}

	portHTTP, err := config.GetValue(config.Key_PortHTTP)
	if err != nil {
		return architect.AppSettings{}, err
	}

	portSwagger, err := config.GetValue(config.Key_PortSwagger)
	if err != nil {
		return architect.AppSettings{}, err
	}

	portGRPC, err := config.GetValue(config.Key_PortGRPC)
	if err != nil {
		return architect.AppSettings{}, err
	}

	return architect.AppSettings{
		LogLevel:    logLevel.String(),
		Host:        host.String(),
		PortHTTP:    uint(portHTTP.Int32()),
		PortSwagger: uint(portSwagger.Int32()),
		PortGRPC:    uint(portGRPC.Int32()),
		SwaggerFS:   swagger.SRC,
	}, nil
}
