package dhtms

import "github.com/andrerrcosta2/vstr/pb"

type Nms struct {
	Id   []byte
	Type string
	Data []byte
}

func NewNms(id []byte, typ string, data []byte) *Nms {
	return &Nms{
		Id:   id,
		Type: typ,
		Data: data,
	}
}

func (n *Nms) Pbf() *pb.Nms {
	return &pb.Nms{
		Id:   n.Id,
		Type: n.Type,
		Data: n.Data,
	}
}
