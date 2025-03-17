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

func TestFsCallbacks_Create(t *testing.T) {
	t.Parallel()

	f := aferomock.NopFile(t)

	testCases := []struct {
		scenario       string
		mockFs         aferomock.FsMocker
		fsCallbacks    aferomock.FsCallbacks
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
			scenario: "overridden - error",
			mockFs:   aferomock.NopFs,
			fsCallbacks: aferomock.FsCallbacks{
				CreateFunc: func(string) (afero.File, error) {
					return nil, errors.New("create error")
				},
			},
			expectedError: "create error",
		},
		{
			scenario: "overridden - success",
			mockFs:   aferomock.NopFs,
			fsCallbacks: aferomock.FsCallbacks{
				CreateFunc: func(string) (afero.File, error) {
					return f, nil
				},
			},
			expectedResult: f,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := aferomock.OverrideFs(tc.mockFs(t), tc.fsCallbacks)
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

func TestFsCallbacks_Mkdir(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        aferomock.FsMocker
		fsCallbacks   aferomock.FsCallbacks
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
			scenario: "overridden - error",
			mockFs:   aferomock.NopFs,
			fsCallbacks: aferomock.FsCallbacks{
				MkdirFunc: func(string, os.FileMode) error {
					return errors.New("mkdir error")
				},
			},
			expectedError: "mkdir error",
		},
		{
			scenario: "overridden - success",
			mockFs:   aferomock.NopFs,
			fsCallbacks: aferomock.FsCallbacks{
				MkdirFunc: func(string, os.FileMode) error {
					return nil
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := aferomock.OverrideFs(tc.mockFs(t), tc.fsCallbacks)
			err := fs.Mkdir("test", os.ModePerm)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestFsCallbacks_MkdirAll(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        aferomock.FsMocker
		fsCallbacks   aferomock.FsCallbacks
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
			scenario: "overridden - error",
			mockFs:   aferomock.NopFs,
			fsCallbacks: aferomock.FsCallbacks{
				MkdirAllFunc: func(string, os.FileMode) error {
					return errors.New("mkdir all error")
				},
			},
			expectedError: "mkdir all error",
		},
		{
			scenario: "overridden - success",
			mockFs:   aferomock.NopFs,
			fsCallbacks: aferomock.FsCallbacks{
				MkdirAllFunc: func(string, os.FileMode) error {
					return nil
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := aferomock.OverrideFs(tc.mockFs(t), tc.fsCallbacks)
			err := fs.MkdirAll("path/test", os.ModePerm)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestFsCallbacks_Open(t *testing.T) {
	t.Parallel()

	f := aferomock.NopFile(t)

	testCases := []struct {
		scenario       string
		mockFs         aferomock.FsMocker
		fsCallbacks    aferomock.FsCallbacks
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
			scenario: "overridden - error",
			mockFs:   aferomock.NopFs,
			fsCallbacks: aferomock.FsCallbacks{
				OpenFunc: func(string) (afero.File, error) {
					return nil, errors.New("create error")
				},
			},
			expectedError: "create error",
		},
		{
			scenario: "overridden - success",
			mockFs:   aferomock.NopFs,
			fsCallbacks: aferomock.FsCallbacks{
				OpenFunc: func(string) (afero.File, error) {
					return f, nil
				},
			},
			expectedResult: f,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := aferomock.OverrideFs(tc.mockFs(t), tc.fsCallbacks)
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

func TestFsCallbacks_OpenFile(t *testing.T) {
	t.Parallel()

	f := aferomock.NopFile(t)

	testCases := []struct {
		scenario       string
		mockFs         aferomock.FsMocker
		fsCallbacks    aferomock.FsCallbacks
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
			scenario: "overridden - error",
			mockFs:   aferomock.NopFs,
			fsCallbacks: aferomock.FsCallbacks{
				OpenFileFunc: func(string, int, os.FileMode) (afero.File, error) {
					return nil, errors.New("open file error")
				},
			},
			expectedError: "open file error",
		},
		{
			scenario: "overridden - success",
			mockFs:   aferomock.NopFs,
			fsCallbacks: aferomock.FsCallbacks{
				OpenFileFunc: func(string, int, os.FileMode) (afero.File, error) {
					return f, nil
				},
			},
			expectedResult: f,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := aferomock.OverrideFs(tc.mockFs(t), tc.fsCallbacks)
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

func TestFsCallbacks_Remove(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        aferomock.FsMocker
		fsCallbacks   aferomock.FsCallbacks
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
			scenario: "overridden - error",
			mockFs:   aferomock.NopFs,
			fsCallbacks: aferomock.FsCallbacks{
				RemoveFunc: func(string) error {
					return errors.New("remove error")
				},
			},
			expectedError: "remove error",
		},
		{
			scenario: "overridden - success",
			mockFs:   aferomock.NopFs,
			fsCallbacks: aferomock.FsCallbacks{
				RemoveFunc: func(string) error {
					return nil
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := aferomock.OverrideFs(tc.mockFs(t), tc.fsCallbacks)
			err := fs.Remove("test.txt")

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestFsCallbacks_RemoveAll(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        aferomock.FsMocker
		fsCallbacks   aferomock.FsCallbacks
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
			scenario: "overridden - error",
			mockFs:   aferomock.NopFs,
			fsCallbacks: aferomock.FsCallbacks{
				RemoveAllFunc: func(string) error {
					return errors.New("remove all error")
				},
			},
			expectedError: "remove all error",
		},
		{
			scenario: "overridden - success",
			mockFs:   aferomock.NopFs,
			fsCallbacks: aferomock.FsCallbacks{
				RemoveAllFunc: func(string) error {
					return nil
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := aferomock.OverrideFs(tc.mockFs(t), tc.fsCallbacks)
			err := fs.RemoveAll("path/test")

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestFsCallbacks_Rename(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        aferomock.FsMocker
		fsCallbacks   aferomock.FsCallbacks
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
			scenario: "overridden - error",
			mockFs:   aferomock.NopFs,
			fsCallbacks: aferomock.FsCallbacks{
				RenameFunc: func(string, string) error {
					return errors.New("rename error")
				},
			},
			expectedError: "rename error",
		},
		{
			scenario: "overridden - success",
			mockFs:   aferomock.NopFs,
			fsCallbacks: aferomock.FsCallbacks{
				RenameFunc: func(string, string) error {
					return nil
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := aferomock.OverrideFs(tc.mockFs(t), tc.fsCallbacks)
			err := fs.Rename("oldname", "newname")

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestFsCallbacks_Stat(t *testing.T) {
	t.Parallel()

	fi := aferomock.NopFileInfo(t)

	testCases := []struct {
		scenario       string
		mockFs         aferomock.FsMocker
		fsCallbacks    aferomock.FsCallbacks
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
			scenario: "overridden - error",
			mockFs:   aferomock.NopFs,
			fsCallbacks: aferomock.FsCallbacks{
				StatFunc: func(string) (os.FileInfo, error) {
					return nil, errors.New("stat error")
				},
			},
			expectedError: "stat error",
		},
		{
			scenario: "overridden - success",
			mockFs:   aferomock.NopFs,
			fsCallbacks: aferomock.FsCallbacks{
				StatFunc: func(string) (os.FileInfo, error) {
					return fi, nil
				},
			},
			expectedResult: fi,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := aferomock.OverrideFs(tc.mockFs(t), tc.fsCallbacks)
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

func TestFsCallbacks_Name(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		fsCallbacks    aferomock.FsCallbacks
		expectedResult string
	}{
		{
			scenario:       "upstream",
			expectedResult: "aferomock.Fs",
		},
		{
			scenario: "overridden",
			fsCallbacks: aferomock.FsCallbacks{
				NameFunc: func() string {
					return "overridden"
				},
			},
			expectedResult: "overridden",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := aferomock.OverrideFs(aferomock.NopFs(t), tc.fsCallbacks)

			actual := fs.Name()

			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestFsCallbacks_Chmod(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        aferomock.FsMocker
		fsCallbacks   aferomock.FsCallbacks
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
			scenario: "overridden - error",
			mockFs:   aferomock.NopFs,
			fsCallbacks: aferomock.FsCallbacks{
				ChmodFunc: func(string, os.FileMode) error {
					return errors.New("chmod error")
				},
			},
			expectedError: "chmod error",
		},
		{
			scenario: "overridden - success",
			mockFs:   aferomock.NopFs,
			fsCallbacks: aferomock.FsCallbacks{
				ChmodFunc: func(string, os.FileMode) error {
					return nil
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := aferomock.OverrideFs(tc.mockFs(t), tc.fsCallbacks)
			err := fs.Chmod("test.txt", os.ModePerm)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestFsCallbacks_Chown(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        aferomock.FsMocker
		fsCallbacks   aferomock.FsCallbacks
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
			scenario: "overridden - error",
			mockFs:   aferomock.NopFs,
			fsCallbacks: aferomock.FsCallbacks{
				ChownFunc: func(string, int, int) error {
					return errors.New("chown error")
				},
			},
			expectedError: "chown error",
		},
		{
			scenario: "overridden - success",
			mockFs:   aferomock.NopFs,
			fsCallbacks: aferomock.FsCallbacks{
				ChownFunc: func(string, int, int) error {
					return nil
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := aferomock.OverrideFs(tc.mockFs(t), tc.fsCallbacks)
			err := fs.Chown("test.txt", 501, 501)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestFsCallbacks_Chtimes(t *testing.T) {
	t.Parallel()

	ts := time.Now()

	testCases := []struct {
		scenario      string
		mockFs        aferomock.FsMocker
		fsCallbacks   aferomock.FsCallbacks
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
			scenario: "overridden - error",
			mockFs:   aferomock.NopFs,
			fsCallbacks: aferomock.FsCallbacks{
				ChtimesFunc: func(string, time.Time, time.Time) error {
					return errors.New("chtimes error")
				},
			},
			expectedError: "chtimes error",
		},
		{
			scenario: "overridden - success",
			mockFs:   aferomock.NopFs,
			fsCallbacks: aferomock.FsCallbacks{
				ChtimesFunc: func(string, time.Time, time.Time) error {
					return nil
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			fs := aferomock.OverrideFs(tc.mockFs(t), tc.fsCallbacks)
			err := fs.Chtimes("test.txt", ts, ts)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestWrapFs(t *testing.T) {
	t.Parallel()

	fs := aferomock.WrapFs(aferomock.NopFs(t), aferomock.WrappedFs{})

	assert.NotNil(t, fs)
	assert.IsType(t, aferomock.FsCallbacks{}, fs)
}
