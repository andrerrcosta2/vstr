// github.com/andrerrcosta2 *** andrerrcosta@gmail.com

package nod

import (
	"bytes"
	"github.com/andrerrcosta2/vstr/dht/crd/cid"
	"github.com/andrerrcosta2/vstr/pb"
	"testing"
)

func TestNewFge(t *testing.T) {
	id := cid.New("test")
	nod := &Nod{
		ID: id,
		Ip: "127.0.0.1",
		Pt: 8080,
	}
	fge := NewFge(id, nod)
	if !bytes.Equal(fge.Strt[:], id[:]) {
		t.Errorf("Expected start ID %v, got %v", id, fge.Strt)
	}
	if fge.Nod != nod {
		t.Errorf("Expected node %v, got %v", nod, fge.Nod)
	}
}

func TestFgePbf(t *testing.T) {
	id := cid.New("test")
	nod := &Nod{
		ID: id,
		Ip: "127.0.0.1",
		Pt: 8080,
	}
	fge := NewFge(id, nod)
	pbf := fge.Pbf()
	if !bytes.Equal(pbf.Str, id[:]) {
		t.Errorf("Expected start ID %v, got %v", id, pbf.Str)
	}
	if pbf.Nod.Ip != nod.Ip || pbf.Nod.Port != nod.Pt {
		t.Errorf("Expected node IP %v and Port %v, got IP %v and Port %v", nod.Ip, nod.Pt, pbf.Nod.Ip, pbf.Nod.Port)
	}
}

func TestNodPbf(t *testing.T) {
	id := cid.New("test")
	nod := &Nod{
		ID: id,
		Ip: "127.0.0.1",
		Pt: 8080,
	}
	pbf := nod.Pbf()
	if !bytes.Equal(pbf.Id, id[:]) {
		t.Errorf("Expected node ID %v, got %v", id, pbf.Id)
	}
	if pbf.Ip != nod.Ip || pbf.Port != nod.Pt {
		t.Errorf("Expected node IP %v and Port %v, got IP %v and Port %v", nod.Ip, nod.Pt, pbf.Ip, pbf.Port)
	}
}

func TestDbf(t *testing.T) {
	id := cid.New("test")
	pbn := pb.Nod{
		Id:   id[:],
		Ip:   "127.0.0.1",
		Port: 8080,
	}
	nod, err := Dbf(&pbn)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !bytes.Equal(nod.ID[:], id[:]) {
		t.Errorf("Expected node ID %v, got %v", id, nod.ID)
	}
	if nod.Ip != pbn.Ip || nod.Pt != pbn.Port {
		t.Errorf("Expected node IP %v and Port %v, got IP %v and Port %v", pbn.Ip, pbn.Port, nod.Ip, nod.Pt)
	}
}
