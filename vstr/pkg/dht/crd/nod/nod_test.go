// github.com/andrerrcosta2 *** andrerrcosta@gmail.com

package nod

import (
	"bytes"
	"github.com/andrerrcosta2/vstr/dht/crd/cid"
	"github.com/andrerrcosta2/vstr/pb"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestNewNod(t *testing.T) {
	ip := net.ParseIP("192.168.1.1")
	pt := uint16(8080)
	id := cid.New(ip, pt)
	nod := NewNod(id, ip, pt)

	assert.Equal(t, id, nod.ID)
	assert.Equal(t, ip, nod.Ip)
	assert.Equal(t, pt, nod.Pt)
	assert.NotNil(t, nod.Fgs)
	assert.Equal(t, cid.M, len(nod.Fgs))
	assert.NotNil(t, nod.Dat)
}

func TestCpn(t *testing.T) {
	id := cid.New(net.IP{127, 0, 0, 1}, 8080)
	nod := &Nod{
		ID: id,
		Ip: net.IP{127, 0, 0, 1},
		Pt: 8080,
	}
	cpn := nod.Cpn(id)
	if cpn != nod {
		t.Errorf("Expected node %v, got %v", nod, cpn)
	}
}

func TestNodPbf(t *testing.T) {
	id := cid.New(net.IP{127, 0, 0, 1}, 8080)
	nod := &Nod{
		ID: id,
		Ip: net.IP{127, 0, 0, 1},
		Pt: 8080,
	}
	pbf := nod.Pbf()
	if !bytes.Equal(pbf.Id, id[:]) {
		t.Errorf("Expected node ID %v, got %v", id, pbf.Id)
	}
	if !nod.Ip.Equal(net.ParseIP(pbf.Ip)) || pbf.Port != int32(nod.Pt) {
		t.Errorf("Expected node IP %v and Port %v, got IP %v and Port %v", nod.Ip, nod.Pt, pbf.Ip, pbf.Port)
	}
}

func TestDbf(t *testing.T) {
	id := cid.New(net.IP{127, 0, 0, 1}, 8080)
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
	if !nod.Ip.Equal(net.ParseIP(pbn.Ip)) || int32(nod.Pt) != pbn.Port {
		t.Errorf("Expected node IP %v and Port %v, got IP %v and Port %v", pbn.Ip, pbn.Port, nod.Ip, nod.Pt)
	}
}
