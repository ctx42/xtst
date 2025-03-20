// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package assert

import (
	"errors"
	"fmt"
	"testing"

	"github.com/ctx42/xtst/internal/affirm"
	"github.com/ctx42/xtst/internal/types"
	"github.com/ctx42/xtst/pkg/check"
	"github.com/ctx42/xtst/pkg/tester"
)

func Test_Error(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t).Close()

		// --- When ---
		have := Error(tspy, errors.New("e0"))

		// --- Then ---
		affirm.True(t, have)
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectError()
		tspy.IgnoreLogs()
		tspy.Close()

		// --- When ---
		have := Error(tspy, nil)

		// --- Then ---
		affirm.False(t, have)
	})

	t.Run("option is passed", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectError()
		tspy.ExpectLogContain("\tpath: pth")
		tspy.Close()

		opt := check.WithPath("pth")

		// --- When ---
		have := Error(tspy, nil, opt)

		// --- Then ---
		affirm.False(t, have)
	})
}

func Test_NoError(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t).Close()

		// --- When ---
		have := NoError(tspy, nil)

		// --- Then ---
		affirm.True(t, have)
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFatal()
		tspy.ExpectLogContain("expected error to be nil")
		tspy.Close()

		// --- When ---
		var have bool
		func() {
			defer func() { _ = recover() }()
			have = NoError(tspy, errors.New("e0"))
		}()

		// --- Then ---
		affirm.False(t, have)
	})

	t.Run("error with option", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFatal()
		tspy.ExpectLogContain("\tpath: pth\n")
		tspy.Close()

		opt := check.WithPath("pth")

		// --- When ---
		var have bool
		func() {
			defer func() { _ = recover() }()
			have = NoError(tspy, errors.New("e0"), opt)
		}()

		// --- Then ---
		affirm.False(t, have)
	})
}

func Test_ErrorIs(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t).Close()

		err0 := errors.New("err0")
		err1 := errors.New("err1")
		err2 := fmt.Errorf("wrap: %w %w", err0, err1)

		// --- When ---
		have := ErrorIs(tspy, err2, err1)

		// --- Then ---
		affirm.True(t, have)
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFatal()
		tspy.ExpectLogContain("expected err to have target in its tree")
		tspy.Close()

		err0 := errors.New("err0")
		err1 := errors.New("err1")

		// --- When ---
		var have bool
		func() {
			defer func() { _ = recover() }()
			have = ErrorIs(tspy, err0, err1)
		}()

		// --- Then ---
		affirm.False(t, have)
	})

	t.Run("error with option", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFatal()
		tspy.ExpectLogContain("\tpath: pth\n")
		tspy.Close()

		err0 := errors.New("err0")
		err1 := errors.New("err1")
		opt := check.WithPath("pth")

		// --- When ---
		var have bool
		func() {
			defer func() { _ = recover() }()
			have = ErrorIs(tspy, err0, err1, opt)
		}()

		// --- Then ---
		affirm.False(t, have)
	})
}

func Test_ErrorAs(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		var target *types.TPtr
		tspy := tester.New(t).Close()

		// --- When ---
		have := ErrorAs(tspy, &types.TPtr{Val: "A"}, &target)

		// --- Then ---
		affirm.True(t, have)
		affirm.Equal(t, "A", target.Val)
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectError()
		tspy.IgnoreLogs()
		tspy.Close()

		var target types.TVal

		// --- When ---
		have := ErrorAs(tspy, &types.TPtr{Val: "A"}, &target)

		// --- Then ---
		affirm.False(t, have)
		affirm.Equal(t, "", target.Val)
	})

	t.Run("error with option", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectError()
		tspy.ExpectLogContain("\tpath: pth\n")
		tspy.Close()

		var target types.TVal
		opt := check.WithPath("pth")

		// --- When ---
		have := ErrorAs(tspy, &types.TPtr{Val: "A"}, &target, opt)

		// --- Then ---
		affirm.False(t, have)
		affirm.Equal(t, "", target.Val)
	})
}

func Test_ErrorEqual(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t).Close()

		// --- When ---
		have := ErrorEqual(tspy, "e0", errors.New("e0"))

		// --- Then ---
		affirm.True(t, have)
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectError()
		tspy.IgnoreLogs()
		tspy.Close()

		// --- When ---
		have := ErrorEqual(tspy, "e1", errors.New("e0"))

		// --- Then ---
		affirm.False(t, have)
	})

	t.Run("error with option", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectError()
		tspy.ExpectLogContain("\tpath: pth\n")
		tspy.Close()

		opt := check.WithPath("pth")

		// --- When ---
		have := ErrorEqual(tspy, "e1", errors.New("e0"), opt)

		// --- Then ---
		affirm.False(t, have)
	})
}

func Test_ErrorContain(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t).Close()

		// --- When ---
		have := ErrorContain(tspy, "def", errors.New("abc def ghi"))

		// --- Then ---
		affirm.True(t, have)
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectError()
		tspy.IgnoreLogs()
		tspy.Close()

		// --- When ---
		have := ErrorContain(tspy, "xyz", errors.New("abc def ghi"))

		// --- Then ---
		affirm.False(t, have)
	})

	t.Run("error with path", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectError()
		tspy.ExpectLogContain("\tpath: pth\n")
		tspy.Close()

		opt := check.WithPath("pth")

		// --- When ---
		have := ErrorContain(tspy, "xyz", errors.New("abc def ghi"), opt)

		// --- Then ---
		affirm.False(t, have)
	})
}

func Test_ErrorRegexp(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t).Close()

		// --- When ---
		have := ErrorRegexp(tspy, "^abc", errors.New("abc def ghi"))

		// --- Then ---
		affirm.True(t, have)
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectError()
		tspy.IgnoreLogs()
		tspy.Close()

		// --- When ---
		have := ErrorRegexp(tspy, "abc$", errors.New("abc def ghi"))

		// --- Then ---
		affirm.False(t, have)
	})

	t.Run("error with option", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectError()
		tspy.ExpectLogContain("\t  path: pth\n")
		tspy.Close()

		opt := check.WithPath("pth")

		// --- When ---
		have := ErrorRegexp(tspy, "abc$", errors.New("abc def ghi"), opt)

		// --- Then ---
		affirm.False(t, have)
	})

	t.Run("invalid regex", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectError()
		tspy.IgnoreLogs()
		tspy.Close()

		// --- When ---
		have := ErrorRegexp(tspy, "[a-z", errors.New("abc def ghi"))

		// --- Then ---
		affirm.False(t, have)
	})
}
