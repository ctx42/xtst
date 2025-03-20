// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package assert

import (
	"testing"

	"github.com/ctx42/xtst/internal/affirm"
	"github.com/ctx42/xtst/internal/types"
	"github.com/ctx42/xtst/pkg/check"
	"github.com/ctx42/xtst/pkg/tester"
)

func Test_Same(t *testing.T) {
	t.Run("success pointers", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t).Close()
		ptr0 := &types.TPtr{Val: "0"}

		// --- When ---
		have := Same(tspy, ptr0, ptr0)

		// --- Then ---
		affirm.True(t, have)
	})

	t.Run("error want is value", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectError()
		tspy.IgnoreLogs()
		tspy.Close()

		w := types.TPtr{Val: "0"}
		h := &types.TPtr{Val: "0"}

		// --- When ---
		have := Same(tspy, w, h)

		// --- Then ---
		affirm.False(t, have)
	})

	t.Run("error have is value", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectError()
		tspy.IgnoreLogs()
		tspy.Close()

		w := &types.TPtr{Val: "0"}
		h := types.TPtr{Val: "0"}

		// --- When ---
		have := Same(tspy, w, h)

		// --- Then ---
		affirm.False(t, have)
	})

	t.Run("error not same pointers", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectError()
		tspy.IgnoreLogs()
		tspy.Close()

		ptr0 := &types.TPtr{Val: "0"}
		ptr1 := &types.TPtr{Val: "1"}

		// --- When ---
		have := Same(tspy, ptr0, ptr1)

		// --- Then ---
		affirm.False(t, have)
	})

	t.Run("error with option", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectError()
		tspy.ExpectLogContain("\ttrail: type.field\n")
		tspy.Close()

		ptr0 := &types.TPtr{Val: "0"}
		ptr1 := &types.TPtr{Val: "1"}

		opt := check.WithTrail("type.field")

		// --- When ---
		have := Same(tspy, ptr0, ptr1, opt)

		// --- Then ---
		affirm.False(t, have)
	})
}

func Test_NotSame(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t).Close()

		ptr0 := &types.TPtr{Val: "0"}
		ptr1 := &types.TPtr{Val: "1"}

		// --- When ---
		have := NotSame(tspy, ptr0, ptr1)

		// --- Then ---
		affirm.True(t, have)
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectError()
		tspy.IgnoreLogs()
		tspy.Close()

		ptr0 := &types.TPtr{Val: "0"}

		// --- When ---
		have := NotSame(tspy, ptr0, ptr0)

		// --- Then ---
		affirm.False(t, have)
	})

	t.Run("error with option", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectError()
		tspy.ExpectLogContain("\ttrail: type.field\n")
		tspy.Close()

		ptr0 := &types.TPtr{Val: "0"}

		opt := check.WithTrail("type.field")

		// --- When ---
		have := NotSame(tspy, ptr0, ptr0, opt)

		// --- Then ---
		affirm.False(t, have)
	})
}
