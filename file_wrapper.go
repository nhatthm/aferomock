package aferomock

import (
	"io/fs"

	"github.com/spf13/afero"
)

var _ fs.File = (*WrappedFile)(nil)

// WrappedFile is a callback-based mock for afero.File.
type WrappedFile struct {
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
func (f *WrappedFile) Close() error {
	return f.CloseFunc()
}

// Name satisfies the afero.File interface.
func (f *WrappedFile) Name() string {
	return f.NameFunc()
}

// Read satisfies the afero.File interface.
func (f *WrappedFile) Read(p []byte) (int, error) {
	return f.ReadFunc(p)
}

// ReadAt satisfies the afero.File interface.
func (f *WrappedFile) ReadAt(p []byte, off int64) (int, error) {
	return f.ReadAtFunc(p, off)
}

// Readdir satisfies the afero.File interface.
func (f *WrappedFile) Readdir(count int) ([]fs.FileInfo, error) {
	return f.ReaddirFunc(count)
}

// Readdirnames satisfies the afero.File interface.
func (f *WrappedFile) Readdirnames(n int) ([]string, error) {
	return f.ReaddirnamesFunc(n)
}

// Seek satisfies the afero.File interface.
func (f *WrappedFile) Seek(offset int64, whence int) (int64, error) {
	return f.SeekFunc(offset, whence)
}

// Stat satisfies the afero.File interface.
func (f *WrappedFile) Stat() (fs.FileInfo, error) {
	return f.StatFunc()
}

// Sync satisfies the afero.File interface.
func (f *WrappedFile) Sync() error {
	return f.SyncFunc()
}

// Truncate satisfies the afero.File interface.
func (f *WrappedFile) Truncate(size int64) error {
	return f.TruncateFunc(size)
}

// Write satisfies the afero.File interface.
func (f *WrappedFile) Write(p []byte) (int, error) {
	return f.WriteFunc(p)
}

// WriteAt satisfies the afero.File interface.
func (f *WrappedFile) WriteAt(p []byte, off int64) (int, error) {
	return f.WriteAtFunc(p, off)
}

// WriteString satisfies the afero.File interface.
func (f *WrappedFile) WriteString(s string) (int, error) {
	return f.WriteStringFunc(s)
}

// WrapFile wraps an afero.File with a WrappedFile.
func WrapFile(file afero.File, wrappedFile WrappedFile) *WrappedFile { //nolint: cyclop,dupl
	if wrappedFile.CloseFunc == nil {
		wrappedFile.CloseFunc = file.Close
	}

	if wrappedFile.NameFunc == nil {
		wrappedFile.NameFunc = file.Name
	}

	if wrappedFile.ReadFunc == nil {
		wrappedFile.ReadFunc = file.Read
	}

	if wrappedFile.ReadAtFunc == nil {
		wrappedFile.ReadAtFunc = file.ReadAt
	}

	if wrappedFile.ReaddirFunc == nil {
		wrappedFile.ReaddirFunc = file.Readdir
	}

	if wrappedFile.ReaddirnamesFunc == nil {
		wrappedFile.ReaddirnamesFunc = file.Readdirnames
	}

	if wrappedFile.SeekFunc == nil {
		wrappedFile.SeekFunc = file.Seek
	}

	if wrappedFile.StatFunc == nil {
		wrappedFile.StatFunc = file.Stat
	}

	if wrappedFile.SyncFunc == nil {
		wrappedFile.SyncFunc = file.Sync
	}

	if wrappedFile.TruncateFunc == nil {
		wrappedFile.TruncateFunc = file.Truncate
	}

	if wrappedFile.WriteFunc == nil {
		wrappedFile.WriteFunc = file.Write
	}

	if wrappedFile.WriteAtFunc == nil {
		wrappedFile.WriteAtFunc = file.WriteAt
	}

	if wrappedFile.WriteStringFunc == nil {
		wrappedFile.WriteStringFunc = file.WriteString
	}

	return &wrappedFile
}
