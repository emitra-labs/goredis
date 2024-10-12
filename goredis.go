package goredis

import (
	"context"

	"github.com/emitra-labs/common/validator"
	"github.com/redis/go-redis/v9"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	URL string `env:"GOREDIS_URL" validate:"required"`
}

var Client *redis.Client

func Open() {
	var config Config

	// Load config from environment variables
	err := envconfig.Process(context.Background(), &config)
	if err != nil {
		panic(err)
	}

	// Validate config
	err = validator.Validate(config)
	if err != nil {
		panic(err)
	}

	opts, err := redis.ParseURL(config.URL)
	if err != nil {
		panic(err)
	}

	Client = redis.NewClient(opts)
}

func Close() error {
	return Client.Close()
}
