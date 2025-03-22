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

func Test_TimeEqual(t *testing.T) {
	t.Run("equal both time.Time", func(t *testing.T) {
		// --- Given ---
		want := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		have := time.Date(2000, 1, 2, 4, 4, 5, 0, types.WAW)

		// --- When ---
		err := TimeEqual(want, have)

		// --- Then ---
		affirm.Nil(t, err)
		affirm.True(t, want.Equal(have))
	})

	t.Run("not equal both time.Time", func(t *testing.T) {
		// --- Given ---
		want := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		have := time.Date(2000, 1, 2, 4, 4, 6, 0, types.WAW)

		// --- When ---
		err := TimeEqual(want, have)

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
		err := TimeEqual("2000-01-02T04:04:05+01:00", have)

		// --- Then ---
		affirm.Nil(t, err)
	})

	t.Run("not equal want is string", func(t *testing.T) {
		// --- Given ---
		have := time.Date(2000, 1, 2, 2, 4, 4, 0, time.UTC)

		// --- When ---
		err := TimeEqual("2000-01-02T03:04:05+01:00", have)

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
		err := TimeEqual(want, "2000-01-02T03:04:05+01:00")

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
		err := TimeEqual("2022-02-18", time.Now())

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "[want] failed to parse time:\n" +
			"\tformat: 2006-01-02T15:04:05.999999999Z07:00\n" +
			"\t value: 2022-02-18"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("invalid have date format", func(t *testing.T) {
		// --- When ---
		err := TimeEqual(time.Now(), "2022-02-18")

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
		err := TimeEqual(want, have, opt)

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

		opts     []Option
		have     any
		haveType timeRep
		want     time.Time
		wantTZ   *time.Location
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
			affirm.Equal(t, tc.haveType, haveType)
			affirm.True(t, tc.want.Equal(haveTim))
			affirm.Equal(t, tc.wantTZ.String(), haveTim.Location().String())
		})
	}
}
