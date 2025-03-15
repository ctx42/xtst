// Copyright (c) 2025 Rafal Zajac
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

	num := len(keys)
	prn.write("{").nli(num)
	dmp.cfg.PrintType = false // Don't print types for map values.
	for i, key := range keys {
		last := i == num-1

		prn.tab(lvl)
		prn.write(dmp.Dump(lvl, key))
		prn.write(":").space()

		sub := dmp.Dump(lvl, val.MapIndex(key))
		dmp.cfg.PrintType = true
		prn.write(sub)
		prn.comma(last).sep(last).nl()
	}
	prn.tab(lvl - 1).write("}")

	return prn.String()
}
