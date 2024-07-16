package udp

import (
	"github.com/andrerrcosta2/vstr/nwk/dlr"
	"github.com/andrerrcosta2/vstr/nwk/nwkc"
	"net"
	"time"
)

type UdpDlr struct{}

func (d *UdpDlr) Dl(addr string) (dlr.Conn, error) {
	return net.Dial("net", addr)
}

func (d *UdpDlr) Dlt(addr string, tout time.Duration) (dlr.Conn, error) {
	return net.DialTimeout("udp", addr, tout)
}

func (d *UdpDlr) ChkNwk(tout time.Duration) bool {
	_, err := net.DialTimeout("udp", nwkc.RFA, tout)
	return err == nil
}
