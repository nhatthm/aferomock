package aferomock_test

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.nhat.io/aferomock"
)

func TestFs_Create(t *testing.T) {
	t.Parallel()

	f := aferomock.NopFile(t)

	testCases := []struct {
		scenario       string
		mockFs         aferomock.FsMocker
		expectedResult afero.File
		expectedError  string
	}{
		{
			scenario: "callback error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Create", "test.txt").
					Return(func(string) (afero.File, error) {
						return nil, errors.New("callback error")
					})
			}),
			expectedError: "callback error",
		},
		{
			scenario: "callback success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Create", "test.txt").
					Return(func(string) (afero.File, error) {
						return f, nil
					})
			}),
			expectedResult: f,
		},
		{
			scenario: "callback for only error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Create", "test.txt").
					Return(nil, func(string) error {
						return errors.New("callback error")
					})
			}),
			expectedError: "callback error",
		},
		{
			scenario: "callback for only result",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Create", "test.txt").
					Return(func(string) afero.File {
						return f
					}, nil)
			}),
			expectedResult: f,
		},
		{
			scenario: "error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Create", "test.txt").
					Return(nil, errors.New("create error"))
			}),
			expectedError: "create error",
		},
		{
			scenario: "success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Create", "test.txt").
					Return(f, nil)
			}),
			expectedResult: f,
		},
	}

	for _, tc := range testCases {
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

func TestFs_Create_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFs(func(fs *aferomock.Fs) { //nolint: gosec
			fs.On("Create", mock.Anything)
		})(t).Create("") //nolint: errcheck
	})
}

func TestFs_Mkdir(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        aferomock.FsMocker
		expectedError string
	}{
		{
			scenario: "callback",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Mkdir", "test", os.ModePerm).
					Return(func(string, os.FileMode) error {
						return errors.New("callback error")
					})
			}),
			expectedError: "callback error",
		},
		{
			scenario: "error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Mkdir", "test", os.ModePerm).
					Return(errors.New("mkdir error"))
			}),
			expectedError: "mkdir error",
		},
		{
			scenario: "success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Mkdir", "test", os.ModePerm).
					Return(nil)
			}),
		},
	}

	for _, tc := range testCases {
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

func TestFs_Mkdir_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFs(func(fs *aferomock.Fs) { //nolint: gosec
			fs.On("Mkdir", mock.Anything, mock.Anything)
		})(t).Mkdir("", 0) //nolint: errcheck
	})
}

func TestFs_MkdirAll(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        aferomock.FsMocker
		expectedError string
	}{
		{
			scenario: "callback",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("MkdirAll", "path/test", os.ModePerm).
					Return(func(string, os.FileMode) error {
						return errors.New("callback error")
					})
			}),
			expectedError: "callback error",
		},
		{
			scenario: "error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("MkdirAll", "path/test", os.ModePerm).
					Return(errors.New("mkdir all error"))
			}),
			expectedError: "mkdir all error",
		},
		{
			scenario: "success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("MkdirAll", "path/test", os.ModePerm).
					Return(nil)
			}),
		},
	}

	for _, tc := range testCases {
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

func TestFs_MkdirAll_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFs(func(fs *aferomock.Fs) { //nolint: gosec
			fs.On("MkdirAll", mock.Anything, mock.Anything)
		})(t).MkdirAll("", 0) //nolint: errcheck
	})
}

func TestFs_Open(t *testing.T) {
	t.Parallel()

	f := aferomock.NopFile(t)

	testCases := []struct {
		scenario       string
		mockFs         aferomock.FsMocker
		expectedResult afero.File
		expectedError  string
	}{
		{
			scenario: "callback error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Open", "test.txt").
					Return(func(string) (afero.File, error) {
						return nil, errors.New("callback error")
					})
			}),
			expectedError: "callback error",
		},
		{
			scenario: "callback success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Open", "test.txt").
					Return(func(string) (afero.File, error) {
						return f, nil
					})
			}),
			expectedResult: f,
		},
		{
			scenario: "callback for only error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Open", "test.txt").
					Return(nil, func(string) error {
						return errors.New("callback error")
					})
			}),
			expectedError: "callback error",
		},
		{
			scenario: "callback for only result",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Open", "test.txt").
					Return(func(string) afero.File {
						return f
					}, nil)
			}),
			expectedResult: f,
		},
		{
			scenario: "error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Open", "test.txt").
					Return(nil, errors.New("create error"))
			}),
			expectedError: "create error",
		},
		{
			scenario: "success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Open", "test.txt").
					Return(f, nil)
			}),
			expectedResult: f,
		},
	}

	for _, tc := range testCases {
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

func TestFs_Open_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFs(func(fs *aferomock.Fs) { //nolint: gosec
			fs.On("Open", mock.Anything)
		})(t).Open("") //nolint: errcheck
	})
}

func TestFs_OpenFile(t *testing.T) {
	t.Parallel()

	f := aferomock.NopFile(t)

	testCases := []struct {
		scenario       string
		mockFs         aferomock.FsMocker
		expectedResult afero.File
		expectedError  string
	}{
		{
			scenario: "callback error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("OpenFile", "test.txt", 0, os.ModePerm).
					Return(func(string, int, os.FileMode) (afero.File, error) {
						return nil, errors.New("callback error")
					})
			}),
			expectedError: "callback error",
		},
		{
			scenario: "callback success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("OpenFile", "test.txt", 0, os.ModePerm).
					Return(func(string, int, os.FileMode) (afero.File, error) {
						return f, nil
					})
			}),
			expectedResult: f,
		},
		{
			scenario: "callback for only error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("OpenFile", "test.txt", 0, os.ModePerm).
					Return(nil, func(string, int, os.FileMode) error {
						return errors.New("callback error")
					})
			}),
			expectedError: "callback error",
		},
		{
			scenario: "callback for only result",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("OpenFile", "test.txt", 0, os.ModePerm).
					Return(func(string, int, os.FileMode) afero.File {
						return f
					}, nil)
			}),
			expectedResult: f,
		},
		{
			scenario: "error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("OpenFile", "test.txt", 0, os.ModePerm).
					Return(nil, errors.New("open file error"))
			}),
			expectedError: "open file error",
		},
		{
			scenario: "success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("OpenFile", "test.txt", 0, os.ModePerm).
					Return(f, nil)
			}),
			expectedResult: f,
		},
	}

	for _, tc := range testCases {
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

func TestFs_OpenFile_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFs(func(fs *aferomock.Fs) { //nolint: gosec
			fs.On("OpenFile", mock.Anything, mock.Anything, mock.Anything)
		})(t).OpenFile("", 0, 0) //nolint: errcheck
	})
}

func TestFs_Remove(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        aferomock.FsMocker
		expectedError string
	}{
		{
			scenario: "callback",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Remove", "test.txt").
					Return(func(string) error {
						return errors.New("callback error")
					})
			}),
			expectedError: "callback error",
		},
		{
			scenario: "error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Remove", "test.txt").
					Return(errors.New("remove error"))
			}),
			expectedError: "remove error",
		},
		{
			scenario: "success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Remove", "test.txt").
					Return(nil)
			}),
		},
	}

	for _, tc := range testCases {
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

func TestFs_Remove_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFs(func(fs *aferomock.Fs) { //nolint: gosec
			fs.On("Remove", mock.Anything)
		})(t).Remove("") //nolint: errcheck
	})
}

func TestFs_RemoveAll(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        aferomock.FsMocker
		expectedError string
	}{
		{
			scenario: "callback",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("RemoveAll", "path/test").
					Return(func(string) error {
						return errors.New("callback error")
					})
			}),
			expectedError: "callback error",
		},
		{
			scenario: "error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("RemoveAll", "path/test").
					Return(errors.New("remove all error"))
			}),
			expectedError: "remove all error",
		},
		{
			scenario: "success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("RemoveAll", "path/test").
					Return(nil)
			}),
		},
	}

	for _, tc := range testCases {
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

func TestFs_RemoveAll_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFs(func(fs *aferomock.Fs) { //nolint: gosec
			fs.On("RemoveAll", mock.Anything)
		})(t).RemoveAll("") //nolint: errcheck
	})
}

func TestFs_Rename(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        aferomock.FsMocker
		expectedError string
	}{
		{
			scenario: "callback",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Rename", "oldname", "newname").
					Return(func(string, string) error {
						return errors.New("callback error")
					})
			}),
			expectedError: "callback error",
		},
		{
			scenario: "error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Rename", "oldname", "newname").
					Return(errors.New("rename error"))
			}),
			expectedError: "rename error",
		},
		{
			scenario: "success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Rename", "oldname", "newname").
					Return(nil)
			}),
		},
	}

	for _, tc := range testCases {
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

func TestFs_Rename_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFs(func(fs *aferomock.Fs) { //nolint: gosec
			fs.On("Rename", mock.Anything, mock.Anything)
		})(t).Rename("", "") //nolint: errcheck
	})
}

func TestFs_Stat(t *testing.T) {
	t.Parallel()

	fi := aferomock.NopFileInfo(t)

	testCases := []struct {
		scenario       string
		mockFs         aferomock.FsMocker
		expectedResult os.FileInfo
		expectedError  string
	}{
		{
			scenario: "callback error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Stat", "test.txt").
					Return(func(string) (os.FileInfo, error) {
						return nil, errors.New("callback error")
					})
			}),
			expectedError: "callback error",
		},
		{
			scenario: "callback success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Stat", "test.txt").
					Return(func(string) (os.FileInfo, error) {
						return fi, nil
					})
			}),
			expectedResult: fi,
		},
		{
			scenario: "callback for only error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Stat", "test.txt").
					Return(nil, func(string) error {
						return errors.New("callback error")
					})
			}),
			expectedError: "callback error",
		},
		{
			scenario: "callback for only result",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Stat", "test.txt").
					Return(func(string) os.FileInfo {
						return fi
					}, nil)
			}),
			expectedResult: fi,
		},
		{
			scenario: "error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Stat", "test.txt").
					Return(nil, errors.New("stat error"))
			}),
			expectedError: "stat error",
		},
		{
			scenario: "success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Stat", "test.txt").
					Return(fi, nil)
			}),
			expectedResult: fi,
		},
	}

	for _, tc := range testCases {
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

func TestFs_Stat_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFs(func(fs *aferomock.Fs) { //nolint: gosec
			fs.On("Stat", mock.Anything)
		})(t).Stat("") //nolint: errcheck
	})
}

func TestFs_Name(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "aferomock.Fs", aferomock.NopFs(t).Name())
}

func TestFs_Name_Callback(t *testing.T) {
	t.Parallel()

	fs := aferomock.NewFs(t)

	fs.On("Name").
		Return(func() string {
			return "callback"
		})

	assert.Equal(t, "callback", fs.Name())
}

func TestFs_Name_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		fs := aferomock.NewFs(t) //nolint: gosec

		fs.On("Name") //nolint: errcheck

		fs.Name()
	})
}

func TestFs_Chmod(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        aferomock.FsMocker
		expectedError string
	}{
		{
			scenario: "callback",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Chmod", "test.txt", os.ModePerm).
					Return(func(string, os.FileMode) error {
						return errors.New("callback error")
					})
			}),
			expectedError: "callback error",
		},
		{
			scenario: "error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Chmod", "test.txt", os.ModePerm).
					Return(errors.New("chmod error"))
			}),
			expectedError: "chmod error",
		},
		{
			scenario: "success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Chmod", "test.txt", os.ModePerm).
					Return(nil)
			}),
		},
	}

	for _, tc := range testCases {
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

func TestFs_Chmod_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFs(func(fs *aferomock.Fs) { //nolint: gosec
			fs.On("Chmod", mock.Anything, mock.Anything)
		})(t).Chmod("", 0) //nolint: errcheck
	})
}

func TestFs_Chown(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        aferomock.FsMocker
		expectedError string
	}{
		{
			scenario: "callback",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Chown", "test.txt", 501, 501).
					Return(func(string, int, int) error {
						return errors.New("callback error")
					})
			}),
			expectedError: "callback error",
		},
		{
			scenario: "error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Chown", "test.txt", 501, 501).
					Return(errors.New("chown error"))
			}),
			expectedError: "chown error",
		},
		{
			scenario: "success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Chown", "test.txt", 501, 501).
					Return(nil)
			}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := tc.mockFs(t)
			err := fs.Chown("test.txt", 501, 501)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestFs_Chown_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFs(func(fs *aferomock.Fs) { //nolint: gosec
			fs.On("Chown", mock.Anything, mock.Anything, mock.Anything)
		})(t).Chown("", 0, 0) //nolint: errcheck
	})
}

func TestFs_Chtimes(t *testing.T) {
	t.Parallel()

	ts := time.Now()

	testCases := []struct {
		scenario      string
		mockFs        aferomock.FsMocker
		expectedError string
	}{
		{
			scenario: "callback",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Chtimes", "test.txt", ts, ts).
					Return(func(string, time.Time, time.Time) error {
						return errors.New("callback error")
					})
			}),
			expectedError: "callback error",
		},
		{
			scenario: "error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Chtimes", "test.txt", ts, ts).
					Return(errors.New("chtimes error"))
			}),
			expectedError: "chtimes error",
		},
		{
			scenario: "success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Chtimes", "test.txt", ts, ts).
					Return(nil)
			}),
		},
	}

	for _, tc := range testCases {
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

func TestFs_Chtimes_NoReturnValuePanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		aferomock.MockFs(func(fs *aferomock.Fs) { //nolint: gosec
			fs.On("Chtimes", mock.Anything, mock.Anything, mock.Anything)
		})(t).Chtimes("", time.Time{}, time.Time{}) //nolint: errcheck
	})
}
