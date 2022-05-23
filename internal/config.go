package internal

import "github.com/ilyakaznacheev/cleanenv"

// Conf is to store our service configuration
type Conf struct {
	Port      int    `env:"PORT" env-default:"8080"`
	Host      string `env:"HOST" env-default:"localhost"`
	RedisHost string `env:"REDIS_HOST" env-default:"localhost"`
	RedisPort string `env:"REDIS_PORT" env-default:"6379"`
	RedisPass string `env:"REDIS_PASS" env-default:""`
	RedisDB   int    `env:"REDIS_DB" env-default:"0"`
}

// NewConf return new Conf instance from env
func NewConf() (*Conf, error) {
	var conf Conf

	err := cleanenv.ReadEnv(&conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
