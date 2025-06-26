package alex

// RedisConfigOptions holds the configuration options for connecting to a Redis cache system.
// It includes the address of the Redis server, an optional password for authentication,
// and the database number to select within the Redis instance.
// This struct is used as input for building the final RedisConfig.
type RedisConfigOptions struct {
	Addr     string // Addr is the address of the Redis server (e.g., "localhost:6379").
	Password string // Password is the optional authentication password for the Redis server.
	DB       int    // DB is the database number to be selected within the Redis instance (default is 0).
}

// RedisConfigOptionsBuilder provides a builder pattern for constructing RedisConfigOptions.
// It accumulates option functions that can be applied to configure a RedisConfigOptions instance.
// This builder implements the builderutil.Lister interface to work with the functional options pattern.
type RedisConfigOptionsBuilder struct {
	Opts []func(*RedisConfigOptions) error // Opts contains the list of option functions to be applied
}

// SetAddr configures the Redis server address for the connection.
// It appends an option function that sets the Addr field of RedisConfigOptions.
//
// Parameters:
//   - addr: The Redis server address (e.g., "localhost:6379", "redis.example.com:6379")
//
// Returns:
//   - *RedisConfigOptionsBuilder: The builder instance for method chaining
func (b *RedisConfigOptionsBuilder) SetAddr(addr string) *RedisConfigOptionsBuilder {
	b.Opts = append(b.Opts, func(o *RedisConfigOptions) error {
		o.Addr = addr
		return nil
	})
	return b
}

// SetPassword configures the authentication password for the Redis connection.
// It appends an option function that sets the Password field of RedisConfigOptions.
//
// Parameters:
//   - password: The authentication password for the Redis server
//
// Returns:
//   - *RedisConfigOptionsBuilder: The builder instance for method chaining
func (b *RedisConfigOptionsBuilder) SetPassword(password string) *RedisConfigOptionsBuilder {
	b.Opts = append(b.Opts, func(o *RedisConfigOptions) error {
		o.Password = password
		return nil
	})
	return b
}

// SetDB configures the Redis database number to select.
// It appends an option function that sets the DB field of RedisConfigOptions.
// Redis databases are numbered starting from 0.
//
// Parameters:
//   - db: The database number to select (typically 0-15 in default Redis configurations)
//
// Returns:
//   - *RedisConfigOptionsBuilder: The builder instance for method chaining
func (b *RedisConfigOptionsBuilder) SetDB(db int) *RedisConfigOptionsBuilder {
	b.Opts = append(b.Opts, func(o *RedisConfigOptions) error {
		o.DB = db
		return nil
	})
	return b
}

// List returns the slice of option functions accumulated by the builder.
// This method implements the builderutil.Lister interface, allowing the builder
// to be used with the builderutil.Build function.
//
// Returns:
//   - []func(*RedisConfigOptions) error: A slice of option functions that can be applied to configure RedisConfigOptions
func (b *RedisConfigOptionsBuilder) List() []func(*RedisConfigOptions) error {
	return b.Opts
}

// NewRedisConfigOptionsBuilder creates and returns a new instance of RedisConfigOptionsBuilder.
// This function provides a convenient way to initialize the builder for creating Redis configuration options.
//
// Returns:
//   - *RedisConfigOptionsBuilder: A new instance of RedisConfigOptionsBuilder ready to be configured
//
// Example:
//
//	builder := NewRedisConfigOptionsBuilder()
//	config, err := NewRedisConfig(builder.SetAddr("localhost:6379").SetPassword("secret").SetDB(1))
func NewRedisConfigOptionsBuilder() *RedisConfigOptionsBuilder {
	return &RedisConfigOptionsBuilder{}
}

// RedisConfig represents the final Redis configuration used for establishing connections.
// This struct is created from RedisConfigOptions after validation and contains all the necessary
// parameters for connecting to a Redis server.
type RedisConfig struct {
	Addr     string // Addr is the address of the Redis server (e.g., "localhost:6379").
	Password string // Password is the optional authentication password for the Redis server.
	DB       int    // DB is the database number to be selected within the Redis instance (default is 0).
}
