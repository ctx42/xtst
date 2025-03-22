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

// ErrTimeType is returned when value of type "any" is not [time.Time], string,
// int64.
var ErrTimeType = fmt.Errorf("not supported time type")

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

// getTime returns [time.Time] based on "tim". For values that need to be
// parsed or interpreted it always returns time in UTC.
//
// When "tim" is of type:
//   - time.Time  - the "tim" is returned as it is.
//   - string     - the "tim" is parsed using "format".
//   - int, int64 - the "tim" is treated as Unix Timestamp.
//
// For other types function returns [ErrTimeType].
func getTime(tim any, opt Options) (time.Time, error) {
	switch val := tim.(type) {
	case time.Time:
		return val, nil

	case string:
		have, err := time.Parse(opt.TimeFormat, val)
		if err == nil {
			return have.UTC(), nil
		}

		var pe *time.ParseError
		if errors.As(err, &pe) {
			msg := notice.New("failed to parse time").
				Trail(opt.Trail).
				Append("format", "%s", opt.TimeFormat).
				Append("value", "%q", pe.Value)
			if pe.Message != "" {
				msg = msg.Append("error", "%s", strings.Trim(pe.Message, " :"))
			}
			err = msg
		}
		return time.Time{}, err

	case int:
		return time.Unix(int64(val), 0).UTC(), nil

	case int64:
		return time.Unix(val, 0).UTC(), nil

	default:
		msg := notice.New("failed to parse time").
			Trail(opt.Trail).
			Append("cause", "%s", ErrTimeType).
			Wrap(ErrTimeType)
		return time.Time{}, msg
	}
}
