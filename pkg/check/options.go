// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"strconv"
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

// WithTrailLog is [Check] option turning on collection of checked paths. The
// paths are added to the provided slice.
func WithTrailLog(list *[]string) Option {
	return func(ops Options) Options {
		ops.TrailLog = list
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

// WithOptions is [Check] option which passes all options.
func WithOptions(src Options) Option {
	return func(ops Options) Options {
		ops.DumpCfg = src.DumpCfg
		ops.TimeFormat = src.TimeFormat
		ops.Recent = src.Recent
		ops.Trail = src.Trail
		ops.TrailLog = src.TrailLog
		ops.now = src.now
		ops.skipType = src.skipType
		return ops
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

	// List of non-skipped trails.
	TrailLog *[]string

	// Function used to get current time. Used preliminary to inject clock in
	// tests of checks and assertions using [time.Now].
	now func() time.Time

	// In cases of nested structs you do not want to add field type to the
	// trail. When it's true the type argument in structTrail is ignored.
	skipType bool
}

// DefaultOptions returns default [Options].
func DefaultOptions(opts ...Option) Options {
	ops := Options{
		DumpCfg: dump.NewConfig(
			dump.WithTimeFormat(DumpTimeFormat),
			dump.WithMaxDepth(DumpDepth),
		),
		Recent:     RecentDuration,
		TimeFormat: ParseTimeFormat,
		now:        time.Now,
	}
	return ops.set(opts)
}

// set sets [Options] from slice of [Option] functions.
func (ops Options) set(opts []Option) Options {
	dst := ops
	for _, opt := range opts {
		dst = opt(dst)
	}
	return dst
}

// logTrail logs non-empty [Options.Trail] to [Options.TrailLog].
func (ops Options) logTrail() Options {
	if ops.TrailLog != nil && ops.Trail != "" {
		*ops.TrailLog = append(*ops.TrailLog, ops.Trail)
	}
	return ops
}

// structTrail updates [Options.Trail] with struct type and/or field name
// considering already existing trail.
//
// Example trails:
//
//	Type.Field
//	Type.Field.Field
//	Type.Field[1].Field
//	Type.Field["A"].Field
func (ops Options) structTrail(typeName, fldName string) Options {
	if ops.skipType {
		typeName = ""
	}
	ops.skipType = false
	next := typeName
	if typeName == "" {
		next = fldName
	}
	if next == "" {
		next = typeName
	}
	if ops.Trail == "" {
		ops.Trail = next
		return ops
	}
	if typeName != "" && ops.Trail[len(ops.Trail)-1] == ']' {
		next = fldName
	}
	if next != "" {
		ops.Trail += "." + next
		return ops
	}
	return ops
}

// mapTrail updates [Options.Trail] with map value path considering already
// existing trail.
//
// Example trails:
//
//	map[1]
//	["A"]map[1]
//	[1]map["A"]
//	field["A"]
func (ops Options) mapTrail(key string) Options {
	if ops.Trail == "" {
		ops.Trail = "map"
	}
	if ops.Trail[len(ops.Trail)-1] == ']' {
		ops.Trail += "map"
	}
	ops.Trail += "[" + key + "]"
	return ops
}

// arrTrail updates [Options.Trail] with slice or array index considering
// already existing trail.
//
// Example trails:
//
//	arr[1]
func (ops Options) arrTrail(idx int) Options {
	ops.Trail += "[" + strconv.Itoa(idx) + "]"
	return ops
}
