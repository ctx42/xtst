// Copyright (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package dump

import (
	"reflect"
	"testing"

	"github.com/ctx42/xtst/internal/affirm"
)

func Test_sampleDumper_tabular(t *testing.T) {
	tt := []struct {
		testN string

		val  any
		want string
	}{
		{"bool true", true, "true"},
		{"bool false", false, "false"},
		{"int", 123, "123"},
		{"int8", int8(123), "123"},
		{"int16", int16(123), "123"},
		{"int32", int32(123), "123"},
		{"int64", int64(123), "123"},
		{"uint", uint(123), "123"},
		{"uint16", uint16(123), "123"},
		{"uint32", uint32(123), "123"},
		{"uint64", uint64(123), "123"},
		{"uintptr", uintptr(123), "123"},
		{"float32", float32(12.3), "12.3"},
		{"float32 very small", float32(0.00000000000003), "0.00000000000003"},
		{"float64", 12.3, "12.3"},
		{"float64 very small", 0.00000000000003, "0.00000000000003"},
		{"string", "string", `"string"`},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have := simpleDumper(Dump{}, 0, reflect.ValueOf(tc.val))

			// --- Then ---
			affirm.Equal(t, tc.want, have)
		})
	}
}

func Test_sampleDumper(t *testing.T) {
	t.Run("string with Config.Flat false", func(t *testing.T) {
		// --- Given ---
		dmp := New(NewConfig())

		// --- When ---
		have := simpleDumper(dmp, 0, reflect.ValueOf("str0\nstr1\n"))

		// --- Then ---
		affirm.Equal(t, "\"str0\nstr1\n\"", have)
	})

	t.Run("string with Config.Flat true", func(t *testing.T) {
		// --- Given ---
		dmp := New(NewConfig(Flat))

		// --- When ---
		have := simpleDumper(dmp, 0, reflect.ValueOf("str0\nstr1\n"))

		// --- Then ---
		affirm.Equal(t, "\"str0\\nstr1\\n\"", have)
	})
}
