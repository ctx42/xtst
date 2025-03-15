// Copyright (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package notice

// Row represents [Notice] row.
type Row struct {
	Name   string
	Format string
	Args   []any
}

// NewRow is constructor function for [Row].
func NewRow(name, format string, args ...any) Row {
	return Row{Name: name, Format: format, Args: args}
}

// Path is convenience constructor function for [Row] with name `path`.
func Path(name string) Row {
	return Row{Name: fieldPath, Format: name}
}
