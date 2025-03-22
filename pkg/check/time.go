// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/ctx42/xtst/pkg/notice"
)

// Sentinel errors.
var (
	// ErrTimeType is returned when time representation is not supported.
	ErrTimeType = fmt.Errorf("not supported time type")

	// ErrTimeParse is used when date parsing fails for whatever reason.
	ErrTimeParse = fmt.Errorf("time parsing")

	// ErrDurType is returned when duration representation is not supported.
	ErrDurType = fmt.Errorf("not supported duration type")

	// ErrDurParse is used when duration parsing fails for whatever reason.
	ErrDurParse = fmt.Errorf("duration parsing")
)

// timeRep is time representation.
type timeRep string

// The time representations the [Time] supports.
const (
	timeTypeTim   timeRep = "tim-tim"
	timeTypeStr   timeRep = "tim-string"
	timeTypeInt   timeRep = "tim-int"
	timeTypeInt64 timeRep = "tim-int64"
)

// durRep is duration representation.
type durRep string

// The duration representations the [Duration] supports.
const (
	durTypeDur   durRep = "dur-dur"
	durTypeStr   durRep = "dur-str"
	durTypeInt   durRep = "dur-int"
	durTypeInt64 durRep = "dur-int64"
)

// Time checks "want" and "have" dates are equal. Returns nil if they are,
// otherwise returns an error with a message indicating the expected and actual
// values.
//
// The "want" and "have" might be date representations in form of string, int,
// int64 or [time.Time]. For string representations the [Options.TimeFormat] is
// used during parsing and the returned date is always in UTC. The int and
// int64 types are interpreted as Unix Timestamp and the date returned is also
// in UTC.
func Time(want, have any, opts ...Option) error {
	wTim, wTyp, err := getTime(want, opts...)
	if err != nil {
		return notice.From(err, "want")
	}
	hTim, hTyp, err := getTime(have, opts...)
	if err != nil {
		return notice.From(err, "have")
	}
	if err = timeEqual(wTim, hTim, opts...); err != nil {
		if !errors.Is(err, ErrTimeParse) && wTyp == timeTypeStr {
			err = notice.From(err).Want("%s", want)
		}
		if !errors.Is(err, ErrTimeParse) && hTyp == timeTypeStr {
			err = notice.From(err).Have("%s", have)
		}
		return err
	}
	return nil
}

// TimeExact checks "want" and "have" dates are equal and are in the same
// timezone. Returns nil they are, otherwise returns an error with a message
// indicating the expected and actual values.
//
// The "want" and "have" might be date representations in form of string, int,
// int64 or [time.Time]. For string representations the [Options.TimeFormat] is
// used during parsing and the returned date is always in UTC. The int and
// int64 types are interpreted as Unix Timestamp and the date returned is also
// in UTC.
func TimeExact(want, have any, opts ...Option) error {
	wTim, wTyp, err := getTime(want, opts...)
	if err != nil {
		return notice.From(err, "want")
	}
	hTim, hTyp, err := getTime(have, opts...)
	if err != nil {
		return notice.From(err, "have")
	}
	if err = timeEqual(wTim, hTim, opts...); err != nil {
		if !errors.Is(err, ErrTimeParse) && wTyp == timeTypeStr {
			err = notice.From(err).Want("%s", want)
		}
		if !errors.Is(err, ErrTimeParse) && hTyp == timeTypeStr {
			err = notice.From(err).Have("%s", have)
		}
		return err
	}

	return Zone(wTim.Location(), hTim.Location(), opts...)
}

// Within checks "want" and "have" dates are equal "within" given duration.
// Returns nil if they are, otherwise returns an error with a message
// indicating the expected and actual values.
//
// The "want" and "have" might be date representations in form of string, int,
// int64 or [time.Time]. For string representations the [Options.TimeFormat] is
// used during parsing and the returned date is always in UTC. The int and
// int64 types are interpreted as Unix Timestamp and the date returned is also
// in UTC.
//
// The "within" might be duration representation in form of string, int, int64
// or [time.Duration].
func Within(want, within, have any, opts ...Option) error {
	wTim, _, err := getTime(want, opts...)
	if err != nil {
		return notice.From(err, "want")
	}
	hTim, _, err := getTime(have, opts...)
	if err != nil {
		return notice.From(err, "have")
	}

	diel := hTim.In(wTim.Location())
	diff := wTim.Sub(diel)

	dur, _, err := getDur(within, opts...)
	if err != nil {
		return err
	}
	if math.Abs(float64(diff)) <= math.Abs(float64(dur)) {
		return nil
	}

	wantFmt, haveFmt := FormatDates(wTim, hTim)
	ops := DefaultOptions().set(opts)
	return notice.New("expected dates to be within").
		Trail(ops.Trail).
		Want("%s", wantFmt).
		Have("%s", haveFmt).
		Append("max diff", "%s", dur).
		Append("have diff", "%s", diff.String())
}

// timeEqual checks "want" and have dates are equal. Returns nil if they are,
// otherwise returns an error with a message indicating the expected and actual
// values.
func timeEqual(want, have time.Time, opts ...Option) error {
	if want.Equal(have) {
		return nil
	}

	ops := DefaultOptions().set(opts)
	wantFmt, haveFmt := FormatDates(want, have)
	diff := want.Sub(have)
	return notice.New("expected equal dates").
		Trail(ops.Trail).
		Want("%s", wantFmt).
		Have("%s", haveFmt).
		Append("diff", "%s", diff.String())
}

// Zone checks "want" and "have" timezones are equal. Returns nil if they are,
// otherwise returns an error with a message indicating the expected and actual
// values.
func Zone(want, have *time.Location, opts ...Option) error {
	if want == nil {
		ops := DefaultOptions().set(opts)
		return notice.New("expected timezone").
			Trail(ops.Trail).
			Append("which", "want").
			Want("<not-nil>").
			Have("<nil>")
	}
	if have == nil {
		ops := DefaultOptions().set(opts)
		return notice.New("expected timezone").
			Trail(ops.Trail).
			Append("which", "have").
			Want("<not-nil>").
			Have("<nil>")
	}
	if want.String() == have.String() {
		return nil
	}

	ops := DefaultOptions().set(opts)
	return notice.New("expected same timezone").
		Trail(ops.Trail).
		Want("%s", want.String()).
		Have("%s", have.String())
}

// Duration checks "want" and "have" durations are equal. Returns nil if they
// are, otherwise returns an error with a message indicating the expected and
// actual values.
//
// The "want" and "have" might be duration representation in form of string,
// int, int64 or [time.Duration].
func Duration(want, have any, opts ...Option) error {
	wDur, _, err := getDur(want, opts...)
	if err != nil {
		return notice.From(err, "want")
	}
	hDur, _, err := getDur(have, opts...)
	if err != nil {
		return notice.From(err, "have")
	}

	if wDur == hDur {
		return nil
	}
	ops := DefaultOptions().set(opts)
	return notice.New("expected equal time durations").
		Trail(ops.Trail).
		Want("%s", wDur.String()).
		Have("%s", hDur.String())
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

// getTime returns date represented by "tim". The "tim" might be date
// representation in form of string, int, int64 or [time.Time]. For string
// representations the [Options.TimeFormat] is used during parsing and the
// returned date is always in UTC. The int and int64 types are interpreted as
// Unix Timestamp and the date returned is also in UTC.
//
// When error is returned it will always have [ErrTimeParse], [ErrTimeType] in
// its chain.
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
				Wrap(ErrTimeParse)
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

// getDur returns duration represented by "dur". The "dur" might be duration
// represented by string, int, int64 or [time.Duration].
//
// When error is returned it will always have [ErrDurParse], [ErrDurType] in
// its chain.
func getDur(dur any, opts ...Option) (time.Duration, durRep, error) {
	switch val := dur.(type) {
	case time.Duration:
		return val, durTypeDur, nil

	case string:
		have, err := time.ParseDuration(val)
		if err == nil {
			return have, durTypeStr, nil
		}

		ops := DefaultOptions().set(opts)
		msg := notice.New("failed to parse duration").
			Trail(ops.Trail).
			Append("value", "%s", dur).
			Wrap(ErrDurParse)
		return 0, durTypeStr, msg

	case int:
		return time.Duration(val), durTypeInt, nil

	case int64:
		return time.Duration(val), durTypeInt64, nil

	default:
		ops := DefaultOptions().set(opts)
		msg := notice.New("failed to parse duration").
			Trail(ops.Trail).
			Append("cause", "%s", ErrDurType).
			Wrap(ErrDurType)
		return 0, "", msg
	}
}
