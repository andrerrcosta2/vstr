package nod

import (
	"github.com/andrerrcosta2/vstr/dht/crd/cid"
	"github.com/andrerrcosta2/vstr/pb"
)

type Fge struct {
	Strt cid.Id
	Nod  *Nod
}

func NewFge(str cid.Id, nod *Nod) *Fge {
	return &Fge{
		Strt: str,
		Nod:  nod,
	}
}

func (f *Fge) Pbf() *pb.Fge {
	return &pb.Fge{
		Str: f.Strt[:],
		Nod: f.Nod.Pbf(),
	}
}
