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

	// DefaultIndent is default additional indent when dumping values.
	DefaultIndent = 0

	// DefaultTabWith is default tab width in spaces.
	DefaultTabWith = 2
)

// Package wide configuration.
var (
	// TimeFormat is configurable format for dumping [time.Time] values.
	TimeFormat = DefaultTimeFormat

	// Depth is configurable depth when dumping values recursively.
	Depth = DefaultDepth

	// Indent is configurable additional indent when dumping values.
	Indent = DefaultIndent

	// TabWidth is configurable tab width in spaces.
	TabWidth = DefaultTabWith
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

// WithDumper adds custom [Dumper] to the config.
func WithDumper(typ any, dumper Dumper) func(cfg *Config) {
	return func(cfg *Config) { cfg.Dumpers[reflect.TypeOf(typ)] = dumper }
}

// WithMaxDepth is option for [NewConfig] which controls maximum nesting when
// bumping recursive types.
func WithMaxDepth(maximum int) func(cfg *Config) {
	return func(cfg *Config) { cfg.MaxDepth = maximum }
}

// WithIndent is option for [NewConfig] which how much additional indentation
// to apply to dumped values.
func WithIndent(n int) func(cfg *Config) {
	return func(cfg *Config) { cfg.Indent = n }
}

// WithTabWidth is option for [NewConfig] setting tab width in spaces.
func WithTabWidth(n int) func(cfg *Config) {
	return func(cfg *Config) { cfg.TabWidth = n }
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
	// The depth is also used to properly indent values being dumped.
	MaxDepth int

	// How much additional indentation to apply to values being dumped.
	Indent int

	// Default tab with in spaces.
	TabWidth int
}

// NewConfig returns new instance of [Config] with default values.
func NewConfig(opts ...Option) Config {
	cfg := Config{
		TimeFormat: TimeFormat,
		PrintType:  true,
		UseAny:     true,
		Dumpers:    make(map[reflect.Type]Dumper),
		MaxDepth:   Depth,
		Indent:     Indent,
		TabWidth:   TabWidth,
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
