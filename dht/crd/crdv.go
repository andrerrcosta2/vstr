// github.com/andrerrcosta2 *** andrerrcosta@gmail.com

package crd

import (
	"errors"
	"fmt"
	"github.com/andrerrcosta2/vstr/dht/cmn/dhtcmx"
	"github.com/andrerrcosta2/vstr/dht/cmn/dhtms"
	"github.com/andrerrcosta2/vstr/dht/crd/cid"
	"github.com/andrerrcosta2/vstr/dht/crd/fly"
	"github.com/andrerrcosta2/vstr/dht/crd/lut"
	"github.com/andrerrcosta2/vstr/dht/crd/nod"
	"github.com/andrerrcosta2/vstr/dht/crd/nwu"
	"github.com/andrerrcosta2/vstr/dht/dhtcfg"
	"github.com/andrerrcosta2/vstr/nmm/hchk"
	"github.com/andrerrcosta2/vstr/nmm/srcfg"
	"github.com/andrerrcosta2/vstr/nwk"
	"github.com/andrerrcosta2/vstr/srd/ret"
	"sync"
)

type crdv struct {
	cfg  dhtcfg.CrdCfg
	nod  nod.Nod
	msr  dhtcmx.Nmr
	lut  lut.LUT
	urgs *nwu.Urstk
	hsts *hchk.Hchks[*nwu.Ureg]
}

func New(cfg dhtcfg.CrdCfg, msr dhtcmx.Nmr) *crdv {
	urgs := nwu.NewUrstk(100)
	return &crdv{
		cfg:  cfg,
		nod:  nod.Nod{},
		msr:  msr,
		lut:  lut.NewCnlt(),
		urgs: urgs,
		hsts: hchk.NewHchks[*nwu.Ureg](urgs),
	}
}

func (c *crdv) St(key string, value []byte) {
	c.nod.Dat[key] = value
}

func (c *crdv) Rt(key string) ([]byte, bool) {
	dat, ok := c.nod.Dat[key]
	return dat, ok
}

func (c *crdv) fs(id cid.Id) (*nod.Nod, error) {
	n := c.nod.Ffe(id)
	if n != nil {
		return n, nil
	}
	return c.fws(id)
}

func (c *crdv) fp(id cid.Id) (*nod.Nod, error) {
	if c.nod.Suc.ID.Eq(id) {
		return &c.nod, nil
	}
	return c.fwp(c.nod.Pre.ID)
}

// fws
func (c *crdv) fws(id cid.Id) (*nod.Nod, error) {
	cpn := c.nod.Cpn(id)
	if cpn == nil {
		return nil, fmt.Errorf("fgs not initialized\n")
	}
	out, err := fly.Rsc(c.msr, id[:], cpn)

	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crdv) fwp(id cid.Id) (*nod.Nod, error) {
	if c.nod.Pre == nil {
		return nil, errors.New("no pre found\n")
	}
	out, err := fly.Rpc(c.msr, id[:], c.nod.Pre)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *crdv) bsp() (*nod.Nod, error) {
	return ret.RtSh(uint(len(c.cfg.Btln)), c.cfg.JnDl, c.cfg.Btln, func(n nwk.NwkAddr) (*nod.Nod, error) {
		out, err := fly.Rsc(c.msr, c.nod.ID[:], n)
		if err != nil {
			return nil, err
		}
		return out, nil
	})
}

func (c *crdv) ift() {
	wg := sync.WaitGroup{}
	for i := 0; i < cid.M; i++ {
		wg.Add(1)
		go func(i uint) {
			defer wg.Done()
			c.ife(i)
		}(uint(i))
	}
	wg.Wait()
}

func (c *crdv) ife(i uint) {
	str := c.nod.ID.Strt(i)
	suc, err := c.fs(str)
	if err != nil {
		c.urgs.Add(nwu.UFT_NNFE, err, str)
	} else {
		c.nod.Fgs[i] = nod.NewFge(str, suc)
	}
}

func (c *crdv) uft() {
	wg := sync.WaitGroup{}
	for i := range c.nod.Fgs {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c.ufe(c.nod.Fgs[i])
		}(i)
	}
	wg.Wait()
}

func (c *crdv) ufe(f *nod.Fge) {
	if f.Nod != nil {
		if err := c.msr.Hchk(f.Nod); err == nil {
			return
		}
	}
	suc, err := c.fs(f.Strt)
	if err != nil {
		f.Nod = nil
		c.urgs.Add(nwu.UFT_NNFE, err, f.Strt)
	} else {
		f.Nod = suc
		c.urgs.Add(nwu.UFT_NNFR, nil, f.Strt)
	}
}

func (c *crdv) unet() {
	_, err := c.msr.Qs(dhtms.NewNms(c.nod.ID[:], nwu.UNET, nil), c.nod.Suc)
	c.urgs.Add(nwu.UNET_ERR, err, nil)
}

/**
JN
*/

func (c *crdv) Jn() error {
	c.hsts.Upd(hchk.ST_INIT, nil)
	n, err := c.bsp()

	if err != nil {
		c.urgs.Add(nwu.JN_BNNF, err, nil)
		c.hsts.Upd(hchk.ST_DWN, err)
		return err
	}

	if n != nil {
		c.lut.Add(n.Ip, n.Pt)
		c.nod.Suc = n
		// TODO: c.nod.Pre
	} else {
		c.nod.Suc = &c.nod
		c.nod.Pre = &c.nod
	}

	c.ift()
	c.unet()
	c.hsts.Upd(hchk.ST_RLS, nil)
	return nil
}

/**
STB
*/

func (c *crdv) Hbt() {
	c.hsts.Upd(hchk.ST_UPD, nil)
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := c.msr.Hchk(c.nod.Suc); err != nil {
			c.urgs.Add(nwu.STB_DSNU, err, nil)
		}
	}()
	go func() {
		defer wg.Done()
		if err := c.msr.Hchk(c.nod.Pre); err != nil {
			c.urgs.Add(nwu.STB_DPNU, err, nil)
		}
	}()
	wg.Wait()

	c.uft()
	c.unet()
	c.hsts.Upd(hchk.ST_RLS, nil)
}

/**
RCV
*/

func (c *crdv) Rcv(s srcfg.Srcfg) {
	c.hsts.Upd(hchk.ST_UPD, nil)
	defer c.hsts.Upd(hchk.ST_RLS, nil)

	c.dpsu(s)
	c.dpnu(s)
	c.ift()
	c.unet()
}

func (c *crdv) dpsu(s srcfg.Srcfg) {
	if c.urgs.Ctn(nwu.STB_DSNU) {
		stf := func(a, b *nod.Fge) bool {
			return !a.Strt.Btw(c.nod.ID, b.Strt)
		}
		pms, err := ret.RtSt(uint(s.RtAtt), s.RtDl, c.nod.Fgs, stf, func(n *nod.Fge) (*nod.Nod, error) {
			return fly.Rsc(c.msr, c.nod.ID[:], n.Nod)
		})

		if err != nil {
			c.urgs.Add(nwu.STB_DSNU, err, nil)
		} else {
			c.nod.Suc = pms
			c.urgs.RK(nwu.STB_DSNU)
		}
	}
}

func (c *crdv) dpnu(s srcfg.Srcfg) {
	if c.urgs.Ctn(nwu.STB_DPNU) {
		stf := func(a, b *nod.Fge) bool {
			return !a.Strt.Btw(c.nod.ID, b.Strt)
		}
		pms, err := ret.RtSt(uint(s.RtAtt), s.RtDl, c.nod.Fgs, stf, func(n *nod.Fge) (*nod.Nod, error) {
			return fly.Rpc(c.msr, c.nod.ID[:], n.Nod)
		})

		if err != nil {
			c.urgs.Add(nwu.STB_DPNU, err, nil)
		} else {
			c.nod.Suc = pms
			c.urgs.RK(nwu.STB_DPNU)
		}
	}
}
