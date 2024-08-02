// Andre R. R. Costa *** github.com/andrerrcosta2

package lut

import (
	"errors"
	"github.com/andrerrcosta2/vstr/dht/crd/cid"
	"github.com/andrerrcosta2/vstr/dht/crd/nod"
	"net"
	"sync"
)

type Cnlt struct {
	mt sync.Mutex
	tb map[string]*nod.Nod
}

func Gid(ip net.IP, pt uint16) cid.Id {
	return cid.New(ip, pt)
}

func (t *Cnlt) Add(ip net.IP, pt uint16) cid.Id {
	id := Gid(ip, pt)
	n := &nod.Nod{
		ID: id,
		Ip: ip,
		Pt: pt,
	}
	t.mt.Lock()
	t.tb[id.String()] = n
	t.mt.Unlock()
	return id
}

func (t *Cnlt) Get(id cid.Id) (*nod.Nod, error) {
	t.mt.Lock()
	defer t.mt.Unlock()
	nod, ok := t.tb[id.String()]
	if !ok {
		return nil, errors.New("nod not found")
	}
	return nod, nil
}

func NewCnlt() *Cnlt {
	return &Cnlt{
		tb: make(map[string]*nod.Nod),
	}
}
