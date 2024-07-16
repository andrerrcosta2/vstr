// Andre R. R. Costa *** github.com/andrerrcosta2 *** andrerrcosta@gmail.com

package nod

import (
	"fmt"
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

type Nod struct {
	ID  cid.Id
	Ip  string
	Pt  int32
	Fgs []*Fge
	Pre *Nod
	Suc *Nod
	Dat map[string][]byte
}

func NewNod(id cid.Id, ip string, pt int32) *Nod {
	return &Nod{
		ID:  id,
		Ip:  ip,
		Pt:  pt,
		Fgs: make([]*Fge, cid.M),
		Dat: make(map[string][]byte),
	}
}

func (n *Nod) Cpn(id cid.Id) *Nod {
	var cpn cid.Id
	var out *Nod
	var drt = false

	for _, e := range n.Fgs {
		if id.Gt(e.Strt) && (!drt || e.Strt.Gt(cpn)) {
			cpn = e.Strt
			out = e.Nod
			drt = true
		}
	}
	return out
}

func (n *Nod) Ffe(id cid.Id) *Nod {
	for _, e := range n.Fgs {
		if id.Btw(e.Strt, e.Nod.ID) {
			return e.Nod
		}
	}
	return nil
}

func (n *Nod) Hfe() *Nod {
	var hid cid.Id
	var out *Nod

	for _, entry := range n.Fgs {
		if entry.Strt.Btw(hid, n.ID) {
			hid = entry.Strt
			out = entry.Nod
		}
	}
	return out
}

func (n *Nod) Fgh() bool {
	if n.Fgs == nil {
		return false
	}
	if len(n.Fgs) < cid.M {
		return false
	}
	for _, fg := range n.Fgs {
		if !fg.Strt.Vld() || fg.Nod == nil {
			return false
		}
	}
	return true

}

func (n *Nod) Gfaddr() string {
	return fmt.Sprintf("%s:%d", n.Ip, n.Pt)
}

func (n *Nod) Pbf() *pb.Nod {
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

func Dbf(n *pb.Nod) (*Nod, error) {
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
