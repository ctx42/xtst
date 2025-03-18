// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package notice

import (
	"errors"
	"testing"

	"github.com/ctx42/xtst/internal/affirm"
)

func Test_Lines(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		// --- When ---
		have := Lines(1, "")

		// --- Then ---
		affirm.Equal(t, "", have)
	})

	t.Run("one line", func(t *testing.T) {
		// --- When ---
		have := Lines(1, "abc")

		// --- Then ---
		affirm.Equal(t, "\t>| abc", have)
	})

	t.Run("multiple lines", func(t *testing.T) {
		// --- When ---
		have := Lines(1, "abc\ndef\nghi")

		// --- Then ---
		want := "" +
			"\t>| abc\n" +
			"\t>| def\n" +
			"\t>| ghi"
		affirm.Equal(t, want, have)
	})

	t.Run("no ident", func(t *testing.T) {
		// --- When ---
		have := Lines(0, "abc\ndef\nghi")

		// --- Then ---
		want := "" +
			">| abc\n" +
			">| def\n" +
			">| ghi"
		affirm.Equal(t, want, have)
	})
}

func Test_Unwrap(t *testing.T) {
	t.Run("unwrap multiple", func(t *testing.T) {
		// --- Given ---
		err0 := errors.New("e0")
		err1 := errors.New("e1")
		ers := errors.Join(err0, err1)

		// --- When ---
		have := Unwrap(ers)

		// --- Then ---
		affirm.DeepEqual(t, []error{err0, err1}, have)
	})

	t.Run("unwrap not multi error", func(t *testing.T) {
		// --- Given ---
		err := errors.New("e0")

		// --- When ---
		have := Unwrap(err)

		// --- Then ---
		affirm.DeepEqual(t, []error{err}, have)
	})

	t.Run("unwrap nil", func(t *testing.T) {
		// --- When ---
		have := Unwrap(nil)

		// --- Then ---
		affirm.NotNil(t, have)
		affirm.Equal(t, 0, len(have))
	})
}
