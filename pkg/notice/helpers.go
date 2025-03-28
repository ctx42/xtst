// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package notice

import (
	"strings"
)

// Lines takes an indent value as number of spaces and a string of lines, and
// returns the lines formatted with the specified indent.
func Lines(indent int, lns string) string {
	if lns == "" {
		return ""
	}
	out := strings.TrimSpace(lns)
	rows := strings.Split(out, "\n")
	for i, lin := range rows {
		ind := strings.Repeat(" ", indent)
		rows[i] = ind + ">| " + lin
	}
	return strings.Join(rows, "\n")
}

// Unwrap unwraps joined errors. Returns nil if err is nil, unwraps only
// non-nil errors.
func Unwrap(err error) []error {
	if err == nil {
		return nil
	}
	var ers []error
	if es, ok := err.(interface{ Unwrap() []error }); ok {
		for _, e := range es.Unwrap() {
			ers = append(ers, e)
		}
	} else {
		ers = append(ers, err)
	}
	return ers
}
