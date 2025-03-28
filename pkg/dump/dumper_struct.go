// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package dump

import (
	"reflect"
	"strings"
)

// mapDumper requires val to be dereferenced representation of [reflect.Struct]
// and returns its string representation in format defined by [Dump]
// configuration.
func structDumper(dmp Dump, lvl int, val reflect.Value) string {
	prn := newPrinter(dmp.cfg)
	prn.tab(dmp.cfg.Indent + lvl)
	vTyp := val.Type()

	num := val.NumField() // Total number of fields.
	lastPrivate := false
	prn.write("{").nli(num)

	for i := 0; i < num; i++ {
		last := i == num-1

		fld := vTyp.Field(i)
		if !fld.IsExported() {
			lastPrivate = last
			continue
		}

		// Field name.
		prn.tab(dmp.cfg.Indent + lvl + 1)
		prn.write(fld.Name)
		prn.write(":").space()

		// Field value.
		dmp.cfg.PrintType = true
		sub := dmp.value(lvl+1, val.Field(i))
		sub = strings.TrimLeft(sub, "\t")

		prn.write(sub)
		prn.comma(last).sep(last).nl()
	}

	if lastPrivate && dmp.cfg.Flat {
		return strings.TrimRight(prn.String(), ",") + "}"
	}
	prn.tab(dmp.cfg.Indent + lvl).write("}")

	return prn.String()
}
