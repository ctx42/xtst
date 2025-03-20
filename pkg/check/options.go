// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"github.com/ctx42/xtst/pkg/dump"
)

// Check is signature for generic check function comparing two arguments
// returning error if they are not. The returned error might be one or more
// errors joined with [errors.Join].
type Check func(want, have any, opts ...Option) error

// SingleCheck is signature for generic check function checking single value.
// Returns error if value does not match expectations. The returned error might
// be one or more errors joined with [errors.Join].
type SingleCheck func(have any, opts ...Option) error

// Option represents [Check] and [SingleCheck] option.
type Option func(Options) Options

// WithTrail is [Check] option setting initial field/element/key breadcrumb
// trail.
func WithTrail(pth string) Option {
	return func(ops Options) Options {
		ops.Trail = pth
		return ops
	}
}

// WithDump is [Check] and [SingleCheck] option setting [dump.Config] options.
func WithDump(dopts ...dump.Option) Option {
	return func(copts Options) Options {
		for _, opt := range dopts {
			opt(&copts.DumpConfig)
		}
		return copts
	}
}

type DumpConfig = dump.Config

// Options represents options used by [Check] and [SingleCheck] functions.
type Options struct {
	DumpConfig // Dump configuration.

	// Field/element/key breadcrumb trail which uniquely describes nested
	// object being checked.
	Trail string
}

// DefaultOptions returns default [Options].
func DefaultOptions() Options { return Options{DumpConfig: dump.NewConfig()} }

// set sets [Options] from slice of [Option] functions.
func (ops Options) set(opts []Option) Options {
	dst := ops
	for _, opt := range opts {
		dst = opt(dst)
	}
	return dst
}
