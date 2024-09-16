package aferomock

import (
	"io/fs"
	"testing"

	"github.com/spf13/afero"
)

var _ afero.Fs = (*Fs)(nil)

// FsMocker is Fs mocker.
type FsMocker func(tb testing.TB) *Fs

// NoMockFs is no mock Fs.
// Deprecated: use NopFs instead.
var NoMockFs = NopFs

// NopFs is no mock Fs.
var NopFs = MockFs()

// MockFs creates Fs mock with cleanup to ensure all the expectations are met.
func MockFs(mocks ...func(fs *Fs)) FsMocker {
	return func(tb testing.TB) *Fs {
		tb.Helper()

		fs := NewFs(tb)

		for _, m := range mocks {
			m(fs)
		}

		fs.On("Name").Maybe().
			Return("aferomock.Fs")

		return fs
	}
}

var _ afero.File = (*File)(nil)

// FileMocker is File mocker.
type FileMocker func(tb testing.TB) *File

// NopFile is no mock File.
var NopFile = MockFile()

// MockFile creates File mock with cleanup to ensure all the expectations are met.
func MockFile(mocks ...func(f *File)) FileMocker {
	return func(tb testing.TB) *File {
		tb.Helper()

		f := NewFile(tb)

		for _, m := range mocks {
			m(f)
		}

		return f
	}
}

var _ fs.FileInfo = (*FileInfo)(nil)

// FileInfoMocker is FileInfo mocker.
type FileInfoMocker func(tb testing.TB) *FileInfo

// NoMockFileInfo is no mock FileInfo.
// Deprecated: use NopFileInfo instead.
var NoMockFileInfo = NopFileInfo

// NopFileInfo is no mock FileInfo.
var NopFileInfo = MockFileInfo()

// MockFileInfo creates FileInfo mock with cleanup to ensure all the expectations are met.
func MockFileInfo(mocks ...func(fi *FileInfo)) FileInfoMocker {
	return func(tb testing.TB) *FileInfo {
		tb.Helper()

		fi := NewFileInfo(tb)

		for _, m := range mocks {
			m(fi)
		}

		return fi
	}
}
