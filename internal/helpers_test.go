// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package internal

import (
	"testing"

	"github.com/ctx42/testing/internal/affirm"
	"github.com/ctx42/testing/internal/types"
)

func Test_Same_tabular(t *testing.T) {
	// Pointers.
	ptr0 := &types.TPtr{Val: "0"}
	ptr1 := &types.TPtr{Val: "1"}

	// Interfaces.
	var itfPtr0, itfPtr1 types.TItf
	itfPtr0, itfPtr1 = &types.TPtr{Val: "0"}, &types.TPtr{Val: "1"}

	// Functions.
	fn0 := func() {}
	fn1 := func() {}
	type TFn0 func()
	type TFn1 func()
	var fnNil0 TFn0
	var fnNil1 TFn1
	var fnA, fnB TFn0

	// Slices.
	s0 := []int{1, 2, 3}
	s1 := []int{1, 2, 3}
	var sNil0 []int
	var sNil1 []float64
	type Ts []int
	var sA, sB Ts

	// Arrays.
	a0 := [2]int{1, 2}
	a1 := [2]int{1, 2}
	var aNil0 [2]int
	var aNil1 [2]float64
	type Ta []int
	var aA, aB Ta

	// Maps.
	m0 := map[string]int{"a": 1, "b": 2}
	m1 := map[string]int{"a": 1, "b": 2}
	var mNil0 map[string]int
	var mNil1 map[string]float64
	type Tm map[string]int
	var mA, mB Tm

	// Channels.
	c0 := make(chan int)
	c1 := make(chan int)
	var cNil0 chan int
	type Tc chan int
	var cNilA chan byte
	var cA, cB Tc

	tt := []struct {
		testN string

		want any
		have any
		same bool
	}{
		{"ptr same", ptr0, ptr0, true},
		{"ptr not same", ptr0, ptr1, false},
		{"ptr different types not same", &types.TPtr{}, &types.TVal{}, false},
		{"prt nil both", nil, nil, false},

		{"itf ptr same", itfPtr0, itfPtr0, true},
		{"itf ptr not same", itfPtr0, itfPtr1, false},
		{"obj val not same", types.TVal{}, types.TVal{}, false},

		{"func same", fn0, fn0, true},
		{"func not same", fn0, fn1, false},
		{"func nil want", nil, fn1, false},
		{"func nil type want", fnNil0, fn1, false},
		{"func nil have", fn0, nil, false},
		{"func nil type have", fn0, fnNil0, false},
		{"func nil type both", fnNil0, fnNil0, true},
		{"func nil different type both", fnNil0, fnNil1, false},
		{"func nil same type both", fnA, fnB, true},

		{"slice same", s0, s0, true},
		{"slice not same", s0, s1, false},
		{"slice nil want", nil, s1, false},
		{"slice nil var want", sNil0, s1, false},
		{"slice nil have", s0, nil, false},
		{"slice nil var have", s0, sNil0, false},
		{"slice nil var both", sNil0, sNil0, true},
		{"slice nil different type both", sNil0, sNil1, false},
		{"slice nil same type both", sA, sB, true},

		{"array same", a0, a0, false},
		{"array not same", a0, a1, false},
		{"array nil want", nil, a1, false},
		{"array nil var want", aNil0, a1, false},
		{"array nil have", a0, nil, false},
		{"array nil var have", a0, aNil0, false},
		{"array nil var both", aNil0, aNil0, false},
		{"array nil different type both", aNil0, aNil1, false},
		{"array nil same type both", aA, aB, true},

		{"map same", m0, m0, true},
		{"map not same", m0, m1, false},
		{"map nil want", nil, s1, false},
		{"map nil var want", mNil0, m1, false},
		{"map nil have", m0, nil, false},
		{"map nil var have", m0, mNil0, false},
		{"map nil both", nil, nil, false},
		{"map nil var both", mNil0, mNil0, true},
		{"map nil different type both", mNil0, mNil1, false},
		{"map nil same type both", mA, mB, true},

		{"chanel same", c0, c0, true},
		{"chanel not same", c0, c1, false},
		{"chanel nil want", nil, c1, false},
		{"chanel nil var want", mNil0, c1, false},
		{"chanel nil have", c0, nil, false},
		{"chanel nil var have", c0, cNil0, false},
		{"chanel nil both", nil, nil, false},
		{"chanel nil var both", cNil0, cNil0, true},
		{"chanel nil different type both", cNil0, cNilA, false},
		{"chanel nil same type both", cA, cB, true},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have := Same(tc.want, tc.have)

			// --- Then ---
			affirm.DeepEqual(t, tc.same, have)
		})
	}
}

func Test_Len_tabular(t *testing.T) {
	ch := make(chan int, 10)
	ch <- 0
	t.Cleanup(func() { <-ch; close(ch) })

	tt := []struct {
		testN string

		val  any
		want int
		ok   bool
	}{
		{"success string", "abc", 3, true},
		{"success slice", []int{1, 2}, 2, true},
		{"success array", [...]int{1, 2}, 2, true},
		{"success channel", ch, 1, true},

		{"error not supported", 123, 0, false},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			haveLen, haveOK := Len(tc.val)

			// --- Then ---
			affirm.Equal(t, tc.want, haveLen)
			affirm.Equal(t, tc.ok, haveOK)
		})
	}
}

func Test_DidPanic(t *testing.T) {
	t.Run("panicked", func(t *testing.T) {
		// --- Given ---
		fn := func() { panic("panic") }

		// --- When ---
		did, val, stack := DidPanic(fn)

		// --- Then ---
		affirm.True(t, did)
		affirm.Equal(t, "panic", val)
		affirm.True(t, stack != "")
	})

	t.Run("no panic", func(t *testing.T) {
		// --- Given ---
		fn := func() {}

		// --- When ---
		did, val, stack := DidPanic(fn)

		// --- Then ---
		affirm.False(t, did)
		affirm.Nil(t, val)
		affirm.Equal(t, 0, len(stack))
	})
}
