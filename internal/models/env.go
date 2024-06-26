package models

import "effective_mobile_test_task/pkg/misc"

type EnvConfig struct {
	Production        bool   `env:"PRODUCTION"`
	AppConfigPath     string `env:"APP_CONFIG_PATH"`
	LogLevel          string `env:"LOG_LEVEL"`
	MaxConns          int32  `env:"MAX_CONNS"`
	MinConns          int32  `env:"MIN_CONNS"`
	MaxConnLifetime   int64  `env:"MAX_CONN_LIFE_TIME"`
	MaxConnIdleTime   int64  `env:"MAX_CONNIDLE_TIME"`
	HealthCheckPeriod int64  `env:"HEALTH_CHECK_PERIOD"`
	EmApi             string `env:"EM_API"`
}

func NewEnvConfig() (*EnvConfig, error) {
	var conf EnvConfig
	err := misc.ParseEnv(&conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
