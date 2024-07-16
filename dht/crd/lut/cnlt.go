// Andre R. R. Costa *** github.com/andrerrcosta2

package lut

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/andrerrcosta2/vstr/dht/crd/cid"
	"github.com/andrerrcosta2/vstr/dht/crd/nod"
	"sync"
)

type Cnlt struct {
	mt sync.Mutex
	tb map[string]*nod.Nod
}

func Gid(ip string, pt int32) cid.Id {
	addr := fmt.Sprintf("%s:%d", ip, pt)
	return sha1.Sum([]byte(addr))
}

func (t *Cnlt) Add(ip string, pt int32) cid.Id {
	id := Gid(ip, pt)
	n := &nod.Nod{
		ID:  id,
		Ip:  ip,
		Pt:  pt,
		Fgs: make([]nod.Fge, cid.M),
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
