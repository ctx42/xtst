package check

// Check is signature for generic check function comparing two arguments
// returning error if they are not. The returned error might be one or more
// errors joined with [errors.Join].
type Check func(want, have any, opts ...Option) error

// SingleCheck is signature for generic check function checking single value.
// Returns error if value does not match expectations. The returned error might
// be one or more errors joined with [errors.Join].
type SingleCheck func(have any, opts ...Option) error

// Option represents [Check] option.
type Option func(Options) Options

// WithPath is [Check] option setting initial field/element/key path.
func WithPath(pth string) Option {
	return func(ops Options) Options {
		ops.Path = pth
		return ops
	}
}

// Options represents options used by [Check] and [SingleCheck] functions.
type Options struct {
	// Field/element/key path which uniquely describes nested object being
	// checked.
	Path string
}

// DefaultOptions returns default [Options].
func DefaultOptions() Options { return Options{} }

// set sets [Options] from slice of [Option] functions.
func (ops Options) set(opts []Option) Options {
	dst := ops
	for _, opt := range opts {
		dst = opt(dst)
	}
	return dst
}
