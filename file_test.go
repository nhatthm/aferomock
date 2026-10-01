package aferomock_test

import (
	"errors"
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"go.nhat.io/aferomock"
)

func TestFile_Close(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFile      aferomock.FileMocker
		expectedError error
	}{
		{
			scenario: "callback",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Close().
					RunAndReturn(func() error {
						return errors.New("callback")
					})
			}),
			expectedError: errors.New("callback"),
		},
		{
			scenario: "error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Close().
					Return(errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Close().
					Return(nil)
			}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual := tc.mockFile(t).Close()

			require.Equal(t, tc.expectedError, actual)
		})
	}
}

func TestFile_Close_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFile(func(f *aferomock.File) { //nolint: gosec
			f.EXPECT().Close()
		})(t).Close() //nolint: errcheck
	})
}

func TestFile_Name(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		expectedResult string
	}{
		{
			scenario: "callback",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Name().
					RunAndReturn(func() string {
						return "callback"
					})
			}),
			expectedResult: "callback",
		},
		{
			scenario: "no name",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Name().
					Return("")
			}),
		},
		{
			scenario: "has name",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Name().
					Return("name")
			}),
			expectedResult: "name",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual := tc.mockFile(t).Name()

			require.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestFile_Name_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFile(func(f *aferomock.File) { //nolint: gosec
			f.EXPECT().Name()
		})(t).Name() //nolint: errcheck
	})
}

func TestFile_Read(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		expectedResult int
		expectedError  error
	}{
		{
			scenario: "callback error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Read([]byte("hello")).
					RunAndReturn(func([]byte) (int, error) {
						return 0, errors.New("callback")
					})
			}),
			expectedError: errors.New("callback"),
		},
		{
			scenario: "callback success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Read([]byte("hello")).
					RunAndReturn(func(b []byte) (int, error) {
						return len(b), nil
					})
			}),
			expectedResult: 5,
		},
		{
			scenario: "error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Read([]byte("hello")).
					Return(0, errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Read([]byte("hello")).
					Return(5, nil)
			}),
			expectedResult: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := tc.mockFile(t).Read([]byte("hello"))

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestFile_Read_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFile(func(f *aferomock.File) { //nolint: gosec
			f.EXPECT().Read(mock.Anything)
		})(t).Read(nil) //nolint: errcheck
	})
}

func TestFile_ReadAt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		expectedResult int
		expectedError  error
	}{
		{
			scenario: "callback error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().ReadAt([]byte("hello"), int64(1)).
					RunAndReturn(func([]byte, int64) (int, error) {
						return 0, errors.New("callback")
					})
			}),
			expectedError: errors.New("callback"),
		},
		{
			scenario: "callback success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().ReadAt([]byte("hello"), int64(1)).
					RunAndReturn(func(b []byte, offset int64) (int, error) {
						return len(b[offset:]), nil
					})
			}),
			expectedResult: 4,
		},
		{
			scenario: "error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().ReadAt([]byte("hello"), int64(1)).
					Return(0, errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().ReadAt([]byte("hello"), int64(1)).
					Return(4, nil)
			}),
			expectedResult: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := tc.mockFile(t).ReadAt([]byte("hello"), 1)

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestFile_ReadAt_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFile(func(f *aferomock.File) { //nolint: gosec
			f.EXPECT().ReadAt(mock.Anything, mock.Anything)
		})(t).ReadAt(nil, 0) //nolint: errcheck
	})
}

func TestFile_Readdir(t *testing.T) {
	t.Parallel()

	fi := aferomock.NopFileInfo(t)

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		expectedResult []fs.FileInfo
		expectedError  error
	}{
		{
			scenario: "callback error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Readdir(1).
					RunAndReturn(func(int) ([]fs.FileInfo, error) {
						return nil, errors.New("callback")
					})
			}),
			expectedError: errors.New("callback"),
		},
		{
			scenario: "callback success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Readdir(1).
					RunAndReturn(func(int) ([]fs.FileInfo, error) {
						return []fs.FileInfo{fi}, nil
					})
			}),
			expectedResult: []fs.FileInfo{fi},
		},
		{
			scenario: "error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Readdir(1).
					Return(nil, errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Readdir(1).
					Return([]fs.FileInfo{fi}, nil)
			}),
			expectedResult: []fs.FileInfo{fi},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := tc.mockFile(t).Readdir(1)

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestFile_Readdir_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFile(func(f *aferomock.File) { //nolint: gosec
			f.EXPECT().Readdir(mock.Anything)
		})(t).Readdir(0) //nolint: errcheck
	})
}

func TestFile_Readdirnames(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		expectedResult []string
		expectedError  error
	}{
		{
			scenario: "callback error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Readdirnames(1).
					RunAndReturn(func(int) ([]string, error) {
						return nil, errors.New("callback")
					})
			}),
			expectedError: errors.New("callback"),
		},
		{
			scenario: "callback success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Readdirnames(1).
					RunAndReturn(func(int) ([]string, error) {
						return []string{"foobar"}, nil
					})
			}),
			expectedResult: []string{"foobar"},
		},
		{
			scenario: "error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Readdirnames(1).
					Return(nil, errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Readdirnames(1).
					Return([]string{"foobar"}, nil)
			}),
			expectedResult: []string{"foobar"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := tc.mockFile(t).Readdirnames(1)

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestFile_Readdirnames_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFile(func(f *aferomock.File) { //nolint: gosec
			f.EXPECT().Readdirnames(mock.Anything)
		})(t).Readdirnames(0) //nolint: errcheck
	})
}

func TestFile_Seek(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		expectedResult int64
		expectedError  error
	}{
		{
			scenario: "callback error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Seek(int64(64), 10).
					RunAndReturn(func(int64, int) (int64, error) {
						return 0, errors.New("callback")
					})
			}),
			expectedError: errors.New("callback"),
		},
		{
			scenario: "callback success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Seek(int64(64), 10).
					RunAndReturn(func(int64, int) (int64, error) {
						return 10, nil
					})
			}),
			expectedResult: 10,
		},
		{
			scenario: "error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Seek(int64(64), 10).
					Return(int64(0), errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Seek(int64(64), 10).
					Return(int64(10), nil)
			}),
			expectedResult: int64(10),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := tc.mockFile(t).Seek(64, 10)

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestFile_Seek_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFile(func(f *aferomock.File) { //nolint: gosec
			f.EXPECT().Seek(mock.Anything, mock.Anything)
		})(t).Seek(0, 0) //nolint: errcheck
	})
}

func TestFile_Stat(t *testing.T) {
	t.Parallel()

	fi := aferomock.NopFileInfo(t)

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		expectedResult fs.FileInfo
		expectedError  error
	}{
		{
			scenario: "callback error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Stat().
					RunAndReturn(func() (fs.FileInfo, error) {
						return nil, errors.New("callback")
					})
			}),
			expectedError: errors.New("callback"),
		},
		{
			scenario: "callback success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Stat().
					RunAndReturn(func() (fs.FileInfo, error) {
						return fi, nil
					})
			}),
			expectedResult: fi,
		},
		{
			scenario: "error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Stat().
					Return(nil, errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Stat().
					Return(fi, nil)
			}),
			expectedResult: fi,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := tc.mockFile(t).Stat()

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestFile_Stat_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFile(func(f *aferomock.File) { //nolint: gosec
			f.EXPECT().Stat()
		})(t).Stat() //nolint: errcheck
	})
}

func TestFile_Sync(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFile      aferomock.FileMocker
		expectedError error
	}{
		{
			scenario: "callback",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Sync().
					RunAndReturn(func() error {
						return errors.New("callback")
					})
			}),
			expectedError: errors.New("callback"),
		},
		{
			scenario: "error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Sync().
					Return(errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Sync().
					Return(nil)
			}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual := tc.mockFile(t).Sync()

			require.Equal(t, tc.expectedError, actual)
		})
	}
}

func TestFile_Sync_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFile(func(f *aferomock.File) { //nolint: gosec
			f.EXPECT().Sync()
		})(t).Sync() //nolint: errcheck
	})
}

func TestFile_Truncate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFile      aferomock.FileMocker
		expectedError error
	}{
		{
			scenario: "callback error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Truncate(int64(64)).
					RunAndReturn(func(int64) error {
						return errors.New("callback")
					})
			}),
			expectedError: errors.New("callback"),
		},
		{
			scenario: "callback success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Truncate(int64(64)).
					RunAndReturn(func(int64) error {
						return nil
					})
			}),
		},
		{
			scenario: "error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Truncate(int64(64)).
					Return(errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Truncate(int64(64)).
					Return(nil)
			}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			err := tc.mockFile(t).Truncate(64)

			require.Equal(t, err, tc.expectedError)
		})
	}
}

func TestFile_Truncate_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFile(func(f *aferomock.File) { //nolint: gosec
			f.EXPECT().Truncate(mock.Anything)
		})(t).Truncate(0) //nolint: errcheck
	})
}

func TestFile_Write(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		expectedResult int
		expectedError  error
	}{
		{
			scenario: "callback error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Write([]byte("hello")).
					RunAndReturn(func([]byte) (int, error) {
						return 0, errors.New("callback")
					})
			}),
			expectedError: errors.New("callback"),
		},
		{
			scenario: "callback success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Write([]byte("hello")).
					RunAndReturn(func(b []byte) (int, error) {
						return len(b), nil
					})
			}),
			expectedResult: 5,
		},
		{
			scenario: "error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Write([]byte("hello")).
					Return(0, errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().Write([]byte("hello")).
					Return(5, nil)
			}),
			expectedResult: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := tc.mockFile(t).Write([]byte("hello"))

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestFile_Write_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFile(func(f *aferomock.File) { //nolint: gosec
			f.EXPECT().Write(mock.Anything)
		})(t).Write(nil) //nolint: errcheck
	})
}

func TestFile_WriteAt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		expectedResult int
		expectedError  error
	}{
		{
			scenario: "callback error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().WriteAt([]byte("hello"), int64(1)).
					RunAndReturn(func([]byte, int64) (int, error) {
						return 0, errors.New("callback")
					})
			}),
			expectedError: errors.New("callback"),
		},
		{
			scenario: "callback success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().WriteAt([]byte("hello"), int64(1)).
					RunAndReturn(func(b []byte, offset int64) (int, error) {
						return len(b[offset:]), nil
					})
			}),
			expectedResult: 4,
		},
		{
			scenario: "error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().WriteAt([]byte("hello"), int64(1)).
					Return(0, errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().WriteAt([]byte("hello"), int64(1)).
					Return(4, nil)
			}),
			expectedResult: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := tc.mockFile(t).WriteAt([]byte("hello"), 1)

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestFile_WriteAt_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFile(func(f *aferomock.File) { //nolint: gosec
			f.EXPECT().WriteAt(mock.Anything, mock.Anything)
		})(t).WriteAt(nil, 0) //nolint: errcheck
	})
}

func TestFile_WriteString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		expectedResult int
		expectedError  error
	}{
		{
			scenario: "callback error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().WriteString("hello").
					RunAndReturn(func(string) (int, error) {
						return 0, errors.New("callback")
					})
			}),
			expectedError: errors.New("callback"),
		},
		{
			scenario: "callback success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().WriteString("hello").
					RunAndReturn(func(b string) (int, error) {
						return len(b), nil
					})
			}),
			expectedResult: 5,
		},
		{
			scenario: "error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().WriteString("hello").
					Return(0, errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.EXPECT().WriteString("hello").
					Return(5, nil)
			}),
			expectedResult: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := tc.mockFile(t).WriteString("hello")

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestFile_WriteString_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFile(func(f *aferomock.File) { //nolint: gosec
			f.EXPECT().WriteString(mock.Anything)
		})(t).WriteString("") //nolint: errcheck
	})
}
