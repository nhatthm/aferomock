package aferomock

import (
	"os"
	"testing"
	"time"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// FsMocker is Fs mocker.
type FsMocker func(tb testing.TB) *Fs

// NoMockFs is no mock Fs.
var NoMockFs = MockFs()

var _ afero.Fs = (*Fs)(nil)

// Fs is a afero.Fs.
type Fs struct {
	mock.Mock
}

// Create satisfies afero.Fs.
func (f *Fs) Create(name string) (afero.File, error) {
	ret := f.Called(name)

	ret1 := ret.Get(0)
	ret2 := ret.Error(1)

	if ret1 == nil {
		return nil, ret2
	}

	return ret1.(afero.File), ret.Error(1)
}

// Mkdir satisfies afero.Fs.
func (f *Fs) Mkdir(name string, perm os.FileMode) error {
	return f.Called(name, perm).Error(0)
}

// MkdirAll satisfies afero.Fs.
func (f *Fs) MkdirAll(path string, perm os.FileMode) error {
	return f.Called(path, perm).Error(0)
}

// Open satisfies afero.Fs.
func (f *Fs) Open(name string) (afero.File, error) {
	ret := f.Called(name)

	ret1 := ret.Get(0)
	ret2 := ret.Error(1)

	if ret1 == nil {
		return nil, ret2
	}

	return ret1.(afero.File), ret.Error(1)
}

// OpenFile satisfies afero.Fs.
func (f *Fs) OpenFile(name string, flag int, perm os.FileMode) (afero.File, error) {
	ret := f.Called(name, flag, perm)

	ret1 := ret.Get(0)
	ret2 := ret.Error(1)

	if ret1 == nil {
		return nil, ret2
	}

	return ret1.(afero.File), ret.Error(1)
}

// Remove satisfies afero.Fs.
func (f *Fs) Remove(name string) error {
	return f.Called(name).Error(0)
}

// RemoveAll satisfies afero.Fs.
func (f *Fs) RemoveAll(path string) error {
	return f.Called(path).Error(0)
}

// Rename satisfies afero.Fs.
func (f *Fs) Rename(oldname, newname string) error {
	return f.Called(oldname, newname).Error(0)
}

// Stat satisfies afero.Fs.
func (f *Fs) Stat(name string) (os.FileInfo, error) {
	ret := f.Called(name)

	ret1 := ret.Get(0)
	ret2 := ret.Error(1)

	if ret1 == nil {
		return nil, ret2
	}

	return ret1.(os.FileInfo), ret.Error(1)
}

// Name satisfies afero.Fs.
func (f *Fs) Name() string {
	return "aferomock.Fs"
}

// Chmod satisfies afero.Fs.
func (f *Fs) Chmod(name string, mode os.FileMode) error {
	return f.Called(name, mode).Error(0)
}

// Chown satisfies afero.Fs.
func (f *Fs) Chown(name string, uid, gid int) error {
	return f.Called(name, uid, gid).Error(0)
}

// Chtimes satisfies afero.Fs.
func (f *Fs) Chtimes(name string, atime time.Time, mtime time.Time) error {
	return f.Called(name, atime, mtime).Error(0)
}

// NewFs mocks afero.Fs interface.
func NewFs(mocks ...func(fs *Fs)) *Fs {
	fs := &Fs{}

	for _, m := range mocks {
		m(fs)
	}

	return fs
}

// MockFs creates Fs mock with cleanup to ensure all the expectations are met.
func MockFs(mocks ...func(fs *Fs)) FsMocker {
	return func(tb testing.TB) *Fs {
		tb.Helper()

		fs := NewFs(mocks...)

		tb.Cleanup(func() {
			assert.True(tb, fs.Mock.AssertExpectations(tb))
		})

		return fs
	}
}
