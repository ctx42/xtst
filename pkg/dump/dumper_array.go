// Copyright (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package dump

import (
	"reflect"
	"strings"
)

// arrayDumper requires val to be dereferenced representation of [reflect.Array]
// or [reflect.Slice] and returns its string representation in format defined
// by [Dump] configuration.
func arrayDumper(dmp Dump, lvl int, val reflect.Value) string {
	prn := newPrinter(dmp.cfg)

	if dmp.cfg.PrintType {
		valTypStr := val.Type().String()
		if dmp.cfg.UseAny {
			switch {
			case valTypStr == "interface{}":
				valTypStr = "any"
			case strings.HasSuffix(valTypStr, "]interface {}"):
				valTypStr = strings.Replace(valTypStr, "interface {}", "any", 1)
			}
		}
		prn.write(valTypStr)
	}

	num := val.Len()
	prn.write("{").nli(num)
	dmp.cfg.PrintType = false // Don't print types for array elements.
	for i := 0; i < num; i++ {
		last := i == num-1

		sub := dmp.Dump(lvl, val.Index(i))
		if strings.HasPrefix(sub, "{") {
			prn.tab(lvl)
		}
		prn.write(sub)
		prn.comma(last).sep(last).nl()
	}
	prn.tab(lvl - 1).write("}")

	return prn.String()
}
