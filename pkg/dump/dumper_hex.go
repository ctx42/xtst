// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package dump

import (
	"fmt"
	"reflect"
)

// hexPtrDumper is a generic hex dumper for pointers. It expects val to
// represent one of the kinds:
//
//   - [reflect.Uint8]
//   - [reflect.Uintptr]
//   - [reflect.UnsafePointer]
//
// Returns [valErrUsage] ("<dump-usage-error>") string if kind cannot be
// matched. It requires val to be dereferenced value and returns its string
// representation in format defined by [Dump] configuration.
func hexPtrDumper(_ Dump, _ int, val reflect.Value) string {
	switch val.Kind() {
	case reflect.Uint8:
		return fmt.Sprintf("0x%x", val.Interface())
	case reflect.Uintptr:
		return fmt.Sprintf("<0x%x>", val.Uint())
	case reflect.UnsafePointer:
		return fmt.Sprintf("<0x%x>", val.Pointer())
	default:
		return valErrUsage
	}
}
