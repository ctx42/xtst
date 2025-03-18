// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package dump

import (
	"reflect"
	"testing"
	"time"

	"github.com/ctx42/xtst/internal/affirm"
)

func Test_GetDurDumper_tabular(t *testing.T) {
	tt := []struct {
		testN string

		format string
		val    time.Duration
		want   string
	}{
		{"empty format", "", 1500 * time.Millisecond, `"1.5s"`},
		{"DurAsString", DurAsString, 1500 * time.Millisecond, `"1.5s"`},
		{"DurAsSeconds", DurAsSeconds, 1500 * time.Millisecond, "1.5"},
		{"unsupported", "abc", 1500 * time.Millisecond, `"1.5s"`},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			dmp := GetDurDumper(tc.format)

			// --- Then ---
			have := dmp(Dump{}, 0, reflect.ValueOf(tc.val))
			affirm.Equal(t, tc.want, have)
		})
	}
}

func Test_DurDumperString(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		dur := 1500 * time.Millisecond
		val := reflect.ValueOf(dur)

		// --- When ---
		have := DurDumperString(Dump{}, 0, val)

		// --- Then ---
		affirm.Equal(t, `"1.5s"`, have)
	})

	t.Run("zero value", func(t *testing.T) {
		// --- Given ---
		dur := 0 * time.Second
		val := reflect.ValueOf(dur)

		// --- When ---
		have := DurDumperString(Dump{}, 0, val)

		// --- Then ---
		affirm.Equal(t, `"0s"`, have)
	})
}

func Test_DurDumperSeconds(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		dur := 1500 * time.Millisecond
		val := reflect.ValueOf(dur)

		// --- When ---
		have := DurDumperSeconds(Dump{}, 0, val)

		// --- Then ---
		affirm.Equal(t, "1.5", have)
	})

	t.Run("zero value", func(t *testing.T) {
		// --- Given ---
		dur := 0 * time.Second
		val := reflect.ValueOf(dur)

		// --- When ---
		have := DurDumperSeconds(Dump{}, 0, val)

		// --- Then ---
		affirm.Equal(t, "0", have)
	})
}
