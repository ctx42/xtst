// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package dump

import (
	"reflect"
	"time"
)

// Package wide default configuration.
const (
	// DefaultTimeFormat is default format for parsing time strings.
	DefaultTimeFormat = time.RFC3339Nano

	// DefaultDepth is default depth when dumping values recursively.
	DefaultDepth = 6
)

// Package wide configuration.
var (
	// TimeFormat is configurable format for dumping [time.Time] values.
	TimeFormat = DefaultTimeFormat

	// Depth is configurable depth when dumping values recursively.
	Depth = DefaultDepth
)

// Option represents [NewConfig] option.
type Option func(*Config)

// WithFlat is option for [NewConfig] which makes [Dump] display values in one line.
func WithFlat(cfg *Config) { cfg.Flat = true }

// WithCompact is option for [NewConfig] which makes [Dump] display values without
// unnecessary whitespaces.
func WithCompact(cfg *Config) { cfg.Compact = true }

// WithPtrAddr is option for [NewConfig] which makes [Dump] display pointer
// addresses.
func WithPtrAddr(cfg *Config) { cfg.PtrAddr = true }

// WithTimeFormat is option for [NewConfig] which makes [Dump] display [time.Time]
// using given format. The format might be standard Go time formating layout or
// one of the custom values - see [Config.TimeFormat] for more details.
func WithTimeFormat(format string) func(cfg *Config) {
	return func(cfg *Config) { cfg.TimeFormat = format }
}

// WithPrintType is option for [NewConfig] which makes [Dump] print types.
func WithPrintType(cfg *Config) { cfg.PrintType = true }

// WithDumper adds custom [Dumper] to the config.
func WithDumper(typ any, dumper Dumper) func(cfg *Config) {
	return func(cfg *Config) { cfg.Dumpers[reflect.TypeOf(typ)] = dumper }
}

// WithDepth is option for [NewConfig] which controls maximum nesting when bumping
// recursive types.
func WithDepth(maximum int) func(cfg *Config) {
	return func(cfg *Config) { cfg.Depth = maximum }
}

// Config represents [Dump] configuration.
type Config struct {
	// Display values on one line.
	Flat bool

	// Do not use any indents or whitespace separators.
	Compact bool

	// Controls how [time.Time] is formated.
	//
	// Aside form Go time formating layouts following custom formats are
	// available:
	//
	//  - [TimeAsUnix] - Unix timestamp,
	//
	// By default (empty value) [time.RFC3339Nano] is used.
	TimeFormat string

	// Controls how [time.Duration] is formated.
	//
	// Supports formats:
	//
	//  - [DurAsString]
	//  - [DurAsSeconds]
	DurationFormat string

	// Show pointer addresses.
	PtrAddr bool

	// Print types.
	PrintType bool

	// Use "any" instead of "interface{}".
	UseAny bool

	// Custom type dumpers.
	//
	// By default, dumpers for types:
	//   - [time.Time]
	//   - [time.Duration]
	//   - [time.Location]
	//
	// are automatically registered.
	Dumpers map[reflect.Type]Dumper

	// Controls maximum nesting when dumping recursive types.
	Depth int
}

// NewConfig returns new instance of [Config] with default values.
func NewConfig(opts ...Option) Config {
	cfg := Config{
		TimeFormat: TimeFormat,
		PrintType:  true,
		UseAny:     true,
		Dumpers:    make(map[reflect.Type]Dumper),
		Depth:      Depth,
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	if _, ok := cfg.Dumpers[typTime]; !ok {
		cfg.Dumpers[typTime] = GetTimeDumper(cfg.TimeFormat)
	}

	if _, ok := cfg.Dumpers[typLocation]; !ok {
		cfg.Dumpers[typLocation] = zoneDumper
	}

	if _, ok := cfg.Dumpers[typDur]; !ok {
		cfg.Dumpers[typDur] = GetDurDumper(cfg.DurationFormat)
	}

	return cfg
}
