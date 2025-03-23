// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"errors"
	"fmt"
	"math"
	"strconv"
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
	wTim, wStr, _, err := getTime(want, opts...)
	if err != nil {
		return notice.From(err, "want")
	}
	hTim, hStr, _, err := getTime(have, opts...)
	if err != nil {
		return notice.From(err, "have")
	}
	if wTim.Equal(hTim) {
		return nil
	}

	diff := wTim.Sub(hTim)
	ops := DefaultOptions().set(opts)
	wantFmt, haveFmt := formatDates(wTim, wStr, hTim, hStr)
	return notice.New("expected equal dates").
		Trail(ops.Trail).
		Want("%s", wantFmt).
		Have("%s", haveFmt).
		Append("diff", "%s", diff.String())
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
	wTim, wStr, _, err := getTime(want, opts...)
	if err != nil {
		return notice.From(err, "want")
	}
	hTim, hStr, _, err := getTime(have, opts...)
	if err != nil {
		return notice.From(err, "have")
	}

	if !wTim.Equal(hTim) {
		diff := wTim.Sub(hTim)
		ops := DefaultOptions().set(opts)
		wantFmt, haveFmt := formatDates(wTim, wStr, hTim, hStr)
		return notice.New("expected equal dates").
			Trail(ops.Trail).
			Want("%s", wantFmt).
			Have("%s", haveFmt).
			Append("diff", "%s", diff.String())
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
	wTim, wStr, _, err := getTime(want, opts...)
	if err != nil {
		return notice.From(err, "want")
	}
	hTim, hStr, _, err := getTime(have, opts...)
	if err != nil {
		return notice.From(err, "have")
	}
	dur, durStr, _, err := getDur(within, opts...)
	if err != nil {
		return err
	}

	diff := wTim.Sub(hTim.In(wTim.Location()))
	if math.Abs(float64(diff)) <= math.Abs(float64(dur)) {
		return nil
	}

	wantFmt, haveFmt := formatDates(wTim, wStr, hTim, hStr)
	ops := DefaultOptions().set(opts)
	return notice.New("expected dates to be within").
		Trail(ops.Trail).
		Want("%s", wantFmt).
		Have("%s", haveFmt).
		Append("max diff", "%s", durStr).
		Append("have diff", "%s", diff.String())
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
	wDur, wStr, _, err := getDur(want, opts...)
	if err != nil {
		return notice.From(err, "want")
	}
	hDur, hStr, _, err := getDur(have, opts...)
	if err != nil {
		return notice.From(err, "have")
	}

	if wDur == hDur {
		return nil
	}
	ops := DefaultOptions().set(opts)
	return notice.New("expected equal time durations").
		Trail(ops.Trail).
		Want("%s", wStr).
		Have("%s", hStr)
}

// formatDates formats two dates for comparison in an error message.
//
// Example:
//
//	2000-01-02T03:04:05Z ( 2000-01-02T03:04:05Z      )
//	2001-01-02T02:04:05Z ( 2001-01-02T03:04:05+01:00 )
func formatDates(
	wTim time.Time, wTimStr string,
	hTim time.Time, hTimStr string,
) (string, string) {

	wTimUTC := wTim.In(time.UTC).Format(time.RFC3339Nano)
	hTimUTC := hTim.In(time.UTC).Format(time.RFC3339Nano)

	wTimStrLen := len(wTimStr)
	hTimStrLen := len(hTimStr)

	var wTimPad, hTimPad string
	if wTimStrLen < hTimStrLen {
		wTimPad = strings.Repeat(" ", hTimStrLen-wTimStrLen)
	}
	if hTimStrLen < wTimStrLen {
		hTimPad = strings.Repeat(" ", wTimStrLen-hTimStrLen)
	}

	var want, have string
	if wTimUTC == wTimStr {
		want = fmt.Sprintf("%s", wTimUTC)
	} else {
		want = fmt.Sprintf("%s ( %s %s)", wTimUTC, wTimStr, wTimPad)

	}

	if hTimUTC == hTimStr {
		have = fmt.Sprintf("%s", hTimUTC)
	} else {
		have = fmt.Sprintf("%s ( %s %s)", hTimUTC, hTimStr, hTimPad)
	}

	return want, have
}

// getTime returns date represented by "tim", its string representation and
// type of the argument passed. The "tim" might be date representation in form
// of string, int, int64 or [time.Time]. For string representations the
// [Options.TimeFormat] is used during parsing and the returned date is always
// in UTC. The int and int64 types are interpreted as Unix Timestamp and the
// date returned is also in UTC.
//
// When error is returned it will always have [ErrTimeParse], [ErrTimeType] in
// its chain.
func getTime(tim any, opts ...Option) (time.Time, string, timeRep, error) {
	ops := DefaultOptions().set(opts)
	switch val := tim.(type) {
	case time.Time:
		return val, val.Format(time.RFC3339Nano), timeTypeTim, nil

	case string:
		have, err := time.Parse(ops.TimeFormat, val)
		if err == nil {
			return have.UTC(), val, timeTypeStr, nil
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
		return time.Time{}, val, timeTypeStr, err

	case int:
		str := strconv.Itoa(val)
		return time.Unix(int64(val), 0).UTC(), str, timeTypeInt, nil

	case int64:
		str := strconv.FormatInt(val, 10)
		return time.Unix(val, 0).UTC(), str, timeTypeInt64, nil

	default:
		str := fmt.Sprintf("%v", val)
		msg := notice.New("failed to parse time").
			Trail(ops.Trail).
			Append("cause", "%s", ErrTimeType).
			Wrap(ErrTimeType)
		return time.Time{}, str, "", msg
	}
}

// getDur returns duration represented by "dur", its string representation and
// type of the argument passed. The "dur" might be duration represented by
// string, int, int64 or [time.Duration].
//
// When error is returned it will always have [ErrDurParse], [ErrDurType] in
// its chain.
func getDur(dur any, opts ...Option) (time.Duration, string, durRep, error) {
	switch val := dur.(type) {
	case time.Duration:
		return val, val.String(), durTypeDur, nil

	case string:
		have, err := time.ParseDuration(val)
		if err == nil {
			return have, val, durTypeStr, nil
		}

		ops := DefaultOptions().set(opts)
		msg := notice.New("failed to parse duration").
			Trail(ops.Trail).
			Append("value", "%s", dur).
			Wrap(ErrDurParse)
		return 0, val, durTypeStr, msg

	case int:
		str := strconv.Itoa(val)
		return time.Duration(val), str, durTypeInt, nil

	case int64:
		str := fmt.Sprintf("%v", val)
		return time.Duration(val), str, durTypeInt64, nil

	default:
		str := fmt.Sprintf("%v", val)
		ops := DefaultOptions().set(opts)
		msg := notice.New("failed to parse duration").
			Trail(ops.Trail).
			Append("cause", "%s", ErrDurType).
			Wrap(ErrDurType)
		return 0, str, "", msg
	}
}
