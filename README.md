# Alex

A Go package providing configuration and utility functions for Redis cache operations with a clean builder pattern and functional options.

## Features

- 🏗️ **Builder Pattern**: Easy-to-use builder for constructing Redis configurations
- ✅ **Validation**: Built-in validation for Redis connection parameters
- 🔧 **Functional Options**: Flexible configuration using functional options pattern
- 📦 **Type Safety**: Separate input/output types for better API design
- ⚡ **Zero Dependencies**: Minimal external dependencies

## Installation

```bash
go get github.com/your-username/alex
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/your-username/alex"
)

func main() {
    // Create a builder
    builder := alex.NewRedisConfigOptionsBuilder()
    
    // Build Redis configuration
    config, err := alex.NewRedisConfig(
        builder.SetAddr("localhost:6379").
               SetPassword("mypassword").
               SetDB(1),
    )
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Redis Config: %+v\n", config)
}
```

## API Reference

### Core Types

#### `RedisConfigOptions`
Input struct used for building Redis configurations:

```go
type RedisConfigOptions struct {
    Addr     string // Redis server address (e.g., "localhost:6379")
    Password string // Optional authentication password
    DB       int    // Database number (default: 0)
}
```

#### `RedisConfig`
Final Redis configuration struct returned after validation:

```go
type RedisConfig struct {
    Addr     string // Redis server address
    Password string // Authentication password
    DB       int    // Database number
}
```

### Builder

#### `RedisConfigOptionsBuilder`
Builder pattern for constructing Redis configurations:

```go
builder := alex.NewRedisConfigOptionsBuilder()
```

##### Methods

- **`SetAddr(addr string)`** - Set Redis server address
- **`SetPassword(password string)`** - Set authentication password  
- **`SetDB(db int)`** - Set database number
- **`List()`** - Get list of option functions (implements `builderutil.Lister`)

### Functions

#### `NewRedisConfigOptionsBuilder()`
Creates a new builder instance:

```go
builder := alex.NewRedisConfigOptionsBuilder()
```

#### `NewRedisConfig(opts ...builderutil.Lister[RedisConfigOptions])`
Creates and validates a Redis configuration:

```go
config, err := alex.NewRedisConfig(builder)
```

**Validation Rules:**
- Configuration options must not be nil
- Redis address is required and cannot be empty
- Database number must be greater than or equal to 0

### Constants

#### `TimeoutDefault`
Default timeout value in seconds:

```go
const TimeoutDefault = 1
```

## Usage Examples

### Basic Configuration

```go
builder := alex.NewRedisConfigOptionsBuilder()
config, err := alex.NewRedisConfig(
    builder.SetAddr("localhost:6379"),
)
if err != nil {
    log.Fatal(err)
}
```

### Full Configuration

```go
builder := alex.NewRedisConfigOptionsBuilder()
config, err := alex.NewRedisConfig(
    builder.SetAddr("redis.example.com:6379").
           SetPassword("secretpassword").
           SetDB(2),
)
if err != nil {
    log.Fatal(err)
}

// Use the config...
fmt.Printf("Connecting to: %s\n", config.Addr)
fmt.Printf("Database: %d\n", config.DB)
```

### Method Chaining

```go
config, err := alex.NewRedisConfig(
    alex.NewRedisConfigOptionsBuilder().
        SetAddr("localhost:6379").
        SetPassword("mypass").
        SetDB(1),
)
```

### Error Handling

```go
config, err := alex.NewRedisConfig(
    alex.NewRedisConfigOptionsBuilder().
        SetDB(-1), // Invalid DB number
)
if err != nil {
    // Handle validation error
    fmt.Printf("Configuration error: %v\n", err)
    return
}
```

## Architecture

The package uses a two-tier architecture:

1. **Input Layer** (`RedisConfigOptions`): Flexible configuration options that can be built using the builder pattern
2. **Output Layer** (`RedisConfig`): Validated, immutable configuration ready for use

This separation provides:
- **Type Safety**: Clear distinction between configuration input and final config
- **Validation**: Input is validated before creating the final configuration
- **Immutability**: Final configuration cannot be accidentally modified

## Validation

The `NewRedisConfig` function performs the following validations:

| Rule | Error Message |
|------|---------------|
| Options not nil | "redis config options is nil" |
| Address required | "redis address is required" |
| DB >= 0 | "redis database must be greater than 0" |

## Dependencies

- `github.com/zeroxsolutions/strike/builderutil` - For functional options pattern

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Changelog

### v0.0.1
- Initial release
- Builder pattern for Redis configuration
- Validation for Redis connection parameters
- Functional options support

```shell
go get github.com/zeroxsolutions/alex
```