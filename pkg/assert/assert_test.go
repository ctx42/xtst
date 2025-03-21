// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package assert

import (
	"testing"

	"github.com/ctx42/xtst/internal/affirm"
	"github.com/ctx42/xtst/pkg/check"
	"github.com/ctx42/xtst/pkg/tester"
)

func Test_Count(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.Close()

		// --- When ---
		have := Count(tspy, 2, "ab", "ab cd ab")

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
		have := Count(tspy, 1, 123, "ab cd ef")

		// --- Then ---
		affirm.False(t, have)
	})

	t.Run("error with option", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectError()
		tspy.ExpectLogContain("\ttrail: type.field")
		tspy.Close()

		opt := check.WithTrail("type.field")

		// --- When ---
		have := Count(tspy, 1, 123, "ab cd ef", opt)

		// --- Then ---
		affirm.False(t, have)
	})
}

func Test_SameType(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t).Close()

		// --- When ---
		have := SameType(tspy, true, true)

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
		have := SameType(tspy, 1, uint(1))

		// --- Then ---
		affirm.False(t, have)
	})

	t.Run("error with options", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFatal()
		tspy.ExpectLogContain("\ttrail: type.field\n")
		tspy.Close()

		defer func() { _ = recover() }()
		opt := check.WithTrail("type.field")

		// --- When ---
		have := SameType(tspy, 1, uint(1), opt)

		// --- Then ---
		affirm.False(t, have)
	})
}
