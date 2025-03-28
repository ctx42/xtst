// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package dump

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/ctx42/xtst/internal/affirm"
	"github.com/ctx42/xtst/internal/types"
)

func Test_hexPtrDumper_tabular(t *testing.T) {
	sPtr := &types.TPtr{Val: "a"}

	tt := []struct {
		testN string

		val    any
		indent int
		level  int
		want   string
	}{
		{"uintptr", uintptr(1234), 0, 0, "<0x4d2>"},
		{"byte", byte(123), 0, 0, "0x7b"},
		{"usage error", 1234, 0, 0, valErrUsage},
		{"unsafe pointer", unsafe.Pointer(sPtr), 0, 0, fmt.Sprintf("<%p>", sPtr)},

		{"uses indent", 1234, 2, 0, "    " + valErrUsage},
		{"uses level", 1234, 0, 1, "  " + valErrUsage},
		{"uses indent and level", 1234, 2, 1, "      " + valErrUsage},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			dmp := New(NewConfig(WithIndent(tc.indent)))

			// --- When ---
			have := hexPtrDumper(dmp, tc.level, reflect.ValueOf(tc.val))

			// --- Then ---
			affirm.Equal(t, tc.want, have)
		})
	}
}
