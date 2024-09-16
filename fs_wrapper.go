package aferomock

import (
	"io/fs"
	"time"

	"github.com/spf13/afero"
)

var _ afero.Fs = &WrappedFs{}

// WrappedFs is a callback-based mock for afero.Fs.
type WrappedFs struct {
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
func (fs *WrappedFs) Chmod(name string, mode fs.FileMode) error {
	return fs.ChmodFunc(name, mode)
}

// Chown satisfies the afero.Fs interface.
func (fs *WrappedFs) Chown(name string, uid int, gid int) error {
	return fs.ChownFunc(name, uid, gid)
}

// Chtimes satisfies the afero.Fs interface.
func (fs *WrappedFs) Chtimes(name string, atime time.Time, mtime time.Time) error {
	return fs.ChtimesFunc(name, atime, mtime)
}

// Create satisfies the afero.Fs interface.
func (fs *WrappedFs) Create(name string) (afero.File, error) {
	return fs.CreateFunc(name)
}

// Mkdir satisfies the afero.Fs interface.
func (fs *WrappedFs) Mkdir(name string, perm fs.FileMode) error {
	return fs.MkdirFunc(name, perm)
}

// MkdirAll satisfies the afero.Fs interface.
func (fs *WrappedFs) MkdirAll(path string, perm fs.FileMode) error {
	return fs.MkdirAllFunc(path, perm)
}

// Name satisfies the afero.Fs interface.
func (fs *WrappedFs) Name() string {
	return fs.NameFunc()
}

// Open satisfies the afero.Fs interface.
func (fs *WrappedFs) Open(name string) (afero.File, error) {
	return fs.OpenFunc(name)
}

// OpenFile satisfies the afero.Fs interface.
func (fs *WrappedFs) OpenFile(name string, flag int, perm fs.FileMode) (afero.File, error) {
	return fs.OpenFileFunc(name, flag, perm)
}

// Remove satisfies the afero.Fs interface.
func (fs *WrappedFs) Remove(name string) error {
	return fs.RemoveFunc(name)
}

// RemoveAll satisfies the afero.Fs interface.
func (fs *WrappedFs) RemoveAll(path string) error {
	return fs.RemoveAllFunc(path)
}

// Rename satisfies the afero.Fs interface.
func (fs *WrappedFs) Rename(oldname string, newname string) error {
	return fs.RenameFunc(oldname, newname)
}

// Stat satisfies the afero.Fs interface.
func (fs *WrappedFs) Stat(name string) (fs.FileInfo, error) {
	return fs.StatFunc(name)
}

// WrapFs wraps a afero.Fs with custom implementations.
func WrapFs(fs afero.Fs, wrappedFs WrappedFs) *WrappedFs { //nolint: cyclop,dupl
	if wrappedFs.ChmodFunc == nil {
		wrappedFs.ChmodFunc = fs.Chmod
	}

	if wrappedFs.ChownFunc == nil {
		wrappedFs.ChownFunc = fs.Chown
	}

	if wrappedFs.ChtimesFunc == nil {
		wrappedFs.ChtimesFunc = fs.Chtimes
	}

	if wrappedFs.CreateFunc == nil {
		wrappedFs.CreateFunc = fs.Create
	}

	if wrappedFs.MkdirFunc == nil {
		wrappedFs.MkdirFunc = fs.Mkdir
	}

	if wrappedFs.MkdirAllFunc == nil {
		wrappedFs.MkdirAllFunc = fs.MkdirAll
	}

	if wrappedFs.NameFunc == nil {
		wrappedFs.NameFunc = fs.Name
	}

	if wrappedFs.OpenFunc == nil {
		wrappedFs.OpenFunc = fs.Open
	}

	if wrappedFs.OpenFileFunc == nil {
		wrappedFs.OpenFileFunc = fs.OpenFile
	}

	if wrappedFs.RemoveFunc == nil {
		wrappedFs.RemoveFunc = fs.Remove
	}

	if wrappedFs.RemoveAllFunc == nil {
		wrappedFs.RemoveAllFunc = fs.RemoveAll
	}

	if wrappedFs.RenameFunc == nil {
		wrappedFs.RenameFunc = fs.Rename
	}

	if wrappedFs.StatFunc == nil {
		wrappedFs.StatFunc = fs.Stat
	}

	return &wrappedFs
}
