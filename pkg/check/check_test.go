// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"errors"
	"fmt"
	"testing"

	"github.com/ctx42/xtst/internal/affirm"
	"github.com/ctx42/xtst/internal/cases"
	"github.com/ctx42/xtst/internal/types"
)

func Test_Error(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		err := Error(errors.New("e0"))

		// --- Then ---
		affirm.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// --- When ---
		err := Error(nil)

		// --- Then ---
		affirm.NotNil(t, err)
		affirm.Equal(t, "expected non-nil error", err.Error())
	})

	t.Run("error with path option", func(t *testing.T) {
		// --- Given ---
		opt := WithPath("pth")

		// --- When ---
		err := Error(nil, opt)

		// --- Then ---
		affirm.NotNil(t, err)
		affirm.Equal(t, "expected non-nil error:\n\tpath: pth", err.Error())
	})
}

func Test_NoError(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		err := NoError(nil)

		// --- Then ---
		affirm.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// --- When ---
		err := NoError(errors.New("e0"))

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected error to be nil:\n" +
			"\twant: <nil>\n" +
			"\thave: \"e0\""
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("error with path option", func(t *testing.T) {
		// --- Given ---
		opt := WithPath("pth")

		// --- When ---
		err := NoError(errors.New("e0"), opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected error to be nil:\n" +
			"\tpath: pth\n" +
			"\twant: <nil>\n" +
			"\thave: \"e0\""
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("nil interface", func(t *testing.T) {
		// --- Given ---
		var e *types.TPtr

		// --- When ---
		err := NoError(e)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected error to be nil:\n" +
			"\twant: <nil>\n" +
			"\thave: *types.TPtr"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("nil interface with option", func(t *testing.T) {
		// --- Given ---
		var e *types.TPtr
		opt := WithPath("pth")

		// --- When ---
		err := NoError(e, opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected error to be nil:\n" +
			"\tpath: pth\n" +
			"\twant: <nil>\n" +
			"\thave: *types.TPtr"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("empty non-nil error", func(t *testing.T) {
		// --- Given ---
		e := &types.TPtr{}

		// --- When ---
		err := NoError(e)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected error to be nil:\n" +
			"\twant: <nil>\n" +
			"\thave: \"\""
		affirm.Equal(t, wMsg, err.Error())
	})
}

func Test_ErrorIs(t *testing.T) {
	t.Run("error with option", func(t *testing.T) {
		// --- Given ---
		opt := WithPath("field")

		// --- When ---
		have := ErrorIs(errors.New("err0"), errors.New("err0"), opt)

		// --- Then ---
		affirm.NotNil(t, have)
		wMsg := "expected err to have target in its tree:\n" +
			"\tpath: field\n" +
			"\twant: (*errors.errorString) err0\n" +
			"\thave: (*errors.errorString) err0"
		affirm.Equal(t, wMsg, have.Error())
	})
}

func Test_ErrorIs_success_tabular(t *testing.T) {
	err0 := errors.New("err0")
	err1 := errors.New("err1")
	err2 := fmt.Errorf("wrap: %w %w", err0, err1)

	tt := []struct {
		testN string

		err    error
		target error
	}{
		{"nil nil", nil, nil},
		{"same error", err0, err0},
		{"related 0", err2, err0},
		{"related 1", err2, err1},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			err := ErrorIs(tc.err, tc.target)

			// --- Then ---
			affirm.Nil(t, err)
		})
	}
}

func Test_ErrorIs_error_tabular(t *testing.T) {
	err0 := errors.New("err0")
	err1 := errors.New("err1")
	err2 := fmt.Errorf("wrap: %w %w", err0, err1)

	tt := []struct {
		testN string

		have     error
		haveType string
		haveStr  string
		want     error
		wantType string
		wantStr  string
	}{
		{"err nil", err0, "*errors.errorString", "err0", nil, "<nil>", "<nil>"},
		{"nil err", nil, "<nil>", "<nil>", err0, "*errors.errorString", "err0"},
		{"not related", err0, "*errors.errorString", "err0", err1, "*errors.errorString", "err1"},
		{"not related 0", err0, "*errors.errorString", "err0", err2, "*fmt.wrapErrors", "wrap: err0 err1"},
		{"not related 1", err1, "*errors.errorString", "err1", err2, "*fmt.wrapErrors", "wrap: err0 err1"},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			const format = "expected err to have target in its tree:\n" +
				"\twant: (%s) %s\n" +
				"\thave: (%s) %s"

			wantLog := fmt.Sprintf(
				format,
				tc.wantType,
				tc.wantStr,
				tc.haveType,
				tc.haveStr,
			)

			// --- When ---
			err := ErrorIs(tc.have, tc.want)

			// --- Then ---
			affirm.NotNil(t, err)
			affirm.Equal(t, wantLog, err.Error())
		})
	}
}

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
