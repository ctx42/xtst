// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
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
	var str string
	switch val.Kind() {
	case reflect.Func:
		ptrAddr := valAddr
		if dmp.PtrAddr {
			ptr := reflect.ValueOf(val.Pointer())
			ptrAddr = hexPtrDumper(dmp, lvl, ptr)
		}
		str = fmt.Sprintf("%s(%s)", valFunc, ptrAddr)
	default:
		str = valErrUsage
	}

	prn := NewPrinter(dmp)
	return prn.Tab(dmp.Indent + lvl).Write(str).String()
}
