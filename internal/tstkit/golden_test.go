package tstkit

import (
	"testing"

	"github.com/ctx42/testing/internal/affirm"
	"github.com/ctx42/testing/pkg/tester"
)

func Test_Golden(t *testing.T) {
	t.Run("success case 1", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.Close()

		// --- When ---
		have := Golden(tspy, "testdata/golden_case1.txt")

		// --- Then ---
		affirm.Equal(t, "Content #1.\nContent #2.", have)
	})

	t.Run("success case 2", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.Close()

		// --- When ---
		have := Golden(tspy, "testdata/golden_case2.txt")

		// --- Then ---
		affirm.Equal(t, "Content #1.\nContent #2.\n", have)
	})

	t.Run("success case 3", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.Close()

		// --- When ---
		have := Golden(tspy, "testdata/golden_case3.txt")

		// --- Then ---
		affirm.Equal(t, "Content #1.\nContent #2.\n\n", have)
	})

	t.Run("error no marker", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFail()
		tspy.ExpectLogEqual("golden file is missing \"---\" marker")
		tspy.Close()

		var have string

		// --- When ---
		fn := func() { have = Golden(tspy, "testdata/golden_no_marker.txt") }

		// --- Then ---
		msg := affirm.Panic(t, fn)
		affirm.NotNil(t, *msg)
		affirm.Equal(t, "", have)
	})

	t.Run("not existing file", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFail()
		tspy.ExpectLogContain("no such file or directory")
		tspy.ExpectLogContain("testdata/not-existing.txt")
		tspy.Close()

		var have string

		// --- When ---
		fn := func() { have = Golden(tspy, "testdata/not-existing.txt") }

		// --- Then ---
		msg := affirm.Panic(t, fn)
		affirm.NotNil(t, *msg)
		affirm.Equal(t, "", have)
	})
}
