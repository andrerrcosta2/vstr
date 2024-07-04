// Andre R. R. Costa *** github.com/andrerrcosta2

package nod

import (
	"github.com/andrerrcosta/vstr/crd/cid"
	"github.com/andrerrcosta/vstr/pb"
)

type Fge struct {
	strt cid.Id
	nod  *Nod
}

func NewFge(str cid.Id, nod *Nod) *Fge {
	return &Fge{
		strt: str,
		nod:  nod,
	}
}

func (f *Fge) Pbf() *pb.Fge {
	return &pb.Fge{
		Str: f.strt[:],
		Nod: f.nod.Pbf(),
	}
}

type Nod struct {
	ID  cid.Id
	Ip  string
	Pt  int32
	Fgs []Fge
	Pre *Nod
	Suc *Nod
	Dat map[string][]byte
}

// Stb Stab
func (n *Nod) Stb() {
	x := n.Suc.Pre
	if x != nil && x.ID.Btw(n.ID, n.Suc.ID) {
		n.Suc = x
	}
	n.Suc.nfy(n)
}

// Cpn cpn
func (n *Nod) Cpn(nd Nod) *Nod {
	for _, e := range n.Fgs {
		if e.nod.ID.Btw(n.ID, nd.ID) {
			return e.nod
		}
	}
	return n
}

func (n *Nod) Ffe(id cid.Id) *Nod {
	for _, e := range n.Fgs {
		if e.strt.Btw(n.ID, id) {
			return e.nod
		}
	}
	return nil
}

// nfy nfy
func (n *Nod) nfy(t *Nod) {
	// TODO:
}

func (n Nod) Pbf() *pb.Nod {
	pbn := &pb.Nod{
		Id:   n.ID[:],
		Ip:   n.Ip,
		Port: n.Pt,
		Fgs:  make([]*pb.Fge, len(n.Fgs)),
		Dat:  n.Dat,
	}
	for i, fge := range n.Fgs {
		pbn.Fgs[i] = fge.Pbf()
	}
	if n.Pre != nil {
		pbn.Pre = &pb.Nod{
			Id:   n.Pre.ID[:],
			Ip:   n.Pre.Ip,
			Port: n.Pre.Pt,
		}
	}
	if n.Suc != nil {
		pbn.Suc = &pb.Nod{
			Id:   n.Suc.ID[:],
			Ip:   n.Suc.Ip,
			Port: n.Suc.Pt,
		}
	}
	return pbn
}

func Dbf(n pb.Nod) (*Nod, error) {
	id, err := cid.BytesToId(n.Id)
	if err != nil {
		return nil, err
	}
	return &Nod{
		ID: id,
		Ip: n.Ip,
		Pt: n.Port,
	}, nil
}
