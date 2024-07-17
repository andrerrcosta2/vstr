package nod

import (
	"bytes"
	"github.com/andrerrcosta2/vstr/dht/crd/cid"
	"net"
	"testing"
)

func TestNewFge(t *testing.T) {
	id := cid.New(net.IP{127, 0, 0, 1}, 8080)
	nod := &Nod{
		ID: id,
		Ip: net.IP{127, 0, 0, 1},
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
	id := cid.New(net.IP{127, 0, 0, 1}, 8080)
	nod := &Nod{
		ID: id,
		Ip: net.IP{127, 0, 0, 1},
		Pt: 8080,
	}
	fge := NewFge(id, nod)
	pbf := fge.Pbf()
	if !bytes.Equal(pbf.Str, id[:]) {
		t.Errorf("Expected start ID %v, got %v", id, pbf.Str)
	}
	if !nod.Ip.Equal(net.ParseIP(pbf.Nod.Ip)) || pbf.Nod.Port != int32(nod.Pt) {
		t.Errorf("Expected node IP %v and Port %v, got IP %v and Port %v", nod.Ip, nod.Pt, pbf.Nod.Ip, pbf.Nod.Port)
	}
}
