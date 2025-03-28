// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package dump

import (
	"fmt"
	"reflect"
	"testing"
	"time"
	"unsafe"

	"github.com/ctx42/xtst/internal/affirm"
	"github.com/ctx42/xtst/internal/tstkit"
	"github.com/ctx42/xtst/internal/types"
)

func Test_New(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		cfg := Config{
			Flat:           true,
			Compact:        true,
			TimeFormat:     time.DateOnly,
			DurationFormat: DurAsSeconds,
			PtrAddr:        true,
		}

		// --- When ---
		dmp := New(cfg)

		// --- Then ---
		affirm.DeepEqual(t, cfg, dmp.cfg)
	})
}

func Test_Default(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		dmp := Default()

		// --- Then ---
		affirm.False(t, dmp.cfg.Flat)
		affirm.False(t, dmp.cfg.Compact)
		affirm.Equal(t, time.RFC3339Nano, dmp.cfg.TimeFormat)
		affirm.Equal(t, "", dmp.cfg.DurationFormat)
		affirm.False(t, dmp.cfg.PtrAddr)
		affirm.True(t, dmp.cfg.UseAny)
		affirm.True(t, len(dmp.cfg.Dumpers) == 3)
		affirm.Equal(t, 6, dmp.cfg.MaxDepth)

		val, ok := dmp.cfg.Dumpers[typDur]
		affirm.True(t, ok)
		affirm.NotNil(t, val)

		val, ok = dmp.cfg.Dumpers[typLocation]
		affirm.True(t, ok)
		affirm.NotNil(t, val)

		val, ok = dmp.cfg.Dumpers[typTime]
		affirm.True(t, ok)
		affirm.NotNil(t, val)
	})
}

func Test_Dump_Any_Value_smoke_tabular(t *testing.T) {
	var itfNil types.TItf
	var itfVal, itfPtr types.TItf
	var sNil *types.TA
	itfVal = types.TVal{}
	itfPtr = &types.TPtr{}
	sPtr := &types.TPtr{Val: "a"}
	var aAnyNil any

	tt := []struct {
		testN string

		cfg  Config
		v    any
		want string
	}{
		// Simple.
		{"bool true", NewConfig(WithFlat, WithCompact), true, "true"},
		{"int", NewConfig(WithFlat, WithCompact), 123, "123"},
		{"int8", NewConfig(WithFlat, WithCompact), int8(123), "123"},
		{"int16", NewConfig(WithFlat, WithCompact), int16(123), "123"},
		{"int32", NewConfig(WithFlat, WithCompact), int32(123), "123"},
		{"int64", NewConfig(WithFlat, WithCompact), int64(123), "123"},
		{"uint", NewConfig(WithFlat, WithCompact), uint(123), "123"},
		{"uint8", NewConfig(WithFlat, WithCompact), uint8(123), "0x7b"},
		{"byte", NewConfig(WithFlat, WithCompact), byte(123), "0x7b"},
		{"uint16", NewConfig(WithFlat, WithCompact), uint16(123), "123"},
		{"uint32", NewConfig(WithFlat, WithCompact), uint32(123), "123"},
		{"uint64", NewConfig(WithFlat, WithCompact), uint64(123), "123"},
		{"uintptr", NewConfig(WithFlat, WithCompact), uintptr(123), "<0x7b>"},
		{"float32", NewConfig(WithFlat, WithCompact), float32(12.3), "12.3"},
		{"float64", NewConfig(WithFlat, WithCompact), 12.3, "12.3"},
		{"complex64", NewConfig(WithFlat, WithCompact), complex(float32(1), float32(2)), "(1+2i)"},
		{"complex128", NewConfig(WithFlat, WithCompact), complex(3.3, 4.4), "(3.3+4.4i)"},
		{"array", NewConfig(WithFlat, WithCompact), [2]int{}, "[2]int{0,0}"},
		{"chan", NewConfig(WithFlat, WithCompact), make(chan int), "(chan int)(<addr>)"},
		{"func", NewConfig(WithFlat, WithCompact), func() {}, "<func>(<addr>)"},
		{"interface nil", NewConfig(WithFlat, WithCompact), itfNil, valNil},
		{"any nil", NewConfig(WithFlat, WithCompact), aAnyNil, valNil},
		{"interface val", NewConfig(WithFlat, WithCompact), itfVal, `{Val:""}`},
		{"interface ptr", NewConfig(WithFlat, WithCompact), itfPtr, `{Val:""}`},
		{
			"map",
			NewConfig(WithFlat, WithCompact),
			map[string]string{"A": "a", "B": "b"},
			`map[string]string{"A":"a","B":"b"}`,
		},
		{"struct pointer", NewConfig(WithFlat, WithCompact), sPtr, `{Val:"a"}`},
		{"slice", NewConfig(WithFlat, WithCompact), []int{1, 2}, "[]int{1,2}"},
		{"string", NewConfig(WithFlat, WithCompact), "string", `"string"`},
		{"struct", NewConfig(WithFlat, WithCompact), struct{ F0 int }{}, "{F0:0}"},
		{
			"registered",
			NewConfig(WithFlat, WithCompact),
			time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
			`"2000-01-02T03:04:05Z"`,
		},
		{"struct nil", NewConfig(WithFlat, WithCompact), sNil, "nil"},
		{
			"registered",
			NewConfig(WithFlat, WithCompact),
			time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
			`"2000-01-02T03:04:05Z"`,
		},
		{
			"unsafe pointer",
			NewConfig(WithFlat, WithCompact),
			unsafe.Pointer(sPtr),
			fmt.Sprintf("<%p>", sPtr),
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			dmp := New(tc.cfg)

			// --- When ---
			haveAny := dmp.Any(tc.v)
			haveVal := dmp.Value(reflect.ValueOf(tc.v))

			// --- Then ---
			affirm.Equal(t, tc.want, haveAny)
			affirm.Equal(t, tc.want, haveVal)
		})
	}
}

func Test_Dump_Any(t *testing.T) {
	t.Run("nil interface value", func(t *testing.T) {
		// --- Given ---
		var itfNil types.TItf
		dmp := Default()

		// --- When ---
		have := dmp.Any(itfNil)

		// --- Then ---
		affirm.Equal(t, valNil, have)
	})

	t.Run("slice of slices of any", func(t *testing.T) {
		// --- Given ---
		val := [][]any{
			{"str00", 0, "str02"},
			{"str10", 1, "str12"},
			{"str10", 1, nil},
		}
		dmp := New(NewConfig(WithFlat, WithCompact))

		// --- When ---
		have := dmp.Any(val)

		// --- Then ---
		want := `[][]any{{"str00",0,"str02"},{"str10",1,"str12"},{"str10",1,nil}}`
		affirm.Equal(t, want, have)
	})

	t.Run("depth", func(t *testing.T) {
		// --- Given ---
		val := struct {
			S0 struct {
				S1 struct {
					S2 struct {
						S4 struct {
							S5 struct {
								S6 struct{ VAL int }
							}
						}
					}
				}
			}
		}{}
		dmp := New(NewConfig(WithFlat, WithCompact))

		// --- When ---
		have := dmp.Any(val)

		// --- Then ---
		affirm.Equal(t, "{S0:{S1:{S2:{S4:{S5:{S6:{VAL:<...>}}}}}}}", have)
	})

	t.Run("format nested slices", func(t *testing.T) {
		// --- Given ---
		type Node struct {
			Value    int
			Children []*Node
		}

		val := &Node{
			Value: 1,
			Children: []*Node{
				{
					Value: 2,
				},
				{
					Value: 3,
					Children: []*Node{
						{
							Value: 4,
						},
					},
				},
			},
		}

		// --- When ---
		have := Default().Any(val)

		// --- Then ---
		want := tstkit.Golden(t, "testdata/struct_nested.txt")
		affirm.Equal(t, want, have)
	})

	t.Run("format nested slices indented twice", func(t *testing.T) {
		// --- Given ---
		type Node struct {
			Value    int
			Children []*Node
		}

		val := &Node{
			Value: 1,
			Children: []*Node{
				{
					Value: 2,
				},
				{
					Value: 3,
					Children: []*Node{
						{
							Value: 4,
						},
					},
				},
			},
		}
		dmp := New(NewConfig(WithIndent(2)))

		// --- When ---
		have := dmp.Any(val)

		// --- Then ---
		want := tstkit.Golden(t, "testdata/struct_nested_with_indent.txt")
		affirm.Equal(t, want, have)
	})
}
