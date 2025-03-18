// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package dump

import (
	"fmt"
	"reflect"
	"strconv"
)

// simpleDumper is a generic dumper for simple types. It expects val to
// represent one of the kinds:
//
//   - [reflect.Bool]
//   - [reflect.Int]
//   - [reflect.Int8]
//   - [reflect.Int16]
//   - [reflect.Int32]
//   - [reflect.Int64]
//   - [reflect.Uint]
//   - [reflect.Uint8]
//   - [reflect.Uint16]
//   - [reflect.Uint32]
//   - [reflect.Uint64]
//   - [reflect.Uintptr]
//   - [reflect.Float32]
//   - [reflect.Float64]
//   - [reflect.String]
//
// It requires val to be dereferenced value and returns its string
// representation in format defined by [Dump] configuration.
func simpleDumper(dmp Dump, _ int, val reflect.Value) string {
	v := val.Interface()

	var format string
	switch val.Kind() {
	case reflect.String:
		format = `"%v"`
		if dmp.cfg.Flat {
			format = `%#v`
		}

	case reflect.Float32:
		// nolint: forcetypeassert
		return strconv.FormatFloat(float64(v.(float32)), 'f', -1, 32)

	case reflect.Float64:
		// nolint: forcetypeassert
		return strconv.FormatFloat(v.(float64), 'f', -1, 64)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		format = `%d`

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64:
		format = `%d`

	default:
		format = `%v`
	}

	return fmt.Sprintf(format, v)
}
