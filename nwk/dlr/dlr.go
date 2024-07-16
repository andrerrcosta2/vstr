package dlr

import (
	"log"
	"time"
)

type Dlr interface {
	Dl(addr string) (Conn, error)
	Dlt(addr string, tout time.Duration) (Conn, error)
	ChkNwk(tout time.Duration) bool
}

func Cls(conn Conn) {
	if err := conn.Close(); err != nil {
		log.Printf("Error closing connection: %v", err)
	}
}
