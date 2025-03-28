// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package dump

import (
	"fmt"
	"reflect"
)

// chanDumper is a generic dumper for channels. It expects val to represent one
// of the kinds:
//
//   - [reflect.Chan]
//
// Returns [valErrUsage] ("<dump-usage-error>") string if kind cannot be
// matched. It requires val to be dereferenced value and returns its string
// representation in format defined by [Dump] configuration.
func chanDumper(dmp Dump, lvl int, val reflect.Value) string {
	var str string
	switch val.Kind() {
	case reflect.Chan:
		ptrAddr := valAddr
		if dmp.cfg.PtrAddr {
			ptr := reflect.ValueOf(val.Pointer())
			ptrAddr = hexPtrDumper(dmp, lvl, ptr)
		}
		str = fmt.Sprintf("(%s)(%s)", val.Type(), ptrAddr)
	default:
		str = valErrUsage
	}

	prn := newPrinter(dmp.cfg)
	return prn.tab(dmp.cfg.Indent + lvl).write(str).String()
}
