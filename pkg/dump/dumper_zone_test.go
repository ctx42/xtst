// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package dump

import (
	"reflect"
	"testing"

	"github.com/ctx42/testing/internal/affirm"
	"github.com/ctx42/testing/internal/types"
)

func Test_zoneDumper_tabular(t *testing.T) {
	tt := []struct {
		testN string

		val    any
		indent int
		level  int
		want   string
	}{
		{
			"timezone",
			*types.WAW,
			0,
			0,
			`"Europe/Warsaw"`,
		},
		{
			"uses indent and level",
			*types.WAW,
			2,
			1,
			"      \"Europe/Warsaw\"",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			dmp := New(WithIndent(tc.indent))

			// --- When ---
			have := zoneDumper(dmp, tc.level, reflect.ValueOf(tc.val))

			// --- Then ---
			affirm.Equal(t, tc.want, have)
		})
	}
}
