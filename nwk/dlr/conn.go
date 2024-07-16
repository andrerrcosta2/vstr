package dlr

import (
	"io"
	"net"
)

type Conn interface {
	io.Closer
	io.Writer
	io.Reader
}

var _ Conn = (net.Conn)(nil)
