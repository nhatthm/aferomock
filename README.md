# Mock for `spf13/afero`

[![GitHub Releases](https://img.shields.io/github/v/release/nhatthm/aferomock)](https://github.com/nhatthm/aferomock/releases/latest)
[![Build Status](https://github.com/nhatthm/aferomock/actions/workflows/test.yaml/badge.svg)](https://github.com/nhatthm/aferomock/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/nhatthm/aferomock/branch/master/graph/badge.svg?token=eTdAgDE2vR)](https://codecov.io/gh/nhatthm/aferomock)
[![Go Report Card](https://goreportcard.com/badge/github.com/nhatthm/aferomock)](https://goreportcard.com/report/github.com/nhatthm/aferomock)
[![GoDevDoc](https://img.shields.io/badge/dev-doc-00ADD8?logo=go)](https://pkg.go.dev/github.com/nhatthm/aferomock)
[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/donate/?hosted_button_id=PJZSGJN57TDJY)

**aferomock** is a mock library for [spf13/afero](https://github.com/spf13/afero)

## Prerequisites

- `Go >= 1.16`

## Install

```bash
go get github.com/nhatthm/aferomock
```

## Examples

```go
package mypackage_test

import (
	"errors"
	"os"
	"testing"

	"github.com/nhatthm/aferomock"
	"github.com/stretchr/testify/assert"
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
		tc := tc
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

If this project help you reduce time to develop, you can give me a cup of coffee :)

### Paypal donation

[![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/donate/?hosted_button_id=PJZSGJN57TDJY)

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;or scan this

<img src="https://user-images.githubusercontent.com/1154587/113494222-ad8cb200-94e6-11eb-9ef3-eb883ada222a.png" width="147px" />
