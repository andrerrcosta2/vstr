package tcp

import (
	"github.com/andrerrcosta2/vstr/vstr/pkg/ntk/dlr"
	"net"
	"time"
)

type TcpDlr struct{}

func (d *TcpDlr) Dl(addr string) (dlr.Conn, error) {
	return net.Dial("net", addr)
}

func (d *TcpDlr) Dlt(addr string, tout time.Duration) (dlr.Conn, error) {
	return net.DialTimeout("net", addr, tout)
}

func (d *TcpDlr) ChkNwk(tout time.Duration) bool {
	conn, err := net.DialTimeout("tcp", "example.com:80", tout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func NewDlr() *TcpDlr {
	return &TcpDlr{}
}
