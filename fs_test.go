package aferomock

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/spf13/afero"
	"github.com/spf13/afero/mem"
	"github.com/stretchr/testify/assert"
)

func TestFs_Create(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFs         FsMocker
		expectedResult afero.File
		expectedError  string
	}{
		{
			scenario: "no error",
			mockFs: MockFs(func(fs *Fs) {
				fs.On("Create", "test.txt").
					Return(&mem.File{}, nil)
			}),
			expectedResult: &mem.File{},
		},
		{
			scenario: "error",
			mockFs: MockFs(func(fs *Fs) {
				fs.On("Create", "test.txt").
					Return(nil, errors.New("create error"))
			}),
			expectedError: "create error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := tc.mockFs(t)
			result, err := fs.Create("test.txt")

			assert.Equal(t, tc.expectedResult, result)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestFs_Mkdir(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        FsMocker
		expectedError string
	}{
		{
			scenario: "no error",
			mockFs: MockFs(func(fs *Fs) {
				fs.On("Mkdir", "test", os.ModePerm).
					Return(nil)
			}),
		},
		{
			scenario: "error",
			mockFs: MockFs(func(fs *Fs) {
				fs.On("Mkdir", "test", os.ModePerm).
					Return(errors.New("mkdir error"))
			}),
			expectedError: "mkdir error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := tc.mockFs(t)
			err := fs.Mkdir("test", os.ModePerm)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestFs_MkdirAll(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        FsMocker
		expectedError string
	}{
		{
			scenario: "no error",
			mockFs: MockFs(func(fs *Fs) {
				fs.On("MkdirAll", "path/test", os.ModePerm).
					Return(nil)
			}),
		},
		{
			scenario: "error",
			mockFs: MockFs(func(fs *Fs) {
				fs.On("MkdirAll", "path/test", os.ModePerm).
					Return(errors.New("mkdir all error"))
			}),
			expectedError: "mkdir all error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := tc.mockFs(t)
			err := fs.MkdirAll("path/test", os.ModePerm)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestFs_Open(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFs         FsMocker
		expectedResult afero.File
		expectedError  string
	}{
		{
			scenario: "no error",
			mockFs: MockFs(func(fs *Fs) {
				fs.On("Open", "test.txt").
					Return(&mem.File{}, nil)
			}),
			expectedResult: &mem.File{},
		},
		{
			scenario: "error",
			mockFs: MockFs(func(fs *Fs) {
				fs.On("Open", "test.txt").
					Return(nil, errors.New("create error"))
			}),
			expectedError: "create error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := tc.mockFs(t)
			result, err := fs.Open("test.txt")

			assert.Equal(t, tc.expectedResult, result)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestFs_OpenFile(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFs         FsMocker
		expectedResult afero.File
		expectedError  string
	}{
		{
			scenario: "no error",
			mockFs: MockFs(func(fs *Fs) {
				fs.On("OpenFile", "test.txt", 0, os.ModePerm).
					Return(&mem.File{}, nil)
			}),
			expectedResult: &mem.File{},
		},
		{
			scenario: "error",
			mockFs: MockFs(func(fs *Fs) {
				fs.On("OpenFile", "test.txt", 0, os.ModePerm).
					Return(nil, errors.New("create error"))
			}),
			expectedError: "create error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := tc.mockFs(t)
			result, err := fs.OpenFile("test.txt", 0, os.ModePerm)

			assert.Equal(t, tc.expectedResult, result)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestFs_Remove(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        FsMocker
		expectedError string
	}{
		{
			scenario: "no error",
			mockFs: MockFs(func(fs *Fs) {
				fs.On("Remove", "test.txt").
					Return(nil)
			}),
		},
		{
			scenario: "error",
			mockFs: MockFs(func(fs *Fs) {
				fs.On("Remove", "test.txt").
					Return(errors.New("remove error"))
			}),
			expectedError: "remove error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := tc.mockFs(t)
			err := fs.Remove("test.txt")

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestFs_RemoveAll(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        FsMocker
		expectedError string
	}{
		{
			scenario: "no error",
			mockFs: MockFs(func(fs *Fs) {
				fs.On("RemoveAll", "path/test").
					Return(nil)
			}),
		},
		{
			scenario: "error",
			mockFs: MockFs(func(fs *Fs) {
				fs.On("RemoveAll", "path/test").
					Return(errors.New("remove all error"))
			}),
			expectedError: "remove all error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := tc.mockFs(t)
			err := fs.RemoveAll("path/test")

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestFs_Rename(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        FsMocker
		expectedError string
	}{
		{
			scenario: "no error",
			mockFs: MockFs(func(fs *Fs) {
				fs.On("Rename", "oldname", "newname").
					Return(nil)
			}),
		},
		{
			scenario: "error",
			mockFs: MockFs(func(fs *Fs) {
				fs.On("Rename", "oldname", "newname").
					Return(errors.New("rename error"))
			}),
			expectedError: "rename error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := tc.mockFs(t)
			err := fs.Rename("oldname", "newname")

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestFs_Stat(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mockFs         FsMocker
		expectedResult os.FileInfo
		expectedError  string
	}{
		{
			scenario: "no error",
			mockFs: MockFs(func(fs *Fs) {
				fs.On("Stat", "test.txt").
					Return(&mem.FileInfo{}, nil)
			}),
			expectedResult: &mem.FileInfo{},
		},
		{
			scenario: "error",
			mockFs: MockFs(func(fs *Fs) {
				fs.On("Stat", "test.txt").
					Return(nil, errors.New("stat error"))
			}),
			expectedError: "stat error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := tc.mockFs(t)
			result, err := fs.Stat("test.txt")

			assert.Equal(t, tc.expectedResult, result)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestFs_Name(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "aferomock.Fs", NoMockFs(t).Name())
}

func TestFs_Chmod(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        FsMocker
		expectedError string
	}{
		{
			scenario: "no error",
			mockFs: MockFs(func(fs *Fs) {
				fs.On("Chmod", "test.txt", os.ModePerm).
					Return(nil)
			}),
		},
		{
			scenario: "error",
			mockFs: MockFs(func(fs *Fs) {
				fs.On("Chmod", "test.txt", os.ModePerm).
					Return(errors.New("chmod error"))
			}),
			expectedError: "chmod error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := tc.mockFs(t)
			err := fs.Chmod("test.txt", os.ModePerm)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestFs_Chown(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        FsMocker
		expectedError string
	}{
		{
			scenario: "no error",
			mockFs: MockFs(func(fs *Fs) {
				fs.On("Chown", "test.txt", 0, 0).
					Return(nil)
			}),
		},
		{
			scenario: "error",
			mockFs: MockFs(func(fs *Fs) {
				fs.On("Chown", "test.txt", 0, 0).
					Return(errors.New("chown error"))
			}),
			expectedError: "chown error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := tc.mockFs(t)
			err := fs.Chown("test.txt", 0, 0)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestFs_Chtimes(t *testing.T) {
	t.Parallel()

	ts := time.Now()

	testCases := []struct {
		scenario      string
		mockFs        FsMocker
		expectedError string
	}{
		{
			scenario: "no error",
			mockFs: MockFs(func(fs *Fs) {
				fs.On("Chtimes", "test.txt", ts, ts).
					Return(nil)
			}),
		},
		{
			scenario: "error",
			mockFs: MockFs(func(fs *Fs) {
				fs.On("Chtimes", "test.txt", ts, ts).
					Return(errors.New("chtimes error"))
			}),
			expectedError: "chtimes error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := tc.mockFs(t)
			err := fs.Chtimes("test.txt", ts, ts)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}
