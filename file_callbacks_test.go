package aferomock_test

import (
	"errors"
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nhat.io/aferomock"
)

func TestFileCallbacks_Close(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFile      aferomock.FileMocker
		fileCallbacks aferomock.FileCallbacks
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
			scenario: "overridden - error",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				CloseFunc: func() error {
					return errors.New("error")
				},
			},
			expectedError: errors.New("error"),
		},
		{
			scenario: "overridden - success",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				CloseFunc: func() error {
					return nil
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual := aferomock.OverrideFile(tc.mockFile(t), tc.fileCallbacks).Close()

			require.Equal(t, tc.expectedError, actual)
		})
	}
}

func TestFileCallbacks_Name(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		fileCallbacks  aferomock.FileCallbacks
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
			scenario: "overridden - no name",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				NameFunc: func() string {
					return ""
				},
			},
		},
		{
			scenario: "overridden - has name",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				NameFunc: func() string {
					return "name"
				},
			},
			expectedResult: "name",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual := aferomock.OverrideFile(tc.mockFile(t), tc.fileCallbacks).Name()

			require.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestFileCallbacks_Read(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		fileCallbacks  aferomock.FileCallbacks
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
			scenario: "overridden - error",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				ReadFunc: func([]byte) (int, error) {
					return 0, errors.New("error")
				},
			},
			expectedError: errors.New("error"),
		},
		{
			scenario: "overridden - success",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				ReadFunc: func(b []byte) (int, error) {
					return len(b), nil
				},
			},
			expectedResult: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := aferomock.OverrideFile(tc.mockFile(t), tc.fileCallbacks).Read([]byte("hello"))

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestFileCallbacks_ReadAt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		fileCallbacks  aferomock.FileCallbacks
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
			scenario: "overridden - error",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				ReadAtFunc: func([]byte, int64) (int, error) {
					return 0, errors.New("error")
				},
			},
			expectedError: errors.New("error"),
		},
		{
			scenario: "overridden - success",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				ReadAtFunc: func(b []byte, off int64) (int, error) {
					return len(b[off:]), nil
				},
			},
			expectedResult: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := aferomock.OverrideFile(tc.mockFile(t), tc.fileCallbacks).ReadAt([]byte("hello"), 1)

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestFileCallbacks_Readdir(t *testing.T) {
	t.Parallel()

	fi := aferomock.NopFileInfo(t)

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		fileCallbacks  aferomock.FileCallbacks
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
			scenario: "overridden - error",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				ReaddirFunc: func(int) ([]fs.FileInfo, error) {
					return nil, errors.New("error")
				},
			},
			expectedError: errors.New("error"),
		},
		{
			scenario: "overridden - success",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				ReaddirFunc: func(int) ([]fs.FileInfo, error) {
					return []fs.FileInfo{fi}, nil
				},
			},
			expectedResult: []fs.FileInfo{fi},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := aferomock.OverrideFile(tc.mockFile(t), tc.fileCallbacks).Readdir(1)

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestFileCallbacks_Readdirnames(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		fileCallbacks  aferomock.FileCallbacks
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
			scenario: "overridden - error",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				ReaddirnamesFunc: func(int) ([]string, error) {
					return nil, errors.New("error")
				},
			},
			expectedError: errors.New("error"),
		},
		{
			scenario: "overridden - success",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				ReaddirnamesFunc: func(int) ([]string, error) {
					return []string{"foobar"}, nil
				},
			},
			expectedResult: []string{"foobar"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := aferomock.OverrideFile(tc.mockFile(t), tc.fileCallbacks).Readdirnames(1)

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestFileCallbacks_Seek(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		fileCallbacks  aferomock.FileCallbacks
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
			scenario: "overridden - error",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				SeekFunc: func(int64, int) (int64, error) {
					return 0, errors.New("error")
				},
			},
			expectedError: errors.New("error"),
		},
		{
			scenario: "overridden - success",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				SeekFunc: func(int64, int) (int64, error) {
					return 10, nil
				},
			},
			expectedResult: int64(10),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := aferomock.OverrideFile(tc.mockFile(t), tc.fileCallbacks).Seek(64, 10)

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestFileCallbacks_Stat(t *testing.T) {
	t.Parallel()

	fi := aferomock.NopFileInfo(t)

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		fileCallbacks  aferomock.FileCallbacks
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
			scenario: "overridden - error",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				StatFunc: func() (fs.FileInfo, error) {
					return nil, errors.New("error")
				},
			},
			expectedError: errors.New("error"),
		},
		{
			scenario: "overridden - success",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				StatFunc: func() (fs.FileInfo, error) {
					return fi, nil
				},
			},
			expectedResult: fi,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := aferomock.OverrideFile(tc.mockFile(t), tc.fileCallbacks).Stat()

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestFileCallbacks_Sync(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFile      aferomock.FileMocker
		fileCallbacks aferomock.FileCallbacks
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
			scenario: "overridden - error",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				SyncFunc: func() error {
					return errors.New("error")
				},
			},
			expectedError: errors.New("error"),
		},
		{
			scenario: "overridden - success",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				SyncFunc: func() error {
					return nil
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual := aferomock.OverrideFile(tc.mockFile(t), tc.fileCallbacks).Sync()

			require.Equal(t, tc.expectedError, actual)
		})
	}
}

func TestFileCallbacks_Truncate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFile      aferomock.FileMocker
		fileCallbacks aferomock.FileCallbacks
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
			scenario: "overridden - error",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				TruncateFunc: func(int64) error {
					return errors.New("error")
				},
			},
			expectedError: errors.New("error"),
		},
		{
			scenario: "overridden - success",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				TruncateFunc: func(int64) error {
					return nil
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			err := aferomock.OverrideFile(tc.mockFile(t), tc.fileCallbacks).Truncate(64)

			require.Equal(t, err, tc.expectedError)
		})
	}
}

func TestFileCallbacks_Write(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		fileCallbacks  aferomock.FileCallbacks
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
			scenario: "overridden - error",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				WriteFunc: func([]byte) (int, error) {
					return 0, errors.New("error")
				},
			},
			expectedError: errors.New("error"),
		},
		{
			scenario: "overridden - success",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				WriteFunc: func(b []byte) (int, error) {
					return len(b), nil
				},
			},
			expectedResult: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := aferomock.OverrideFile(tc.mockFile(t), tc.fileCallbacks).Write([]byte("hello"))

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestFileCallbacks_WriteAt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		fileCallbacks  aferomock.FileCallbacks
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
			scenario: "overridden - error",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				WriteAtFunc: func([]byte, int64) (int, error) {
					return 0, errors.New("error")
				},
			},
			expectedError: errors.New("error"),
		},
		{
			scenario: "overridden - success",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				WriteAtFunc: func(b []byte, off int64) (int, error) {
					return len(b[off:]), nil
				},
			},
			expectedResult: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := aferomock.OverrideFile(tc.mockFile(t), tc.fileCallbacks).WriteAt([]byte("hello"), 1)

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestFileCallbacks_WriteString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFile       aferomock.FileMocker
		fileCallbacks  aferomock.FileCallbacks
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
			scenario: "overridden - error",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				WriteStringFunc: func(string) (int, error) {
					return 0, errors.New("error")
				},
			},
			expectedError: errors.New("error"),
		},
		{
			scenario: "overridden - success",
			mockFile: aferomock.NopFile,
			fileCallbacks: aferomock.FileCallbacks{
				WriteStringFunc: func(s string) (int, error) {
					return len(s), nil
				},
			},
			expectedResult: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual, err := aferomock.OverrideFile(tc.mockFile(t), tc.fileCallbacks).WriteString("hello")

			require.Equal(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}
