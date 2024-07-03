// Andre R. R. Costa *** github.com/andrerrcosta2
package tcp

import (
	"google.golang.org/protobuf/proto"
	"net"
	"sync"
)

type TcpCpmx struct {
	mt sync.Mutex
}

func Spm(msg proto.Message, addr string) (proto.Message, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	buf, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}

	if _, err := conn.Write(buf); err != nil {
		return nil, err
	}

	rbf := make([]byte, 1024)
	rd, err := conn.Read(rbf)
	if err != nil {
		return nil, err
	}

	var res proto.Message
	if err := proto.Unmarshal(rbf[:rd], res); err != nil {
		return nil, err
	}

	return res, nil
}
