package alex

import (
	"errors"

	"github.com/zeroxsolutions/strike/builderutil"
)

// NewMinioConfig creates a new MinioConfig from MinioOption by applying the provided option functions.
// It uses the builderutil package to build the configuration options in a functional options pattern,
// then creates and returns a final MinioConfig instance.
//
// Parameters:
//   - opts: Variable number of option functions that configure the MinioOption
//
// Returns:
//   - *MinioConfig: A pointer to the final Minio configuration instance
//   - error: An error if the configuration building process fails or validation fails
//
// Example:
//
//	builder := NewMinioOption()
//	config, err := NewMinioConfig(builder.SetEndpoint("https://minio.example.com").SetAccessKey("accessKey").SetSecretKey("secretKey").SetUseSSL(true).SetBucketName("bucketName").SetRegion("region"))
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Minio Config: %+v\n", config)
func NewMinioConfig(opts ...builderutil.Lister[MinioOption]) (*MinioConfig, error) {
	options, err := builderutil.Build(opts...)
	if err != nil {
		return nil, err
	}
	if options.Endpoint == "" {
		return nil, errors.New("minio endpoint is required")
	}
	if options.AccessKey == "" {
		return nil, errors.New("minio access key is required")
	}
	if options.SecretKey == "" {
		return nil, errors.New("minio secret key is required")
	}
	if options.BucketName == "" {
		return nil, errors.New("minio bucket name is required")
	}
	return &MinioConfig{
		Endpoint:   options.Endpoint,
		AccessKey:  options.AccessKey,
		SecretKey:  options.SecretKey,
		UseSSL:     options.UseSSL,
		BucketName: options.BucketName,
		Region:     options.Region,
	}, nil
}
