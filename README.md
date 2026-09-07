# Mock for `spf13/afero`

[![GitHub Releases](https://img.shields.io/github/v/release/nhatthm/aferomock)](https://github.com/nhatthm/aferomock/releases/latest)
[![Build Status](https://github.com/nhatthm/aferomock/actions/workflows/test.yaml/badge.svg)](https://github.com/nhatthm/aferomock/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/nhatthm/aferomock/branch/master/graph/badge.svg?token=eTdAgDE2vR)](https://codecov.io/gh/nhatthm/aferomock)
[![GoDevDoc](https://img.shields.io/badge/dev-doc-00ADD8?logo=go)](https://pkg.go.dev/go.nhat.io/aferomock)
[![Donate](https://img.shields.io/badge/%20-Donate-%20?style=flat&logo=githubsponsors&color=E5E4E2)](http://donate.nhat.me)

**aferomock** is a mock library for [spf13/afero](https://github.com/spf13/afero)

## Prerequisites

- `Go >= 1.25`

## Install

```bash
go get go.nhat.io/aferomock
```

## Examples

```go
package mypackage_test

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.nhat.io/aferomock"
)

func TestMyPackage(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mockFs        aferomock.FsMocker
		expectedError string
	}{
		{
			scenario: "no error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.MkdirAll("highway/to/hell", os.ModePerm).Return(nil)
			}),
		},
		{
			scenario: "error",
			mockFs: aferomock.MockFs(func(fs *aferomock.Fs) {
				fs.MkdirAll("highway/to/hell", os.ModePerm).Return(errors.New("mkdir error"))
			}),
			expectedError: "mkdir error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			err := tc.mockFs(t).MkdirAll("highway/to/hell")

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}
```

## Donation

If this project saved you some development time, buy me a cup of coffee :)

[![donate](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](http://donate.nhat.me)

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;or scan this

<img src="https://github.com/nhatthm/donate.nhat.me/blob/master/images/qr_sponsor.png" width="147px" />

[<sub><sup>[table of contents]</sup></sub>](#table-of-contents)
