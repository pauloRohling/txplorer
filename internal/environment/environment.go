package environment

import (
	"errors"
	"time"
)

type Config struct {
	Server struct {
		Port int32 `yaml:"port" env:"SERVER_PORT" env-default:"8080"`
	} `yaml:"server"`
	Security struct {
		Secret          string        `yaml:"secret" env:"SECURITY_SECRET"`
		TokenExpiration time.Duration `yaml:"token-expiration" env:"SECURITY_TOKEN_EXPIRATION" env-default:"24h"`
	} `yaml:"security"`
	Database struct {
		Host     string `yaml:"host" env:"DATABASE_HOST"`
		Port     int32  `yaml:"port" env:"DATABASE_PORT"`
		Name     string `yaml:"name" env:"DATABASE_NAME"`
		User     string `yaml:"user" env:"DATABASE_USER"`
		Password string `yaml:"password" env:"DATABASE_PASSWORD"`
		SSL      *bool  `yaml:"ssl" env:"DATABASE_SSL"`
		Pool     struct {
			MaxOpenConns    int           `yaml:"max-open-conns" env:"DATABASE_MAX_OPEN_CONNS" env-default:"100"`
			MaxIdleConns    int           `yaml:"max-idle-conns" env:"DATABASE_MAX_IDLE_CONNS" env-default:"10"`
			ConnMaxLifetime time.Duration `yaml:"conn-max-lifetime" env:"DATABASE_CONN_MAX_LIFETIME" env-default:"5m"`
			ConnMaxIdleTime time.Duration `yaml:"conn-max-idle-time" env:"DATABASE_CONN_MAX_IDLE_TIME" env-default:"30s"`
		}
	} `yaml:"database"`
}

func (config *Config) validateDefaultValues() {
	if config.Database.SSL == nil {
		config.Database.SSL = new(bool)
		*config.Database.SSL = true
	}
}

func (config *Config) validateRequiredFields() error {
	if env.Database.Host == "" {
		return errors.New("database host is required")
	}

	if env.Database.Port == 0 {
		return errors.New("database port is required")
	}

	if env.Database.Name == "" {
		return errors.New("database name is required")
	}

	if env.Database.User == "" {
		return errors.New("database user is required")
	}

	if env.Database.Password == "" {
		return errors.New("database password is required")
	}

	if env.Security.Secret == "" {
		return errors.New("security secret is required")
	}

	return nil
}
