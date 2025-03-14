// Copyright (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package affirm

import (
	"errors"
	"io"
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
	t.Run("success", func(t *testing.T) {
		ti := &testing.T{}
		have := Panic(ti, "abc", func() { panic("abc") })
		if !have || ti.Failed() {
			t.Error("expected success")
		}
	})

	t.Run("not matching message", func(t *testing.T) {
		ti := &testing.T{}
		have := Panic(ti, "xyz", func() { panic("abc") })
		if have || !ti.Failed() {
			t.Error("expected failure")
		}
	})

	t.Run("no panic", func(t *testing.T) {
		ti := &testing.T{}
		have := Panic(ti, "xyz", func() {})
		if have || !ti.Failed() {
			t.Error("expected failure")
		}
	})

	t.Run("not string panic message", func(t *testing.T) {
		ti := &testing.T{}
		have := Panic(ti, "xyz", func() { panic(io.ErrShortWrite) })
		if have || !ti.Failed() {
			t.Error("expected failure")
		}
	})
}
