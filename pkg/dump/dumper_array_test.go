// Copyright (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package dump

import (
	"reflect"
	"testing"
	"time"

	"github.com/ctx42/xtst/internal/affirm"
	"github.com/ctx42/xtst/internal/types"
)

func Test_arrayDumper_tabular(t *testing.T) {
	tt := []struct {
		testN string

		cfg  Config
		val  any
		want string
	}{
		{
			"flat & compact array empty",
			NewConfig(Flat, Compact),
			[2]int{},
			"[2]int{0,0}",
		},
		{
			"flat & compact array empty any",
			NewConfig(Flat, Compact),
			[2]any{},
			"[2]any{nil,nil}",
		},
		{
			"flat & compact array of int",
			NewConfig(Flat, Compact),
			[...]int{1, 2},
			"[2]int{1,2}",
		},
		{
			"flat & compact array of float32",
			NewConfig(Flat, Compact),
			[...]float32{1.1, 2.2},
			"[2]float32{1.1,2.2}",
		},
		{
			"compact array",
			NewConfig(Compact),
			[2]int{},
			"[2]int{\n0,\n0,\n}",
		},
		{
			"compact array of int",
			NewConfig(Compact),
			[...]int{1, 2},
			"[2]int{\n1,\n2,\n}",
		},
		{
			"compact array of float32",
			NewConfig(Compact),
			[...]float32{1.1, 2.2},
			"[2]float32{\n1.1,\n2.2,\n}",
		},
		{
			"default array",
			NewConfig(),
			[2]int{},
			"[2]int{\n0,\n0,\n}",
		},
		{
			"default array of int",
			NewConfig(),
			[...]int{1, 2},
			"[2]int{\n1,\n2,\n}",
		},
		{
			"default array of float32",
			NewConfig(),
			[...]float32{1.1, 2.2},
			"[2]float32{\n1.1,\n2.2,\n}",
		},
		{
			"array of times",
			NewConfig(),
			[...]time.Time{
				time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
				time.Date(2000, 1, 2, 3, 4, 5, 0, types.WAW),
			},
			"[2]time.Time{\n\"2000-01-02T03:04:05Z\",\n\"2000-01-02T03:04:05+01:00\",\n}",
		},
		{
			"array of times formated as Unix timestamps",
			NewConfig(TimeFormat(TimeAsUnix)),
			[...]time.Time{
				time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
				time.Date(2000, 1, 2, 3, 4, 5, 0, types.WAW),
			},
			"[2]time.Time{\n946782245,\n946778645,\n}",
		},
		{
			"array of integer type constants",
			NewConfig(),
			[...]types.IntType{0, 1},
			"[2]types.IntType{\n0,\n1,\n}",
		},
		{
			"array of map[string]int",
			NewConfig(Flat, Compact),
			[...]map[string]int{
				{"A": 1},
				{"b": 2},
			},
			`[2]map[string]int{{"A":1},{"b":2}}`,
		},
		{
			"array of map[string]int print type",
			NewConfig(Flat, Compact, PrintType),
			[...]map[string]int{
				{"A": 1},
				{"b": 2},
			},
			`[2]map[string]int{{"A":1},{"b":2}}`,
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
