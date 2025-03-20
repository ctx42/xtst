// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"testing"

	"github.com/ctx42/xtst/internal/affirm"
	"github.com/ctx42/xtst/internal/cases"
)

func Test_Nil(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		err := Nil(nil)

		// --- Then ---
		affirm.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// --- When ---
		err := Nil(42)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected value to be nil:\n" +
			"\twant: <nil>\n" +
			"\thave: 42"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("error with path option", func(t *testing.T) {
		// --- Given ---
		opt := WithPath("pth")

		// --- When ---
		err := Nil(42, opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected value to be nil:\n" +
			"\tpath: pth\n" +
			"\twant: <nil>\n" +
			"\thave: 42"
		affirm.Equal(t, wMsg, err.Error())
	})
}

func Test_Nil_ZENValues(t *testing.T) {
	for _, tc := range cases.ZENValues() {
		t.Run("Nil "+tc.Desc, func(t *testing.T) {
			// --- When ---
			have := Nil(tc.Val)

			// --- Then ---
			if tc.IsNil && have != nil {
				format := "expected nil error:\n\thave: %#v"
				t.Errorf(format, have)
			}
			if !tc.IsNil && have == nil {
				format := "expected not-nil error:\n\thave: %#v"
				t.Errorf(format, have)
			}
		})
	}
}

func Test_NotNil(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		err := NotNil(42)

		// --- Then ---
		affirm.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// --- When ---
		err := NotNil(nil)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected non-nil value"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("error with path option", func(t *testing.T) {
		// --- Given ---
		opt := WithPath("pth")

		// --- When ---
		err := NotNil(nil, opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected non-nil value:\n\tpath: pth"
		affirm.Equal(t, wMsg, err.Error())
	})
}
