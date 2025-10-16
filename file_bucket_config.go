package alex

// FileBucketOption represents the configuration options for a file bucket.
// It includes the base path of the file bucket.
type FileBucketOption struct {
	BasePath string // BasePath is the base path of the file bucket.
}

// FileBucketOptionBuilder provides a builder pattern for constructing FileBucketOption.
// It accumulates option functions that can be applied to configure a FileBucketOption instance.
// This builder implements the builderutil.Lister interface to work with the functional options pattern.
type FileBucketOptionBuilder struct {
	Opts []func(*FileBucketOption) error // Opts contains the list of option functions to be applied
}

// SetBasePath configures the base path for the file bucket.
// It appends an option function that sets the BasePath field of FileBucketOption.
//
// Parameters:
//   - basePath: The base path of the file bucket
//
// Returns:
//   - *FileBucketOptionBuilder: The builder instance for method chaining
//
// Example:
//
//	builder := NewFileBucketOption()
//	config, err := NewFileBucketOption(builder.SetBasePath("basePath"))
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("File Bucket Option: %+v\n", config)
func (builder *FileBucketOptionBuilder) SetBasePath(basePath string) *FileBucketOptionBuilder {
	builder.Opts = append(builder.Opts, func(args *FileBucketOption) error {
		args.BasePath = basePath
		return nil
	})
	return builder
}

// List returns the slice of option functions accumulated by the builder.
// This method implements the builderutil.Lister interface, allowing the builder
// to be used with the builderutil.Build function.
//
// Returns:
//   - []func(*FileBucketOption) error: A slice of option functions that can be applied to configure FileBucketOption
func (builder *FileBucketOptionBuilder) List() []func(*FileBucketOption) error {
	return builder.Opts
}

// NewFileBucketOption creates and returns a new instance of FileBucketOptionBuilder.
// This function provides a convenient way to initialize the builder for creating FileBucket configuration options.
//
// Returns:
//   - *FileBucketOptionBuilder: A new instance of FileBucketOptionBuilder ready to be configured
//
// Example:
//
//	builder := NewFileBucketOption()
func NewFileBucketOption() *FileBucketOptionBuilder {
	return &FileBucketOptionBuilder{}
}

// FileBucketConfig represents the final FileBucket configuration used for using a file bucket.
// This struct is created from FileBucketOption after validation and contains all the necessary
// parameters for using a file bucket.
type FileBucketConfig struct {
	BasePath string // BasePath is the base path of the file bucket.
}
