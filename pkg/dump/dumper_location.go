// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package dump

import (
	"reflect"
	"time"
)

// TODO(rz): rename to zoneDumper

// locationDumper requires val to be dereferenced representation of
// [reflect.Struct] which can be cast to [time.Location] and returns its
// string representation in format defined by [Dump] configuration.
func locationDumper(dmp Dump, lvl int, val reflect.Value) string {
	loc := val.Interface().(time.Location) // nolint: forcetypeassert
	val = reflect.ValueOf((&loc).String())
	return simpleDumper(dmp, lvl, val)
}
