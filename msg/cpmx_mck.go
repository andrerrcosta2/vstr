// github.com/andrerrcosta2 *** andrerrcosta@gmail.com

package msg

import (
	"github.com/stretchr/testify/mock"
	"net"
	"time"
)

type mdlr struct {
	mock.Mock
	Conn net.Conn
	Err  error
}

func (d *mdlr) Dial(addr string) (net.Conn, error) {
	return d.Conn, d.Err
}

type mcon struct {
	mock.Mock
	rd []byte
	wd []byte
}

func (c mcon) Read(b []byte) (n int, err error) {
	copy(b, c.rd)
	return len(c.rd), nil
}

func (c mcon) Write(b []byte) (n int, err error) {
	c.wd = b
	return len(b), nil
}

func (c mcon) Close() error {
	return nil
}

func (c mcon) LocalAddr() net.Addr                { return nil }
func (c mcon) RemoteAddr() net.Addr               { return nil }
func (c mcon) SetDeadline(t time.Time) error      { return nil }
func (c mcon) SetReadDeadline(t time.Time) error  { return nil }
func (c mcon) SetWriteDeadline(t time.Time) error { return nil }
