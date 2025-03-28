// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package dump

import (
	"reflect"
	"testing"

	"github.com/ctx42/xtst/internal/affirm"
	"github.com/ctx42/xtst/internal/types"
)

func Test_arrayDumper_tabular(t *testing.T) {
	var nilArr [2]int

	tt := []struct {
		testN string

		cfg  Config
		val  any
		want string
	}{
		{
			"default",
			NewConfig(),
			[2]int{0, 1},
			"[2]int{\n  0,\n  1,\n}",
		},
		{
			"nil array",
			NewConfig(),
			nilArr,
			"[2]int{\n  0,\n  0,\n}",
		},
		{
			"default with indent",
			NewConfig(WithIndent(2)),
			[2]int{0, 1},
			"    [2]int{\n      0,\n      1,\n    }",
		},
		{
			"flat array",
			NewConfig(WithFlat),
			[2]int{0, 1},
			"[2]int{0, 1}",
		},
		{
			"flat and compact array",
			NewConfig(WithFlat, WithCompact),
			[2]int{0, 1},
			"[2]int{0,1}",
		},
		{
			"flat array empty int",
			NewConfig(WithFlat),
			[2]int{},
			"[2]int{0, 0}",
		},
		{
			"flat slice empty",
			NewConfig(WithFlat),
			[]int{},
			"[]int{}",
		},
		{
			"flat array empty any",
			NewConfig(WithFlat),
			[2]any{},
			"[2]any{nil, nil}",
		},
		{
			"flat array of map[string]int",
			NewConfig(WithFlat),
			[...]map[string]int{
				{"A": 1},
				{"b": 2},
			},
			`[2]map[string]int{{"A": 1}, {"b": 2}}`,
		},
		{
			"array of map[int]int",
			NewConfig(),
			[...]map[int]int{
				{1: 10},
				{2: 20},
			},
			"[2]map[int]int{\n  {\n    1: 10,\n  },\n  {\n    2: 20,\n  },\n}",
		},
		{
			"array of structs",
			NewConfig(),
			[]types.T1{{Int: 1}, {Int: 2}},
			"[]types.T1{\n" +
				"  {\n    Int: 1,\n    T1: nil,\n  },\n" +
				"  {\n    Int: 2,\n    T1: nil,\n  },\n" +
				"}",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			dmp := New(tc.cfg)

			// --- When ---
			have := arrayDumper(dmp, 0, reflect.ValueOf(tc.val))

			// --- Then ---
			affirm.Equal(t, tc.want, have)
		})
	}
}
