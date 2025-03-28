// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package dump

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
)

// mapDumper requires val to be dereferenced representation of [reflect.Map]
// and returns its string representation in format defined by [Dump]
// configuration.
//
// nolint: cyclop
func mapDumper(dmp Dump, lvl int, val reflect.Value) string {
	prn := newPrinter(dmp.cfg)
	prn.tab(dmp.cfg.Indent + lvl)

	if dmp.cfg.PrintType {
		keyTyp := val.Type().Key()
		valTyp := val.Type().Elem()
		valTypStr := strings.ReplaceAll(valTyp.String(), " ", "")
		if valTypStr == "interface{}" && dmp.cfg.UseAny {
			valTypStr = "any"
		}
		str := fmt.Sprintf("map[%s]%s", keyTyp.String(), valTypStr)
		prn.write(str)
	}

	keys := val.MapKeys()
	slices.SortStableFunc(keys, valueCmp)

	if val.IsNil() {
		return prn.write("(nil)").String()
	}

	num := val.Len()
	prn.write("{").nli(num)

	dmp.cfg.PrintType = false // Don't print types for map values.
	for i, key := range keys {
		last := i == num-1

		sub := dmp.value(lvl+1, key)
		prn.write(sub)
		prn.write(":").space()

		sub = dmp.value(lvl+1, val.MapIndex(key))
		sub = strings.TrimLeft(sub, "\t")

		dmp.cfg.PrintType = true
		prn.write(sub)
		prn.comma(last).sep(last).nl()
	}
	prn.tab(dmp.cfg.Indent + lvl).write("}")

	return prn.String()
}
