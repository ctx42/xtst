// Copyright (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package notice

import (
	"testing"

	"github.com/ctx42/xtst/internal/affirm"
)

func Test_NewRow(t *testing.T) {
	t.Run("without args", func(t *testing.T) {
		// --- When ---
		have := NewRow("name", "format")

		// --- Then ---
		affirm.Equal(t, "name", have.Name)
		affirm.Equal(t, "format", have.Format)
		affirm.NotNil(t, have.Args)
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

func Test_Path(t *testing.T) {
	// --- When ---
	have := Path("name")

	// --- Then ---
	affirm.Equal(t, "path", have.Name)
	affirm.Equal(t, "name", have.Format)
	affirm.NotNil(t, have.Args)
	affirm.Equal(t, 0, len(have.Args))
}
