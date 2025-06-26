// Package alex provides configuration and utility functions for Redis cache operations.
// It includes builders and configuration structures to simplify Redis connection setup.
package alex

import (
	"errors"

	"github.com/zeroxsolutions/strike/builderutil"
)

// NewRedisConfig creates a new RedisConfig from RedisConfigOptions by applying the provided option functions.
// It uses the builderutil package to build the configuration options in a functional options pattern,
// then creates and returns a final RedisConfig instance.
// The function performs validation to ensure the configuration is valid before returning.
//
// Validation rules:
//   - Configuration options must not be nil
//   - Redis address (Addr) is required and cannot be empty
//   - Database number (DB) must be greater than or equal to 0
//
// Parameters:
//   - opts: Variable number of option functions that configure the RedisConfigOptions
//
// Returns:
//   - *RedisConfig: A pointer to the final Redis configuration instance
//   - error: An error if the configuration building process fails or validation fails
//
// Example:
//
//	builder := NewRedisConfigOptions()
//	config, err := NewRedisConfig(builder.SetAddr("localhost:6379").SetDB(0))
//	if err != nil {
//	    log.Fatal(err)
//	}
func NewRedisConfig(opts ...builderutil.Lister[RedisConfigOptions]) (*RedisConfig, error) {
	options, err := builderutil.Build(opts...)
	if err != nil {
		return nil, err
	}
	if options == nil {
		return nil, errors.New("redis config options is nil")
	}
	if options.Addr == "" {
		return nil, errors.New("redis address is required")
	}
	if options.DB < 0 {
		return nil, errors.New("redis database must be greater than 0")
	}
	return &RedisConfig{
		Addr:     options.Addr,
		Password: options.Password,
		DB:       options.DB,
	}, nil
}
