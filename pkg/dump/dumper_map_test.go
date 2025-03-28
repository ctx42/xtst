// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package dump

import (
	"reflect"
	"testing"

	"github.com/ctx42/xtst/internal/affirm"
	"github.com/ctx42/xtst/internal/tstkit"
	"github.com/ctx42/xtst/internal/types"
)

func Test_mapDumper_tabular(t *testing.T) {
	var nilMap map[string]int

	tt := []struct {
		testN string

		cfg  Config
		val  any
		want string
	}{
		{
			"empty map",
			NewConfig(WithFlat),
			map[string]int{},
			`map[string]int{}`,
		},
		{
			"nil map",
			NewConfig(),
			nilMap,
			`map[string]int(nil)`,
		},
		{
			"default map[int]int",
			NewConfig(),
			map[int]int{1: 10, 2: 20},
			"map[int]int{\n  1: 10,\n  2: 20,\n}",
		},
		{
			"default map[int]int ith indent",
			NewConfig(WithIndent(2)),
			map[int]int{1: 10, 2: 20},
			"    map[int]int{\n      1: 10,\n      2: 20,\n    }",
		},
		{
			"flat map[int]int",
			NewConfig(WithFlat),
			map[int]int{1: 10, 2: 20},
			"map[int]int{1: 10, 2: 20}",
		},
		{
			"flat and compact map[int]int",
			NewConfig(WithFlat, WithCompact),
			map[int]int{1: 10, 2: 20},
			"map[int]int{1:10,2:20}",
		},
		{
			"flat map[int]types.T1",
			NewConfig(WithFlat, WithCompact, WithTimeFormat(TimeAsUnix)),
			map[int]types.T1{0: {Int: 0}, 1: {Int: 1}},
			"map[int]types.T1{0:{Int:0,T1:nil},1:{Int:1,T1:nil}}",
		},
		{
			"default map[int]types.T1",
			NewConfig(WithTimeFormat(TimeAsUnix)),
			map[int]types.T1{0: {Int: 0}, 1: {Int: 1}},
			tstkit.Golden(t, "testdata/map_of_structs.txt"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			dmp := New(tc.cfg)

			// --- When ---
			have := mapDumper(dmp, 0, reflect.ValueOf(tc.val))

			// --- Then ---
			affirm.Equal(t, tc.want, have)
		})
	}
}
