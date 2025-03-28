// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"reflect"
	"testing"
	"time"

	"github.com/ctx42/xtst/internal"
	"github.com/ctx42/xtst/internal/affirm"
	"github.com/ctx42/xtst/pkg/dump"
)

func Test_WithTrail(t *testing.T) {
	// --- Given ---
	ops := Options{}

	// --- When ---
	have := WithTrail("type.field")(ops)

	// --- Then ---
	affirm.Equal(t, "", ops.Trail)
	affirm.Equal(t, "type.field", have.Trail)
}

func Test_WithTimeFormat(t *testing.T) {
	// --- Given ---
	ops := Options{}

	// --- When ---
	have := WithTimeFormat(time.RFC3339)(ops)

	// --- Then ---
	affirm.Equal(t, time.RFC3339, have.TimeFormat)
}

func Test_WithRecent(t *testing.T) {
	// --- Given ---
	ops := Options{}

	// --- When ---
	have := WithRecent(time.Second)(ops)

	// --- Then ---
	affirm.Equal(t, time.Second, have.Recent)
}

func Test_WithDump(t *testing.T) {
	// --- Given ---
	ops := Options{}

	// --- When ---
	have := WithDump(dump.WithMaxDepth(100))(ops)

	// --- Then ---
	affirm.Equal(t, 0, ops.DumpCfg.MaxDepth)
	affirm.Equal(t, 100, have.DumpCfg.MaxDepth)
}

func Test_DefaultOptions(t *testing.T) {
	// --- When ---
	have := DefaultOptions()

	// --- Then ---
	affirm.Equal(t, false, have.DumpCfg.PtrAddr)
	affirm.Equal(t, DefaultDumpTimeFormat, have.DumpCfg.TimeFormat)

	affirm.Equal(t, DefaultParseTimeFormat, have.TimeFormat)
	affirm.Equal(t, DefaultRecentDuration, have.Recent)
	affirm.Equal(t, "", have.Trail)
	affirm.True(t, internal.Same(time.Now, have.now))
	affirm.Equal(t, 5, reflect.ValueOf(have).NumField())
}

func Test_Options_set(t *testing.T) {
	// --- Given ---
	ops := Options{}

	// --- When ---
	have := ops.set([]Option{WithTrail("type.field")})

	// --- Then ---
	affirm.Equal(t, "", ops.Trail)
	affirm.Equal(t, "type.field", have.Trail)
}
