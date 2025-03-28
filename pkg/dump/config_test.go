// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package dump

import (
	"reflect"
	"testing"
	"time"

	"github.com/ctx42/xtst/internal/affirm"
)

func Test_WithFlat(t *testing.T) {
	// --- Given ---
	cfg := &Config{}

	// --- When ---
	WithFlat(cfg)

	// --- Then ---
	affirm.True(t, cfg.Flat)
}

func Test_WithCompact(t *testing.T) {
	// --- Given ---
	cfg := &Config{}

	// --- When ---
	WithCompact(cfg)

	// --- Then ---
	affirm.True(t, cfg.Compact)
}

func Test_WithPtrAddr(t *testing.T) {
	// --- Given ---
	cfg := &Config{}

	// --- When ---
	WithPtrAddr(cfg)

	// --- Then ---
	affirm.True(t, cfg.PtrAddr)
}

func Test_WithTimeFormat(t *testing.T) {
	// --- Given ---
	cfg := &Config{}

	// --- When ---
	opt := WithTimeFormat(TimeAsUnix)

	// --- Then ---
	opt(cfg)
	affirm.Equal(t, TimeAsUnix, cfg.TimeFormat)
}

func Test_WithMaxDepth(t *testing.T) {
	// --- Given ---
	cfg := &Config{}

	// --- When ---
	opt := WithMaxDepth(10)

	// --- Then ---
	opt(cfg)
	affirm.Equal(t, 10, cfg.MaxDepth)
}

func Test_WithIndent(t *testing.T) {
	// --- Given ---
	cfg := &Config{}

	// --- When ---
	opt := WithIndent(10)

	// --- Then ---
	opt(cfg)
	affirm.Equal(t, 10, cfg.Indent)
}

func Test_WithTabWidth(t *testing.T) {
	// --- Given ---
	cfg := &Config{}

	// --- When ---
	opt := WithTabWidth(10)

	// --- Then ---
	opt(cfg)
	affirm.Equal(t, 10, cfg.TabWidth)
}

func Test_WithDumper(t *testing.T) {
	// --- Given ---
	cfg := Config{Dumpers: make(map[reflect.Type]Dumper)}

	// --- When ---
	WithDumper(time.Time{}, GetTimeDumper(time.Kitchen))(&cfg)

	// --- Then ---
	dmp := New(cfg).Any(time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC))
	affirm.Equal(t, `"3:04AM"`, dmp)
}

func Test_NewConfig(t *testing.T) {
	// --- When ---
	have := NewConfig()

	// --- Then ---
	affirm.False(t, have.Flat)
	affirm.False(t, have.Compact)
	affirm.Equal(t, TimeFormat, have.TimeFormat)
	affirm.Equal(t, "", have.DurationFormat)
	affirm.False(t, have.PtrAddr)
	affirm.True(t, have.UseAny)
	affirm.True(t, len(have.Dumpers) == 3)
	affirm.Equal(t, DefaultDepth, have.MaxDepth)
	affirm.Equal(t, DefaultIndent, have.Indent)
	affirm.Equal(t, DefaultTabWith, have.TabWidth)

	val, ok := have.Dumpers[typDur]
	affirm.True(t, ok)
	affirm.NotNil(t, val)

	val, ok = have.Dumpers[typLocation]
	affirm.True(t, ok)
	affirm.NotNil(t, val)

	val, ok = have.Dumpers[typTime]
	affirm.True(t, ok)
	affirm.NotNil(t, val)
}
