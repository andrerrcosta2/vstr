// Andre R. R. Costa *** github.com/andrerrcosta2

package dhtcmx

import (
	"github.com/andrerrcosta2/vstr/dht/cmn/cmcod"
	"github.com/andrerrcosta2/vstr/dht/cmn/dhtms"
	"github.com/andrerrcosta2/vstr/msg/tpl"
	"github.com/andrerrcosta2/vstr/nmm/hchk"
	"github.com/andrerrcosta2/vstr/nwk"
	"sync"
)

type Dnmr struct {
	tpl  tpl.Tpl[*dhtms.Nms, *dhtms.Nms]
	hchk hchk.Chckr[nwk.NwkAddr]
}

func (t *Dnmr) Qs(msg *dhtms.Nms, n nwk.NwkAddr) (*dhtms.Nms, error) {
	res, err := t.tpl.Sm(msg, n.Gfaddr())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *Dnmr) Qm(msg *dhtms.Nms, nds ...nwk.NwkAddr) []*dhtms.Nms {
	var wg sync.WaitGroup
	rch := make(chan *dhtms.Nms, len(nds))

	for _, nd := range nds {
		wg.Add(1)
		go func(nd nwk.NwkAddr) {
			defer wg.Done()
			res, err := t.Qs(msg, nd)
			if err != nil {
				rch <- &dhtms.Nms{
					Type: cmcod.RES_GCE,
					Data: []byte(err.Error()),
				}
			}
			rch <- res
		}(nd)
	}
	wg.Wait()
	close(rch)

	opt := make([]*dhtms.Nms, 0, len(nds))
	for r := range rch {
		opt = append(opt, r)
	}
	return opt
}

func (t *Dnmr) Qms(n nwk.NwkAddr, msgs ...*dhtms.Nms) []*dhtms.Nms {
	var wg sync.WaitGroup
	rch := make(chan *dhtms.Nms, len(msgs))

	for _, m := range msgs {
		wg.Add(1)
		go func(m *dhtms.Nms) {
			defer wg.Done()
			res, err := t.Qs(m, n)
			if err != nil {
				rch <- &dhtms.Nms{
					Type: cmcod.RES_GCE,
					Data: []byte(err.Error()),
				}
				return
			}
			rch <- res
		}(m)
	}
	wg.Wait()
	close(rch)

	opt := make([]*dhtms.Nms, 0, len(msgs))
	for r := range rch {
		opt = append(opt, r)
	}
	return opt
}

func (t *Dnmr) Hchk(n nwk.NwkAddr) error {
	return t.hchk.Chk(n)
}

func NewDnmr(tpl tpl.Tpl[*dhtms.Nms, *dhtms.Nms], chckr hchk.Chckr[nwk.NwkAddr]) *Dnmr {
	return &Dnmr{
		tpl:  tpl,
		hchk: chckr,
	}
}
