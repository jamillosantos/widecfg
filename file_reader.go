package widecfg

import (
	"errors"
	"os"
)

var (
	ErrNotOpen = errors.New("not open")
)

type FileReader struct {
	fName string
	file  *os.File
}

func NewFileReader(fname string) Reader {
	return &FileReader{
		fName: fname,
	}
}

func (reader *FileReader) Open() error {
	f, err := os.Open(reader.fName)
	if err != nil {
		return err
	}
	reader.file = f
	return nil
}

func (reader *FileReader) Read(p []byte) (n int, err error) {
	if reader.file == nil {
		return 0, ErrNotOpen
	}
	return reader.file.Read(p)
}

func (reader *FileReader) Close() error {
	return reader.file.Close()
}
