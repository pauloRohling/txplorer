package env

import (
	"errors"
	"time"
)

type Environment struct {
	Server struct {
		Port int32 `yaml:"port" envconfig:"SERVER_PORT"`
	} `yaml:"server"`
	Security struct {
		Secret          string        `yaml:"secret" envconfig:"SECURITY_SECRET"`
		TokenExpiration time.Duration `yaml:"token-expiration" envconfig:"SECURITY_TOKEN_EXPIRATION"`
	} `yaml:"security"`
	Database struct {
		Host     string `yaml:"host" envconfig:"DATABASE_HOST"`
		Port     int32  `yaml:"port" envconfig:"DATABASE_PORT"`
		Name     string `yaml:"name" envconfig:"DATABASE_NAME"`
		User     string `yaml:"user" envconfig:"DATABASE_USER"`
		Password string `yaml:"password" envconfig:"DATABASE_PASSWORD"`
		SSL      bool   `yaml:"ssl" envconfig:"DATABASE_SSL"`
		Pool     struct {
			MaxOpenConns    int           `yaml:"max-open-conns" envconfig:"DATABASE_MAX_OPEN_CONNS"`
			MaxIdleConns    int           `yaml:"max-idle-conns" envconfig:"DATABASE_MAX_IDLE_CONNS"`
			ConnMaxLifetime time.Duration `yaml:"conn-max-lifetime" envconfig:"DATABASE_CONN_MAX_LIFETIME"`
			ConnMaxIdleTime time.Duration `yaml:"conn-max-idle-time" envconfig:"DATABASE_CONN_MAX_IDLE_TIME"`
		}
	} `yaml:"database"`
}

func (env *Environment) Validate() error {
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

	if env.Security.TokenExpiration == 0 {
		env.Security.TokenExpiration = time.Hour * 24
	}

	if env.Server.Port == 0 {
		env.Server.Port = 8080
	}

	if env.Database.Pool.MaxOpenConns == 0 {
		env.Database.Pool.MaxOpenConns = 100
	}

	if env.Database.Pool.MaxIdleConns == 0 {
		env.Database.Pool.MaxIdleConns = 10
	}

	if env.Database.Pool.ConnMaxLifetime == 0 {
		env.Database.Pool.ConnMaxLifetime = 5 * time.Minute
	}

	if env.Database.Pool.ConnMaxIdleTime == 0 {
		env.Database.Pool.ConnMaxIdleTime = time.Second * 30
	}

	return nil
}
