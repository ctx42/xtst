// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package dump

import (
	"fmt"
	"reflect"
)

// complexDumper is a generic dumper for complex values. It expects val to
// represent one of the kinds:
//
//   - [reflect.Complex64]
//   - [reflect.Complex128]
//
// Returns [valErrUsage] ("<dump-usage-error>") string if kind cannot be
// matched. It requires val to be dereferenced value and returns its string
// representation in format defined by [Dump] configuration.
func complexDumper(_ Dump, _ int, val reflect.Value) string {
	switch val.Kind() {
	case reflect.Complex64, reflect.Complex128:
		return fmt.Sprintf("%v", val.Interface())
	default:
		return valErrUsage
	}
}
