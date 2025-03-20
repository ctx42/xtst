// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"testing"

	"github.com/ctx42/xtst/internal/affirm"
)

func Test_FileExist(t *testing.T) {
	t.Run("exists", func(t *testing.T) {
		// --- When ---
		err := FileExist("testdata/file.txt")

		// --- Then ---
		affirm.Nil(t, err)
	})

	t.Run("error does not exist", func(t *testing.T) {
		// --- Given ---
		opt := WithTrail("type.field")

		// --- When ---
		err := FileExist("testdata/not_existing.txt", opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected path to an existing file:\n" +
			"\ttrail: type.field\n" +
			"\t path: testdata/not_existing.txt"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("error is a directory", func(t *testing.T) {
		// --- Given ---
		opt := WithTrail("type.field")

		// --- When ---
		err := FileExist("testdata/dir", opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected path to be existing file:\n" +
			"\ttrail: type.field\n" +
			"\t path: testdata/dir"
		affirm.Equal(t, wMsg, err.Error())
	})
}

func Test_NoFileExist(t *testing.T) {
	t.Run("exists", func(t *testing.T) {
		// --- When ---
		err := NoFileExist("testdata/file.txt")

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected path to not existing file:\n" +
			"\tpath: testdata/file.txt"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("does not exist", func(t *testing.T) {
		// --- When ---
		err := NoFileExist("testdata/not_existing.txt")

		// --- Then ---
		affirm.Nil(t, err)
	})

	t.Run("is a directory", func(t *testing.T) {
		// --- Given ---
		opt := WithTrail("type.field")

		// --- When ---
		err := NoFileExist("testdata/dir", opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected path to be not existing file:\n" +
			"\ttrail: type.field\n" +
			"\t path: testdata/dir"
		affirm.Equal(t, wMsg, err.Error())
	})
}

func Test_DirExist(t *testing.T) {
	t.Run("exists", func(t *testing.T) {
		// --- When ---
		err := DirExist("testdata/dir")

		// --- Then ---
		affirm.Nil(t, err)
	})

	t.Run("does not exist", func(t *testing.T) {
		// --- Given ---
		opt := WithTrail("type.field")

		// --- When ---
		err := DirExist("testdata/not_existing_dir", opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected path to an existing directory:\n" +
			"\ttrail: type.field\n" +
			"\t path: testdata/not_existing_dir"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("is a file", func(t *testing.T) {
		// --- Given ---
		opt := WithTrail("type.field")

		// --- When ---
		err := DirExist("testdata/file.txt", opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected path to be existing directory:\n" +
			"\ttrail: type.field\n" +
			"\t path: testdata/file.txt"
		affirm.Equal(t, wMsg, err.Error())
	})
}

func Test_NoDirExist(t *testing.T) {
	t.Run("exists", func(t *testing.T) {
		// --- Given ---
		opt := WithTrail("type.field")

		// --- When ---
		err := NoDirExist("testdata/dir", opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected path to not existing directory:\n" +
			"\ttrail: type.field\n" +
			"\t path: testdata/dir"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("does not exist", func(t *testing.T) {
		// --- When ---
		err := NoDirExist("testdata/not_existing_dir")

		// --- Then ---
		affirm.Nil(t, err)
	})

	t.Run("is a file", func(t *testing.T) {
		// --- Given ---
		opt := WithTrail("type.field")

		// --- When ---
		err := NoDirExist("testdata/file.txt", opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected path to be not existing directory:\n" +
			"\ttrail: type.field\n" +
			"\t path: testdata/file.txt"
		affirm.Equal(t, wMsg, err.Error())
	})
}
