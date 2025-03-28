// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package dump

import (
	"testing"

	"github.com/ctx42/xtst/internal/affirm"
)

func Test_newPrinter(t *testing.T) {
	// --- Given ---
	cfg := NewConfig()

	// --- When ---
	have := newPrinter(cfg)

	// --- Then ---
	affirm.NotNil(t, have.buf)
}

func Test_printer_nli(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		// --- Given ---
		cfg := Config{Flat: false}

		// --- When ---
		have := newPrinter(cfg).nli(1)

		// --- Then ---
		affirm.Equal(t, "\n", have.String())
	})

	t.Run("default with zero count", func(t *testing.T) {
		// --- Given ---
		cfg := Config{Flat: false}

		// --- When ---
		have := newPrinter(cfg).nli(0)

		// --- Then ---
		affirm.Equal(t, "", have.String())
	})

	t.Run("flat", func(t *testing.T) {
		// --- Given ---
		cfg := Config{Flat: true}

		// --- When ---
		have := newPrinter(cfg).nli(1)

		// --- Then ---
		affirm.Equal(t, "", have.String())
	})
}

func Test_printer_nl(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		// --- Given ---
		cfg := Config{Flat: false}

		// --- When ---
		have := newPrinter(cfg).nl()

		// --- Then ---
		affirm.Equal(t, "\n", have.String())
	})

	t.Run("flat", func(t *testing.T) {
		// --- Given ---
		cfg := Config{Flat: true}

		// --- When ---
		have := newPrinter(cfg).nl()

		// --- Then ---
		affirm.Equal(t, "", have.String())
	})
}

func Test_printer_comma(t *testing.T) {
	t.Run("not last and not flat", func(t *testing.T) {
		// --- Given ---
		cfg := Config{Flat: false}

		// --- When ---
		have := newPrinter(cfg).comma(false)

		// --- Then ---
		affirm.Equal(t, ",", have.String())
	})

	t.Run("not last and flat", func(t *testing.T) {
		// --- Given ---
		cfg := Config{Flat: true}

		// --- When ---
		have := newPrinter(cfg).comma(false)

		// --- Then ---
		affirm.Equal(t, ",", have.String())
	})

	t.Run("last and not flat", func(t *testing.T) {
		// --- Given ---
		cfg := Config{Flat: false}

		// --- When ---
		have := newPrinter(cfg).comma(true)

		// --- Then ---
		affirm.Equal(t, ",", have.String())
	})

	t.Run("last and flat", func(t *testing.T) {
		// --- Given ---
		cfg := Config{Flat: true}

		// --- When ---
		have := newPrinter(cfg).comma(true)

		// --- Then ---
		affirm.Equal(t, "", have.String())
	})
}

func Test_printer_tab(t *testing.T) {
	t.Run("default and positive n", func(t *testing.T) {
		// --- Given ---
		cfg := Config{Flat: false, TabWidth: 2}

		// --- When ---
		have := newPrinter(cfg).tab(2)

		// --- Then ---
		affirm.Equal(t, "    ", have.String())
	})

	t.Run("default and negative n", func(t *testing.T) {
		// --- Given ---
		cfg := Config{Flat: false}

		// --- When ---
		have := newPrinter(cfg).tab(-2)

		// --- Then ---
		affirm.Equal(t, "", have.String())
	})

	t.Run("flat", func(t *testing.T) {
		// --- Given ---
		cfg := Config{Flat: true}

		// --- When ---
		have := newPrinter(cfg).tab(2)

		// --- Then ---
		affirm.Equal(t, "", have.String())
	})
}

func Test_printer_space(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		// --- Given ---
		cfg := Config{Compact: false}

		// --- When ---
		have := newPrinter(cfg).space()

		// --- Then ---
		affirm.Equal(t, " ", have.String())
	})

	t.Run("compact", func(t *testing.T) {
		// --- Given ---
		cfg := Config{Compact: true}

		// --- When ---
		have := newPrinter(cfg).space()

		// --- Then ---
		affirm.Equal(t, "", have.String())
	})
}

func Test_printer_sep(t *testing.T) {
	t.Run("default and not last", func(t *testing.T) {
		// --- Given ---
		cfg := Config{Compact: false}

		// --- When ---
		have := newPrinter(cfg).sep(false)

		// --- Then ---
		affirm.Equal(t, "", have.String())
	})

	t.Run("default and last", func(t *testing.T) {
		// --- Given ---
		cfg := Config{Compact: false}

		// --- When ---
		have := newPrinter(cfg).sep(true)

		// --- Then ---
		affirm.Equal(t, "", have.String())
	})

	t.Run("compact and flat and last", func(t *testing.T) {
		// --- Given ---
		cfg := Config{Compact: true, Flat: true}

		// --- When ---
		have := newPrinter(cfg).sep(true)

		// --- Then ---
		affirm.Equal(t, "", have.String())
	})

	t.Run("compact and not flat and last", func(t *testing.T) {
		// --- Given ---
		cfg := Config{Compact: true, Flat: false}

		// --- When ---
		have := newPrinter(cfg).sep(true)

		// --- Then ---
		affirm.Equal(t, "", have.String())
	})

	t.Run("not compact and not flat and last", func(t *testing.T) {
		// --- Given ---
		cfg := Config{Compact: true, Flat: false}

		// --- When ---
		have := newPrinter(cfg).sep(true)

		// --- Then ---
		affirm.Equal(t, "", have.String())
	})

	t.Run("not compact and flat and not last", func(t *testing.T) {
		// --- Given ---
		cfg := Config{Compact: false, Flat: true}

		// --- When ---
		have := newPrinter(cfg).sep(false)

		// --- Then ---
		affirm.Equal(t, " ", have.String())
	})
}

func Test_printer_write_String(t *testing.T) {
	// --- Given ---
	cfg := NewConfig()

	// --- When ---
	have := newPrinter(cfg).write("test")

	// --- Then ---
	affirm.Equal(t, "test", have.String())
}
