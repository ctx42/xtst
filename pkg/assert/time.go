// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package assert

import (
	"time"

	"github.com/ctx42/xtst/pkg/check"
	"github.com/ctx42/xtst/pkg/tester"
)

// Time asserts both arguments are dates and are equal. The "want" and "have"
// might be date representations in form of string, int, int64 or [time.Time].
// For string representations the [check.Options.TimeFormat] is used during
// parsing and the returned date is always in UTC. The int and int64 types are
// interpreted as Unix Timestamp and the date returned is also in UTC. Returns
// true if they are, otherwise marks the test as failed, writes error message
// to test log and returns false.
func Time(t tester.T, want, have any, opts ...check.Option) bool {
	t.Helper()
	if e := check.Time(want, have, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}

// TimeExact asserts dates are equal and have the same timezone. The "want" and
// "have" might be date representations in form of string, int, int64 or
// [time.Time]. For string representations the [check.Options.TimeFormat] is
// used during parsing and the returned date is always in UTC. The int and
// int64 types are interpreted as Unix Timestamp and the date returned is also
// in UTC. Returns nil if dates are the same, otherwise it returns an error
// with a message indicating the expected and actual values.
func TimeExact(t tester.T, want, have any, opts ...check.Option) bool {
	t.Helper()
	if e := check.TimeExact(want, have, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}

// Within asserts "have" date is "within" duration of "want" date. Returns true
// if it's, otherwise marks the test as failed, writes error message to test
// log and returns false.
//
// The "want" and "have" might be date representations in form of string, int,
// int64 or [time.Time]. For string representations the
// [check.Options.TimeFormat] is used during parsing and the returned date is
// always in UTC. The int and int64 types are interpreted as Unix Timestamp and
// the date returned is also in UTC.
func Within(t tester.T, want, within, have any, opts ...check.Option) bool {
	t.Helper()
	if e := check.Within(want, within, have, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}

// Zone asserts timezones are equal. Returns true if they are, otherwise marks
// the test as failed, writes error message to test log and returns false.
func Zone(t tester.T, want, have *time.Location, opts ...check.Option) bool {
	t.Helper()
	if e := check.Zone(want, have, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}

// Duration asserts durations are equal. Returns true if they are, otherwise
// marks the test as failed, writes error message to test log and returns false.
//
// The "want" and "have" durations might be duration representations in form of
// string, int, int64 or [time.Duration].
func Duration(t tester.T, want, have any, opts ...check.Option) bool {
	t.Helper()
	if e := check.Duration(want, have, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}
