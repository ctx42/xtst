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

func Test_WithOptions(t *testing.T) {
	// --- Given ---
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
			Depth: 6,
		},
		TimeFormat: time.RFC3339,
		Recent:     123,
		Trail:      "trail",
		now:        time.Now,
		skipType:   true,
	}

	// --- When ---
	have := WithOptions(ops)(Options{})

	// --- Then ---
	affirm.True(t, internal.Same(ops.DumpCfg.Dumpers, have.DumpCfg.Dumpers))
	affirm.True(t, internal.Same(ops.TrailLog, have.TrailLog))
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
		affirm.True(t, internal.Same(time.Now, have.now))
		affirm.False(t, have.skipType)
		affirm.Equal(t, 7, reflect.ValueOf(have).NumField())
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
		affirm.True(t, internal.Same(time.Now, have.now))
		affirm.False(t, have.skipType)
		affirm.Equal(t, 7, reflect.ValueOf(have).NumField())
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

	t.Run("does not logTrail() empty paths", func(t *testing.T) {
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

func Test_Options_skipTrail(t *testing.T) {
	t.Run("skipType is reset to false", func(t *testing.T) {
		// --- Given ---
		ops := Options{skipType: true}

		// --- When ---
		have := ops.structTrail("type", "field")

		// --- Then ---
		affirm.False(t, have.skipType)
	})
}

func Test_Options_structTrail_tabular(t *testing.T) {
	tt := []struct {
		testN string

		trail    string
		skipNext bool
		typName  string
		fldName  string
		want     string
	}{
		{"empty path with field", "", false, "", "field", "field"},
		{"empty path with type", "", false, "type", "", "type"},
		{"path with type", "path", false, "type", "", "path.type"},
		{"path with field", "path", false, "", "field", "path.field"},
		{"path with empty type and field", "path", false, "", "", "path"},
		{"path with index", "path[1]", false, "", "", "path[1]"},
		{"path with index and type", "path[1]", false, "type", "", "path[1]"},
		{"path with index and field", "path[1]", false, "", "field", "path[1].field"},
		{
			"path with index type and field",
			"path[1]",
			false,
			"type",
			"field",
			"path[1].field",
		},
		{
			"skip type",
			"path",
			true,
			"type",
			"field",
			"path.field",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			ops := Options{Trail: tc.trail, skipType: tc.skipNext}

			// --- When ---
			have := ops.structTrail(tc.typName, tc.fldName)

			// --- Then ---
			affirm.Equal(t, tc.want, have.Trail)
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
		{"empty path with key", "", "key", "map[key]"},
		{"path ends with index", "[1]", "key", "[1]map[key]"},
		{"path ends with index", "[1]", "key", "[1]map[key]"},
		{"not empty path", "field", "key", "field[key]"},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			ops := Options{Trail: tc.trail}

			// --- When ---
			have := ops.mapTrail(tc.key)

			// --- Then ---
			affirm.Equal(t, tc.want, have.Trail)
		})
	}
}

func Test_Options_arrTrail_tabular(t *testing.T) {
	tt := []struct {
		testN string

		trail string
		key   int
		want  string
	}{
		{"empty path with key", "", 1, "[1]"},
		{"path ends with index", "[1]", 2, "[1][2]"},
		{"not empty path", "field", 1, "field[1]"},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			ops := Options{Trail: tc.trail}

			// --- When ---
			have := ops.arrTrail(tc.key)

			// --- Then ---
			affirm.Equal(t, tc.want, have.Trail)
		})
	}
}
