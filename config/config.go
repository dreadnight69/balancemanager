package config

import (
	"github.com/pkg/errors"
	"os"
)

func init() {
	_ = os.Setenv("DB_HOST", "0.0.0.0")
	_ = os.Setenv("DB_PORT", "5432")
	_ = os.Setenv("DB_NAME", "balancemanager")
	_ = os.Setenv("DB_USER", "postgres")
	_ = os.Setenv("DB_PASSWORD", "admin")
	_ = os.Setenv("LISTEN_HOST", "0.0.0.0")
	_ = os.Setenv("LISTEN_PORT", "8001")
	_ = os.Setenv("SERVER_READ_TIMEOUT_SECONDS", "5")
	_ = os.Setenv("SERVER_WRITE_TIMEOUT_SECONDS", "5")

}

type Config struct {
	Http *Http `json:"http"`
	DB   *DB   `json:"db"`
}

func New() Config {
	return Config{
		Http: &Http{},
		DB:   &DB{},
	}
}

func (c *Config) Init() error {

	if err := c.Http.Init(); err != nil {
		return errors.Wrap(err, "Init: unable to get env for HTTP")
	}
	if err := c.DB.Init(); err != nil {
		return errors.Wrap(err, "Init: unable to get env for DB")
	}
	return nil
}
