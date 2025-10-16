package alex

import (
	"errors"

	"github.com/zeroxsolutions/strike/builderutil"
)

// NewFileBucketConfig creates a new FileBucketConfig from FileBucketOption by applying the provided option functions.
// It uses the builderutil package to build the configuration options in a functional options pattern,
// then creates and returns a final FileBucketConfig instance.
//
// Parameters:
//   - opts: Variable number of option functions that configure the FileBucketOption
//
// Returns:
//   - *FileBucketConfig: A pointer to the final FileBucket configuration instance
//   - error: An error if the configuration building process fails or validation fails
//
// Example:
//
//	builder := NewFileBucketOption()
//	config, err := NewFileBucketConfig(builder.SetBasePath("basePath"))
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("File Bucket Config: %+v\n", config)
func NewFileBucketConfig(opts ...builderutil.Lister[FileBucketOption]) (*FileBucketConfig, error) {
	options, err := builderutil.Build(opts...)
	if err != nil {
		return nil, err
	}
	if options.BasePath == "" {
		return nil, errors.New("file bucket base path is required")
	}
	return &FileBucketConfig{
		BasePath: options.BasePath,
	}, nil
}
