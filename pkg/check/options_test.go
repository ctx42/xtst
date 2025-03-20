// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"reflect"
	"testing"

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

func Test_WithDump(t *testing.T) {
	// --- Given ---
	ops := Options{}

	// --- When ---
	have := WithDump(dump.Depth(100))(ops)

	// --- Then ---
	affirm.Equal(t, 0, ops.Depth)
	affirm.Equal(t, 100, have.Depth)
}

func Test_DefaultOptions(t *testing.T) {
	// --- When ---
	have := DefaultOptions()

	// --- Then ---
	affirm.Equal(t, "", have.Trail)
	affirm.Equal(t, false, have.PtrAddr)
	affirm.Equal(t, 2, reflect.ValueOf(have).NumField())
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
