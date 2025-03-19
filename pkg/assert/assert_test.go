package assert

import (
	"errors"
	"testing"

	"github.com/ctx42/xtst/internal/affirm"
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
		tspy.ExpectLogContain("\tpath: field\n")
		tspy.Close()

		opt := check.WithPath("field")

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

func Test_Nil(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t).Close()

		// --- When ---
		have := Nil(tspy, nil)

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
		have := Nil(tspy, 42)

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
		have := Nil(tspy, 42, opt)

		// --- Then ---
		affirm.False(t, have)
	})
}

func Test_NotNil(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t).Close()

		// --- When ---
		have := NotNil(tspy, 42)

		// --- Then ---
		affirm.True(t, have)
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFatal()
		tspy.IgnoreLogs()
		tspy.Close()

		defer func() { _ = recover() }()

		// --- When ---
		var have bool
		func() {
			defer func() { _ = recover() }()
			have = NotNil(tspy, nil)
		}()

		// --- Then ---
		affirm.False(t, have)
	})

	t.Run("option is passed", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFail()
		tspy.ExpectLogContain("\tpath: pth")
		tspy.Close()

		defer func() { _ = recover() }()
		opt := check.WithPath("pth")

		// --- When ---
		NotNil(tspy, nil, opt)
	})
}
