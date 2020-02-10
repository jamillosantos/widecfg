package testingutils

type ErrorReader struct {
	OpenErr  error
	ReadErr  error
	CloseErr error
}

func (reader *ErrorReader) Open() error {
	return reader.OpenErr
}

func (reader *ErrorReader) Read(p []byte) (n int, err error) {
	return 0, reader.ReadErr
}

func (reader *ErrorReader) Close() error {
	return reader.CloseErr
}
