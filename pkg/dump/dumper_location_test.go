// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package dump

import (
	"reflect"
	"testing"

	"github.com/ctx42/xtst/internal/affirm"
	"github.com/ctx42/xtst/internal/types"
)

func Test_locationDumper_tabular(t *testing.T) {
	tt := []struct {
		testN string

		cfg  Config
		val  any
		want string
	}{
		{
			"timezone",
			NewConfig(),
			*types.WAW,
			`"Europe/Warsaw"`,
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			dmp := New(tc.cfg)

			// --- When ---
			have := locationDumper(dmp, 0, reflect.ValueOf(tc.val))

			// --- Then ---
			affirm.Equal(t, tc.want, have)
		})
	}
}
