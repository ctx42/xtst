// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"os"

	"github.com/ctx42/xtst/pkg/notice"
)

// FileExist checks "pth" points to an existing file. Returns an error if the
// path points to a filesystem entry which is not a file or there is an error
// when trying to check the path. On success, it returns nil.
func FileExist(pth string, opts ...Option) error {
	inf, err := os.Lstat(pth)
	ops := DefaultOptions().set(opts)
	if err != nil {
		if os.IsNotExist(err) {
			return notice.New("expected path to an existing file").
				Trail(ops.Trail).
				Append("path", "%s", pth)
		}
		return notice.New("expected os.Lstat to succeed").
			Trail(ops.Trail).
			Append("path", "%s", pth).
			Append("error", "%s", err)
	}
	if inf.IsDir() {
		return notice.New("expected path to be existing file").
			Trail(ops.Trail).
			Append("path", "%s", pth)
	}
	return nil
}

// NoFileExist checks "pth" points to not existing file. Returns an error if
// the path points to an existing filesystem entry. On success, it returns nil.
func NoFileExist(pth string, opts ...Option) error {
	inf, err := os.Lstat(pth)
	ops := DefaultOptions().set(opts)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return notice.New("expected os.Lstat to succeed").
			Trail(ops.Trail).
			Append("path", "%s", pth).
			Append("error", "%s", err)
	}
	if inf.IsDir() {
		return notice.New("expected path to be not existing file").
			Trail(ops.Trail).
			Append("path", "%s", pth)
	}
	return notice.New("expected path to not existing file").
		Trail(ops.Trail).
		Append("path", "%s", pth)
}

// Content declares type constraint for file content.
type Content interface {
	string | []byte
}

// DirExist checks "pth" points to an existing directory. It fails if the path
// points to a filesystem entry which is not a directory or there is an error
// when trying to check the path. When it fails it returns an error with a
// detailed message indicating the expected and actual values.
func DirExist(pth string, opts ...Option) error {
	inf, err := os.Lstat(pth)
	ops := DefaultOptions().set(opts)
	if err != nil {
		if os.IsNotExist(err) {
			return notice.New("expected path to an existing directory").
				Trail(ops.Trail).
				Append("path", "%s", pth)
		}
		return notice.New("expected os.Lstat to succeed").
			Trail(ops.Trail).
			Append("path", "%s", pth).
			Append("error", "%s", err)
	}
	if !inf.IsDir() {
		return notice.New("expected path to be existing directory").
			Trail(ops.Trail).
			Append("path", "%s", pth)
	}
	return nil
}

// NoDirExist checks "pth" points to not existing directory. It fails if the
// path points to an existing filesystem entry. When it fails it returns an
// error with a detailed message indicating the expected and actual values.
func NoDirExist(pth string, opts ...Option) error {
	inf, err := os.Lstat(pth)
	ops := DefaultOptions().set(opts)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return notice.New("expected os.Lstat to succeed").
			Trail(ops.Trail).
			Append("path", "%s", pth).
			Append("error", "%s", err)
	}
	if !inf.IsDir() {
		return notice.New("expected path to be not existing directory").
			Trail(ops.Trail).
			Append("path", "%s", pth)
	}
	return notice.New("expected path to not existing directory").
		Trail(ops.Trail).
		Append("path", "%s", pth)
}
