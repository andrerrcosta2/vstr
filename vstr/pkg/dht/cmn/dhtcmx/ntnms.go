package dhtcmx

import (
	"github.com/andrerrcosta2/vstr/vstr/pkg/dht/cmn/dhtms"
	"github.com/andrerrcosta2/vstr/vstr/pkg/ntk/dlr"
	"github.com/andrerrcosta2/vstr/vstr/pkg/pb"
	"google.golang.org/protobuf/proto"
)

type ntns struct {
	dlr dlr.Dlr
}

func (s *ntns) Sm(n *dhtms.Nms, addr string) (*dhtms.Nms, error) {
	conn, err := s.dlr.Dl(addr)
	if err != nil {
		return nil, err
	}

	defer dlr.Cls(conn)

	buf, err := proto.Marshal(&pb.Nms{
		Type: n.Type,
		Id:   n.Id,
		Data: n.Data,
	})

	if err != nil {
		return nil, err
	}

	if _, err = conn.Write(buf); err != nil {
		return nil, err
	}

	rbf := make([]byte, 1024)
	rd, err := conn.Read(rbf)
	if err != nil {
		return nil, err
	}

	res := &pb.Nms{}
	if err := proto.Unmarshal(rbf[:rd], res); err != nil {
		return nil, err
	}

	return &dhtms.Nms{
		Id:   res.Id,
		Type: res.Type,
		Data: res.Data,
	}, nil
}

func NewNtns(dlr dlr.Dlr) *ntns {
	return &ntns{
		dlr: dlr,
	}
}
