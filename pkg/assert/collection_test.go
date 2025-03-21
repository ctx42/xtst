// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package assert

import (
	"testing"

	"github.com/ctx42/xtst/internal/affirm"
	"github.com/ctx42/xtst/pkg/check"
	"github.com/ctx42/xtst/pkg/tester"
)

func Test_Len(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t).Close()

		// --- When ---
		have := Len(tspy, 2, []int{0, 1})

		// --- Then ---
		affirm.True(t, have)
	})

	t.Run("fatal when want is less than actual length", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFatal()
		tspy.IgnoreLogs()
		tspy.Close()

		defer func() { _ = recover() }()

		// --- When ---
		have := Len(tspy, 3, []int{0, 1})

		// --- Then ---
		affirm.False(t, have)
	})

	t.Run("error when want is greater than actual length", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectError()
		tspy.IgnoreLogs()
		tspy.Close()

		// --- When ---
		have := Len(tspy, 1, []int{0, 1})

		// --- Then ---
		affirm.False(t, have)
	})

	t.Run("error with option", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectError()
		tspy.ExpectLogContain("\ttrail: type.field\n")
		tspy.Close()

		opt := check.WithTrail("type.field")

		// --- When ---
		have := Len(tspy, 1, []int{0, 1}, opt)

		// --- Then ---
		affirm.False(t, have)
	})
}
