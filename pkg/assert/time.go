// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package assert

import (
	"time"

	"github.com/ctx42/xtst/pkg/check"
	"github.com/ctx42/xtst/pkg/tester"
)

// Time asserts "want" and "have" dates are equal. Returns true if they are,
// otherwise marks the test as failed, writes error message to test log and
// returns false.
//
// The "want" and "have" might be date representations in form of string, int,
// int64 or [time.Time]. For string representations the [Options.TimeFormat] is
// used during parsing and the returned date is always in UTC. The int and
// int64 types are interpreted as Unix Timestamp and the date returned is also
// in UTC.
func Time(t tester.T, want, have any, opts ...check.Option) bool {
	t.Helper()
	if e := check.Time(want, have, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}

// TimeExact asserts "want" and "have" dates are equal and are in the same
// timezone. Returns true if they are, otherwise marks the test as failed,
// writes error message to test log and returns false.
//
// The "want" and "have" might be date representations in form of string, int,
// int64 or [time.Time]. For string representations the [Options.TimeFormat] is
// used during parsing and the returned date is always in UTC. The int and
// int64 types are interpreted as Unix Timestamp and the date returned is also
// in UTC.
func TimeExact(t tester.T, want, have any, opts ...check.Option) bool {
	t.Helper()
	if e := check.TimeExact(want, have, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}

// Before asserts "date" is before "mark". Returns true if it's, otherwise
// marks the test as failed, writes error message to test log and returns false.
//
// The "date" and "mark" might be date representations in form of string, int,
// int64 or [time.Time]. For string representations the
// [check.Options.TimeFormat] is used during parsing and the returned date is
// always in UTC. The int and int64 types are interpreted as Unix Timestamp and
// the date returned is also in UTC.
func Before(t tester.T, date, mark any, opts ...check.Option) bool {
	t.Helper()
	if e := check.Before(date, mark, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}

// After asserts "date" is after "mark". Returns true if it's, otherwise
// marks the test as failed, writes error message to test log and returns false.
//
// The "date" and "mark" might be date representations in form of string, int,
// int64 or [time.Time]. For string representations the
// [check.Options.TimeFormat] is used during parsing and the returned date is
// always in UTC. The int and int64 types are interpreted as Unix Timestamp and
// the date returned is also in UTC.
func After(t tester.T, date, mark time.Time, opts ...check.Option) bool {
	t.Helper()
	if e := check.After(date, mark, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}

// EqualOrBefore asserts "date" is equal or before "mark". Returns true if it's,
// otherwise marks the test as failed, writes error message to test log and
// returns false.
//
// The "date" and "mark" might be date representations in form of string, int,
// int64 or [time.Time]. For string representations the
// [check.Options.TimeFormat] is used during parsing and the returned date is
// always in UTC. The int and int64 types are interpreted as Unix Timestamp and
// the date returned is also in UTC.
func EqualOrBefore(t tester.T, date, mark time.Time, opts ...check.Option) bool {
	t.Helper()
	if e := check.EqualOrBefore(date, mark, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}

// EqualOrAfter asserts "date" is equal or after "mark". Returns true if it's,
// otherwise marks the test as failed, writes error message to test log and
// returns false.
//
// The "date" and "mark" might be date representations in form of string, int,
// int64 or [time.Time]. For string representations the
// [check.Options.TimeFormat] is used during parsing and the returned date is
// always in UTC. The int and int64 types are interpreted as Unix Timestamp and
// the date returned is also in UTC.
func EqualOrAfter(t tester.T, date, mark any, opts ...check.Option) bool {
	t.Helper()
	if e := check.EqualOrAfter(date, mark, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}

// Within asserts "want" and "have" dates are equal "within" given duration.
// Returns true if they are, otherwise marks the test as failed, writes error
// message to test log and returns false.
//
// The "want" and "have" might be date representations in form of string, int,
// int64 or [time.Time]. For string representations the [Options.TimeFormat] is
// used during parsing and the returned date is always in UTC. The int and
// int64 types are interpreted as Unix Timestamp and the date returned is also
// in UTC.
//
// The "within" might be duration representation in form of string, int, int64
// or [time.Duration].
func Within(t tester.T, want, within, have any, opts ...check.Option) bool {
	t.Helper()
	if e := check.Within(want, within, have, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}

// Recent asserts "have" is within [check.Options.Recent] from [time.Now].
// Returns nil if it is, otherwise marks the test as failed, writes error
// message to test log and returns false.
//
// The "have" may represent date in form of a string, int, int64 or [time.Time].
// For string representations the [check.Options.TimeFormat] is used during
// parsing and the returned date is always in UTC. The int and int64 types are
// interpreted as Unix Timestamp and the date returned is also in UTC.
func Recent(t tester.T, have any, opts ...check.Option) bool {
	t.Helper()
	if e := check.Recent(have, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}

// Zone asserts "want" and "have" timezones are equal. Returns true if they are,
// otherwise marks the test as failed, writes error message to test log and
// returns false.
func Zone(t tester.T, want, have *time.Location, opts ...check.Option) bool {
	t.Helper()
	if e := check.Zone(want, have, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}

// Duration asserts "want" and "have" durations are equal. Returns true if they
// are, otherwise marks the test as failed, writes error message to test log and
// returns false.
//
// The "want" and "have" might be duration representation in form of string,
// int, int64 or [time.Duration].
func Duration(t tester.T, want, have any, opts ...check.Option) bool {
	t.Helper()
	if e := check.Duration(want, have, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}
