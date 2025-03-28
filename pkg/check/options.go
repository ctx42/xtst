// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"time"

	"github.com/ctx42/xtst/pkg/dump"
)

// Package wide default configuration.
const (
	// DefaultParseTimeFormat is default format for dumping [time.Time] values.
	DefaultParseTimeFormat = time.RFC3339Nano

	// DefaultRecentDuration is default duration when comparing recent dates.
	DefaultRecentDuration = 10 * time.Second

	// DefaultDumpTimeFormat is default format for parsing time strings.
	DefaultDumpTimeFormat = time.RFC3339Nano

	// DefaultDumpDepth is default depth when dumping values recursively in log
	// messages.
	DefaultDumpDepth = 6
)

// Package wide configuration.
var (
	// ParseTimeFormat is configurable format for parsing time strings.
	ParseTimeFormat = DefaultParseTimeFormat

	// RecentDuration is configurable duration when comparing recent dates.
	RecentDuration = DefaultRecentDuration

	// DumpTimeFormat is configurable format for dumping [time.Time] values.
	DumpTimeFormat = DefaultDumpTimeFormat

	// DumpDepth is configurable depth when dumping values in log messages.
	DumpDepth = DefaultDumpDepth
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

// WithTimeFormat is [Check] option setting time format when parsing dates.
func WithTimeFormat(format string) Option {
	return func(ops Options) Options {
		ops.TimeFormat = format
		return ops
	}
}

// WithRecent is [Check] option setting duration used to compare recent dates.
func WithRecent(recent time.Duration) Option {
	return func(ops Options) Options {
		ops.Recent = recent
		return ops
	}
}

// WithDump is [Check] and [SingleCheck] option setting [dump.Config] options.
func WithDump(optsD ...dump.Option) Option {
	return func(optsC Options) Options {
		for _, opt := range optsD {
			opt(&optsC.DumpCfg)
		}
		return optsC
	}
}

// Options represents options used by [Check] and [SingleCheck] functions.
type Options struct {
	// Dump configuration.
	DumpCfg dump.Config

	// Time format when parsing time strings (default: [time.RFC3339]).
	TimeFormat string

	// Duration when comparing recent dates.
	Recent time.Duration

	// Field/element/key breadcrumb trail being checked.
	Trail string

	// Function used to get current time. Used preliminary to inject clock in
	// tests of checks and assertions using [time.Now].
	now func() time.Time
}

// DefaultOptions returns default [Options].
func DefaultOptions() Options {
	return Options{
		DumpCfg: dump.NewConfig(
			dump.WithTimeFormat(DumpTimeFormat),
			dump.WithMaxDepth(DumpDepth),
		),
		Recent:     RecentDuration,
		TimeFormat: ParseTimeFormat,
		now:        time.Now,
	}
}

// set sets [Options] from slice of [Option] functions.
func (ops Options) set(opts []Option) Options {
	dst := ops
	for _, opt := range opts {
		dst = opt(dst)
	}
	return dst
}
