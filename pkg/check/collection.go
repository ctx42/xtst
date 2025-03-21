// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"github.com/ctx42/xtst/internal"
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
