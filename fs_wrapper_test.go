package aferomock_test

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"

	"go.nhat.io/aferomock"
)

func TestWrappedFs_Create(t *testing.T) {
	t.Parallel()

	f := aferomock.NopFile(t)

	testCases := []struct {
		scenario       string
		mockFs         aferomock.FsMocker
		wrappedFs      aferomock.WrappedFs
		expectedResult afero.File
		expectedError  string
	}{
		{
			scenario: "upstream - error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Create", "test.txt").
					Return(nil, errors.New("create error"))
			}),
			expectedError: "create error",
		},
		{
			scenario: "upstream - success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Create", "test.txt").
					Return(f, nil)
			}),
			expectedResult: f,
		},
		{
			scenario: "wrapped - error",
			mockFs:   aferomock.NopFs,
			wrappedFs: aferomock.WrappedFs{
				CreateFunc: func(string) (afero.File, error) {
					return nil, errors.New("create error")
				},
			},
			expectedError: "create error",
		},
		{
			scenario: "wrapped - success",
			mockFs:   aferomock.NopFs,
			wrappedFs: aferomock.WrappedFs{
				CreateFunc: func(string) (afero.File, error) {
					return f, nil
				},
			},
			expectedResult: f,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := aferomock.WrapFs(tc.mockFs(t), tc.wrappedFs)
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

func TestWrappedFs_Mkdir(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        aferomock.FsMocker
		wrappedFs     aferomock.WrappedFs
		expectedError string
	}{
		{
			scenario: "upstream - error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Mkdir", "test", os.ModePerm).
					Return(errors.New("mkdir error"))
			}),
			expectedError: "mkdir error",
		},
		{
			scenario: "upstream - success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Mkdir", "test", os.ModePerm).
					Return(nil)
			}),
		},
		{
			scenario: "wrapped - error",
			mockFs:   aferomock.NopFs,
			wrappedFs: aferomock.WrappedFs{
				MkdirFunc: func(string, os.FileMode) error {
					return errors.New("mkdir error")
				},
			},
			expectedError: "mkdir error",
		},
		{
			scenario: "wrapped - success",
			mockFs:   aferomock.NopFs,
			wrappedFs: aferomock.WrappedFs{
				MkdirFunc: func(string, os.FileMode) error {
					return nil
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := aferomock.WrapFs(tc.mockFs(t), tc.wrappedFs)
			err := fs.Mkdir("test", os.ModePerm)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestWrappedFs_MkdirAll(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        aferomock.FsMocker
		wrappedFs     aferomock.WrappedFs
		expectedError string
	}{
		{
			scenario: "upstream - error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("MkdirAll", "path/test", os.ModePerm).
					Return(errors.New("mkdir all error"))
			}),
			expectedError: "mkdir all error",
		},
		{
			scenario: "upstream - success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("MkdirAll", "path/test", os.ModePerm).
					Return(nil)
			}),
		},
		{
			scenario: "wrapped - error",
			mockFs:   aferomock.NopFs,
			wrappedFs: aferomock.WrappedFs{
				MkdirAllFunc: func(string, os.FileMode) error {
					return errors.New("mkdir all error")
				},
			},
			expectedError: "mkdir all error",
		},
		{
			scenario: "wrapped - success",
			mockFs:   aferomock.NopFs,
			wrappedFs: aferomock.WrappedFs{
				MkdirAllFunc: func(string, os.FileMode) error {
					return nil
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := aferomock.WrapFs(tc.mockFs(t), tc.wrappedFs)
			err := fs.MkdirAll("path/test", os.ModePerm)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestWrappedFs_Open(t *testing.T) {
	t.Parallel()

	f := aferomock.NopFile(t)

	testCases := []struct {
		scenario       string
		mockFs         aferomock.FsMocker
		wrappedFs      aferomock.WrappedFs
		expectedResult afero.File
		expectedError  string
	}{
		{
			scenario: "upstream - error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Open", "test.txt").
					Return(nil, errors.New("create error"))
			}),
			expectedError: "create error",
		},
		{
			scenario: "upstream - success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Open", "test.txt").
					Return(f, nil)
			}),
			expectedResult: f,
		},
		{
			scenario: "wrapped - error",
			mockFs:   aferomock.NopFs,
			wrappedFs: aferomock.WrappedFs{
				OpenFunc: func(string) (afero.File, error) {
					return nil, errors.New("create error")
				},
			},
			expectedError: "create error",
		},
		{
			scenario: "wrapped - success",
			mockFs:   aferomock.NopFs,
			wrappedFs: aferomock.WrappedFs{
				OpenFunc: func(string) (afero.File, error) {
					return f, nil
				},
			},
			expectedResult: f,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := aferomock.WrapFs(tc.mockFs(t), tc.wrappedFs)
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

func TestWrappedFs_OpenFile(t *testing.T) {
	t.Parallel()

	f := aferomock.NopFile(t)

	testCases := []struct {
		scenario       string
		mockFs         aferomock.FsMocker
		wrappedFs      aferomock.WrappedFs
		expectedResult afero.File
		expectedError  string
	}{
		{
			scenario: "upstream - error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("OpenFile", "test.txt", 0, os.ModePerm).
					Return(nil, errors.New("open file error"))
			}),
			expectedError: "open file error",
		},
		{
			scenario: "upstream - success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("OpenFile", "test.txt", 0, os.ModePerm).
					Return(f, nil)
			}),
			expectedResult: f,
		},
		{
			scenario: "wrapped - error",
			mockFs:   aferomock.NopFs,
			wrappedFs: aferomock.WrappedFs{
				OpenFileFunc: func(string, int, os.FileMode) (afero.File, error) {
					return nil, errors.New("open file error")
				},
			},
			expectedError: "open file error",
		},
		{
			scenario: "wrapped - success",
			mockFs:   aferomock.NopFs,
			wrappedFs: aferomock.WrappedFs{
				OpenFileFunc: func(string, int, os.FileMode) (afero.File, error) {
					return f, nil
				},
			},
			expectedResult: f,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := aferomock.WrapFs(tc.mockFs(t), tc.wrappedFs)
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

func TestWrappedFs_Remove(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        aferomock.FsMocker
		wrappedFs     aferomock.WrappedFs
		expectedError string
	}{
		{
			scenario: "upstream - error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Remove", "test.txt").
					Return(errors.New("remove error"))
			}),
			expectedError: "remove error",
		},
		{
			scenario: "upstream - success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Remove", "test.txt").
					Return(nil)
			}),
		},
		{
			scenario: "wrapped - error",
			mockFs:   aferomock.NopFs,
			wrappedFs: aferomock.WrappedFs{
				RemoveFunc: func(string) error {
					return errors.New("remove error")
				},
			},
			expectedError: "remove error",
		},
		{
			scenario: "wrapped - success",
			mockFs:   aferomock.NopFs,
			wrappedFs: aferomock.WrappedFs{
				RemoveFunc: func(string) error {
					return nil
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := aferomock.WrapFs(tc.mockFs(t), tc.wrappedFs)
			err := fs.Remove("test.txt")

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestWrappedFs_RemoveAll(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        aferomock.FsMocker
		wrappedFs     aferomock.WrappedFs
		expectedError string
	}{
		{
			scenario: "upstream - error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("RemoveAll", "path/test").
					Return(errors.New("remove all error"))
			}),
			expectedError: "remove all error",
		},
		{
			scenario: "upstream - success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("RemoveAll", "path/test").
					Return(nil)
			}),
		},
		{
			scenario: "wrapped - error",
			mockFs:   aferomock.NopFs,
			wrappedFs: aferomock.WrappedFs{
				RemoveAllFunc: func(string) error {
					return errors.New("remove all error")
				},
			},
			expectedError: "remove all error",
		},
		{
			scenario: "wrapped - success",
			mockFs:   aferomock.NopFs,
			wrappedFs: aferomock.WrappedFs{
				RemoveAllFunc: func(string) error {
					return nil
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := aferomock.WrapFs(tc.mockFs(t), tc.wrappedFs)
			err := fs.RemoveAll("path/test")

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestWrappedFs_Rename(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        aferomock.FsMocker
		wrappedFs     aferomock.WrappedFs
		expectedError string
	}{
		{
			scenario: "upstream - error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Rename", "oldname", "newname").
					Return(errors.New("rename error"))
			}),
			expectedError: "rename error",
		},
		{
			scenario: "upstream - success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Rename", "oldname", "newname").
					Return(nil)
			}),
		},
		{
			scenario: "wrapped - error",
			mockFs:   aferomock.NopFs,
			wrappedFs: aferomock.WrappedFs{
				RenameFunc: func(string, string) error {
					return errors.New("rename error")
				},
			},
			expectedError: "rename error",
		},
		{
			scenario: "wrapped - success",
			mockFs:   aferomock.NopFs,
			wrappedFs: aferomock.WrappedFs{
				RenameFunc: func(string, string) error {
					return nil
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := aferomock.WrapFs(tc.mockFs(t), tc.wrappedFs)
			err := fs.Rename("oldname", "newname")

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestWrappedFs_Stat(t *testing.T) {
	t.Parallel()

	fi := aferomock.NopFileInfo(t)

	testCases := []struct {
		scenario       string
		mockFs         aferomock.FsMocker
		wrappedFs      aferomock.WrappedFs
		expectedResult os.FileInfo
		expectedError  string
	}{
		{
			scenario: "upstream - error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Stat", "test.txt").
					Return(nil, errors.New("stat error"))
			}),
			expectedError: "stat error",
		},
		{
			scenario: "upstream - success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Stat", "test.txt").
					Return(fi, nil)
			}),
			expectedResult: fi,
		},
		{
			scenario: "wrapped - error",
			mockFs:   aferomock.NopFs,
			wrappedFs: aferomock.WrappedFs{
				StatFunc: func(string) (os.FileInfo, error) {
					return nil, errors.New("stat error")
				},
			},
			expectedError: "stat error",
		},
		{
			scenario: "wrapped - success",
			mockFs:   aferomock.NopFs,
			wrappedFs: aferomock.WrappedFs{
				StatFunc: func(string) (os.FileInfo, error) {
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

			fs := aferomock.WrapFs(tc.mockFs(t), tc.wrappedFs)
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

func TestWrappedFs_Name(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		wrappedFs      aferomock.WrappedFs
		expectedResult string
	}{
		{
			scenario:       "upstream",
			expectedResult: "aferomock.Fs",
		},
		{
			scenario: "wrapped",
			wrappedFs: aferomock.WrappedFs{
				NameFunc: func() string {
					return "wrapped"
				},
			},
			expectedResult: "wrapped",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := aferomock.WrapFs(aferomock.NopFs(t), tc.wrappedFs)

			actual := fs.Name()

			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestWrappedFs_Chmod(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        aferomock.FsMocker
		wrappedFs     aferomock.WrappedFs
		expectedError string
	}{
		{
			scenario: "upstream - error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Chmod", "test.txt", os.ModePerm).
					Return(errors.New("chmod error"))
			}),
			expectedError: "chmod error",
		},
		{
			scenario: "upstream - success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Chmod", "test.txt", os.ModePerm).
					Return(nil)
			}),
		},
		{
			scenario: "wrapped - error",
			mockFs:   aferomock.NopFs,
			wrappedFs: aferomock.WrappedFs{
				ChmodFunc: func(string, os.FileMode) error {
					return errors.New("chmod error")
				},
			},
			expectedError: "chmod error",
		},
		{
			scenario: "wrapped - success",
			mockFs:   aferomock.NopFs,
			wrappedFs: aferomock.WrappedFs{
				ChmodFunc: func(string, os.FileMode) error {
					return nil
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := aferomock.WrapFs(tc.mockFs(t), tc.wrappedFs)
			err := fs.Chmod("test.txt", os.ModePerm)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestWrappedFs_Chown(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        aferomock.FsMocker
		wrappedFs     aferomock.WrappedFs
		expectedError string
	}{
		{
			scenario: "upstream - error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Chown", "test.txt", 501, 501).
					Return(errors.New("chown error"))
			}),
			expectedError: "chown error",
		},
		{
			scenario: "upstream - success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Chown", "test.txt", 501, 501).
					Return(nil)
			}),
		},
		{
			scenario: "wrapped - error",
			mockFs:   aferomock.NopFs,
			wrappedFs: aferomock.WrappedFs{
				ChownFunc: func(string, int, int) error {
					return errors.New("chown error")
				},
			},
			expectedError: "chown error",
		},
		{
			scenario: "wrapped - success",
			mockFs:   aferomock.NopFs,
			wrappedFs: aferomock.WrappedFs{
				ChownFunc: func(string, int, int) error {
					return nil
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := aferomock.WrapFs(tc.mockFs(t), tc.wrappedFs)
			err := fs.Chown("test.txt", 501, 501)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestWrappedFs_Chtimes(t *testing.T) {
	t.Parallel()

	ts := time.Now()

	testCases := []struct {
		scenario      string
		mockFs        aferomock.FsMocker
		wrappedFs     aferomock.WrappedFs
		expectedError string
	}{
		{
			scenario: "upstream - error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Chtimes", "test.txt", ts, ts).
					Return(errors.New("chtimes error"))
			}),
			expectedError: "chtimes error",
		},
		{
			scenario: "upstream - success",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.On("Chtimes", "test.txt", ts, ts).
					Return(nil)
			}),
		},
		{
			scenario: "wrapped - error",
			mockFs:   aferomock.NopFs,
			wrappedFs: aferomock.WrappedFs{
				ChtimesFunc: func(string, time.Time, time.Time) error {
					return errors.New("chtimes error")
				},
			},
			expectedError: "chtimes error",
		},
		{
			scenario: "wrapped - success",
			mockFs:   aferomock.NopFs,
			wrappedFs: aferomock.WrappedFs{
				ChtimesFunc: func(string, time.Time, time.Time) error {
					return nil
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := aferomock.WrapFs(tc.mockFs(t), tc.wrappedFs)
			err := fs.Chtimes("test.txt", ts, ts)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}
