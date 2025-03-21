// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"reflect"
	"testing"
	"time"

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

func Test_WithDump(t *testing.T) {
	// --- Given ---
	ops := Options{}

	// --- When ---
	have := WithDump(dump.Depth(100))(ops)

	// --- Then ---
	affirm.Equal(t, 0, ops.DumpCfg.Depth)
	affirm.Equal(t, 100, have.DumpCfg.Depth)
}

func Test_DefaultOptions(t *testing.T) {
	// --- When ---
	have := DefaultOptions()

	// --- Then ---
	affirm.Equal(t, false, have.DumpCfg.PtrAddr)
	affirm.Equal(t, time.RFC3339Nano, have.DumpCfg.TimeFormat)

	affirm.Equal(t, time.RFC3339Nano, have.TimeFormat)
	affirm.Equal(t, "", have.Trail)
	affirm.Equal(t, 3, reflect.ValueOf(have).NumField())
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
