package aferomock

import (
	"io/fs"
	"time"
)

var _ fs.FileInfo = (*FileInfoCallbacks)(nil)

// FileInfoCallbacks is a callback-based mock for fs.FileInfo.
type FileInfoCallbacks struct {
	NameFunc    func() string
	SizeFunc    func() int64
	ModeFunc    func() fs.FileMode
	ModTimeFunc func() time.Time
	IsDirFunc   func() bool
	SysFunc     func() interface{}
}

// Name satisfies the fs.FileInfo interface.
func (f FileInfoCallbacks) Name() string {
	return f.NameFunc()
}

// Size satisfies the fs.FileInfo interface.
func (f FileInfoCallbacks) Size() int64 {
	return f.SizeFunc()
}

// Mode satisfies the fs.FileInfo interface.
func (f FileInfoCallbacks) Mode() fs.FileMode {
	return f.ModeFunc()
}

// ModTime satisfies the fs.FileInfo interface.
func (f FileInfoCallbacks) ModTime() time.Time {
	return f.ModTimeFunc()
}

// IsDir satisfies the fs.FileInfo interface.
func (f FileInfoCallbacks) IsDir() bool {
	return f.IsDirFunc()
}

// Sys satisfies the fs.FileInfo interface.
func (f FileInfoCallbacks) Sys() interface{} {
	return f.SysFunc()
}

// OverrideFileInfo overrides the fs.FileInfo methods with the provided callbacks.
func OverrideFileInfo(fi fs.FileInfo, c FileInfoCallbacks) FileInfoCallbacks {
	if c.NameFunc == nil {
		c.NameFunc = fi.Name
	}

	if c.SizeFunc == nil {
		c.SizeFunc = fi.Size
	}

	if c.ModeFunc == nil {
		c.ModeFunc = fi.Mode
	}

	if c.ModTimeFunc == nil {
		c.ModTimeFunc = fi.ModTime
	}

	if c.IsDirFunc == nil {
		c.IsDirFunc = fi.IsDir
	}

	if c.SysFunc == nil {
		c.SysFunc = fi.Sys
	}

	return c
}
