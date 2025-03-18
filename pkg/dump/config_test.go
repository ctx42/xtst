// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package dump

import (
	"reflect"
	"testing"
	"time"

	"github.com/ctx42/xtst/internal/affirm"
)

func Test_Flat(t *testing.T) {
	// --- Given ---
	cfg := &Config{}

	// --- When ---
	Flat(cfg)

	// --- Then ---
	affirm.True(t, cfg.Flat)
}

func Test_Compact(t *testing.T) {
	// --- Given ---
	cfg := &Config{}

	// --- When ---
	Compact(cfg)

	// --- Then ---
	affirm.True(t, cfg.Compact)
}

func Test_PtrAddr(t *testing.T) {
	// --- Given ---
	cfg := &Config{}

	// --- When ---
	PtrAddr(cfg)

	// --- Then ---
	affirm.True(t, cfg.PtrAddr)
}

func Test_TimeFormat(t *testing.T) {
	// --- Given ---
	cfg := &Config{}

	// --- When ---
	opt := TimeFormat(TimeAsUnix)

	// --- Then ---
	opt(cfg)
	affirm.Equal(t, TimeAsUnix, cfg.TimeFormat)
}

func Test_Depth(t *testing.T) {
	// --- Given ---
	cfg := &Config{}

	// --- When ---
	opt := Depth(10)

	// --- Then ---
	opt(cfg)
	affirm.Equal(t, 10, cfg.Depth)
}

func Test_PrintType(t *testing.T) {
	// --- Given ---
	cfg := &Config{}

	// --- When ---
	PrintType(cfg)

	// --- Then ---
	affirm.True(t, cfg.PrintType)
}

func Test_WithDumper(t *testing.T) {
	// --- Given ---
	cfg := Config{Dumpers: make(map[reflect.Type]Dumper)}

	// --- When ---
	WithDumper(time.Time{}, GetTimeDumper(time.Kitchen))(&cfg)

	// --- Then ---
	dmp := New(cfg).DumpAny(time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC))
	affirm.Equal(t, `"3:04AM"`, dmp)
}

func Test_NewConfig(t *testing.T) {
	// --- When ---
	have := NewConfig()

	// --- Then ---
	affirm.False(t, have.Flat)
	affirm.False(t, have.Compact)
	affirm.Equal(t, "", have.TimeFormat)
	affirm.Equal(t, "", have.DurationFormat)
	affirm.False(t, have.PtrAddr)
	affirm.True(t, have.UseAny)
	affirm.True(t, len(have.Dumpers) == 3)
	affirm.Equal(t, 6, have.Depth)

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
