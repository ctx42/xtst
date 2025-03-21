// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"testing"

	"github.com/ctx42/xtst/internal/affirm"
)

func Test_Len(t *testing.T) {
	t.Run("error with option", func(t *testing.T) {
		// --- Given ---
		opt := WithTrail("type.field")

		// --- When ---
		err := Len(1, []int{1, 2}, opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected []int length:\n" +
			"\ttrail: type.field\n" +
			"\t want: 1\n" +
			"\t have: 2"
		affirm.Equal(t, wMsg, err.Error())
	})
}

func Test_Len_success_tabular(t *testing.T) {
	ch := make(chan int, 4)
	ch <- 0
	ch <- 1
	ch <- 2
	t.Cleanup(func() { <-ch; <-ch; <-ch; close(ch) })

	tt := []struct {
		testN string

		val  any
		want int
	}{
		{"int empty success", []int{}, 0},
		{"int success", []int{1}, 1},
		{"map empty success", map[string]int{}, 0},
		{"map success", map[string]int{"A": 1}, 1},
		{"channel success", ch, 3},
		{"string success", "abc", 3},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			err := Len(tc.want, tc.val)

			// --- Then ---
			affirm.Nil(t, err)
		})
	}
}

func Test_Len_error_tabular(t *testing.T) {
	ch := make(chan int, 4)
	ch <- 0
	ch <- 1
	ch <- 2
	t.Cleanup(func() { <-ch; <-ch; <-ch; close(ch) })

	tt := []struct {
		testN string

		val  any
		want int
		wMsg string
	}{
		{"int empty fail", []int{}, 1, "expected []int length:\n\twant: 1\n\thave: 0"},
		{"int fail", []int{1}, 2, "expected []int length:\n\twant: 2\n\thave: 1"},
		{"map empty fail", map[string]int{}, 1, "expected map[string]int length:\n\twant: 1\n\thave: 0"},
		{"map fail", map[string]int{"A": 1}, 2, "expected map[string]int length:\n\twant: 2\n\thave: 1"},
		{"invalid type", 1, 0, "cannot execute len(int)"},
		{"chan fail", ch, 4, "expected chan int length:\n\twant: 4\n\thave: 3"},
		{"string fail", "abc", 4, "expected string length:\n\twant: 4\n\thave: 3"},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			err := Len(tc.want, tc.val)

			// --- Then ---
			affirm.NotNil(t, err)
			affirm.Equal(t, tc.wMsg, err.Error())
		})
	}
}
