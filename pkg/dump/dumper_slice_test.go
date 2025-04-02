// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package dump

import (
	"reflect"
	"testing"

	"github.com/ctx42/xtst/internal/affirm"
)

func Test_sliceDumper_tabular(t *testing.T) {
	tt := []struct {
		testN string

		dmp  Dump
		val  any
		want string
	}{
		{
			"nil slice",
			New(WithFlat, WithCompact),
			[]int(nil),
			"nil",
		},
		{
			"flat & compact slice of int",
			New(WithFlat, WithCompact),
			[]int{1, 2},
			"[]int{1,2}",
		},
		{
			"default slice of int",
			New(),
			[]int{1, 2},
			"[]int{\n  1,\n  2,\n}",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have := sliceDumper(tc.dmp, 0, reflect.ValueOf(tc.val))

			// --- Then ---
			affirm.Equal(t, tc.want, have)
		})
	}
}

func Test_sliceDumper(t *testing.T) {
	t.Run("slice of any", func(t *testing.T) {
		// --- Given ---
		val := []any{"str0", 1, "str2"}
		dmp := New(WithFlat, WithCompact)

		// --- When ---
		have := sliceDumper(dmp, 0, reflect.ValueOf(val))

		// --- Then ---
		affirm.Equal(t, `[]any{"str0",1,"str2"}`, have)
	})
}
