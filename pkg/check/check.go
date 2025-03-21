// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

// Package check provides equality toolkit used by assert package.
package check

import (
	"strings"

	"github.com/ctx42/xtst/pkg/notice"
)

// Count checks there is "count" occurrences of "what" in "where". Returns nil
// if it's, otherwise it returns an error with a message indicating the
// expected and actual values.
func Count(count int, what, where any, opts ...Option) error {
	ops := DefaultOptions().set(opts)

	if src, ok := where.(string); ok {
		var ok bool
		var subT string
		if subT, ok = what.(string); !ok {
			const mHeader = "expected argument \"what\" to be string got %T"
			return notice.New(mHeader, what).
				Trail(ops.Trail)
		}
		haveCnt := strings.Count(src, subT)
		if count == haveCnt {
			return nil
		}
		return notice.New("expected string to contain substrings").
			Trail(ops.Trail).
			Append("want count", "%d", count).
			Append("have count", "%d", haveCnt).
			Append("what", "%q", what).
			Append("where", "%q", where)
	}
	return notice.New("unsupported \"where\" type: %T", where).
		Trail(ops.Trail)
}
