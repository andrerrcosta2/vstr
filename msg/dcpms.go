// Andre R.R. Costa *** github.com/andrerrcosta2 *** andrerrcosta@gmail.com

package msg

import (
	"fmt"
	"github.com/andrerrcosta2/vstr/nwk/dlr"
	"google.golang.org/protobuf/proto"
)

type dcpmx struct {
	dlr dlr.Dlr
}

func NewDcpmx(dlr dlr.Dlr) *dcpmx {
	return &dcpmx{
		dlr: dlr,
	}
}

func (t *dcpmx) Spm(msg proto.Message, addr string) (proto.Message, error) {
	if t.dlr == nil {
		return nil, fmt.Errorf("dlr is nil\n")
	}
	if msg == nil {
		return nil, fmt.Errorf("msg is nil\n")
	}
	if addr == "" {
		return nil, fmt.Errorf("addr is empty\n")
	}

	conn, err := t.dlr.Dl(addr)
	if err != nil {
		return nil, err
	}

	defer dlr.Cls(conn)

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
