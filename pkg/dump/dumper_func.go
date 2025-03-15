// Copyright (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package dump

import (
	"fmt"
	"reflect"
)

// funcDumper is a generic dumper for functions. It expects val to represent
// one of the kinds:
//
//   - [reflect.Func]
//
// Returns [valErrUsage] ("<dump-usage-error>") string if kind cannot be
// matched. It requires val to be dereferenced value and returns its string
// representation in format defined by [Dump] configuration.
func funcDumper(dmp Dump, lvl int, val reflect.Value) string {
	if val.Kind() == reflect.Func {
		ptrAddr := valAddr
		if dmp.cfg.PtrAddr {
			ptr := reflect.ValueOf(val.Pointer())
			ptrAddr = hexPtrDumper(dmp, lvl, ptr)
		}
		return fmt.Sprintf("%s(%s)", valFunc, ptrAddr)
	}
	return valErrUsage
}
