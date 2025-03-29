// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"testing"
	"time"

	"github.com/ctx42/xtst/internal/affirm"
	"github.com/ctx42/xtst/internal/cases"
	"github.com/ctx42/xtst/internal/types"
)

// TODO(rz): make very detailed code review of this file.

func Test_Equal(t *testing.T) {
	t.Run("success structs", func(t *testing.T) {
		// --- Given ---
		s0 := struct {
			Val int
			Now func() time.Time
		}{
			Val: 1,
			Now: time.Now,
		}

		s1 := struct {
			Val int
			Now func() time.Time
		}{
			Val: 1,
			Now: time.Now,
		}

		// --- When ---
		err := Equal(s0, s1)

		// --- Then ---
		affirm.Nil(t, err)
	})

	t.Run("error structs", func(t *testing.T) {
		// --- Given ---
		s0 := struct {
			Val int
			Now func() time.Time
		}{
			Val: 1,
			Now: time.Now,
		}

		s1 := struct {
			Val int
			Now func() time.Time
		}{
			Val: 1,
			Now: func() time.Time { return time.Now() }, // nolint: gocritic
		}

		// --- When ---
		err := Equal(s0, s1)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected values to be equal:\n" +
			"  trail: Now\n" +
			"   want: <func>(<addr>)\n" +
			"   have: <func>(<addr>)"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("error simple", func(t *testing.T) {
		// --- When ---
		err := Equal(42, 44)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected values to be equal:\n" +
			"  want: 42\n" +
			"  have: 44"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("error with option", func(t *testing.T) {
		// --- Given ---
		opt := WithTrail("type.field")

		// --- When ---
		err := Equal(42, 44, opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected values to be equal:\n" +
			"  trail: type.field\n" +
			"   want: 42\n" +
			"   have: 44"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("error different types", func(t *testing.T) {
		// --- When ---
		err := Equal(42, int64(42))

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected values to be equal:\n" +
			"       want: 42\n" +
			"       have: 42\n" +
			"  want type: int\n" +
			"  have type: int64"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("error not printable byte", func(t *testing.T) {
		// --- When ---
		err := Equal(byte(1), byte(2))

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected values to be equal:\n" +
			"  want: 0x1\n" +
			"  have: 0x2"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("error printable byte", func(t *testing.T) {
		// --- When ---
		err := Equal(byte('B'), byte('A'))

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected values to be equal:\n" +
			"  want: 0x42 ('B')\n" +
			"  have: 0x41 ('A')"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("error nested slice of value types", func(t *testing.T) {
		// --- Given ---
		s0 := types.TNested{STA: []types.TA{{Int: 42}}}
		s1 := types.TNested{STA: []types.TA{{Int: 44}}}

		// --- When ---
		err := Equal(s0, s1)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected values to be equal:\n" +
			"  trail: TNested.STA[0].Int\n" +
			"   want: 42\n" +
			"   have: 44"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("error nested slice of pointer types", func(t *testing.T) {
		// --- Given ---
		s0 := types.TNested{STAp: []*types.TA{{Int: 42}}}
		s1 := types.TNested{STAp: []*types.TA{{Int: 44}}}

		// --- When ---
		err := Equal(s0, s1)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected values to be equal:\n" +
			"  trail: TNested.STAp[0].Int\n" +
			"   want: 42\n" +
			"   have: 44"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("error deep nested", func(t *testing.T) {
		// --- Given ---
		s0 := types.TNested{STAp: []*types.TA{{TAp: &types.TA{Int: 42}}}}
		s1 := types.TNested{STAp: []*types.TA{{TAp: &types.TA{Int: 44}}}}

		// --- When ---
		err := Equal(s0, s1)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected values to be equal:\n" +
			"  trail: TNested.STAp[0].TAp.Int\n" +
			"   want: 42\n" +
			"   have: 44"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("error map", func(t *testing.T) {
		// --- Given ---
		m0 := map[string]int{"A": 1}
		m1 := map[string]int{"A": 2}

		// --- When ---
		err := Equal(m0, m1)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected values to be equal:\n" +
			"  trail: map[\"A\"]\n" +
			"   want: 1\n" +
			"   have: 2"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("multiple errors struct", func(t *testing.T) {
		// --- Given ---
		s0 := types.TNested{STA: []types.TA{{Int: 42, Str: "abc"}}}
		s1 := types.TNested{STA: []types.TA{{Int: 44, Str: "xyz"}}}

		// --- When ---
		err := Equal(s0, s1)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "" +
			"expected values to be equal:\n" +
			"  trail: TNested.STA[0].Int\n" +
			"   want: 42\n" +
			"   have: 44\n" +
			" ---\n" +
			"  trail: TNested.STA[0].Str\n" +
			"   want: \"abc\"\n" +
			"   have: \"xyz\""
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("multiple errors slice", func(t *testing.T) {
		// --- Given ---
		s0 := []int{1, 2}
		s1 := []int{2, 3}

		// --- When ---
		err := Equal(s0, s1)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "" +
			"expected values to be equal:\n" +
			"  trail: [0]\n" +
			"   want: 1\n" +
			"   have: 2\n" +
			" ---\n" +
			"  trail: [1]\n" +
			"   want: 2\n" +
			"   have: 3"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("multiple errors map", func(t *testing.T) {
		// --- Given ---
		m0 := map[string]int{"A": 1, "B": 2}
		m1 := map[string]int{"A": 2, "B": 3}

		// --- When ---
		err := Equal(m0, m1)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "" +
			"expected values to be equal:\n" +
			"  trail: map[\"A\"]\n" +
			"   want: 1\n" +
			"   have: 2\n" +
			" ---\n" +
			"  trail: map[\"B\"]\n" +
			"   want: 2\n" +
			"   have: 3"
		affirm.Equal(t, wMsg, err.Error())
	})
}

func Test_Equal_EqualCases_tabular(t *testing.T) {
	for _, tc := range cases.EqualCases() {
		t.Run("Equal "+tc.Desc, func(t *testing.T) {
			// --- When ---
			have := Equal(tc.Val0, tc.Val1)

			// --- Then ---
			if tc.AreEqual && have != nil {
				format := "expected nil error:\n  have: %#v"
				t.Errorf(format, have)
			}
			if !tc.AreEqual && have == nil {
				format := "expected not-nil error:\n  have: %#v"
				t.Errorf(format, have)
			}
		})
	}
}

func Test_equalError(t *testing.T) {
	t.Run("without path", func(t *testing.T) {
		// --- Given ---
		ops := DefaultOptions()

		// --- When ---
		err := equalError(42, 44, ops)

		// --- Then ---
		wMsg := "expected values to be equal:\n" +
			"  want: 42\n" +
			"  have: 44"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("with path", func(t *testing.T) {
		// --- Given ---
		ops := DefaultOptions(WithTrail("type.field"))

		// --- When ---
		err := equalError(42, 44, ops)

		// --- Then ---
		wMsg := "expected values to be equal:\n" +
			"  trail: type.field\n" +
			"   want: 42\n" +
			"   have: 44"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("printable byte", func(t *testing.T) {
		// --- Given ---
		w := byte('A')
		h := byte('B')
		ops := DefaultOptions()

		// --- When ---
		err := equalError(w, h, ops)

		// --- Then ---
		wMsg := "expected values to be equal:\n" +
			"  want: 0x41 ('A')\n" +
			"  have: 0x42 ('B')"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("different types", func(t *testing.T) {
		// --- Given ---
		w := byte('A')
		h := 42
		ops := DefaultOptions()

		// --- When ---
		err := equalError(w, h, ops)

		// --- Then ---
		wMsg := "expected values to be equal:\n" +
			"       want: 0x41 ('A')\n" +
			"       have: 42\n" +
			"  want type: uint8\n" +
			"  have type: int"
		affirm.Equal(t, wMsg, err.Error())
	})
}
