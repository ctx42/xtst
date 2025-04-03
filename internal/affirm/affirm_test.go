// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package affirm

import (
	"errors"
	"runtime"
	"testing"
)

func Test_True(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		ti := &testing.T{}

		// --- When ---
		have := True(ti, true)

		// --- Then ---
		if !have || ti.Failed() {
			t.Error("expected success")
		}
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		ti := &testing.T{}

		// --- When ---
		have := True(ti, false)

		// --- Then ---
		if have || !ti.Failed() {
			t.Error("expected failure")
		}
	})
}

func Test_False(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		ti := &testing.T{}

		// --- When ---
		have := False(ti, false)

		// --- Then ---
		if !have || ti.Failed() {
			t.Error("expected success")
		}
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		ti := &testing.T{}

		// --- When ---
		have := False(ti, true)

		// --- Then ---
		if have || !ti.Failed() {
			t.Error("expected failure")
		}
	})
}

func Test_Equal(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		ti := &testing.T{}

		// --- When ---
		have := Equal(ti, 42, 42)

		// --- Then ---
		if !have || ti.Failed() {
			t.Error("expected success")
		}
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		ti := &testing.T{}

		// --- When ---
		have := Equal(ti, 42, 44)

		// --- Then ---
		if have || !ti.Failed() {
			t.Error("expected failure")
		}
	})
}

func Test_DeepEqual(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		ti := &testing.T{}

		// --- When ---
		have := DeepEqual(ti, []int{42}, []int{42})

		// --- Then ---
		if !have || ti.Failed() {
			t.Error("expected success")
		}
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		ti := &testing.T{}

		// --- When ---
		have := DeepEqual(ti, []int{42}, []int{44})

		// --- Then ---
		if have || !ti.Failed() {
			t.Error("expected failure")
		}
	})
}

func Test_Nil(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		ti := &testing.T{}

		// --- When ---
		have := Nil(ti, nil)

		// --- Then ---
		if !have || ti.Failed() {
			t.Error("expected success")
		}
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		ti := &testing.T{}

		// --- When ---
		have := Nil(ti, errors.New("m0"))

		// --- Then ---
		if have || !ti.Failed() {
			t.Error("expected failure")
		}
	})
}

func Test_NotNil(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		ti := &testing.T{}

		// --- When ---
		have := NotNil(ti, errors.New("m0"))

		// --- Then ---
		if !have || ti.Failed() {
			t.Error("expected success")
		}
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		ti := &testing.T{}

		// --- When ---
		have := NotNil(ti, nil)

		// --- Then ---
		if have || !ti.Failed() {
			t.Error("expected failure")
		}
	})
}

func Test_Panic(t *testing.T) {
	t.Run("success string message", func(t *testing.T) {
		// --- Given ---
		ti := &testing.T{}

		// --- When ---
		have := Panic(ti, func() { panic("abc") })

		// --- Then ---
		if ti.Failed() {
			t.Error("expected not failed test")
		}
		if *have != "abc" {
			t.Errorf("expected panic message 'abc', got '%s'", *have)
		}
	})

	t.Run("success error", func(t *testing.T) {
		// --- Given ---
		ti := &testing.T{}

		// --- When ---
		have := Panic(ti, func() { panic(errors.New("abc")) })

		// --- Then ---
		if ti.Failed() {
			t.Error("expected not failed test")
		}
		if *have != "abc" {
			t.Errorf("expected panic message 'abc', got '%s'", *have)
		}
	})

	t.Run("success other type", func(t *testing.T) {
		// --- Given ---
		ti := &testing.T{}

		// --- When ---
		have := Panic(ti, func() { panic(123) })

		// --- Then ---
		if ti.Failed() {
			t.Error("expected not failed test")
		}
		if *have != "123" {
			t.Errorf("expected panic message '123', got '%s'", *have)
		}
	})

	t.Run("success function panics with nil", func(t *testing.T) {
		// --- Given ---
		ti := &testing.T{}

		// --- When ---
		have := Panic(ti, func() { panic(nil) }) // nolint: govet

		// --- Then ---

		if ti.Failed() {
			t.Error("expected not failed test")
		}
		want := (&runtime.PanicNilError{}).Error()
		if *have != want {
			wMsg := "expected panic message:\n" +
				"want: %s\n" +
				"have: %s\n"
			t.Errorf(wMsg, want, *have)
		}
	})

	t.Run("error function does not panic", func(t *testing.T) {
		// --- Given ---
		ti := &testing.T{}

		// --- When ---
		have := Panic(ti, func() {})

		// --- Then ---
		if !ti.Failed() {
			t.Error("expected failed test")
		}
		if have != nil {
			t.Errorf("expected empty panic message, got '%s'", *have)
		}
	})
}
