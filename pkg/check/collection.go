// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"github.com/ctx42/xtst/internal"
	"github.com/ctx42/xtst/pkg/dump"
	"github.com/ctx42/xtst/pkg/notice"
)

// Len checks "have" has "want" elements. Returns nil if it has, otherwise it
// returns an error with a message indicating the expected and actual values.
func Len(want int, have any, opts ...Option) error {
	cnt, ok := internal.Len(have)
	if !ok {
		return notice.New("cannot execute len(%T)", have)
	}
	if want != cnt {
		ops := DefaultOptions().set(opts)
		msg := notice.New("expected %T length", have).
			Trail(ops.Trail).
			Want("%d", want).
			Have("%d", cnt)
		return msg
	}
	return nil
}

// Has checks slice has "want" value. Returns nil if it does, otherwise it
// returns an error with a message indicating the expected and actual values.
func Has[T comparable](want T, bag []T, opts ...Option) error {
	for _, got := range bag {
		if want == got {
			return nil
		}
	}
	ops := DefaultOptions().set(opts)
	return notice.New("expected slice to have a value").
		Trail(ops.Trail).
		Want("%#v", want).
		Append("slice", "%#v", bag)
}

// HasNo checks slice does not have "want" value. Returns nil if it doesn't,
// otherwise it returns an error with a message indicating the expected and
// actual values.
func HasNo[T comparable](want T, set []T, opts ...Option) error {
	for i, got := range set {
		if want == got {
			ops := DefaultOptions().set(opts)
			dmp := dump.New(ops.DumpConfig)
			return notice.New("expected slice not to have value").
				Trail(ops.Trail).
				Want("%#v", want).
				Append("index", "%d", i).
				Append("slice", "%s", dmp.DumpAny(set))
		}
	}
	return nil
}
