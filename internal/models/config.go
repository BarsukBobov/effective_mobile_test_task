package models

import (
	"effective_mobile_test_task/pkg/misc"
)

type PostgreSQL struct {
	Dsn string `json:"dsn"`
}

type AppConfig struct {
	PostgreSQL          PostgreSQL `json:"postgresql"`
	DbmateMigrationsDir string     `json:"dbmate_migrations_dir"`
}

func NewAppConfig(dbConfig string) (*AppConfig, error) {
	return misc.JsonToStruct[AppConfig](dbConfig)
}
