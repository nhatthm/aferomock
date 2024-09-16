package aferomock_test

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/aferomock"
)

func TestFileInfo_Name(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario     string
		mockFileInfo aferomock.FileInfoMocker
		expected     string
	}{
		{
			scenario: "callback",
			mockFileInfo: aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
				fi.On("Name").Return(func() string {
					return "callback"
				})
			}),
			expected: "callback",
		},
		{
			scenario: "no name",
			mockFileInfo: aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
				fi.On("Name").Return("")
			}),
		},
		{
			scenario: "has name",
			mockFileInfo: aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
				fi.On("Name").Return("name")
			}),
			expected: "name",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			result := tc.mockFileInfo(t).Name()

			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestFileInfo_Name_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
			fi.On("Name")
		})(t).Name()
	})
}

func TestFileInfo_Size(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario     string
		mockFileInfo aferomock.FileInfoMocker
		expected     int64
	}{
		{
			scenario: "callback",
			mockFileInfo: aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
				fi.On("Size").Return(func() int64 {
					return 10
				})
			}),
			expected: 10,
		},
		{
			scenario: "int64",
			mockFileInfo: aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
				fi.On("Size").Return(int64(20))
			}),
			expected: 20,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			result := tc.mockFileInfo(t).Size()

			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestFileInfo_Size_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
			fi.On("Size")
		})(t).Size()
	})
}

func TestFileInfo_Mode(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario     string
		mockFileInfo aferomock.FileInfoMocker
		expected     os.FileMode
	}{
		{
			scenario: "callback",
			mockFileInfo: aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
				fi.On("Mode").Return(func() os.FileMode {
					return os.FileMode(0o644)
				})
			}),
			expected: 0o644,
		},
		{
			scenario: "filemode",
			mockFileInfo: aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
				fi.On("Mode").Return(os.FileMode(0o777))
			}),
			expected: 0o777,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			result := tc.mockFileInfo(t).Mode()

			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestFileInfo_Mode_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
			fi.On("Mode")
		})(t).Mode()
	})
}

func TestFileInfo_ModTime(t *testing.T) {
	t.Parallel()

	ts := time.Now()

	testCases := []struct {
		scenario     string
		mockFileInfo aferomock.FileInfoMocker
		expected     time.Time
	}{
		{
			scenario: "callback",
			mockFileInfo: aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
				fi.On("ModTime").Return(func() time.Time {
					return ts
				})
			}),
			expected: ts,
		},
		{
			scenario: "empty",
			mockFileInfo: aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
				fi.On("ModTime").Return(time.Time{})
			}),
		},
		{
			scenario: "not empty",
			mockFileInfo: aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
				fi.On("ModTime").Return(ts)
			}),
			expected: ts,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			result := tc.mockFileInfo(t).ModTime()

			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestFileInfo_ModTime_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
			fi.On("ModTime")
		})(t).ModTime()
	})
}

func TestFileInfo_IsDir(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario     string
		mockFileInfo aferomock.FileInfoMocker
		expected     bool
	}{
		{
			scenario: "callback",
			mockFileInfo: aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
				fi.On("IsDir").Return(func() bool {
					return true
				})
			}),
			expected: true,
		},
		{
			scenario: "false",
			mockFileInfo: aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
				fi.On("IsDir").Return(false)
			}),
		},
		{
			scenario: "true",
			mockFileInfo: aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
				fi.On("IsDir").Return(true)
			}),
			expected: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			result := tc.mockFileInfo(t).IsDir()

			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestFileInfo_IsDir_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
			fi.On("IsDir")
		})(t).IsDir()
	})
}

func TestFileInfo_Sys(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario     string
		mockFileInfo aferomock.FileInfoMocker
		expected     interface{}
	}{
		{
			scenario: "callback",
			mockFileInfo: aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
				fi.On("Sys").Return(func() interface{} {
					return &struct{}{}
				})
			}),
			expected: &struct{}{},
		},
		{
			scenario: "header",
			mockFileInfo: aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
				fi.On("Sys").Return(&struct{}{})
			}),
			expected: &struct{}{},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			result := tc.mockFileInfo(t).Sys()

			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestFileInfo_Sys_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFileInfo(func(fi *aferomock.FileInfo) {
			fi.On("Sys")
		})(t).Sys()
	})
}
