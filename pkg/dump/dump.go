// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

// Package dump can render string representation of any type.
package dump

import (
	"reflect"
	"time"
)

// Strings used by dump package to indicate special values.
const (
	valNil        = "nil"       // The [reflect.Value] is nil.
	valAddr       = "<addr>"    // The [reflect.Value] is an address.
	valFunc       = "<func>"    // The [reflect.Value] is a function.
	valChan       = "<chan>"    // The [reflect.Value] is a channel.
	valInvalid    = "<invalid>" // The [reflect.Value] is invalid.
	valMaxNesting = "<...>"     // The maximum nesting reached.

	// The [reflect.Value] is unexpected in given context.
	valErrUsage = "<dump-usage-error>"
)

// Types for built-in dumpers.
var (
	typDur      = reflect.TypeOf(time.Duration(0))
	typLocation = reflect.TypeOf(time.Location{})
	typTime     = reflect.TypeOf(time.Time{})
)

var valueOfNil = reflect.ValueOf(nil)

// Dumper represents function signature for value dumpers.
type Dumper func(dmp Dump, level int, val reflect.Value) string

// Dump implements logic for dumping values and types.
type Dump struct {
	cfg Config // Configuration.
}

// New returns new instance of [Dump].
func New(cfg Config) Dump { return Dump{cfg: cfg} }

// DefaultDump is convenience method returning [Dump] instance with default
// configuration.
func DefaultDump() Dump { return New(NewConfig()) }

// DumpAny is a convenience method dumping given argument as a string.
func (dmp Dump) DumpAny(val any) string {
	return dmp.Dump(0, reflect.ValueOf(val))
}

// Dump dumps given value as a string.
//
// nolint: cyclop
func (dmp Dump) Dump(lvl int, val reflect.Value) string {
	if lvl > dmp.cfg.Depth {
		return valMaxNesting
	}

	var str string // One or more lines representing passed value.

	valKnd := val.Kind()
	if valKnd != reflect.Invalid {
		if fn, ok := dmp.cfg.Dumpers[val.Type()]; ok {
			return fn(dmp, lvl, val)
		}
	}

	switch valKnd {
	case reflect.Invalid:
		str = valInvalid
		if valueOfNil == val { // nolint: govet
			str = valNil
		}

	case reflect.Bool, reflect.Int:
		str = simpleDumper(dmp, lvl, val)

	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		str = simpleDumper(dmp, lvl, val)

	case reflect.Uint:
		str = simpleDumper(dmp, lvl, val)

	case reflect.Uint16, reflect.Uint32, reflect.Uint64:
		str = simpleDumper(dmp, lvl, val)

	case reflect.Uint8:
		str = hexPtrDumper(dmp, lvl, val)

	case reflect.Uintptr:
		str = hexPtrDumper(dmp, lvl, val)

	case reflect.Float32, reflect.Float64:
		str = simpleDumper(dmp, lvl, val)

	case reflect.Complex64, reflect.Complex128:
		str = complexDumper(dmp, lvl, val)

	case reflect.Array:
		str = arrayDumper(dmp, lvl, val)

	case reflect.Chan:
		str = chanDumper(dmp, lvl, val)

	case reflect.Func:
		str = funcDumper(dmp, lvl, val)

	case reflect.Interface:
		str = dmp.Dump(lvl, val.Elem())

	case reflect.Map:
		str = mapDumper(dmp, lvl+1, val)

	case reflect.Pointer:
		if val.IsNil() {
			str = valNil
		} else {
			str = dmp.Dump(lvl, val.Elem())
		}

	case reflect.Slice:
		str = sliceDumper(dmp, lvl+1, val)

	case reflect.String:
		str = simpleDumper(dmp, lvl, val)

	case reflect.Struct:
		str = structDumper(dmp, lvl+1, val)

	case reflect.UnsafePointer:
		str = hexPtrDumper(dmp, lvl, val)
	}

	return str
}
