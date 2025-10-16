package alex

// MinioOption represents the configuration options for a Minio client.
// It includes the endpoint, access key, secret key, use SSL, bucket name, and location.
type MinioOption struct {
	Endpoint   string // Endpoint is the URL of the Minio server (e.g., "https://minio.example.com").
	AccessKey  string // AccessKey is the access key for the Minio server.
	SecretKey  string // SecretKey is the secret key for the Minio server.
	UseSSL     bool   // UseSSL is a flag indicating whether to use SSL for the connection.
	BucketName string // BucketName is the name of the bucket to use.
	Region     string // Region is the region of the bucket to use.
}

// MinioOptionBuilder provides a builder pattern for constructing MinioOption.
// It accumulates option functions that can be applied to configure a MinioOption instance.
// This builder implements the builderutil.Lister interface to work with the functional options pattern.
type MinioOptionBuilder struct {
	Opts []func(*MinioOption) error // Opts contains the list of option functions to be applied
}

// List returns the slice of option functions accumulated by the builder.
// This method implements the builderutil.Lister interface, allowing the builder
// to be used with the builderutil.Build function.
//
// Returns:
//   - []func(*MinioOption) error: A slice of option functions that can be applied to configure MinioOption
func (builder *MinioOptionBuilder) List() []func(*MinioOption) error {
	return builder.Opts
}

// NewMinioOption creates and returns a new instance of MinioOptionBuilder.
// This function provides a convenient way to initialize the builder for creating Minio configuration options.
//
// Returns:
//   - *MinioOptionBuilder: A new instance of MinioOptionBuilder ready to be configured
//
// Example:
//
//	builder := NewMinioOption()
//	config, err := NewMinioOption(builder.SetEndpoint("https://minio.example.com").SetAccessKey("accessKey").SetSecretKey("secretKey").SetUseSSL(true).SetBucketName("bucketName").SetRegion("region"))
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Minio Option: %+v\n", config)
func NewMinioOption() *MinioOptionBuilder {
	return &MinioOptionBuilder{}
}

// SetEndpoint configures the endpoint for the Minio client.
// It appends an option function that sets the Endpoint field of MinioOption.
//
// Parameters:
//   - endpoint: The URL of the Minio server (e.g., "https://minio.example.com")
//
// Returns:
//   - *MinioOptionBuilder: The builder instance for method chaining
//
// Example:
//
//	builder := NewMinioOption()
//	config, err := NewMinioOption(builder.SetEndpoint("https://minio.example.com"))
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Minio Option: %+v\n", config)
func (builder *MinioOptionBuilder) SetEndpoint(endpoint string) *MinioOptionBuilder {
	builder.Opts = append(builder.Opts, func(args *MinioOption) error {
		args.Endpoint = endpoint
		return nil
	})
	return builder
}

// SetAccessKey configures the access key for the Minio client.
// It appends an option function that sets the AccessKey field of MinioOption.
//
// Parameters:
//   - accessKey: The access key for the Minio server
//
// Returns:
//   - *MinioOptionBuilder: The builder instance for method chaining
//
// Example:
//
//	builder := NewMinioOption()
//	config, err := NewMinioOption(builder.SetAccessKey("accessKey"))
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Minio Option: %+v\n", config)
func (builder *MinioOptionBuilder) SetAccessKey(accessKey string) *MinioOptionBuilder {
	builder.Opts = append(builder.Opts, func(args *MinioOption) error {
		args.AccessKey = accessKey
		return nil
	})
	return builder
}

// SetSecretKey configures the secret key for the Minio client.
// It appends an option function that sets the SecretKey field of MinioOption.
//
// Parameters:
//   - secretKey: The secret key for the Minio server
//
// Returns:
//   - *MinioOptionBuilder: The builder instance for method chaining
//
// Example:
//
//	builder := NewMinioOption()
//	config, err := NewMinioOption(builder.SetSecretKey("secretKey"))
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Minio Option: %+v\n", config)
func (builder *MinioOptionBuilder) SetSecretKey(secretKey string) *MinioOptionBuilder {
	builder.Opts = append(builder.Opts, func(args *MinioOption) error {
		args.SecretKey = secretKey
		return nil
	})
	return builder
}

// SetUseSSL configures the use of SSL for the Minio client.
// It appends an option function that sets the UseSSL field of MinioOption.
//
// Parameters:
//   - useSSL: A flag indicating whether to use SSL for the connection
//
// Returns:
//   - *MinioOptionBuilder: The builder instance for method chaining
//
// Example:
//
//	builder := NewMinioOption()
//	config, err := NewMinioOption(builder.SetUseSSL(true))
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Minio Option: %+v\n", config)
func (builder *MinioOptionBuilder) SetUseSSL(useSSL bool) *MinioOptionBuilder {
	builder.Opts = append(builder.Opts, func(args *MinioOption) error {
		args.UseSSL = useSSL
		return nil
	})
	return builder
}

// SetBucketName configures the bucket name for the Minio client.
// It appends an option function that sets the BucketName field of MinioOption.
//
// Parameters:
//   - bucketName: The name of the bucket to use
//
// Returns:
//   - *MinioOptionBuilder: The builder instance for method chaining
//
// Example:
//
//	builder := NewMinioOption()
//	config, err := NewMinioOption(builder.SetBucketName("bucketName"))
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Minio Option: %+v\n", config)
func (builder *MinioOptionBuilder) SetBucketName(bucketName string) *MinioOptionBuilder {
	builder.Opts = append(builder.Opts, func(args *MinioOption) error {
		args.BucketName = bucketName
		return nil
	})
	return builder
}

// SetRegion configures the region for the Minio client.
// It appends an option function that sets the Region field of MinioOption.
//
// Parameters:
//   - region: The region of the bucket to use
//
// Returns:
//   - *MinioOptionBuilder: The builder instance for method chaining
//
// Example:
//
//	builder := NewMinioOption()
//	config, err := NewMinioOption(builder.SetRegion("region"))
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Minio Option: %+v\n", config)
func (builder *MinioOptionBuilder) SetRegion(region string) *MinioOptionBuilder {
	builder.Opts = append(builder.Opts, func(args *MinioOption) error {
		args.Region = region
		return nil
	})
	return builder
}

// MinioConfig represents the final Minio configuration used for establishing connections.
// This struct is created from MinioOption after validation and contains all the necessary
// parameters for connecting to a Minio server.
type MinioConfig struct {
	Endpoint   string // Endpoint is the URL of the Minio server (e.g., "https://minio.example.com").
	AccessKey  string // AccessKey is the access key for the Minio server.
	SecretKey  string // SecretKey is the secret key for the Minio server.
	UseSSL     bool   // UseSSL is a flag indicating whether to use SSL for the connection.
	BucketName string // BucketName is the name of the bucket to use.
	Region     string // Region is the region of the bucket to use.
}
