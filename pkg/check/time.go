// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ctx42/xtst/pkg/notice"
)

// Sentinel errors.
var (
	// ErrTimeType is returned when time representation is not supported
	ErrTimeType = fmt.Errorf("not supported time type")

	// ErrTimeParsing is used when date parsing fails for whatever reason.
	ErrTimeParsing = fmt.Errorf("time parsing")
)

// timeRep is time representation.
type timeRep string

// The time representations the [TimeEqual] supports.
const (
	timeTypeTim   timeRep = "time.Time"
	timeTypeStr   timeRep = "string"
	timeTypeInt   timeRep = "int"
	timeTypeInt64 timeRep = "int64"
)

// TimeEqual checks dates are equal. The "want" and "have" might be date
// representations in form of string, int, int64 or [time.Time]. For string
// representations the [Options.TimeFormat] is used during parsing and the
// returned date is always in UTC. The int and int64 types are interpreted as
// Unix Timestamp and the date returned is also in UTC. Returns nil if dates
// are the same, otherwise it returns an error with a message indicating the
// expected and actual values.
func TimeEqual(want, have any, opts ...Option) error {
	wTim, wTyp, err := getTime(want, opts...)
	if err != nil {
		return notice.From(err, "want")
	}
	hTim, hTyp, err := getTime(have, opts...)
	if err != nil {
		return notice.From(err, "have")
	}
	if err = timeEqual(wTim, hTim, opts...); err != nil {
		if !errors.Is(err, ErrTimeParsing) && wTyp == timeTypeStr {
			err = notice.From(err).Want("%s", want)
		}
		if !errors.Is(err, ErrTimeParsing) && hTyp == timeTypeStr {
			err = notice.From(err).Have("%s", have)
		}
		return err
	}
	return nil
}

// timeEqual is internal implementation of [TimeEqual] which takes field path.
func timeEqual(want, have time.Time, opts ...Option) error {
	ops := DefaultOptions().set(opts)
	if want.Equal(have) {
		return nil
	}
	wantFmt, haveFmt := FormatDates(want, have)
	diff := want.Sub(have)
	return notice.New("expected equal dates").
		Trail(ops.Trail).
		Want("%s", wantFmt).
		Have("%s", haveFmt).
		Append("diff", "%s", diff.String())
}

// FormatDates formats two dates for comparison in an error message.
//
// Example:
//
//	2000-01-02T03:04:05Z (2000-01-02T03:04:05Z)
//	2001-01-02T03:04:05+01:00 (2001-01-02T02:04:05Z)
func FormatDates(tim0, tim1 time.Time, opts ...Option) (string, string) {
	ops := DefaultOptions().set(opts)
	tim0date := tim0.Format(ops.TimeFormat)
	tim1date := tim1.Format(ops.TimeFormat)
	tim0inUTC := tim0.In(time.UTC).Format(ops.TimeFormat)
	tim1inUTC := tim1.In(time.UTC).Format(ops.TimeFormat)
	tim0len := len(tim0date)
	tim1len := len(tim1date)
	var tim0pad, tim1pad string
	if tim0len < tim1len {
		tim0pad = strings.Repeat(" ", tim1len-tim0len)
	}
	if tim1len < tim0len {
		tim1pad = strings.Repeat(" ", tim0len-tim1len)
	}
	ret0 := fmt.Sprintf("%s %s(%s)", tim0date, tim0pad, tim0inUTC)
	ret1 := fmt.Sprintf("%s %s(%s)", tim1date, tim1pad, tim1inUTC)
	return ret0, ret1
}

// getTime returns [time.Time] based on "tim" and the recognized type of the
// argument passed to the function. For values that need to be parsed or
// interpreted it always returns time in UTC.
//
// When "tim" is of type:
//   - time.Time  - the "tim" is returned as it is.
//   - string     - the "tim" is parsed using "format".
//   - int, int64 - the "tim" is treated as Unix Timestamp.
//
// For other types function returns [ErrTimeType].
func getTime(tim any, opts ...Option) (time.Time, timeRep, error) {
	ops := DefaultOptions().set(opts)
	switch val := tim.(type) {
	case time.Time:
		return val, timeTypeTim, nil

	case string:
		have, err := time.Parse(ops.TimeFormat, val)
		if err == nil {
			return have.UTC(), timeTypeStr, nil
		}

		var pe *time.ParseError
		if errors.As(err, &pe) {
			msg := notice.New("failed to parse time").
				Trail(ops.Trail).
				Append("format", "%s", ops.TimeFormat).
				Append("value", "%s", pe.Value).
				Wrap(ErrTimeParsing)
			if pe.Message != "" {
				msg = msg.Append("error", "%s", strings.Trim(pe.Message, " :"))
			}
			err = msg
		}
		return time.Time{}, timeTypeStr, err

	case int:
		return time.Unix(int64(val), 0).UTC(), timeTypeInt, nil

	case int64:
		return time.Unix(val, 0).UTC(), timeTypeInt64, nil

	default:
		msg := notice.New("failed to parse time").
			Trail(ops.Trail).
			Append("cause", "%s", ErrTimeType).
			Wrap(ErrTimeType)
		return time.Time{}, "", msg
	}
}
