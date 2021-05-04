package aferomock

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFileInfo_Name(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario     string
		mockFileInfo FileInfoMocker
		expected     string
	}{
		{
			scenario: "not empty",
			mockFileInfo: MockFileInfo(func(i *FileInfo) {
				i.On("Name").Return("name")
			}),
			expected: "name",
		},
		{
			scenario: "empty",
			mockFileInfo: MockFileInfo(func(i *FileInfo) {
				i.On("Name").Return("")
			}),
			expected: "",
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

func TestFileInfo_Size(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario     string
		mockFileInfo FileInfoMocker
		expected     int64
	}{
		{
			scenario: "int",
			mockFileInfo: MockFileInfo(func(i *FileInfo) {
				i.On("Size").Return(10)
			}),
			expected: 10,
		},
		{
			scenario: "int",
			mockFileInfo: MockFileInfo(func(i *FileInfo) {
				i.On("Size").Return(int64(20))
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

func TestFileInfo_Mode(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario     string
		mockFileInfo FileInfoMocker
		expected     os.FileMode
	}{
		{
			scenario: "int",
			mockFileInfo: MockFileInfo(func(i *FileInfo) {
				i.On("Mode").Return(0777)
			}),
			expected: 0777,
		},
		{
			scenario: "filemode",
			mockFileInfo: MockFileInfo(func(i *FileInfo) {
				i.On("Mode").Return(os.FileMode(0777))
			}),
			expected: 0777,
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

func TestFileInfo_ModTime(t *testing.T) {
	t.Parallel()

	ts := time.Now()

	testCases := []struct {
		scenario     string
		mockFileInfo FileInfoMocker
		expected     time.Time
	}{
		{
			scenario: "empty",
			mockFileInfo: MockFileInfo(func(i *FileInfo) {
				i.On("ModTime").Return(time.Time{})
			}),
		},
		{
			scenario: "not empty",
			mockFileInfo: MockFileInfo(func(i *FileInfo) {
				i.On("ModTime").Return(ts)
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

func TestFileInfo_IsDir(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario     string
		mockFileInfo FileInfoMocker
		expected     bool
	}{
		{
			scenario: "false",
			mockFileInfo: MockFileInfo(func(i *FileInfo) {
				i.On("IsDir").Return(false)
			}),
		},
		{
			scenario: "true",
			mockFileInfo: MockFileInfo(func(i *FileInfo) {
				i.On("IsDir").Return(true)
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

func TestFileInfo_Sys(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario     string
		mockFileInfo FileInfoMocker
		expected     interface{}
	}{
		{
			scenario: "header",
			mockFileInfo: MockFileInfo(func(i *FileInfo) {
				i.On("Sys").Return(&struct{}{})
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
