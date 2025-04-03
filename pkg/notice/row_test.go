// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package notice

import (
	"testing"

	"github.com/ctx42/testing/internal/affirm"
)

func Test_NewRow(t *testing.T) {
	t.Run("without args", func(t *testing.T) {
		// --- When ---
		have := NewRow("name", "format")

		// --- Then ---
		affirm.Equal(t, "name", have.Name)
		affirm.Equal(t, "format", have.Format)
		affirm.Nil(t, have.Args)
		affirm.Equal(t, 0, len(have.Args))
	})

	t.Run("with args", func(t *testing.T) {
		// --- When ---
		have := NewRow("name", "format", "a", "b", "c")

		// --- Then ---
		affirm.Equal(t, "name", have.Name)
		affirm.Equal(t, "format", have.Format)
		affirm.DeepEqual(t, []any{"a", "b", "c"}, have.Args)
	})
}

func Test_Trail(t *testing.T) {
	// --- When ---
	have := Trail("type.field")

	// --- Then ---
	affirm.Equal(t, trail, have.Name)
	affirm.Equal(t, "type.field", have.Format)
	affirm.Nil(t, have.Args)
	affirm.Equal(t, 0, len(have.Args))
}
