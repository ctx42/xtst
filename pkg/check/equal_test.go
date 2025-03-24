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
	t.Run("success", func(t *testing.T) {
		// --- When ---
		err := Equal(42, 42)

		// --- Then ---
		affirm.Nil(t, err)
	})

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

	// TODO(rz):
	// t.Run("error structs", func(t *testing.T) {
	// 	// --- Given ---
	// 	s0 := struct {
	// 		Val int
	// 		Now func() time.Time
	// 	}{
	// 		Val: 1,
	// 		Now: time.Now,
	// 	}
	//
	// 	s1 := struct {
	// 		Val int
	// 		Now func() time.Time
	// 	}{
	// 		Val: 1,
	// 		Now: func() time.Time { return time.Now() }, // nolint: gocritic
	// 	}
	//
	// 	// --- When ---
	// 	err := Equal(s0, s1)
	//
	// 	// --- Then ---
	// 	affirm.NotNil(t, err)
	// 	hMsg := err.Error()
	// 	wMsg := "expected same pointers:\n" +
	// 		"\ttrail: Now\n" +
	// 		"\t want: "
	// 	affirm.True(t, strings.Contains(hMsg, wMsg))
	// })

	t.Run("error", func(t *testing.T) {
		// --- When ---
		err := Equal(42, 44)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected values to be equal:\n" +
			"\twant: 42\n" +
			"\thave: 44"
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
			"\ttrail: type.field\n" +
			"\t want: 42\n" +
			"\t have: 44"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("error different types", func(t *testing.T) {
		// --- When ---
		err := Equal(42, int64(42))

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected values to be equal:\n" +
			"\t     want: 42\n" +
			"\t     have: 42\n" +
			"\twant type: int\n" +
			"\thave type: int64"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("error printable byte", func(t *testing.T) {
		// --- When ---
		err := Equal(byte('B'), byte('A'))

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected values to be equal:\n" +
			"\twant: 0x42 ('B')\n" +
			"\thave: 0x41 ('A')"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("error not printable byte", func(t *testing.T) {
		// --- When ---
		err := Equal(byte(1), byte(2))

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected values to be equal:\n" +
			"\twant: 0x1\n" +
			"\thave: 0x2"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("error nested slice of value types", func(t *testing.T) {
		// --- Given ---
		s0 := types.TNested{STA: []types.TA{{Str: "abc"}}}
		s1 := types.TNested{STA: []types.TA{{Str: "xyz"}}}

		// --- When ---
		err := Equal(s0, s1)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected values to be equal:\n" +
			"\ttrail: TNested.STA[0].Str\n" +
			"\t want: \"abc\"\n" +
			"\t have: \"xyz\""
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("error nested slice of pointer types", func(t *testing.T) {
		// --- Given ---
		s0 := types.TNested{STAp: []*types.TA{{Str: "abc"}}}
		s1 := types.TNested{STAp: []*types.TA{{Str: "xyz"}}}

		// --- When ---
		err := Equal(s0, s1)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected values to be equal:\n" +
			"\ttrail: TNested.STAp[0].Str\n" +
			"\t want: \"abc\"\n" +
			"\t have: \"xyz\""
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("error deep nested", func(t *testing.T) {
		// --- Given ---
		s0 := types.TNested{STAp: []*types.TA{{TAp: &types.TA{Str: "abc"}}}}
		s1 := types.TNested{STAp: []*types.TA{{TAp: &types.TA{Str: "xyz"}}}}

		// --- When ---
		err := Equal(s0, s1)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected values to be equal:\n" +
			"\ttrail: TNested.STAp[0].TAp.Str\n" +
			"\t want: \"abc\"\n" +
			"\t have: \"xyz\""
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
			"\ttrail: map[\"A\"]\n" +
			"\t want: 1\n" +
			"\t have: 2"
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
		wMsg := "expected values to be equal:\n" +
			"\ttrail: [0]\n" +
			"\t want: 1\n" +
			"\t have: 2\n" +
			"expected values to be equal:\n" +
			"\ttrail: [1]\n" +
			"\t want: 2\n" +
			"\t have: 3"
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
		wMsg := "expected values to be equal:\n" +
			"\ttrail: map[\"A\"]\n" +
			"\t want: 1\n" +
			"\t have: 2\n" +
			"expected values to be equal:\n" +
			"\ttrail: map[\"B\"]\n" +
			"\t want: 2\n" +
			"\t have: 3"
		affirm.Equal(t, wMsg, err.Error())
	})
}

func Test_Equal_EqualCases(t *testing.T) {
	for _, tc := range cases.EqualCases() {
		t.Run("Equal "+tc.Desc, func(t *testing.T) {
			// --- When ---
			have := Equal(tc.Val0, tc.Val1)

			// --- Then ---
			if tc.AreEqual && have != nil {
				format := "expected nil error:\n\thave: %#v"
				t.Errorf(format, have)
			}
			if !tc.AreEqual && have == nil {
				format := "expected not-nil error:\n\thave: %#v"
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
			"\twant: 42\n" +
			"\thave: 44"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("with path", func(t *testing.T) {
		// --- Given ---
		ops := DefaultOptions(WithTrail("type.field"))

		// --- When ---
		err := equalError(42, 44, ops)

		// --- Then ---
		wMsg := "expected values to be equal:\n" +
			"\ttrail: type.field\n" +
			"\t want: 42\n" +
			"\t have: 44"
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
			"\twant: 0x41 ('A')\n" +
			"\thave: 0x42 ('B')"
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
			"\t     want: 0x41 ('A')\n" +
			"\t     have: 42\n" +
			"\twant type: uint8\n" +
			"\thave type: int"
		affirm.Equal(t, wMsg, err.Error())
	})
}
