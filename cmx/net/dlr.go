package net

import "net"

type Dlr interface {
	Dial(address string) (net.Conn, error)
}
