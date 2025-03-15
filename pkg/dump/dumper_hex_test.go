// Copyright (c) 2025 Rafal Zajac
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

		val  any
		want string
	}{
		{"uintptr", uintptr(1234), "<0x4d2>"},
		{"byte", byte(123), "0x7b"},
		{"usage error", 1234, valErrUsage},
		{"unsafe pointer", unsafe.Pointer(sPtr), fmt.Sprintf("<%p>", sPtr)},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have := hexPtrDumper(Dump{}, 0, reflect.ValueOf(tc.val))

			// --- Then ---
			affirm.Equal(t, tc.want, have)
		})
	}
}
