// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"errors"
	"testing"
	"time"

	"github.com/ctx42/xtst/internal/affirm"
	"github.com/ctx42/xtst/internal/types"
)

func Test_Time(t *testing.T) {
	t.Run("equal both time.Time", func(t *testing.T) {
		// --- Given ---
		want := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		have := time.Date(2000, 1, 2, 4, 4, 5, 0, types.WAW)

		// --- When ---
		err := Time(want, have)

		// --- Then ---
		affirm.Nil(t, err)
		affirm.True(t, want.Equal(have))
	})

	t.Run("not equal both time.Time", func(t *testing.T) {
		// --- Given ---
		want := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		have := time.Date(2000, 1, 2, 4, 4, 6, 0, types.WAW)

		// --- When ---
		err := Time(want, have)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected equal dates:\n" +
			"\twant: 2000-01-02T03:04:05Z      (2000-01-02T03:04:05Z)\n" +
			"\thave: 2000-01-02T04:04:06+01:00 (2000-01-02T03:04:06Z)\n" +
			"\tdiff: -1s"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("equal want is string", func(t *testing.T) {
		// --- Given ---
		have := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)

		// --- When ---
		err := Time("2000-01-02T04:04:05+01:00", have)

		// --- Then ---
		affirm.Nil(t, err)
	})

	t.Run("not equal want is string", func(t *testing.T) {
		// --- Given ---
		have := time.Date(2000, 1, 2, 2, 4, 4, 0, time.UTC)

		// --- When ---
		err := Time("2000-01-02T03:04:05+01:00", have)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected equal dates:\n" +
			"\twant: 2000-01-02T03:04:05+01:00\n" +
			"\thave: 2000-01-02T02:04:04Z (2000-01-02T02:04:04Z)\n" +
			"\tdiff: 1s"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("not equal have is string", func(t *testing.T) {
		// --- Given ---
		want := time.Date(2000, 1, 2, 2, 4, 4, 0, time.UTC)

		// --- When ---
		err := Time(want, "2000-01-02T03:04:05+01:00")

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected equal dates:\n" +
			"\twant: 2000-01-02T02:04:04Z (2000-01-02T02:04:04Z)\n" +
			"\thave: 2000-01-02T03:04:05+01:00\n" +
			"\tdiff: -1s"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("invalid want date format", func(t *testing.T) {
		// --- When ---
		err := Time("2022-02-18", time.Now())

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "[want] failed to parse time:\n" +
			"\tformat: 2006-01-02T15:04:05.999999999Z07:00\n" +
			"\t value: 2022-02-18"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("invalid have date format", func(t *testing.T) {
		// --- When ---
		err := Time(time.Now(), "2022-02-18")

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "[have] failed to parse time:\n" +
			"\tformat: 2006-01-02T15:04:05.999999999Z07:00\n" +
			"\t value: 2022-02-18"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("log message with trail", func(t *testing.T) {
		// --- Given ---
		want := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		have := time.Date(2000, 1, 2, 4, 4, 6, 0, types.WAW)
		opt := WithTrail("type.field")

		// --- When ---
		err := Time(want, have, opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected equal dates:\n" +
			"\ttrail: type.field\n" +
			"\t want: 2000-01-02T03:04:05Z      (2000-01-02T03:04:05Z)\n" +
			"\t have: 2000-01-02T04:04:06+01:00 (2000-01-02T03:04:06Z)\n" +
			"\t diff: -1s"
		affirm.Equal(t, wMsg, err.Error())
	})
}

func Test_Zone(t *testing.T) {
	t.Run("equal", func(t *testing.T) {
		// --- When ---
		err := Zone(time.UTC, time.UTC)

		// --- Then ---
		affirm.Nil(t, err)
	})

	t.Run("want nil", func(t *testing.T) {
		// --- When ---
		err := Zone(nil, time.UTC)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected timezone:\n" +
			"\twhich: want\n" +
			"\t want: <not-nil>\n" +
			"\t have: <nil>"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("have nil", func(t *testing.T) {
		// --- When ---
		err := Zone(time.UTC, nil)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected timezone:\n" +
			"\twhich: have\n" +
			"\t want: <not-nil>\n" +
			"\t have: <nil>"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("not equal", func(t *testing.T) {
		// --- When ---
		err := Zone(time.UTC, types.WAW)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected same timezone:\n" +
			"\twant: UTC\n" +
			"\thave: Europe/Warsaw"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("log message with trail", func(t *testing.T) {
		// --- Given ---
		opt := WithTrail("type.field")

		// --- When ---
		err := Zone(time.UTC, types.WAW, opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected same timezone:\n" +
			"\ttrail: type.field\n" +
			"\t want: UTC\n" +
			"\t have: Europe/Warsaw"
		affirm.Equal(t, wMsg, err.Error())
	})
}

func Test_FormatDates(t *testing.T) {
	t.Run("default format", func(t *testing.T) {
		// --- Given ---
		tim0 := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		tim1 := time.Date(2001, 1, 2, 3, 4, 5, 0, time.UTC)

		// --- When ---
		have0, have1 := FormatDates(tim0, tim1)

		// --- Then ---
		affirm.Equal(t, "2000-01-02T03:04:05Z (2000-01-02T03:04:05Z)", have0)
		affirm.Equal(t, "2001-01-02T03:04:05Z (2001-01-02T03:04:05Z)", have1)
	})

	t.Run("custom format", func(t *testing.T) {
		// --- Given ---
		tim0 := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		tim1 := time.Date(2001, 1, 2, 3, 4, 5, 0, time.UTC)
		opt := WithTimeFormat(time.ANSIC)

		// --- When ---
		have0, have1 := FormatDates(tim0, tim1, opt)

		// --- Then ---
		affirm.Equal(t, "Sun Jan  2 03:04:05 2000 (Sun Jan  2 03:04:05 2000)", have0)
		affirm.Equal(t, "Tue Jan  2 03:04:05 2001 (Tue Jan  2 03:04:05 2001)", have1)
	})

	t.Run("tim0 date string is longer", func(t *testing.T) {
		// --- Given ---
		tim0 := time.Date(2001, 1, 2, 3, 4, 5, 0, types.WAW)
		tim1 := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)

		// --- When ---
		have0, have1 := FormatDates(tim0, tim1)

		// --- Then ---
		affirm.Equal(t, "2001-01-02T03:04:05+01:00 (2001-01-02T02:04:05Z)", have0)
		affirm.Equal(t, "2000-01-02T03:04:05Z      (2000-01-02T03:04:05Z)", have1)
	})

	t.Run("tim1 date string is longer", func(t *testing.T) {
		// --- Given ---
		tim0 := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		tim1 := time.Date(2001, 1, 2, 3, 4, 5, 0, types.WAW)

		// --- When ---
		have0, have1 := FormatDates(tim0, tim1)

		// --- Then ---
		affirm.Equal(t, "2000-01-02T03:04:05Z      (2000-01-02T03:04:05Z)", have0)
		affirm.Equal(t, "2001-01-02T03:04:05+01:00 (2001-01-02T02:04:05Z)", have1)
	})
}

func Test_getTime(t *testing.T) {
	t.Run("wrong time format", func(t *testing.T) {
		// --- Given ---
		opt := WithTimeFormat(time.RFC3339)

		// --- When ---
		haveTim, haveType, err := getTime("2000-01-02", opt)

		// --- Then ---
		affirm.True(t, haveTim.IsZero())
		affirm.Equal(t, timeTypeStr, haveType)
		wMsg := "failed to parse time:\n" +
			"\tformat: 2006-01-02T15:04:05Z07:00\n" +
			"\t value: 2000-01-02"
		affirm.Equal(t, wMsg, err.Error())
		affirm.True(t, errors.Is(err, ErrTimeParse))
	})

	t.Run("empty option time format", func(t *testing.T) {
		// --- Given ---
		opt := WithTimeFormat("")

		// --- When ---
		have, haveType, err := getTime("2000-01-02", opt)

		// --- Then ---
		affirm.True(t, have.IsZero())
		affirm.Equal(t, timeTypeStr, haveType)
		wMsg := "failed to parse time:\n" +
			"\tformat: \n" +
			"\t value: 2000-01-02\n" +
			"\t error: extra text: \"2000-01-02\""
		affirm.Equal(t, wMsg, err.Error())
		affirm.True(t, errors.Is(err, ErrTimeParse))
	})

	t.Run("log message with trail", func(t *testing.T) {
		// --- Given ---
		opts := []Option{
			WithTimeFormat(time.RFC3339),
			WithTrail("type.field"),
		}

		// --- When ---
		have, haveType, err := getTime("2000-01-02", opts...)

		// --- Then ---
		affirm.True(t, have.IsZero())
		affirm.Equal(t, timeTypeStr, haveType)
		wMsg := "failed to parse time:\n" +
			"\t trail: type.field\n" +
			"\tformat: 2006-01-02T15:04:05Z07:00\n" +
			"\t value: 2000-01-02"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("unsupported type", func(t *testing.T) {
		// --- When ---
		have, haveType, err := getTime(true)

		// --- Then ---
		affirm.True(t, have.IsZero())
		affirm.Equal(t, "", haveType)
		wMsg := "failed to parse time:\n" +
			"\tcause: not supported time type"
		affirm.Equal(t, wMsg, err.Error())
		affirm.True(t, errors.Is(err, ErrTimeType))
	})
}

func Test_getTime_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		opts    []Option
		have    any
		haveRep timeRep
		want    time.Time
		wantTZ  *time.Location
	}{
		{
			"time.Time in UTC",
			nil,
			time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
			timeTypeTim,
			time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
			time.UTC,
		},
		{
			"time.Time in WAW",
			nil,
			time.Date(2000, 1, 2, 3, 4, 5, 0, types.WAW),
			timeTypeTim,
			time.Date(2000, 1, 2, 3, 4, 5, 0, types.WAW),
			types.WAW,
		},
		{
			"RFC3339",
			nil,
			"2000-01-02T03:04:05+01:00",
			timeTypeStr,
			time.Date(2000, 1, 2, 2, 4, 5, 0, time.UTC),
			time.UTC,
		},
		{
			"Unix timestamp int",
			nil,
			946778645,
			timeTypeInt,
			time.Date(2000, 1, 2, 2, 4, 5, 0, time.UTC),
			time.UTC,
		},
		{
			"Unix timestamp int64",
			nil,
			int64(946778645),
			timeTypeInt64,
			time.Date(2000, 1, 2, 2, 4, 5, 0, time.UTC),
			time.UTC,
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			haveTim, haveType, err := getTime(tc.have, tc.opts...)

			// --- Then ---
			affirm.Nil(t, err)
			affirm.Equal(t, tc.haveRep, haveType)
			affirm.True(t, tc.want.Equal(haveTim))
			affirm.Equal(t, tc.wantTZ.String(), haveTim.Location().String())
		})
	}
}

func Test_getDur(t *testing.T) {
	t.Run("invalid string duration", func(t *testing.T) {
		// --- When ---
		haveDur, haveRep, err := getDur("abc")

		// --- Then ---
		affirm.Equal(t, time.Duration(0), haveDur)
		affirm.Equal(t, durTypeStr, haveRep)
		wMsg := "failed to parse duration:\n\tvalue: abc"
		affirm.Equal(t, wMsg, err.Error())
		affirm.True(t, errors.Is(err, ErrDurParse))
	})

	t.Run("error unsupported type", func(t *testing.T) {
		// --- When ---
		haveDur, haveRep, err := getDur(true)

		// --- Then ---
		affirm.Equal(t, time.Duration(0), haveDur)
		affirm.Equal(t, "", haveRep)
		wMsg := "failed to parse duration:\n" +
			"\tcause: not supported duration type"
		affirm.Equal(t, wMsg, err.Error())
		affirm.True(t, errors.Is(err, ErrDurType))
	})

	t.Run("log message with trail", func(t *testing.T) {
		// --- Given ---
		opt := WithTrail("type.field")

		// --- When ---
		haveDur, haveRep, err := getDur(true, opt)

		// --- Then ---
		affirm.Equal(t, time.Duration(0), haveDur)
		affirm.Equal(t, "", haveRep)
		wMsg := "failed to parse duration:\n" +
			"\ttrail: type.field\n" +
			"\tcause: not supported duration type"
		affirm.Equal(t, wMsg, err.Error())
		affirm.True(t, errors.Is(err, ErrDurType))
	})
}

func Test_getDur_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have    any
		haveRep durRep
		want    time.Duration
	}{
		{"time.Duration", time.Second, durTypeDur, time.Second},
		{"time.Duration as string", "1s", durTypeStr, time.Second},
		{"time.Duration as int", 12345678, durTypeInt, time.Duration(12345678)},
		{
			"time.Duration as int",
			int64(12345678),
			durTypeInt64,
			time.Duration(12345678),
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			haveDur, haveRep, err := getDur(tc.have)

			// --- Then ---
			affirm.Nil(t, err)
			affirm.Equal(t, tc.haveRep, haveRep)
			affirm.Equal(t, tc.want, haveDur)
		})
	}
}
