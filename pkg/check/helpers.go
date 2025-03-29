// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/ctx42/xtst/pkg/notice"
)

// Types for some of the built-in types.
var (
	typTime       = reflect.TypeOf(time.Time{})
	typTimeLoc    = reflect.TypeOf(time.Location{})  // TODO(rz): do we need both?
	typTimeLocPtr = reflect.TypeOf(&time.Location{}) // TODO(rz): do we need both?
)

// typeString returns type of the value as a string.
func typeString(val reflect.Value) string {
	if !val.IsValid() {
		return "<invalid>"
	}
	return val.Type().String()
}

// isPrintableChar returns true if "v" is a printable ASCII character.
func isPrintableChar(v byte) bool {
	return v >= 32 && v <= 126
}

// valToString returns string representation of the value.
//
// nolint: cyclop
func valToString(key reflect.Value) string {
	switch key.Kind() {

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", key.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return fmt.Sprintf("%d", key.Uint())

	case reflect.Uint64, reflect.Uintptr:
		return fmt.Sprintf("%d", key.Uint())

	case reflect.Float32:
		return strconv.FormatFloat(key.Float(), 'f', -1, 32)

	case reflect.Float64:
		return strconv.FormatFloat(key.Float(), 'f', -1, 64)

	case reflect.String:
		return fmt.Sprintf("%q", key.String())

	case reflect.Bool:
		return fmt.Sprintf("%v", key.Bool())

	case reflect.Struct:
		// For structs, we'll just print the type name
		// since the actual value might be complex.
		pkg := ""
		typ := key.Type()
		if pkgPath := typ.PkgPath(); pkgPath != "" {
			parts := strings.Split(pkgPath, "/")
			pkg = parts[len(parts)-1] + "."
		}
		return fmt.Sprintf("%s%s", pkg, typ.Name())

	case reflect.Ptr:
		if key.IsNil() {
			return "<nil>"
		} else {
			return "*" + valToString(key.Elem())
		}

	case reflect.Complex64:
		return strconv.FormatComplex(key.Complex(), 'f', -1, 64)

	case reflect.Complex128:
		return strconv.FormatComplex(key.Complex(), 'f', -1, 64)

	case reflect.Array:
		return "<array>"

	case reflect.UnsafePointer:
		return fmt.Sprintf("<%p>", key.UnsafePointer())

	default:
		return "<invalid>"
	}
}

// wrap wraps error in multiError if it's an error joined with [errors.Join].
func wrap(err error) error {
	if errs, ok := err.(interface{ Unwrap() []error }); ok {
		return multiError{ers: errs.Unwrap()}
	}
	return err
}

// multiError is a decorator for multi [notice.Notice] error messages.
type multiError struct{ ers []error }

func (e multiError) Error() string {
	var prev string
	var msg *notice.Notice
	for i, err := range e.ers {
		if errors.As(err, &msg) {
			if prev == msg.Header {
				prev = msg.Header
				msg.Header = notice.ContinuationHeader
			} else {
				prev = msg.Header
				if i > 0 {
					msg.Header = "\n" + msg.Header
				}
			}
			continue
		}
		prev = ""
	}
	return errors.Join(e.ers...).Error()
}

func (e multiError) Unwrap() []error { return e.ers }
