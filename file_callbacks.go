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
func (f *FileCallbacks) Close() error {
	return f.CloseFunc()
}

// Name satisfies the afero.File interface.
func (f *FileCallbacks) Name() string {
	return f.NameFunc()
}

// Read satisfies the afero.File interface.
func (f *FileCallbacks) Read(p []byte) (int, error) {
	return f.ReadFunc(p)
}

// ReadAt satisfies the afero.File interface.
func (f *FileCallbacks) ReadAt(p []byte, off int64) (int, error) {
	return f.ReadAtFunc(p, off)
}

// Readdir satisfies the afero.File interface.
func (f *FileCallbacks) Readdir(count int) ([]fs.FileInfo, error) {
	return f.ReaddirFunc(count)
}

// Readdirnames satisfies the afero.File interface.
func (f *FileCallbacks) Readdirnames(n int) ([]string, error) {
	return f.ReaddirnamesFunc(n)
}

// Seek satisfies the afero.File interface.
func (f *FileCallbacks) Seek(offset int64, whence int) (int64, error) {
	return f.SeekFunc(offset, whence)
}

// Stat satisfies the afero.File interface.
func (f *FileCallbacks) Stat() (fs.FileInfo, error) {
	return f.StatFunc()
}

// Sync satisfies the afero.File interface.
func (f *FileCallbacks) Sync() error {
	return f.SyncFunc()
}

// Truncate satisfies the afero.File interface.
func (f *FileCallbacks) Truncate(size int64) error {
	return f.TruncateFunc(size)
}

// Write satisfies the afero.File interface.
func (f *FileCallbacks) Write(p []byte) (int, error) {
	return f.WriteFunc(p)
}

// WriteAt satisfies the afero.File interface.
func (f *FileCallbacks) WriteAt(p []byte, off int64) (int, error) {
	return f.WriteAtFunc(p, off)
}

// WriteString satisfies the afero.File interface.
func (f *FileCallbacks) WriteString(s string) (int, error) {
	return f.WriteStringFunc(s)
}

// OverrideFile overrides the file methods with the given callbacks.
func OverrideFile(file afero.File, callbacks FileCallbacks) *FileCallbacks { //nolint: cyclop,dupl
	if callbacks.CloseFunc == nil {
		callbacks.CloseFunc = file.Close
	}

	if callbacks.NameFunc == nil {
		callbacks.NameFunc = file.Name
	}

	if callbacks.ReadFunc == nil {
		callbacks.ReadFunc = file.Read
	}

	if callbacks.ReadAtFunc == nil {
		callbacks.ReadAtFunc = file.ReadAt
	}

	if callbacks.ReaddirFunc == nil {
		callbacks.ReaddirFunc = file.Readdir
	}

	if callbacks.ReaddirnamesFunc == nil {
		callbacks.ReaddirnamesFunc = file.Readdirnames
	}

	if callbacks.SeekFunc == nil {
		callbacks.SeekFunc = file.Seek
	}

	if callbacks.StatFunc == nil {
		callbacks.StatFunc = file.Stat
	}

	if callbacks.SyncFunc == nil {
		callbacks.SyncFunc = file.Sync
	}

	if callbacks.TruncateFunc == nil {
		callbacks.TruncateFunc = file.Truncate
	}

	if callbacks.WriteFunc == nil {
		callbacks.WriteFunc = file.Write
	}

	if callbacks.WriteAtFunc == nil {
		callbacks.WriteAtFunc = file.WriteAt
	}

	if callbacks.WriteStringFunc == nil {
		callbacks.WriteStringFunc = file.WriteString
	}

	return &callbacks
}
