package aferomock

import (
	"io/fs"
	"time"

	"github.com/spf13/afero"
)

var _ afero.Fs = &FsCallbacks{}

// WrappedFs is a type alias for FsCallbacks.
// Deprecated: Use FsCallbacks instead.
type WrappedFs = FsCallbacks

// FsCallbacks is a callback-based mock for afero.Fs.
type FsCallbacks struct {
	ChmodFunc     func(name string, mode fs.FileMode) error
	ChownFunc     func(name string, uid int, gid int) error
	ChtimesFunc   func(name string, atime time.Time, mtime time.Time) error
	CreateFunc    func(name string) (afero.File, error)
	MkdirFunc     func(name string, perm fs.FileMode) error
	MkdirAllFunc  func(path string, perm fs.FileMode) error
	NameFunc      func() string
	OpenFunc      func(name string) (afero.File, error)
	OpenFileFunc  func(name string, flag int, perm fs.FileMode) (afero.File, error)
	RemoveFunc    func(name string) error
	RemoveAllFunc func(path string) error
	RenameFunc    func(oldname string, newname string) error
	StatFunc      func(name string) (fs.FileInfo, error)
}

// Chmod satisfies the afero.Fs interface.
func (fs FsCallbacks) Chmod(name string, mode fs.FileMode) error {
	return fs.ChmodFunc(name, mode)
}

// Chown satisfies the afero.Fs interface.
func (fs FsCallbacks) Chown(name string, uid int, gid int) error {
	return fs.ChownFunc(name, uid, gid)
}

// Chtimes satisfies the afero.Fs interface.
func (fs FsCallbacks) Chtimes(name string, atime time.Time, mtime time.Time) error {
	return fs.ChtimesFunc(name, atime, mtime)
}

// Create satisfies the afero.Fs interface.
func (fs FsCallbacks) Create(name string) (afero.File, error) {
	return fs.CreateFunc(name)
}

// Mkdir satisfies the afero.Fs interface.
func (fs FsCallbacks) Mkdir(name string, perm fs.FileMode) error {
	return fs.MkdirFunc(name, perm)
}

// MkdirAll satisfies the afero.Fs interface.
func (fs FsCallbacks) MkdirAll(path string, perm fs.FileMode) error {
	return fs.MkdirAllFunc(path, perm)
}

// Name satisfies the afero.Fs interface.
func (fs FsCallbacks) Name() string {
	return fs.NameFunc()
}

// Open satisfies the afero.Fs interface.
func (fs FsCallbacks) Open(name string) (afero.File, error) {
	return fs.OpenFunc(name)
}

// OpenFile satisfies the afero.Fs interface.
func (fs FsCallbacks) OpenFile(name string, flag int, perm fs.FileMode) (afero.File, error) {
	return fs.OpenFileFunc(name, flag, perm)
}

// Remove satisfies the afero.Fs interface.
func (fs FsCallbacks) Remove(name string) error {
	return fs.RemoveFunc(name)
}

// RemoveAll satisfies the afero.Fs interface.
func (fs FsCallbacks) RemoveAll(path string) error {
	return fs.RemoveAllFunc(path)
}

// Rename satisfies the afero.Fs interface.
func (fs FsCallbacks) Rename(oldname string, newname string) error {
	return fs.RenameFunc(oldname, newname)
}

// Stat satisfies the afero.Fs interface.
func (fs FsCallbacks) Stat(name string) (fs.FileInfo, error) {
	return fs.StatFunc(name)
}

// WrapFs wraps a afero.Fs with custom callbacks.
// Deprecated: Use OverrideFs instead.
func WrapFs(fs afero.Fs, callbacks WrappedFs) FsCallbacks {
	return OverrideFs(fs, callbacks)
}

// OverrideFs overrides a afero.Fs with custom callbacks.
func OverrideFs(fs afero.Fs, c FsCallbacks) FsCallbacks { //nolint: cyclop,dupl
	if c.ChmodFunc == nil {
		c.ChmodFunc = fs.Chmod
	}

	if c.ChownFunc == nil {
		c.ChownFunc = fs.Chown
	}

	if c.ChtimesFunc == nil {
		c.ChtimesFunc = fs.Chtimes
	}

	if c.CreateFunc == nil {
		c.CreateFunc = fs.Create
	}

	if c.MkdirFunc == nil {
		c.MkdirFunc = fs.Mkdir
	}

	if c.MkdirAllFunc == nil {
		c.MkdirAllFunc = fs.MkdirAll
	}

	if c.NameFunc == nil {
		c.NameFunc = fs.Name
	}

	if c.OpenFunc == nil {
		c.OpenFunc = fs.Open
	}

	if c.OpenFileFunc == nil {
		c.OpenFileFunc = fs.OpenFile
	}

	if c.RemoveFunc == nil {
		c.RemoveFunc = fs.Remove
	}

	if c.RemoveAllFunc == nil {
		c.RemoveAllFunc = fs.RemoveAll
	}

	if c.RenameFunc == nil {
		c.RenameFunc = fs.Rename
	}

	if c.StatFunc == nil {
		c.StatFunc = fs.Stat
	}

	return c
}
