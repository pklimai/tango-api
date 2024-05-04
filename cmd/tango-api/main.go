package main

import (
	"gitlab.com/zigal0-group/nica/tango-api/config"
	"gitlab.com/zigal0-group/nica/tango-api/internal/api/tango_api_service_impl"
	"gitlab.com/zigal0-group/nica/tango-api/internal/generated/swagger"
	"gitlab.com/zigal0/architect"
	"gitlab.com/zigal0/architect/pkg/logger"
)

func main() {
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

	err = a.Run(
		tango_api_service_impl.New(),
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
