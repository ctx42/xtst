// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/ctx42/xtst/internal/affirm"
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

	t.Run("log message with trail", func(t *testing.T) {
		// --- Given ---
		opt := WithTrail("type.field")

		// --- When ---
		err := Error(nil, opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected non-nil error:\n\ttrail: type.field"
		affirm.Equal(t, wMsg, err.Error())
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

	t.Run("log message with trail", func(t *testing.T) {
		// --- Given ---
		opt := WithTrail("type.field")

		// --- When ---
		err := NoError(errors.New("e0"), opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected error to be nil:\n" +
			"\ttrail: type.field\n" +
			"\t want: <nil>\n" +
			"\t have: \"e0\""
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

	t.Run("log message with trail", func(t *testing.T) {
		// --- Given ---
		var e *types.TPtr
		opt := WithTrail("type.field")

		// --- When ---
		err := NoError(e, opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected error to be nil:\n" +
			"\ttrail: type.field\n" +
			"\t want: <nil>\n" +
			"\t have: *types.TPtr"
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
	t.Run("log message with trail", func(t *testing.T) {
		// --- Given ---
		opt := WithTrail("type.field")

		// --- When ---
		have := ErrorIs(errors.New("err0"), errors.New("err0"), opt)

		// --- Then ---
		affirm.NotNil(t, have)
		wMsg := "expected err to have target in its tree:\n" +
			"\ttrail: type.field\n" +
			"\t want: (*errors.errorString) err0\n" +
			"\t have: (*errors.errorString) err0"
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

func Test_ErrorAs(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		var target *types.TPtr

		// --- When ---
		err := ErrorAs(&types.TPtr{Val: "A"}, &target)

		// --- Then ---
		affirm.Nil(t, err)
		affirm.Equal(t, "A", target.Val)
	})

	t.Run("nil", func(t *testing.T) {
		// --- Given ---
		var target *types.TPtr

		// --- When ---
		err := ErrorAs(nil, target)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected non-nil error"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		var target types.TVal

		// --- When ---
		err := ErrorAs(&types.TPtr{Val: "A"}, &target)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected err to have target in its tree:\n" +
			"\twant: (*types.TPtr) &types.TPtr{Val:\"A\"}\n" +
			"\thave: (*types.TVal) &types.TVal{Val:\"\"}"
		affirm.Equal(t, wMsg, err.Error())

		affirm.Equal(t, "", target.Val)
	})

	t.Run("log message with trail", func(t *testing.T) {
		// --- Given ---
		opt := WithTrail("type.field")
		var target types.TVal

		// --- When ---
		err := ErrorAs(&types.TPtr{Val: "A"}, &target, opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected err to have target in its tree:\n" +
			"\ttrail: type.field\n" +
			"\t want: (*types.TPtr) &types.TPtr{Val:\"A\"}\n" +
			"\t have: (*types.TVal) &types.TVal{Val:\"\"}"
		affirm.Equal(t, wMsg, err.Error())

		affirm.Equal(t, "", target.Val)
	})
}

func Test_ErrorEqual(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		err := ErrorEqual("e0", errors.New("e0"))

		// --- Then ---
		affirm.Nil(t, err)
	})

	t.Run("success with format", func(t *testing.T) {
		// --- When ---
		err := ErrorEqual("e00", errors.New("e00"))

		// --- Then ---
		affirm.Nil(t, err)
	})

	t.Run("success with percent in want", func(t *testing.T) {
		// --- When ---
		err := ErrorEqual("e0%", errors.New("e0%"))

		// --- Then ---
		affirm.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// --- When ---
		err := ErrorEqual("e1", errors.New("e0"))

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected error message to be:\n" +
			"\twant: \"e1\"\n" +
			"\thave: \"e0\""
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("log message with trail", func(t *testing.T) {
		// --- Given ---
		opt := WithTrail("type.field")

		// --- When ---
		err := ErrorEqual("e1", errors.New("e0"), opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected error message to be:\n" +
			"\ttrail: type.field\n" +
			"\t want: \"e1\"\n" +
			"\t have: \"e0\""
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("nil error", func(t *testing.T) {
		// --- When ---
		err := ErrorEqual("e1", nil)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected error message to be:\n" +
			"\twant: \"e1\"\n" +
			"\thave: <nil>"
		affirm.Equal(t, wMsg, err.Error())
	})
}

func Test_ErrorContain(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		err := ErrorContain("def", errors.New("abc def ghi"))

		// --- Then ---
		affirm.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// --- When ---
		err := ErrorContain("xyz", errors.New("abc def ghi"))

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected error message to contain:\n" +
			"\twant: \"xyz\"\n" +
			"\thave: \"abc def ghi\""
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("log message with trail", func(t *testing.T) {
		// --- Given ---
		opt := WithTrail("type.field")

		// --- When ---
		err := ErrorContain("xyz", errors.New("abc def ghi"), opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected error message to contain:\n" +
			"\ttrail: type.field\n" +
			"\t want: \"xyz\"\n" +
			"\t have: \"abc def ghi\""
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("nil error", func(t *testing.T) {
		// --- When ---
		err := ErrorContain("xyz", nil)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected error not to be nil:\n" +
			"\twant: <non-nil>\n" +
			"\thave: <nil>"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("nil log message with trail", func(t *testing.T) {
		// --- Given ---
		opt := WithTrail("type.field")

		// --- When ---
		err := ErrorContain("xyz", nil, opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected error not to be nil:\n" +
			"\ttrail: type.field\n" +
			"\t want: <non-nil>\n" +
			"\t have: <nil>"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("nil interface", func(t *testing.T) {
		// --- Given ---
		var e *types.TPtr

		// --- When ---
		err := ErrorContain("abc", e)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected error not to be nil:\n" +
			"\twant: <non-nil>\n" +
			"\thave: *types.TPtr"
		affirm.Equal(t, wMsg, err.Error())
	})
}

func Test_ErrorRegexp(t *testing.T) {
	t.Run("success string", func(t *testing.T) {
		// --- When ---
		err := ErrorRegexp("def", errors.New("abc def ghi"))

		// --- Then ---
		affirm.Nil(t, err)
	})

	t.Run("success regex", func(t *testing.T) {
		// --- Given ---
		rx := regexp.MustCompile(".* def .*")

		// --- When ---
		err := ErrorRegexp(rx, errors.New("abc def ghi"))

		// --- Then ---
		affirm.Nil(t, err)
	})

	t.Run("success regex", func(t *testing.T) {
		// --- When ---
		err := ErrorRegexp("ghi$", errors.New("abc def ghi"))

		// --- Then ---
		affirm.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// --- When ---
		err := ErrorRegexp("^xyz", errors.New("abc def ghi"))

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected error message to match regexp:\n" +
			"\tregexp: ^xyz\n" +
			"\t  have: \"abc def ghi\""
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("log message with trail", func(t *testing.T) {
		// --- Given ---
		opt := WithTrail("type.field")

		// --- When ---
		err := ErrorRegexp("^xyz", errors.New("abc def ghi"), opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected error message to match regexp:\n" +
			"\t trail: type.field\n" +
			"\tregexp: ^xyz\n" +
			"\t  have: \"abc def ghi\""
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("invalid regexp", func(t *testing.T) {
		// --- When ---
		err := ErrorRegexp("[a-z", errors.New("abc def ghi"))

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected error message to match regexp:\n" +
			"\terror: \"error parsing regexp: missing closing ]: `[a-z`\""
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("nil error", func(t *testing.T) {
		// --- When ---
		err := ErrorRegexp("xyz", nil)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected error not to be nil:\n" +
			"\twant: <non-nil>\n" +
			"\thave: <nil>"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("nil log message with trail", func(t *testing.T) {
		// --- Given ---
		opt := WithTrail("type.field")

		// --- When ---
		err := ErrorRegexp("^ab", nil, opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected error not to be nil:\n" +
			"\ttrail: type.field\n" +
			"\t want: <non-nil>\n" +
			"\t have: <nil>"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("nil interface", func(t *testing.T) {
		// --- Given ---
		var e *types.TPtr

		// --- When ---
		err := ErrorRegexp("^ab", e)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected error not to be nil:\n" +
			"\twant: <non-nil>\n" +
			"\thave: *types.TPtr"
		affirm.Equal(t, wMsg, err.Error())
	})
}
