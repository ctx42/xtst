// Copyright (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package dump

import (
	"reflect"
	"testing"
	"time"

	"github.com/ctx42/xtst/internal/affirm"
	"github.com/ctx42/xtst/internal/tstkit"
	"github.com/ctx42/xtst/internal/types"
)

func Test_mapDumper_tabular(t *testing.T) {
	var nilMap map[string]int
	var nilAnyMap map[string]any

	tt := []struct {
		testN string

		cfg  Config
		val  any
		want string
	}{
		{
			"flat & compact empty map",
			NewConfig(Flat, Compact),
			map[string]int{},
			`map[string]int{}`,
		},
		{
			"flat & compact nil map",
			NewConfig(Flat, Compact),
			nilMap,
			`map[string]int(nil)`,
		},
		{
			"flat & compact nil map with any values",
			NewConfig(Flat, Compact),
			nilAnyMap,
			`map[string]any(nil)`,
		},
		{
			"default empty map",
			NewConfig(),
			make(map[string]int),
			`map[string]int{}`,
		},
		{
			"flat & compact map[string]int",
			NewConfig(Flat, Compact),
			map[string]int{"A": 1, "B": 2},
			`map[string]int{"A":1,"B":2}`,
		},
		{
			"flat & compact map[string]string",
			NewConfig(Flat, Compact),
			map[string]string{"A": "a", "B": "b"},
			`map[string]string{"A":"a","B":"b"}`,
		},
		{
			"flat & compact map[int]int",
			NewConfig(Flat, Compact),
			map[int]int{1: 11, 2: 22},
			"map[int]int{1:11,2:22}",
		},
		{
			"flat map[int]int",
			NewConfig(Flat),
			map[int]int{1: 11, 2: 22},
			"map[int]int{1: 11, 2: 22}",
		},
		{
			"default map[int]int",
			NewConfig(),
			map[int]int{1: 11, 2: 22},
			"map[int]int{\n\t1: 11,\n\t2: 22,\n}",
		},
		{
			"flat map[int]time.Time",
			NewConfig(Flat, TimeFormat(TimeAsUnix)),
			map[int]time.Time{
				1: time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
				2: time.Date(2000, 1, 2, 3, 4, 5, 0, types.WAW),
			},
			"map[int]time.Time{1: 946782245, 2: 946778645}",
		},
		{
			"default map[int]types.T1",
			NewConfig(TimeFormat(TimeAsUnix)),
			map[int]types.T1{0: {Int: 0}, 1: {Int: 1}},
			tstkit.Golden(t, "testdata/map_of_structs.txt"),
		},
		{
			"flat map[int]types.T1",
			NewConfig(Flat, Compact, TimeFormat(TimeAsUnix)),
			map[int]types.T1{0: {Int: 0}, 1: {Int: 1}},
			tstkit.Golden(t, "testdata/map_of_structs_flat_compact.txt"),
		},
		{
			"flat & compact map[string]any with integers",
			NewConfig(Flat, Compact),
			map[string]any{"A": 1, "B": 2},
			"map[string]any{\"A\":1,\"B\":2}",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			dmp := New(tc.cfg)

			// --- When ---
			have := mapDumper(dmp, 1, reflect.ValueOf(tc.val))

			// --- Then ---
			affirm.Equal(t, tc.want, have)
		})
	}
}
