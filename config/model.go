package config

import (
	"github.com/go-redis/redis"
)

var (
	Conf        Config
	Env         string
	RedisClient *redis.Client

	EnvironmentLocal = "LOCAL"
	EnvironmentDev   = "DEV"
	EnvironmentUAT   = "UAT"
	EnvironmentProd  = "PROD"
	ListOfIsland     map[uint64]string

	searchPath = []string{
		"/etc/test_cache_CQRS",
		"$HOME/.test_cache_CQRS",
		".",
	}
	configDefaults = map[string]interface{}{
		"port":       1234,
		"logLevel":   "DEBUG",
		"logFormat":  "text",
		"signString": "supersecret",
	}
	configName = map[string]string{
		"local": "config.local",
		"dev":   "config.dev",
		"uat":   "config.uat",
		"prod":  "config.prod",
		"test":  "config.test",
	}
)

type Config struct {
	Env       string         `mapstructure:"env"`
	Port      int            `mapstructure:"port"`
	LogLevel  string         `mapstructure:"logLevel"`
	LogMode   bool           `mapstructure:"logMode"`
	LogFormat string         `mapstructure:"logFormat"`
	Postgres  PostgresConfig `mapstructure:"postgres"`
	Redis     RedisConfig    `mapstructure:"redis"`
	Rate      int64          `mapstructure:"rate"`
	Interval  string         `mapstructure:"interval"`
}

// PostgresConfig ...
type PostgresConfig struct {
	ConnMaxLifetime    int    `mapstructure:"connectTimeout"`
	MaxOpenConnections int    `mapstructure:"maxOpenConnections"`
	MaxIdleConnections int    `mapstructure:"maxIdleConnections"`
	Host               string `mapstructure:"host"`
	Port               string `mapstructure:"port"`
	Schema             string `mapstructure:"schema"`
	DBName             string `mapstructure:"dbName"`
	User               string `mapstructure:"user"`
	Password           string `mapstructure:"password"`
}

type RedisConfig struct {
	Host        string `mapstructure:"host"`
	Password    string `mapstructure:"password"`
	DB          int    `mapstructure:"db"`
	Port        int    `mapstructure:"port"`
	EnableRedis bool   `mapstructure:"enableRedis"`
}
