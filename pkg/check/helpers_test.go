// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/ctx42/xtst/internal"
	"github.com/ctx42/xtst/internal/affirm"
	"github.com/ctx42/xtst/internal/types"
	"github.com/ctx42/xtst/pkg/notice"
)

func Test_typeString(t *testing.T) {
	tt := []struct {
		testN string

		val  reflect.Value
		want string
	}{
		{"string", reflect.ValueOf("abc"), "string"},
		{"int", reflect.ValueOf(123), "int"},
		{"invalid", reflect.ValueOf(nil), "<invalid>"},
		{"struct", reflect.ValueOf(types.TA{}), "types.TA"},
		{"ptr struct", reflect.ValueOf(&types.TA{}), "*types.TA"},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have := typeString(tc.val)

			// --- Then ---
			affirm.Equal(t, tc.want, have)
		})
	}
}

func Test_isPrintableChar(t *testing.T) {
	for i := 0; i <= 31; i++ {
		if !affirm.False(t, isPrintableChar(byte(i))) {
			t.Logf("expected false for %d", i)
		}
	}
	for i := 32; i <= 126; i++ {
		if !affirm.True(t, isPrintableChar(byte(i))) {
			t.Logf("expected true for %d", i)
		}
	}
	for i := 127; i <= 255; i++ {
		if !affirm.False(t, isPrintableChar(byte(i))) {
			t.Logf("expected false for %d", i)
		}
	}
}

func Test_valToString_tabular(t *testing.T) {
	var itf, nilItf types.TItf
	itf = types.TVal{}
	var ptr, nilPtr *types.TPtr
	ptr = &types.TPtr{}

	tt := []struct {
		testN string

		key  any
		want string
	}{
		{"int", 1, "1"},
		{"int8", int8(8), "8"},
		{"int16", int16(16), "16"},
		{"int32", int32(32), "32"},
		{"int64", int32(64), "64"},

		{"uint", 1, "1"},
		{"uint8", uint8(8), "8"},
		{"uint16", uint16(16), "16"},
		{"uint32", uint32(32), "32"},
		{"uint64", uint32(64), "64"},

		{"uintptr", uintptr(123), "123"},

		{"float32", float32(1.1), "1.1"},
		{"float64", 1.2, "1.2"},

		{"string", "abc", `"abc"`},
		{"bool", true, "true"},

		{"struct", types.TA{}, "types.TA"},
		{"nil interface", nilItf, "<invalid>"},
		{"non-nil interface", itf, "types.TVal"},
		{"nil pointer", nilPtr, "<nil>"},
		{"non-nil pointer", ptr, "*types.TPtr"},

		{"complex64", complex(float32(1.0), float32(2.0)), "(1+2i)"},
		{"complex128", complex(3.0, 4.0), "(3+4i)"},
		{"array", [...]int{1, 2, 3}, "<array>"},
		{"chan", make(chan int), "<invalid>"},
		{"func", func() {}, "<invalid>"},
		{"map", map[string]int{"A": 1}, "<invalid>"},
		{"slice", []int{1, 2, 3}, "<invalid>"},
		{"unsafe pointer", unsafe.Pointer(ptr), fmt.Sprintf("<%p>", ptr)},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			key := reflect.ValueOf(tc.key)

			// --- When ---
			have := valToString(key)

			// --- Then ---
			if tc.want != have {
				format := "expected:\n\twant: %#v\n\thave: %#v"
				t.Errorf(format, tc.want, have)
			}
		})
	}
}

func Test_wrap(t *testing.T) {
	t.Run("single error", func(t *testing.T) {
		// --- Given ---
		msg := notice.New("header").Want("%s", "want").Have("%s", "have")

		// --- When ---
		have := wrap(msg)

		// --- Then ---
		affirm.True(t, internal.Same(msg, have))
	})

	t.Run("joined errors", func(t *testing.T) {
		// --- Given ---
		msg0 := notice.New("header 0").Want("%s", "want 0").Have("%s", "have 0")
		msg1 := notice.New("header 1").Want("%s", "want 1").Have("%s", "have 1")
		msg := errors.Join(msg0, msg1)

		// --- When ---
		have := wrap(msg)

		// --- Then ---
		affirm.False(t, internal.Same(msg, have))
		ers := have.(multiError).Unwrap() // nolint: errorlint
		affirm.Equal(t, 2, len(ers))
		affirm.True(t, internal.Same(msg0, ers[0]))
		affirm.True(t, internal.Same(msg1, ers[1]))
	})

	t.Run("nil error", func(t *testing.T) {
		// --- When ---
		have := wrap(nil)

		// --- Then ---
		affirm.Nil(t, have)
	})
}

func Test_multiError_Error(t *testing.T) {
	t.Run("multiple errors with consecutive headers", func(t *testing.T) {
		// --- Given ---
		msg0 := notice.New("header").Want("%s", "want 0").Have("%s", "have 0")
		msg1 := notice.New("header").Want("%s", "want 1").Have("%s", "have 1")
		me := wrap(errors.Join(msg0, msg1))

		// --- When ---
		have := me.Error()

		// --- Then ---
		wMsg := "" +
			"header:\n" +
			"  want: want 0\n" +
			"  have: have 0\n" +
			" ---\n" +
			"  want: want 1\n" +
			"  have: have 1"
		affirm.Equal(t, wMsg, have)
	})

	t.Run("multiple errors without consecutive headers", func(t *testing.T) {
		// --- Given ---
		msg0 := notice.New("header").Want("%s", "want 0").Have("%s", "have 0")
		msg1 := notice.New("other").Want("%s", "want 1").Have("%s", "have 1")
		msg2 := notice.New("header").Want("%s", "want 2").Have("%s", "have 2")
		me := wrap(errors.Join(msg0, msg1, msg2))

		// --- When ---
		have := me.Error()

		// --- Then ---
		wMsg := "" +
			"header:\n" +
			"  want: want 0\n" +
			"  have: have 0\n" +
			"\n" +
			"other:\n" +
			"  want: want 1\n" +
			"  have: have 1\n" +
			"\n" +
			"header:\n" +
			"  want: want 2\n" +
			"  have: have 2"
		affirm.Equal(t, wMsg, have)
	})

	t.Run("not notice error", func(t *testing.T) {
		// --- Given ---
		msg0 := notice.New("header").Want("%s", "want 0").Have("%s", "have 0")
		msg1 := errors.New("not notice")
		msg2 := notice.New("header").Want("%s", "want 2").Have("%s", "have 2")
		me := wrap(errors.Join(msg0, msg1, msg2))

		// --- When ---
		have := me.Error()

		// --- Then ---
		wMsg := "header:\n" +
			"  want: want 0\n" +
			"  have: have 0\n" +
			"\n" +
			"not notice\n" +
			"\n" +
			"header:\n" +
			"  want: want 2\n" +
			"  have: have 2"
		affirm.Equal(t, wMsg, have)
	})

	t.Run("multiple errors serialized multiple times", func(t *testing.T) {
		// --- Given ---
		msg0 := notice.New("header").Want("%s", "want 0").Have("%s", "have 0")
		msg1 := notice.New("header").Want("%s", "want 1").Have("%s", "have 1")
		me := wrap(errors.Join(msg0, msg1))

		// --- When ---
		have0 := me.Error()
		have1 := me.Error()

		// --- Then ---
		wMsg := "" +
			"header:\n" +
			"  want: want 0\n" +
			"  have: have 0\n" +
			" ---\n" +
			"  want: want 1\n" +
			"  have: have 1"
		affirm.Equal(t, wMsg, have0)
		affirm.Equal(t, have0, have1)
	})
}
