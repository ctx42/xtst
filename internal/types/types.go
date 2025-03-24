// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

// Package types provides example types used in tests.
package types

import (
	"errors"
	"fmt"
	"time"
)

// TODO(rz): make very detailed code review of this file. Check if all the types
//  are required.

// WAW represents Europe/Warsaw timezone.
var WAW *time.Location

func init() {
	var err error
	WAW, err = time.LoadLocation("Europe/Warsaw")
	if err != nil {
		panic(err)
	}
}

// TFuncA is a function used in tests.
func TFuncA() {}

// TFuncB is a function used in tests.
func TFuncB() {}

// /////////////////////////////////////////////////////////////////////////////

// TIntType is type alias used in tests.
type TIntType int

// /////////////////////////////////////////////////////////////////////////////

// TStrType is a type used in tests.
type TStrType string

// /////////////////////////////////////////////////////////////////////////////

// TItf is an interface used in tests.
type TItf interface{ AAA() string }

// /////////////////////////////////////////////////////////////////////////////

// TInt is type used in tests.
type TInt struct{ V int }

// NewTInt returns error when v is not 42.
func NewTInt(v int) (*TInt, error) {
	if v != 42 {
		return nil, errors.New("not cool")
	}
	return &TInt{V: v}, nil
}

// /////////////////////////////////////////////////////////////////////////////

// TA is an example type with fields of different types.
type TA struct {
	Int int
	Str string
	Tim time.Time
	Dur time.Duration
	Loc *time.Location
	TAp *TA

	private int // nolint: unused
}

// /////////////////////////////////////////////////////////////////////////////

// TB is an example type with fields of different types.
type TB struct {
	TA     // Embedded by value.
	TAv TA // Non pointer type.
}

// /////////////////////////////////////////////////////////////////////////////

// TC is an example type with embedded field which is not struct or interface.
type TC struct {
	TD
	Int int
}

// /////////////////////////////////////////////////////////////////////////////

// TD represents type which is another name for string.
type TD string

// /////////////////////////////////////////////////////////////////////////////

// TNested represents type with slices and maps of other types.
type TNested struct {
	SInt    []int
	STA     []TA
	STAp    []*TA
	MStrInt map[string]int
	MStrTyp map[string]TA
	MIntTyp map[int]TA
}

// /////////////////////////////////////////////////////////////////////////////

// TVal implements [TItf], has value receiver methods.
type TVal struct{ Val string } // nolint: errname

// AAA implements TItf.
func (typ TVal) AAA() string { return typ.Val }

// Error implements builtin in Error interface.
func (typ TVal) Error() string { return typ.Val }

// /////////////////////////////////////////////////////////////////////////////

// TPtr implements [TItf], has pointer receiver methods.
type TPtr struct{ Val string } // nolint: errname

// AAA implements TItf.
func (typ *TPtr) AAA() string { return typ.Val }

// Variadic1 is variadic function with one additional argument.
func (typ *TPtr) Variadic1(str string, i ...int) string {
	return fmt.Sprintf("%s %s %v", typ.Val, str, i)
}

// Error implements builtin in error interface.
func (typ *TPtr) Error() string { return typ.Val }

// PS adds two numbers.
func (typ *TPtr) PS(a, b string) string { return a + typ.Val + b }

// /////////////////////////////////////////////////////////////////////////////

// T1 represents nested structure.
type T1 struct {
	Int int
	T1  *T1
}

// /////////////////////////////////////////////////////////////////////////////

const (
	TCBoolA       bool       = true
	TCBoolB       bool       = true
	TCBoolC       bool       = false
	TCStringA     string     = "abc"
	TCStringB     string     = "abc"
	TCStringC     string     = "cba"
	TCIntA        int        = -123
	TCIntB        int        = -123
	TCIntC        int        = -321
	TCInt8A       int8       = -8
	TCInt8B       int8       = -8
	TCInt8C       int8       = -13
	TCInt16A      int16      = -16
	TCInt16B      int16      = -16
	TCInt16C      int16      = -61
	TCInt32A      int32      = -32
	TCInt32B      int32      = -32
	TCInt32C      int32      = -23
	TCInt64A      int64      = -64
	TCInt64B      int64      = -64
	TCInt64C      int64      = -46
	TCUintA       uint       = 123
	TCUintB       uint       = 123
	TCUintC       uint       = 321
	TCUint8A      uint8      = 8
	TCUint8B      uint8      = 8
	TCUint8C      uint8      = 13
	TCUint16A     uint16     = 16
	TCUint16B     uint16     = 16
	TCUint16C     uint16     = 61
	TCUint32A     uint32     = 32
	TCUint32B     uint32     = 32
	TCUint32C     uint32     = 23
	TCUint64A     uint64     = 64
	TCUint64B     uint64     = 64
	TCUint64C     uint64     = 46
	TCUintptrA    uintptr    = 42
	TCUintptrB    uintptr    = 42
	TCUintptrC    uintptr    = 24
	TCFloat32A    float32    = 32.0
	TCFloat32B    float32    = 32.0
	TCFloat32C    float32    = 23.0
	TCFloat64A    float64    = 64.0
	TCFloat64B    float64    = 64.0
	TCFloat64C    float64    = 46.0
	TCComplex64A  complex64  = 6i + 4
	TCComplex64B  complex64  = 6i + 4
	TCComplex64C  complex64  = 4i + 6
	TCComplex128A complex128 = 12i + 8
	TCComplex128B complex128 = 12i + 8
	TCComplex128C complex128 = 8i + 12

	CBoolA       = true
	CBoolB       = true
	CBoolC       = false
	CStringA     = "abc"
	CStringB     = "abc"
	CStringC     = "xyz"
	CIntA        = 123
	CIntB        = 123
	CIntC        = 321
	CFloatA      = 1.23
	CFloatB      = 1.23
	CFloatC      = 3.21
	CComplex64A  = 6i + 4
	CComplex64B  = 6i + 4
	CComplex64C  = 4i + 6
	CComplex128A = 12i + 8
	CComplex128B = 12i + 8
	CComplex128C = 8i + 12
)
