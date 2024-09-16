package aferomock

import (
	"io/fs"

	"github.com/spf13/afero"
)

var _ fs.File = (*FileCallbacks)(nil)

// FileCallbacks is a callback-based mock for afero.File.
type FileCallbacks struct {
	CloseFunc        func() error
	NameFunc         func() string
	ReadFunc         func(p []byte) (int, error)
	ReadAtFunc       func(p []byte, off int64) (int, error)
	ReaddirFunc      func(count int) ([]fs.FileInfo, error)
	ReaddirnamesFunc func(n int) ([]string, error)
	SeekFunc         func(offset int64, whence int) (int64, error)
	StatFunc         func() (fs.FileInfo, error)
	SyncFunc         func() error
	TruncateFunc     func(size int64) error
	WriteFunc        func(p []byte) (int, error)
	WriteAtFunc      func(p []byte, off int64) (int, error)
	WriteStringFunc  func(s string) (int, error)
}

// Close satisfies the afero.File interface.
func (f FileCallbacks) Close() error {
	return f.CloseFunc()
}

// Name satisfies the afero.File interface.
func (f FileCallbacks) Name() string {
	return f.NameFunc()
}

// Read satisfies the afero.File interface.
func (f FileCallbacks) Read(p []byte) (int, error) {
	return f.ReadFunc(p)
}

// ReadAt satisfies the afero.File interface.
func (f FileCallbacks) ReadAt(p []byte, off int64) (int, error) {
	return f.ReadAtFunc(p, off)
}

// Readdir satisfies the afero.File interface.
func (f FileCallbacks) Readdir(count int) ([]fs.FileInfo, error) {
	return f.ReaddirFunc(count)
}

// Readdirnames satisfies the afero.File interface.
func (f FileCallbacks) Readdirnames(n int) ([]string, error) {
	return f.ReaddirnamesFunc(n)
}

// Seek satisfies the afero.File interface.
func (f FileCallbacks) Seek(offset int64, whence int) (int64, error) {
	return f.SeekFunc(offset, whence)
}

// Stat satisfies the afero.File interface.
func (f FileCallbacks) Stat() (fs.FileInfo, error) {
	return f.StatFunc()
}

// Sync satisfies the afero.File interface.
func (f FileCallbacks) Sync() error {
	return f.SyncFunc()
}

// Truncate satisfies the afero.File interface.
func (f FileCallbacks) Truncate(size int64) error {
	return f.TruncateFunc(size)
}

// Write satisfies the afero.File interface.
func (f FileCallbacks) Write(p []byte) (int, error) {
	return f.WriteFunc(p)
}

// WriteAt satisfies the afero.File interface.
func (f FileCallbacks) WriteAt(p []byte, off int64) (int, error) {
	return f.WriteAtFunc(p, off)
}

// WriteString satisfies the afero.File interface.
func (f FileCallbacks) WriteString(s string) (int, error) {
	return f.WriteStringFunc(s)
}

// OverrideFile overrides the afero.File methods with the provided callbacks.
func OverrideFile(file afero.File, c FileCallbacks) FileCallbacks { //nolint: cyclop,dupl
	if c.CloseFunc == nil {
		c.CloseFunc = file.Close
	}

	if c.NameFunc == nil {
		c.NameFunc = file.Name
	}

	if c.ReadFunc == nil {
		c.ReadFunc = file.Read
	}

	if c.ReadAtFunc == nil {
		c.ReadAtFunc = file.ReadAt
	}

	if c.ReaddirFunc == nil {
		c.ReaddirFunc = file.Readdir
	}

	if c.ReaddirnamesFunc == nil {
		c.ReaddirnamesFunc = file.Readdirnames
	}

	if c.SeekFunc == nil {
		c.SeekFunc = file.Seek
	}

	if c.StatFunc == nil {
		c.StatFunc = file.Stat
	}

	if c.SyncFunc == nil {
		c.SyncFunc = file.Sync
	}

	if c.TruncateFunc == nil {
		c.TruncateFunc = file.Truncate
	}

	if c.WriteFunc == nil {
		c.WriteFunc = file.Write
	}

	if c.WriteAtFunc == nil {
		c.WriteAtFunc = file.WriteAt
	}

	if c.WriteStringFunc == nil {
		c.WriteStringFunc = file.WriteString
	}

	return c
}
