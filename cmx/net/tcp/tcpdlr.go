package tcp

import "net"

type TcpDlr struct{}

func (d *TcpDlr) Dial(addr string) (net.Conn, error) {
	return net.Dial("tcp", addr)
}
