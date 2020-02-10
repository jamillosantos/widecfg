package widecfg

import "io"

type Reader interface {
	Open() error
	io.Reader
	io.Closer
}
