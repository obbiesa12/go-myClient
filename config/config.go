package config

import (
	"github.com/joeshaw/envdecode"
)

type Config struct {
	DBConn       string `env:"DB_CONN,required"`
	RedisAddress string `env:"REDIS_ADDRESS,required"`
	S3Name       string `env:"S3_NAME,required"`
	S3Region     string `env:"S3_REGION,required"`
}

func LoadConfig() *Config {
	var cfg Config
	if err := envdecode.Decode(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
