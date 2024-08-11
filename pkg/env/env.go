package env

import (
	"errors"
	"time"
)

type Environment struct {
	Server struct {
		Port int `yml:"port" envconfig:"SERVER_PORT"`
	} `yml:"server"`
	Database struct {
		Host     string `yml:"host" envconfig:"DATABASE_HOST"`
		Port     int    `yml:"port" envconfig:"DATABASE_PORT"`
		Name     string `yml:"name" envconfig:"DATABASE_NAME"`
		User     string `yml:"user" envconfig:"DATABASE_USER"`
		Password string `yml:"password" envconfig:"DATABASE_PASSWORD"`
		SSL      bool   `yml:"ssl" envconfig:"DATABASE_SSL"`
		Pool     struct {
			MaxOpenConns    int           `yml:"max-open-conns" envconfig:"DATABASE_MAX_OPEN_CONNS"`
			MaxIdleConns    int           `yml:"max-idle-conns" envconfig:"DATABASE_MAX_IDLE_CONNS"`
			ConnMaxLifetime time.Duration `yml:"conn-max-lifetime" envconfig:"DATABASE_CONN_MAX_LIFETIME"`
			ConnMaxIdleTime time.Duration `yml:"conn-max-idle-time" envconfig:"DATABASE_CONN_MAX_IDLE_TIME"`
		}
	} `yml:"database"`
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
