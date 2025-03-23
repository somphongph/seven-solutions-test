package configs

import (
	"time"

	"go.uber.org/zap/zapcore"
)

type Configs struct {
	App     App     `mapstructure:"app"`
	Server  Server  `mapstructure:"server"`
	Service Service `mapstructure:"service"`
}

type App struct {
	Name string `mapstructure:"name"`
	Env  string `mapstructure:"env" example:"dev,sit,uat,prod"`
	Port int    `mapstructure:"port"`
}

type Server struct {
	PrettyPrint  bool `mapstructure:"pretty_print"`
	LogLevel     zapcore.Level
	IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

type Service struct {
	Debug bool `mapstructure:"debug"`
}
