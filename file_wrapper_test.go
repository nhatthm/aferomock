package aferomock_test

import (
	"errors"
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nhat.io/aferomock"
)

func TestWrappedFile_Close(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFile      aferomock.FileMocker
		wrappedFile   aferomock.WrappedFile
		expectedError error
	}{
		{
			scenario: "upstream - error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("Close").
					Return(errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "upstream - success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("Close").
					Return(nil)
			}),
		},
		{
			scenario: "wrapped - error",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
				CloseFunc: func() error {
					return errors.New("error")
				},
			},
			expectedError: errors.New("error"),
		},
		{
			scenario: "wrapped - success",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
				CloseFunc: func() error {
					return nil
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual := aferomock.WrapFile(tc.mockFile(t), tc.wrappedFile).Close()

			require.Equal(t, actual, tc.expectedError)
		})
	}
}

func TestWrappedFile_Name(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		wrappedFile    aferomock.WrappedFile
		expectedResult string
	}{
		{
			scenario: "upstream - no name",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("Name").
					Return("")
			}),
		},
		{
			scenario: "upstream - has name",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("Name").
					Return("name")
			}),
			expectedResult: "name",
		},
		{
			scenario: "wrapped - no name",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
				NameFunc: func() string {
					return ""
				},
			},
		},
		{
			scenario: "wrapped - has name",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
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

			actual := aferomock.WrapFile(tc.mockFile(t), tc.wrappedFile).Name()

			require.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestWrappedFile_Read(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		wrappedFile    aferomock.WrappedFile
		expectedResult int
		expectedError  error
	}{
		{
			scenario: "upstream - error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("Read", []byte("hello")).
					Return(0, errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "upstream - success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("Read", []byte("hello")).
					Return(5, nil)
			}),
			expectedResult: 5,
		},
		{
			scenario: "wrapped - error",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
				ReadFunc: func([]byte) (int, error) {
					return 0, errors.New("error")
				},
			},
			expectedError: errors.New("error"),
		},
		{
			scenario: "wrapped - success",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
				ReadFunc: func(b []byte) (int, error) {
					return len(b), nil
				},
			},
			expectedResult: 5,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := aferomock.WrapFile(tc.mockFile(t), tc.wrappedFile).Read([]byte("hello"))

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestWrappedFile_ReadAt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		wrappedFile    aferomock.WrappedFile
		expectedResult int
		expectedError  error
	}{
		{
			scenario: "upstream - error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("ReadAt", []byte("hello"), int64(1)).
					Return(0, errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "upstream - success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("ReadAt", []byte("hello"), int64(1)).
					Return(4, nil)
			}),
			expectedResult: 4,
		},
		{
			scenario: "wrapped - error",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
				ReadAtFunc: func([]byte, int64) (int, error) {
					return 0, errors.New("error")
				},
			},
			expectedError: errors.New("error"),
		},
		{
			scenario: "wrapped - success",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
				ReadAtFunc: func(b []byte, off int64) (int, error) {
					return len(b[off:]), nil
				},
			},
			expectedResult: 4,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := aferomock.WrapFile(tc.mockFile(t), tc.wrappedFile).ReadAt([]byte("hello"), 1)

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestWrappedFile_Readdir(t *testing.T) {
	t.Parallel()

	fi := aferomock.NopFileInfo(t)

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		wrappedFile    aferomock.WrappedFile
		expectedResult []fs.FileInfo
		expectedError  error
	}{
		{
			scenario: "upstream - error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("Readdir", 1).
					Return(nil, errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "upstream - success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("Readdir", 1).
					Return([]fs.FileInfo{fi}, nil)
			}),
			expectedResult: []fs.FileInfo{fi},
		},
		{
			scenario: "wrapped - error",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
				ReaddirFunc: func(int) ([]fs.FileInfo, error) {
					return nil, errors.New("error")
				},
			},
			expectedError: errors.New("error"),
		},
		{
			scenario: "wrapped - success",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
				ReaddirFunc: func(int) ([]fs.FileInfo, error) {
					return []fs.FileInfo{fi}, nil
				},
			},
			expectedResult: []fs.FileInfo{fi},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := aferomock.WrapFile(tc.mockFile(t), tc.wrappedFile).Readdir(1)

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestWrappedFile_Readdirnames(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		wrappedFile    aferomock.WrappedFile
		expectedResult []string
		expectedError  error
	}{
		{
			scenario: "upstream - error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("Readdirnames", 1).
					Return(nil, errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "upstream - success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("Readdirnames", 1).
					Return([]string{"foobar"}, nil)
			}),
			expectedResult: []string{"foobar"},
		},
		{
			scenario: "wrapped - error",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
				ReaddirnamesFunc: func(int) ([]string, error) {
					return nil, errors.New("error")
				},
			},
			expectedError: errors.New("error"),
		},
		{
			scenario: "wrapped - success",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
				ReaddirnamesFunc: func(int) ([]string, error) {
					return []string{"foobar"}, nil
				},
			},
			expectedResult: []string{"foobar"},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := aferomock.WrapFile(tc.mockFile(t), tc.wrappedFile).Readdirnames(1)

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestWrappedFile_Seek(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		wrappedFile    aferomock.WrappedFile
		expectedResult int64
		expectedError  error
	}{
		{
			scenario: "upstream - error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("Seek", int64(64), 10).
					Return(int64(0), errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "upstream - success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("Seek", int64(64), 10).
					Return(int64(10), nil)
			}),
			expectedResult: int64(10),
		},
		{
			scenario: "wrapped - error",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
				SeekFunc: func(int64, int) (int64, error) {
					return 0, errors.New("error")
				},
			},
			expectedError: errors.New("error"),
		},
		{
			scenario: "wrapped - success",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
				SeekFunc: func(int64, int) (int64, error) {
					return 10, nil
				},
			},
			expectedResult: int64(10),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := aferomock.WrapFile(tc.mockFile(t), tc.wrappedFile).Seek(64, 10)

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestWrappedFile_Stat(t *testing.T) {
	t.Parallel()

	fi := aferomock.NopFileInfo(t)

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		wrappedFile    aferomock.WrappedFile
		expectedResult fs.FileInfo
		expectedError  error
	}{
		{
			scenario: "upstream - error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("Stat").
					Return(nil, errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "upstream - success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("Stat").
					Return(fi, nil)
			}),
			expectedResult: fi,
		},
		{
			scenario: "wrapped - error",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
				StatFunc: func() (fs.FileInfo, error) {
					return nil, errors.New("error")
				},
			},
			expectedError: errors.New("error"),
		},
		{
			scenario: "wrapped - success",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
				StatFunc: func() (fs.FileInfo, error) {
					return fi, nil
				},
			},
			expectedResult: fi,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := aferomock.WrapFile(tc.mockFile(t), tc.wrappedFile).Stat()

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestWrappedFile_Sync(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFile      aferomock.FileMocker
		wrappedFile   aferomock.WrappedFile
		expectedError error
	}{
		{
			scenario: "upstream - error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("Sync").
					Return(errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "upstream - success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("Sync").
					Return(nil)
			}),
		},
		{
			scenario: "wrapped - error",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
				SyncFunc: func() error {
					return errors.New("error")
				},
			},
			expectedError: errors.New("error"),
		},
		{
			scenario: "wrapped - success",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
				SyncFunc: func() error {
					return nil
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual := aferomock.WrapFile(tc.mockFile(t), tc.wrappedFile).Sync()

			require.Equal(t, actual, tc.expectedError)
		})
	}
}

func TestWrappedFile_Truncate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFile      aferomock.FileMocker
		wrappedFile   aferomock.WrappedFile
		expectedError error
	}{
		{
			scenario: "upstream - error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("Truncate", int64(64)).
					Return(errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "upstream - success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("Truncate", int64(64)).
					Return(nil)
			}),
		},
		{
			scenario: "wrapped - error",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
				TruncateFunc: func(int64) error {
					return errors.New("error")
				},
			},
			expectedError: errors.New("error"),
		},
		{
			scenario: "wrapped - success",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
				TruncateFunc: func(int64) error {
					return nil
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			err := aferomock.WrapFile(tc.mockFile(t), tc.wrappedFile).Truncate(64)

			require.Equal(t, err, tc.expectedError)
		})
	}
}

func TestWrappedFile_Write(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		wrappedFile    aferomock.WrappedFile
		expectedResult int
		expectedError  error
	}{
		{
			scenario: "upstream - error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("Write", []byte("hello")).
					Return(0, errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "upstream - success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("Write", []byte("hello")).
					Return(5, nil)
			}),
			expectedResult: 5,
		},
		{
			scenario: "wrapped - error",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
				WriteFunc: func([]byte) (int, error) {
					return 0, errors.New("error")
				},
			},
			expectedError: errors.New("error"),
		},
		{
			scenario: "wrapped - success",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
				WriteFunc: func(b []byte) (int, error) {
					return len(b), nil
				},
			},
			expectedResult: 5,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := aferomock.WrapFile(tc.mockFile(t), tc.wrappedFile).Write([]byte("hello"))

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestWrappedFile_WriteAt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		wrappedFile    aferomock.WrappedFile
		expectedResult int
		expectedError  error
	}{
		{
			scenario: "upstream - error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("WriteAt", []byte("hello"), int64(1)).
					Return(0, errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "upstream - success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("WriteAt", []byte("hello"), int64(1)).
					Return(4, nil)
			}),
			expectedResult: 4,
		},
		{
			scenario: "wrapped - error",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
				WriteAtFunc: func([]byte, int64) (int, error) {
					return 0, errors.New("error")
				},
			},
			expectedError: errors.New("error"),
		},
		{
			scenario: "wrapped - success",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
				WriteAtFunc: func(b []byte, off int64) (int, error) {
					return len(b[off:]), nil
				},
			},
			expectedResult: 4,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := aferomock.WrapFile(tc.mockFile(t), tc.wrappedFile).WriteAt([]byte("hello"), 1)

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestWrappedFile_WriteString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		wrappedFile    aferomock.WrappedFile
		expectedResult int
		expectedError  error
	}{
		{
			scenario: "upstream - error",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("WriteString", "hello").
					Return(0, errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "upstream - success",
			mockFile: aferomock.MockFile(func(f *aferomock.File) {
				f.On("WriteString", "hello").
					Return(5, nil)
			}),
			expectedResult: 5,
		},
		{
			scenario: "wrapped - error",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
				WriteStringFunc: func(string) (int, error) {
					return 0, errors.New("error")
				},
			},
			expectedError: errors.New("error"),
		},
		{
			scenario: "wrapped - success",
			mockFile: aferomock.NopFile,
			wrappedFile: aferomock.WrappedFile{
				WriteStringFunc: func(s string) (int, error) {
					return len(s), nil
				},
			},
			expectedResult: 5,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := aferomock.WrapFile(tc.mockFile(t), tc.wrappedFile).WriteString("hello")

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}
