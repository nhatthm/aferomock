package aferomock_test

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/aferomock"
)

func TestFileInfoCallbacks_Name(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario          string
		mockFileInfo      aferomock.FileInfoMocker
		fileInfoCallbacks aferomock.FileInfoCallbacks
		expectedResult    string
	}{
		{
			scenario: "upstream - no name",
			mockFileInfo: aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
				fi.On("Name").
					Return("")
			}),
		},
		{
			scenario: "upstream - has name",
			mockFileInfo: aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
				fi.On("Name").
					Return("name")
			}),
			expectedResult: "name",
		},
		{
			scenario:     "overridden - no name",
			mockFileInfo: aferomock.NopFileInfo,
			fileInfoCallbacks: aferomock.FileInfoCallbacks{
				NameFunc: func() string {
					return ""
				},
			},
		},
		{
			scenario:     "overridden - has name",
			mockFileInfo: aferomock.NopFileInfo,
			fileInfoCallbacks: aferomock.FileInfoCallbacks{
				NameFunc: func() string {
					return "name"
				},
			},
			expectedResult: "name",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual := aferomock.OverrideFileInfo(tc.mockFileInfo(t), tc.fileInfoCallbacks).Name()

			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestFileInfoCallbacks_Size(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario          string
		mockFileInfo      aferomock.FileInfoMocker
		fileInfoCallbacks aferomock.FileInfoCallbacks
		expectedResult    int64
	}{
		{
			scenario: "upstream",
			mockFileInfo: aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
				fi.On("Size").
					Return(int64(10))
			}),
			expectedResult: 10,
		},
		{
			scenario:     "overridden",
			mockFileInfo: aferomock.NopFileInfo,
			fileInfoCallbacks: aferomock.FileInfoCallbacks{
				SizeFunc: func() int64 {
					return 64
				},
			},
			expectedResult: 64,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual := aferomock.OverrideFileInfo(tc.mockFileInfo(t), tc.fileInfoCallbacks).Size()

			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestFileInfoCallbacks_Mode(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario          string
		mockFileInfo      aferomock.FileInfoMocker
		fileInfoCallbacks aferomock.FileInfoCallbacks
		expectedResult    os.FileMode
	}{
		{
			scenario: "upstream",
			mockFileInfo: aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
				fi.On("Mode").
					Return(os.FileMode(10))
			}),
			expectedResult: 10,
		},
		{
			scenario:     "overridden",
			mockFileInfo: aferomock.NopFileInfo,
			fileInfoCallbacks: aferomock.FileInfoCallbacks{
				ModeFunc: func() os.FileMode {
					return 64
				},
			},
			expectedResult: 64,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual := aferomock.OverrideFileInfo(tc.mockFileInfo(t), tc.fileInfoCallbacks).Mode()

			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestFileInfoCallbacks_ModTime(t *testing.T) {
	t.Parallel()

	ts := time.Now()

	testCases := []struct {
		scenario          string
		mockFileInfo      aferomock.FileInfoMocker
		fileInfoCallbacks aferomock.FileInfoCallbacks
		expectedResult    time.Time
	}{
		{
			scenario: "upstream",
			mockFileInfo: aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
				fi.On("ModTime").
					Return(ts)
			}),
			expectedResult: ts,
		},
		{
			scenario:     "overridden",
			mockFileInfo: aferomock.NopFileInfo,
			fileInfoCallbacks: aferomock.FileInfoCallbacks{
				ModTimeFunc: func() time.Time {
					return ts
				},
			},
			expectedResult: ts,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual := aferomock.OverrideFileInfo(tc.mockFileInfo(t), tc.fileInfoCallbacks).ModTime()

			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestFileInfoCallbacks_IsDir(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario          string
		mockFileInfo      aferomock.FileInfoMocker
		fileInfoCallbacks aferomock.FileInfoCallbacks
		expectedResult    bool
	}{
		{
			scenario: "upstream",
			mockFileInfo: aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
				fi.On("IsDir").
					Return(true)
			}),
			expectedResult: true,
		},
		{
			scenario:     "overridden",
			mockFileInfo: aferomock.NopFileInfo,
			fileInfoCallbacks: aferomock.FileInfoCallbacks{
				IsDirFunc: func() bool {
					return true
				},
			},
			expectedResult: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual := aferomock.OverrideFileInfo(tc.mockFileInfo(t), tc.fileInfoCallbacks).IsDir()

			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestFileInfoCallbacks_Sys(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario          string
		mockFileInfo      aferomock.FileInfoMocker
		fileInfoCallbacks aferomock.FileInfoCallbacks
		expectedResult    interface{}
	}{
		{
			scenario: "upstream",
			mockFileInfo: aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
				fi.On("Sys").
					Return(64)
			}),
			expectedResult: 64,
		},
		{
			scenario:     "overridden",
			mockFileInfo: aferomock.NopFileInfo,
			fileInfoCallbacks: aferomock.FileInfoCallbacks{
				SysFunc: func() interface{} {
					return true
				},
			},
			expectedResult: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual := aferomock.OverrideFileInfo(tc.mockFileInfo(t), tc.fileInfoCallbacks).Sys()

			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}
