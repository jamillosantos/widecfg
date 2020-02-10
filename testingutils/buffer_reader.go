package testingutils

import "bytes"

type BuffReader struct {
	Buff *bytes.Buffer
}

func NewBufferReader() *BuffReader {
	buff := &BuffReader{}
	buff.Open()
	return buff
}

func (reader *BuffReader) Open() error {
	reader.Buff = bytes.NewBuffer(nil)
	return nil
}

func (reader *BuffReader) Read(p []byte) (n int, err error) {
	return reader.Buff.Read(p)
}

func (reader *BuffReader) Close() error {
	reader.Buff = nil
	return nil
}

func (reader *BuffReader) Write(p []byte) (int, error) {
	return reader.Buff.Write(p)
}
