package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gitlab.com/zigal0/architect/pkg/logger"
)

// nolint: revive, gosec
const (
	Key_LogLevel key = "LOG_LEVEL"

	Key_Host        key = "HOST"
	Key_PortHTTP    key = "PORT_HTTP"
	Key_PortSwagger key = "PORT_SWAGGER"
	Key_PortGRPC    key = "PORT_GRPC"

	Key_PGDSN key = "PG_DSN"
)

// Errors
const (
	formatErrNoParamInConfig = "no param '%s' in config"

	formatErrInvalidCast = "failed to cast to %s: %w"
)

// Other constanats
const (
	pathToENV = "config/.env"

	base      = 10
	bitSize32 = 32
)

func InitConfig() error {
	err := godotenv.Load(pathToENV)
	if err != nil {
		return fmt.Errorf("fialed to load env: %w", err)
	}

	return nil
}

type key string

func (ck key) String() string {
	return string(ck)
}

func GetValue(configKey key) (Value, error) {
	val, found := os.LookupEnv(configKey.String())
	if !found {
		return value{}, fmt.Errorf(formatErrNoParamInConfig, configKey)
	}

	return value{val: val}, nil
}

type Value interface {
	String() string
	StringSafe() (string, error)

	Int32() int32
	Int32Safe() (int32, error)

	Bool() bool
	BoolSafe() (bool, error)
}

type value struct {
	val string
}

func (v value) String() string {
	return v.val
}

func (v value) StringSafe() (string, error) {
	return v.val, nil
}

func (v value) Int32() int32 {
	castedValue, err := strconv.ParseInt(v.val, base, bitSize32)
	if err != nil {
		logger.Error(fmt.Errorf(formatErrInvalidCast, "int32", err))

		return 0
	}

	return int32(castedValue)
}

func (v value) Int32Safe() (int32, error) {
	castedValue, err := strconv.ParseInt(v.val, base, bitSize32)
	if err != nil {
		return 0, fmt.Errorf(formatErrInvalidCast, "int32", err)
	}

	return int32(castedValue), nil
}

func (v value) Bool() bool {
	castedValue, err := strconv.ParseBool(v.val)
	if err != nil {
		logger.Error(fmt.Errorf(formatErrInvalidCast, "bool", err))

		return false
	}

	return castedValue
}

func (v value) BoolSafe() (bool, error) {
	castedValue, err := strconv.ParseBool(v.val)
	if err != nil {
		return false, fmt.Errorf(formatErrInvalidCast, "bool", err)
	}

	return castedValue, nil
}
