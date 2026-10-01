package aferomock

func init() { //nolint: gochecknoinits
	SetFsDefaultMocks(func(fs *Fs) {
		fs.EXPECT().Name().Maybe().
			Return("aferomock.Fs")
	})
}
