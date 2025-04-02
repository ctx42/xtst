// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"reflect"
	"testing"
	"time"

	"github.com/ctx42/xtst/internal"
	"github.com/ctx42/xtst/internal/affirm"
	"github.com/ctx42/xtst/pkg/dump"
)

func Test_WithTrail(t *testing.T) {
	// --- Given ---
	ops := Options{}

	// --- When ---
	have := WithTrail("type.field")(ops)

	// --- Then ---
	affirm.Equal(t, "", ops.Trail)
	affirm.Equal(t, "type.field", have.Trail)
}

func Test_WithTrailLog(t *testing.T) {
	// --- Given ---
	buf := make([]string, 0)
	ops := Options{}

	// --- When ---
	have := WithTrailLog(&buf)(ops)

	// --- Then ---
	affirm.True(t, internal.Same(&buf, have.TrailLog))
}

func Test_WithTimeFormat(t *testing.T) {
	// --- Given ---
	ops := Options{}

	// --- When ---
	have := WithTimeFormat(time.RFC3339)(ops)

	// --- Then ---
	affirm.Equal(t, time.RFC3339, have.TimeFormat)
}

func Test_WithRecent(t *testing.T) {
	// --- Given ---
	ops := Options{}

	// --- When ---
	have := WithRecent(time.Second)(ops)

	// --- Then ---
	affirm.Equal(t, time.Second, have.Recent)
}

func Test_WithDump(t *testing.T) {
	// --- Given ---
	ops := Options{}

	// --- When ---
	have := WithDump(dump.WithMaxDepth(100))(ops)

	// --- Then ---
	affirm.Equal(t, 0, ops.DumpCfg.MaxDepth)
	affirm.Equal(t, 100, have.DumpCfg.MaxDepth)
}

func Test_WithTypeChecker(t *testing.T) {
	// --- Given ---
	ops := Options{}
	chk := func(want, have any, opts ...Option) error { return nil }

	// --- When ---
	have := WithTypeChecker(123, chk)(ops)

	// --- Then ---
	haveChk, _ := have.TypeCheckers[reflect.TypeOf(123)]
	affirm.True(t, internal.Same(chk, haveChk))
}

func Test_WithTrailChecker(t *testing.T) {
	// --- Given ---
	ops := Options{}
	chk := func(want, have any, opts ...Option) error { return nil }

	// --- When ---
	have := WithTrailChecker("type.field", chk)(ops)

	// --- Then ---
	haveChk, _ := have.TrailCheckers["type.field"]
	affirm.True(t, internal.Same(chk, haveChk))
}

func Test_WithSkipTrail(t *testing.T) {
	// --- Given ---
	ops := Options{}

	// --- When ---
	have := WithSkipTrail("type.field1", "type.field2")(ops)

	// --- Then ---
	affirm.True(t, ops.SkipTrails == nil)
	affirm.DeepEqual(t, []string{"type.field1", "type.field2"}, have.SkipTrails)
}

func Test_WithOptions(t *testing.T) {
	// --- Given ---
	trailLog := make([]string, 0)
	ops := Options{
		DumpCfg: dump.Config{
			Flat:           true,
			Compact:        true,
			TimeFormat:     time.Kitchen,
			DurationFormat: "DurAsString",
			PtrAddr:        true,
			PrintType:      true,
			UseAny:         true,
			Dumpers: map[reflect.Type]dump.Dumper{
				reflect.TypeOf(123): dump.Dumper(nil),
			},
			MaxDepth: 6,
		},
		TimeFormat:    time.RFC3339,
		Recent:        123,
		Trail:         "trail",
		TrailLog:      &trailLog,
		TypeCheckers:  make(map[reflect.Type]Check),
		TrailCheckers: make(map[string]Check),
		now:           time.Now,
	}

	// --- When ---
	have := WithOptions(ops)(Options{})

	// --- Then ---
	affirm.True(t, internal.Same(ops.DumpCfg.Dumpers, have.DumpCfg.Dumpers))
	affirm.True(t, internal.Same(ops.TrailLog, have.TrailLog))
	affirm.True(t, internal.Same(ops.TypeCheckers, have.TypeCheckers))
	affirm.True(t, internal.Same(ops.TrailCheckers, have.TrailCheckers))
	affirm.True(t, internal.Same(ops.now, have.now))

	ops.now = nil
	have.now = nil
	affirm.True(t, reflect.DeepEqual(ops, have))
}

func Test_DefaultOptions(t *testing.T) {
	t.Run("no options", func(t *testing.T) {
		// --- When ---
		have := DefaultOptions()

		// --- Then ---
		affirm.Equal(t, false, have.DumpCfg.PtrAddr)
		affirm.Equal(t, DefaultDumpTimeFormat, have.DumpCfg.TimeFormat)

		affirm.Equal(t, DefaultParseTimeFormat, have.TimeFormat)
		affirm.Equal(t, DefaultRecentDuration, have.Recent)
		affirm.Equal(t, "", have.Trail)
		affirm.True(t, have.TrailLog == nil)
		affirm.True(t, have.TypeCheckers == nil)
		affirm.True(t, have.TrailCheckers == nil)
		affirm.True(t, have.SkipTrails == nil)
		affirm.True(t, internal.Same(time.Now, have.now))
		affirm.Equal(t, 9, reflect.ValueOf(have).NumField())
	})

	t.Run("with options", func(t *testing.T) {
		// --- When ---
		have := DefaultOptions(WithTrail("type.field"))

		// --- Then ---
		affirm.Equal(t, false, have.DumpCfg.PtrAddr)
		affirm.Equal(t, DefaultDumpTimeFormat, have.DumpCfg.TimeFormat)

		affirm.Equal(t, DefaultParseTimeFormat, have.TimeFormat)
		affirm.Equal(t, DefaultRecentDuration, have.Recent)
		affirm.Equal(t, "type.field", have.Trail)
		affirm.True(t, have.TrailLog == nil)
		affirm.True(t, have.TypeCheckers == nil)
		affirm.True(t, have.TrailCheckers == nil)
		affirm.True(t, have.SkipTrails == nil)
		affirm.True(t, internal.Same(time.Now, have.now))
		affirm.Equal(t, 9, reflect.ValueOf(have).NumField())
	})
}

func Test_Options_logTrail(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		list := make([]string, 0)
		ops := Options{Trail: "abc", TrailLog: &list}

		// --- When ---
		ops.logTrail()

		// --- Then ---
		affirm.DeepEqual(t, []string{"abc"}, list)
		affirm.DeepEqual(t, []string{"abc"}, *ops.TrailLog)
	})

	t.Run("does not log empty trails", func(t *testing.T) {
		// --- Given ---
		list := make([]string, 0)
		ops := Options{Trail: "", TrailLog: &list}

		// --- When ---
		ops.logTrail()

		// --- Then ---
		affirm.DeepEqual(t, []string{}, list)
		affirm.DeepEqual(t, []string{}, *ops.TrailLog)
	})

	t.Run("does not panic when nil", func(t *testing.T) {
		// --- Given ---
		ops := Options{Trail: "abc"}

		// --- When ---
		ops.logTrail()
	})
}

func Test_Options_structTrail_tabular(t *testing.T) {
	tt := []struct {
		testN string

		trail   string
		typName string
		fldName string
		want    string
	}{
		{"no trail and type", "", "type", "", "type"},                         // 1
		{"no trail and field", "", "", "field", "field"},                      // 2
		{"no trail and type and field", "", "type", "field", "type.field"},    // 3
		{"trail and type", "trail", "type", "", "trail"},                      // 4
		{"trail and field", "trail", "", "field", "trail.field"},              // 5
		{"trail and type and field", "trail", "type", "field", "trail.field"}, // 6
		{"trail[] and type", "[]", "type", "", "[]"},                          // 7
		{"trail[] and field", "[]", "", "field", "[].field"},                  // 8
		{"trail[] and type and field", "[]", "type", "field", "[].field"},     // 9
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			ops := Options{Trail: tc.trail}

			// --- When ---
			have := ops.structTrail(tc.typName, tc.fldName)

			// --- Then ---
			affirm.Equal(t, ops.Trail, tc.trail)
			affirm.Equal(t, tc.want, have)
		})
	}
}

func Test_Options_mapTrail_tabular(t *testing.T) {
	tt := []struct {
		testN string

		trail string
		key   string
		want  string
	}{
		{"empty trail with key", "", "key", "map[key]"},
		{"trail ends with index", "[1]", "key", "[1]map[key]"},
		{"trail ends with index", "[1]", "key", "[1]map[key]"},
		{"not empty trail", "field", "key", "field[key]"},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			ops := Options{Trail: tc.trail}

			// --- When ---
			have := ops.mapTrail(tc.key)

			// --- Then ---
			affirm.Equal(t, ops.Trail, tc.trail)
			affirm.Equal(t, tc.want, have)
		})
	}
}

func Test_Options_arrTrail_tabular(t *testing.T) {
	tt := []struct {
		testN string

		trail string
		kind  string
		key   int
		want  string
	}{
		{"empty trail with key", "", "", 1, "[1]"},
		{"empty trail with key", "", "kind", 1, "<kind>[1]"},
		{"trail ends with index", "[1]", "", 2, "[1][2]"},
		{"not empty trail", "field", "", 1, "field[1]"},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			ops := Options{Trail: tc.trail}

			// --- When ---
			have := ops.arrTrail(tc.kind, tc.key)

			// --- Then ---
			affirm.Equal(t, ops.Trail, tc.trail)
			affirm.Equal(t, tc.want, have)
		})
	}
}
