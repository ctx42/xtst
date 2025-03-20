// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"reflect"
	"testing"

	"github.com/ctx42/xtst/internal/affirm"
	"github.com/ctx42/xtst/pkg/dump"
)

func Test_WithPath(t *testing.T) {
	// --- Given ---
	ops := Options{}

	// --- When ---
	have := WithPath("pth")(ops)

	// --- Then ---
	affirm.Equal(t, "", ops.Path)
	affirm.Equal(t, "pth", have.Path)
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
	affirm.Equal(t, "", have.Path)
	affirm.Equal(t, false, have.PtrAddr)
	affirm.Equal(t, 2, reflect.ValueOf(have).NumField())
}

func Test_Options_set(t *testing.T) {
	// --- Given ---
	ops := Options{}

	// --- When ---
	have := ops.set([]Option{WithPath("pth")})

	// --- Then ---
	affirm.Equal(t, "", ops.Path)
	affirm.Equal(t, "pth", have.Path)
}
