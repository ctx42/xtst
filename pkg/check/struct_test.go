// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"fmt"
	"testing"
	"time"

	"github.com/ctx42/xtst/internal/affirm"
	"github.com/ctx42/xtst/internal/types"
	"github.com/ctx42/xtst/pkg/must"
)

// TODO(rz): make very detailed code review of this file.
// TODO(rz): now it uses Equal so we need to move all the tests there.

func Test_structEqual(t *testing.T) {
	t.Run("equal by value", func(t *testing.T) {
		// --- Given ---
		want := types.TA{Str: "abc"}
		have := types.TA{Str: "abc"}
		haveList := make([]string, 0)

		// --- When ---
		err := Equal(want, have, WithTrailLog(&haveList))

		// --- Then ---
		affirm.Nil(t, err)
		wantList := []string{
			"TA.Int",
			"TA.Str",
			"TA.Tim",
			"TA.Dur",
			"TA.Loc",
			"TA.TAp",
		}
		affirm.DeepEqual(t, wantList, haveList)
	})

	t.Run("equal by pointer", func(t *testing.T) {
		// --- Given ---
		want := &types.TA{Str: "abc"}
		have := &types.TA{Str: "abc"}
		haveList := make([]string, 0)

		// --- When ---
		err := Equal(want, have, WithTrailLog(&haveList))

		// --- Then ---
		affirm.Nil(t, err)
		wantList := []string{
			"TA.Int",
			"TA.Str",
			"TA.Tim",
			"TA.Dur",
			"TA.Loc",
			"TA.TAp",
		}
		affirm.DeepEqual(t, wantList, haveList)
	})

	t.Run("equal with embedded not struct field", func(t *testing.T) {
		// --- Given ---
		want := types.TC{TD: types.TD("abc"), Int: 123}
		have := types.TC{TD: types.TD("abc"), Int: 123}
		haveList := make([]string, 0)

		// --- When ---
		err := Equal(want, have, WithTrailLog(&haveList))

		// --- Then ---
		affirm.Nil(t, err)
		wantList := []string{"TC.TD", "TC.Int"}
		affirm.DeepEqual(t, wantList, haveList)
	})

	t.Run("not equal with embedded not struct field", func(t *testing.T) {
		// --- Given ---
		want := types.TC{TD: types.TD("abc"), Int: 123}
		have := types.TC{TD: types.TD("xyz"), Int: 123}
		haveList := make([]string, 0)

		// --- When ---
		err := Equal(want, have, WithTrailLog(&haveList))

		// --- Then ---
		wMsg := "expected values to be equal:\n" +
			"  trail: TC.TD\n" +
			"   want: \"abc\"\n" +
			"   have: \"xyz\""
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("if want is a pointer it must not be nil", func(t *testing.T) {
		// --- Given ---
		var want *types.TA
		have := &types.TA{Str: "abc"}

		// --- When ---
		err := Equal(want, have)

		// --- Then ---
		fmt.Println(err.Error()) // TODO(rz):

		wMsg := "expected values to be equal:\n" +
			"  want: nil\n" +
			"  have: \n" +
			"        {\n" +
			"          Int: 0,\n" +
			"          Str: \"abc\",\n" +
			"          Tim: \"0001-01-01T00:00:00Z\",\n" +
			"          Dur: \"0s\",\n" +
			"          Loc: nil,\n" +
			"          TAp: nil,\n" +
			"        }"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("if have is a pointer it must not be nil", func(t *testing.T) {
		// --- Given ---
		want := &types.TA{Str: "abc"}
		var have *types.TA

		// --- When ---
		err := Equal(want, have)

		// --- Then ---
		wMsg := "expected values to be equal:\n" +
			"  want: \n" +
			"        {\n" +
			"          Int: 0,\n" +
			"          Str: \"abc\",\n" +
			"          Tim: \"0001-01-01T00:00:00Z\",\n" +
			"          Dur: \"0s\",\n" +
			"          Loc: nil,\n" +
			"          TAp: nil,\n" +
			"        }\n" +
			"  have: nil"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("not equal base type field", func(t *testing.T) {
		// --- Given ---
		want := types.TA{Str: "abc"}
		have := types.TA{Str: "xyz"}

		// --- When ---
		err := Equal(want, have)

		// --- Then ---
		wMsg := "expected values to be equal:\n" +
			"  trail: TA.Str\n" +
			"   want: \"abc\"\n" +
			"   have: \"xyz\""
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("multiple errors", func(t *testing.T) {
		// --- Given ---
		want := types.TA{Int: 1, Str: "abc"}
		have := types.TA{Int: 2, Str: "xyz"}

		// --- When ---
		err := Equal(want, have)

		// --- Then ---
		wMsg := "expected values to be equal:\n" +
			"  trail: TA.Int\n" +
			"   want: 1\n" +
			"   have: 2\n" +
			"expected values to be equal:\n" +
			"  trail: TA.Str\n" +
			"   want: \"abc\"\n" +
			"   have: \"xyz\""
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("equal time as an argument", func(t *testing.T) {
		// --- Given ---
		tim := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)

		// --- When ---
		err := Equal(tim, tim)

		// --- Then ---
		affirm.Nil(t, err)
	})

	t.Run("not equal time as an argument", func(t *testing.T) {
		// --- Given ---
		tim0 := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		tim1 := time.Date(2001, 1, 2, 3, 4, 5, 0, time.UTC)

		// --- When ---
		err := Equal(tim0, tim1)

		// --- Then ---
		wMsg := "expected equal dates:\n" +
			"  want: 2000-01-02T03:04:05Z\n" +
			"  have: 2001-01-02T03:04:05Z\n" +
			"  diff: -8784h0m0s"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("not equal time as an argument with root set", func(t *testing.T) {
		// --- Given ---
		tim0 := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		tim1 := time.Date(2001, 1, 2, 3, 4, 5, 0, time.UTC)

		// --- When ---
		// TODO(rz): everywhere WithTrail option is used we need to check the
		//  test name. We no longer refer to tit as root or path.
		err := Equal(tim0, tim1, WithTrail("type.field"))

		// --- Then ---
		wMsg := "expected equal dates:\n" +
			"  trail: type.field\n" +
			"   want: 2000-01-02T03:04:05Z\n" +
			"   have: 2001-01-02T03:04:05Z\n" +
			"   diff: -8784h0m0s"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("equal time.Location as an argument", func(t *testing.T) {
		// --- When ---
		err := Equal(types.WAW, types.WAW)

		// --- Then ---
		affirm.Nil(t, err)
	})

	t.Run("not equal time.Location as an argument", func(t *testing.T) {
		// --- Given ---
		loc0 := types.WAW
		loc1 := time.UTC

		// --- When ---
		err := Equal(loc0, loc1)

		// --- Then ---
		wMsg := "expected same timezone:\n" +
			"  want: Europe/Warsaw\n" +
			"  have: UTC"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("not equal time.Location as an argument with root set", func(t *testing.T) {
		// --- Given ---
		loc0 := types.WAW
		loc1 := time.UTC

		// --- When ---
		err := Equal(loc0, loc1, WithTrail("type.field"))

		// --- Then ---
		wMsg := "expected same timezone:\n" +
			"  trail: type.field\n" +
			"   want: Europe/Warsaw\n" +
			"   have: UTC"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("equal time and timezone", func(t *testing.T) {
		// --- Given ---
		tim := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		want := types.TA{Tim: tim}
		have := types.TA{Tim: tim}

		// --- When ---
		err := Equal(want, have)

		// --- Then ---
		affirm.Nil(t, err)
	})

	t.Run("not equal time and equal timezone", func(t *testing.T) {
		// --- Given ---
		wTim := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		hTim := time.Date(2000, 1, 2, 3, 4, 6, 0, time.UTC)
		want := types.TA{Tim: wTim}
		have := types.TA{Tim: hTim}

		// --- When ---
		err := Equal(want, have)

		// --- Then ---
		wMsg := "expected equal dates:\n" +
			"  trail: TA.Tim\n" +
			"   want: 2000-01-02T03:04:05Z\n" +
			"   have: 2000-01-02T03:04:06Z\n" +
			"   diff: -1s"
		affirm.Equal(t, wMsg, err.Error())
	})

	// t.Run("equal time and not equal timezone", func(t *testing.T) {
	// 	// --- Given ---
	// 	wTim := time.Date(2000, 1, 2, 4, 4, 5, 0, types.WAW)
	// 	hTim := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
	// 	want := types.TA{Tim: wTim}
	// 	have := types.TA{Tim: hTim}
	//
	// 	// --- When ---
	// 	err := Equal(want, have)
	//
	//     // TODO(rz): fails because now time check by default uses Time instead
	//     //  of TimeExact.
	//
	// 	// --- Then ---
	// 	wMsg := "expected same timezone:\n" +
	// 		"  trail: TA.Tim\n" +
	// 		"   want: Europe/Warsaw\n" +
	// 		"   have: UTC"
	// 	affirm.Equal(t, wMsg, err.Error())
	// })

	t.Run("equal timezone field", func(t *testing.T) {
		// --- Given ---
		waw := must.Value(time.LoadLocation("Europe/Warsaw"))
		want := types.TA{Loc: waw}
		have := types.TA{Loc: waw}
		fields := make([]string, 0)

		// --- When ---
		err := Equal(want, have, WithTrailLog(&fields))

		// --- Then ---
		affirm.Nil(t, err)
		visited := []string{
			"TA.Int",
			"TA.Str",
			"TA.Tim",
			"TA.Dur",
			"TA.Loc",
			"TA.TAp",
		}
		affirm.DeepEqual(t, visited, fields)
	})

	t.Run("not equal timezone field", func(t *testing.T) {
		// --- Given ---
		want := types.TA{Loc: must.Value(time.LoadLocation("Europe/Warsaw"))}
		have := types.TA{Loc: must.Value(time.LoadLocation("Europe/Zurich"))}

		// --- When ---
		err := Equal(want, have)

		// --- Then ---
		wMsg := "expected same timezone:\n" +
			"  trail: TA.Loc\n" +
			"   want: Europe/Warsaw\n" +
			"   have: Europe/Zurich"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("nil want timezone not nil have timezone", func(t *testing.T) {
		// --- Given ---
		want := types.TA{Loc: nil}
		have := types.TA{Loc: must.Value(time.LoadLocation("Europe/Warsaw"))}

		// --- When ---
		err := Equal(want, have)

		// --- Then ---
		wMsg := "expected same timezone:\n" +
			"  trail: TA.Loc\n" +
			"   want: UTC\n" +
			"   have: Europe/Warsaw"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("not nil want timezone nil have timezone", func(t *testing.T) {
		// --- Given ---
		want := types.TA{Loc: must.Value(time.LoadLocation("Europe/Warsaw"))}
		have := types.TA{Loc: nil}

		// --- When ---
		err := Equal(want, have)

		// --- Then ---
		wMsg := "expected same timezone:\n" +
			"  trail: TA.Loc\n" +
			"   want: Europe/Warsaw\n" +
			"   have: UTC"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("equal nested", func(t *testing.T) {
		// --- Given ---
		want := types.TA{Str: "abc", TAp: &types.TA{Str: "abc"}}
		have := types.TA{Str: "abc", TAp: &types.TA{Str: "abc"}}

		// --- When ---
		err := Equal(want, have)

		// --- Then ---
		affirm.Nil(t, err)
	})

	t.Run("logs field paths", func(t *testing.T) {
		// --- Given ---
		want := types.TA{Str: "abc", TAp: &types.TA{Str: "abc"}}
		haveList := make([]string, 0)

		// --- When ---
		err := Equal(want, want, WithTrailLog(&haveList))

		// --- Then ---
		affirm.Nil(t, err)
		wantList := []string{
			"TA.Int",
			"TA.Str",
			"TA.Tim",
			"TA.Dur",
			"TA.Loc",
			"TA.TAp.Int",
			"TA.TAp.Str",
			"TA.TAp.Tim",
			"TA.TAp.Dur",
			"TA.TAp.Loc",
			"TA.TAp.TAp",
		}
		affirm.DeepEqual(t, wantList, haveList)
	})

	t.Run("not equal nested", func(t *testing.T) {
		// --- Given ---
		want := types.TA{Str: "abc", TAp: &types.TA{Str: "abc"}}
		have := types.TA{Str: "abc", TAp: &types.TA{Str: "xyz"}}

		// --- When ---
		err := Equal(want, have)

		// --- Then ---
		wMsg := "expected values to be equal:\n" +
			"  trail: TA.TAp.Str\n" +
			"   want: \"abc\"\n" +
			"   have: \"xyz\""
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("not equal nested by value", func(t *testing.T) {
		// --- Given ---
		want := types.TB{TAv: types.TA{Str: "abc"}}
		have := types.TB{TAv: types.TA{Str: "xyz"}}

		// --- When ---
		err := Equal(want, have)

		// --- Then ---
		wMsg := "expected values to be equal:\n" +
			"  trail: TB.TAv.Str\n" +
			"   want: \"abc\"\n" +
			"   have: \"xyz\""
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("not equal nested int slice", func(t *testing.T) {
		// --- Given ---
		s0 := types.TNested{SInt: []int{1, 2, 3}}
		s1 := types.TNested{SInt: []int{1, 7, 3}}

		// --- When ---
		err := Equal(s0, s1)

		// --- Then ---
		wMsg := "expected values to be equal:\n" +
			"  trail: TNested.SInt[1]\n" +
			"   want: 2\n" +
			"   have: 7"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("not equal nested object slice", func(t *testing.T) {
		// --- Given ---
		s0 := types.TNested{STA: []types.TA{{Str: "abc"}}}
		s1 := types.TNested{STA: []types.TA{{Str: "xyz"}}}

		// --- When ---
		err := Equal(s0, s1)

		// --- Then ---
		wMsg := "expected values to be equal:\n" +
			"  trail: TNested.STA[0].Str\n" +
			"   want: \"abc\"\n" +
			"   have: \"xyz\""
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("not equal nested object map string key", func(t *testing.T) {
		// --- Given ---
		s0 := types.TNested{MStrTyp: map[string]types.TA{"A": {Int: 0}}}
		s1 := types.TNested{MStrTyp: map[string]types.TA{"A": {Int: 1}}}

		// --- When ---
		err := Equal(s0, s1)

		// --- Then ---
		wMsg := "expected values to be equal:\n" +
			"  trail: TNested.MStrTyp[\"A\"].Int\n" +
			"   want: 0\n" +
			"   have: 1"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("not equal nested object map int key", func(t *testing.T) {
		// --- Given ---
		s0 := types.TNested{MIntTyp: map[int]types.TA{1: {Int: 0}}}
		s1 := types.TNested{MIntTyp: map[int]types.TA{1: {Int: 1}}}

		// --- When ---
		err := Equal(s0, s1)

		// --- Then ---
		wMsg := "expected values to be equal:\n" +
			"  trail: TNested.MIntTyp[1].Int\n" +
			"   want: 0\n" +
			"   have: 1"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("not equal nested object map int key int value", func(t *testing.T) {
		// --- Given ---
		s0 := types.TNested{MStrInt: map[string]int{"A": 0}}
		s1 := types.TNested{MStrInt: map[string]int{"A": 1}}

		// --- When ---
		err := Equal(s0, s1)

		// --- Then ---
		wMsg := "expected values to be equal:\n" +
			"  trail: TNested.MStrInt[\"A\"]\n" +
			"   want: 0\n" +
			"   have: 1"
		affirm.Equal(t, wMsg, err.Error())
	})

	// TODO(rz):
	// t.Run("custom field check", func(t *testing.T) {
	// 	// --- Given ---
	// 	// Different times within 1s.
	// 	want := types.TA{Tim: time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)}
	// 	have := types.TA{Tim: time.Date(2000, 1, 2, 3, 4, 6, 0, time.UTC)}
	//
	// 	within := func(want, have any, _ ...Option) error {
	// 		w := want.(time.Time)
	// 		h := have.(time.Time)
	// 		return TimeWithin(w, "1s", h)
	// 	}
	// 	opt := WithCheck(within, "TA.Tim")
	//
	// 	// --- When ---
	// 	err := Equal(want, have, opt)
	//
	// 	// --- Then ---
	// 	affirm.Nil(t, err)
	// })
	//
	// t.Run("error custom field check", func(t *testing.T) {
	// 	// --- Given ---
	// 	// Different times within 2s.
	// 	want := types.TA{Tim: time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)}
	// 	have := types.TA{Tim: time.Date(2000, 1, 2, 3, 4, 7, 0, time.UTC)}
	//
	// 	within := func(want, have any, opts ...Option) error {
	// 		ops := DefaultOptions(opts...)
	// 		w := want.(time.Time)
	// 		h := have.(time.Time)
	// 		return timeWithin(w, "1s", h, ops.Path)
	// 	}
	// 	opt := WithCheck(within, "TA.Tim")
	//
	// 	// --- When ---
	// 	err := Equal(want, have, opt)
	//
	// 	// --- Then ---
	// 	wMsg := "expected dates to be within:\n" +
	// 		"       path: TA.Tim\n" +
	// 		"       want: 2000-01-02T03:04:05Z (2000-01-02T03:04:05Z)\n" +
	// 		"       have: 2000-01-02T03:04:07Z (2000-01-02T03:04:07Z)\n" +
	// 		"   max diff: 1s\n" +
	// 		"  have diff: -2s"
	// 	affirm.Equal(t, wMsg, err.Error())
	// })

	t.Run("multiple field errors", func(t *testing.T) {
		// --- Given ---
		s0 := types.TA{Int: 42, Str: "abc"}
		s1 := types.TA{Int: 44, Str: "xyz"}

		// --- When ---
		err := Equal(s0, s1)

		// --- Then ---
		wMsg := "expected values to be equal:\n" +
			"  trail: TA.Int\n" +
			"   want: 42\n" +
			"   have: 44\n" +
			"expected values to be equal:\n" +
			"  trail: TA.Str\n" +
			"   want: \"abc\"\n" +
			"   have: \"xyz\""
		affirm.Equal(t, wMsg, err.Error())
	})

	// TODO(rz):
	// t.Run("skip field", func(t *testing.T) {
	// 	// --- Given ---
	// 	want := types.TA{Str: "abc"}
	// 	have := types.TA{Str: "xyz"}
	//
	// 	// --- When ---
	// 	err := Equal(want, have, WithSkip("TA.Str"))
	//
	// 	// --- Then ---
	// 	affirm.Nil(t, err)
	// })

	t.Run("error field embedded by value", func(t *testing.T) {
		// --- Given ---
		want := types.TB{TA: types.TA{Str: "abc"}}
		have := types.TB{TA: types.TA{Str: "xyz"}}

		// --- When ---
		err := Equal(want, have)

		// --- Then ---
		wMsg := "expected values to be equal:\n" +
			"  trail: TB.TA.Str\n" +
			"   want: \"abc\"\n" +
			"   have: \"xyz\""
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("using trail", func(t *testing.T) {
		// --- Given ---
		want := types.TA{Str: "abc"}
		have := types.TA{Str: "xyz"}

		// --- When ---
		err := Equal(want, have, WithTrail("type.field"))

		// --- Then ---
		wMsg := "expected values to be equal:\n" +
			"  trail: type.field.TA.Str\n" +
			"   want: \"abc\"\n" +
			"   have: \"xyz\""
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("want is not a struct error", func(t *testing.T) {
		// --- Given ---
		want := 42
		have := types.TA{Str: "xyz"}

		// --- When ---
		err := Equal(want, have)

		// --- Then ---
		wMsg := "expected values to be equal:\n" +
			"       want: 42\n" +
			"       have: \n" +
			"             {\n" +
			"               Int: 0,\n" +
			"               Str: \"xyz\",\n" +
			"               Tim: \"0001-01-01T00:00:00Z\",\n" +
			"               Dur: \"0s\",\n" +
			"               Loc: nil,\n" +
			"               TAp: nil,\n" +
			"             }\n" +
			"  want type: int\n" +
			"  have type: types.TA"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("have is not a struct error", func(t *testing.T) {
		// --- Given ---
		want := types.TA{Str: "xyz"}
		have := 42

		// --- When ---
		err := Equal(want, have)

		// --- Then ---
		wMsg := "expected values to be equal:\n" +
			"       want: \n" +
			"             {\n" +
			"               Int: 0,\n" +
			"               Str: \"xyz\",\n" +
			"               Tim: \"0001-01-01T00:00:00Z\",\n" +
			"               Dur: \"0s\",\n" +
			"               Loc: nil,\n" +
			"               TAp: nil,\n" +
			"             }\n" +
			"       have: 42\n" +
			"  want type: types.TA\n" +
			"  have type: int"
		affirm.Equal(t, wMsg, err.Error())
	})
}

// TODO(rz): we use Equal for everything now.
func Test_structEq(t *testing.T) {
	t.Run("want argument must be a structs", func(t *testing.T) {
		// --- Given ---
		want := types.TA{Str: "xyz"}
		have := 42

		// --- When ---
		err := Equal(want, have)

		// --- Then ---
		fmt.Println(err.Error()) // TODO(rz):

		wMsg := "expected values to be equal:\n" +
			"       want: \n" +
			"             {\n" +
			"               Int: 0,\n" +
			"               Str: \"xyz\",\n" +
			"               Tim: \"0001-01-01T00:00:00Z\",\n" +
			"               Dur: \"0s\",\n" +
			"               Loc: nil,\n" +
			"               TAp: nil,\n" +
			"             }\n" +
			"       have: 42\n" +
			"  want type: types.TA\n" +
			"  have type: int"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("have argument must be a structs", func(t *testing.T) {
		// --- Given ---
		want := 42
		have := types.TA{Str: "xyz"}

		// --- When ---
		err := Equal(want, have)

		// --- Then ---
		wMsg := "expected values to be equal:\n" +
			"       want: 42\n" +
			"       have: \n" +
			"             {\n" +
			"               Int: 0,\n" +
			"               Str: \"xyz\",\n" +
			"               Tim: \"0001-01-01T00:00:00Z\",\n" +
			"               Dur: \"0s\",\n" +
			"               Loc: nil,\n" +
			"               TAp: nil,\n" +
			"             }\n" +
			"  want type: int\n" +
			"  have type: types.TA"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("error with path set", func(t *testing.T) {
		// --- Given ---
		want := types.TA{Str: "xyz"}
		have := 42

		// --- When ---
		err := Equal(want, have, WithTrail("type.field"))

		// --- Then ---
		wMsg := "expected values to be equal:\n" +
			"      trail: type.field\n" +
			"       want: \n" +
			"             {\n" +
			"               Int: 0,\n" +
			"               Str: \"xyz\",\n" +
			"               Tim: \"0001-01-01T00:00:00Z\",\n" +
			"               Dur: \"0s\",\n" +
			"               Loc: nil,\n" +
			"               TAp: nil,\n" +
			"             }\n" +
			"       have: 42\n" +
			"  want type: types.TA\n" +
			"  have type: int"
		affirm.Equal(t, wMsg, err.Error())
	})
}
