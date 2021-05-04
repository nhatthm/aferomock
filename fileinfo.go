package aferomock

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// FileInfoMocker is FileInfo mocker.
type FileInfoMocker func(tb testing.TB) *FileInfo

// NoMockFileInfo is no mock FileInfo.
var NoMockFileInfo = MockFileInfo()

var _ os.FileInfo = (*FileInfo)(nil)

// FileInfo is a os.FileInfo.
type FileInfo struct {
	mock.Mock
}

// Name satisfies os.FileInfo.
func (f *FileInfo) Name() string {
	return f.Called().String(0)
}

// Size satisfies os.FileInfo.
func (f *FileInfo) Size() int64 {
	ret := f.Called().Get(0)

	if ret, ok := ret.(int); ok {
		return int64(ret)
	}

	return ret.(int64)
}

// Mode satisfies os.FileInfo.
func (f *FileInfo) Mode() os.FileMode {
	ret := f.Called().Get(0)

	if ret, ok := ret.(int); ok {
		return os.FileMode(ret)
	}

	return ret.(os.FileMode)
}

// ModTime satisfies os.FileInfo.
func (f *FileInfo) ModTime() time.Time {
	return f.Called().Get(0).(time.Time)
}

// IsDir satisfies os.FileInfo.
func (f *FileInfo) IsDir() bool {
	return f.Called().Bool(0)
}

// Sys satisfies os.FileInfo.
func (f *FileInfo) Sys() interface{} {
	return f.Called().Get(0)
}

// NewFileInfo mocks os.FileInfo interface.
func NewFileInfo(mocks ...func(i *FileInfo)) *FileInfo {
	i := &FileInfo{}

	for _, m := range mocks {
		m(i)
	}

	return i
}

// MockFileInfo creates FileInfo mock with cleanup to ensure all the expectations are met.
func MockFileInfo(mocks ...func(i *FileInfo)) FileInfoMocker {
	return func(tb testing.TB) *FileInfo {
		tb.Helper()

		i := NewFileInfo(mocks...)

		tb.Cleanup(func() {
			assert.True(tb, i.Mock.AssertExpectations(tb))
		})

		return i
	}
}
