// Andre R.R. Costa *** github.com/andrerrcosta2 *** andrerrcosta@gmail.com

package cmx

import (
	"fmt"
	"github.com/andrerrcosta2/vstr/cmx/net"
	"google.golang.org/protobuf/proto"
)

type Dcpms struct {
	dlr net.Dlr
}

func NewDcpmx(dlr net.Dlr) *Dcpms {
	return &Dcpms{
		dlr: dlr,
	}
}

func (t *Dcpms) Spm(msg proto.Message, addr string) (proto.Message, error) {
	if t.dlr == nil {
		return nil, fmt.Errorf("dialer is nil\n")
	}
	if msg == nil {
		return nil, fmt.Errorf("message is nil\n")
	}
	if addr == "" {
		return nil, fmt.Errorf("address is empty\n")
	}

	conn, err := t.dlr.Dial(addr)
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

	res := proto.Clone(msg)
	if err := proto.Unmarshal(rbf[:rd], res); err != nil {
		return nil, err
	}

	return res, nil
}
