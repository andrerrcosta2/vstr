// Andre R. R. Costa *** github.com/andrerrcosta2

package crd

import (
	"errors"
	"fmt"
	"github.com/andrerrcosta/vstr/crd/cid"
	"github.com/andrerrcosta/vstr/crd/drl"
	"github.com/andrerrcosta/vstr/crd/lut"
	"github.com/andrerrcosta/vstr/crd/msg"
	"github.com/andrerrcosta/vstr/crd/nod"
	"github.com/andrerrcosta/vstr/pb"
	"google.golang.org/protobuf/proto"
)

type Crdv struct {
	nod *nod.Nod
	msr msg.Nmr
	lut lut.LUT
}

func (c *Crdv) St(key string, value []byte) {
	c.nod.Dat[key] = value
}

func (c *Crdv) Rt(key string) ([]byte, bool) {
	dat, ok := c.nod.Dat[key]
	return dat, ok
}

// Jn jn
func (c *Crdv) Jn(btn *nod.Nod) error {
	n, err := c.bsp(btn)
	if err != nil {
		return err
	}

	if n != nil {
		c.lut.Add(n.Ip, n.Pt)
		c.nod.Suc = n
	} else {
		c.nod.Suc = c.nod
		c.nod.Pre = c.nod
	}
	err = c.ift()
	if err != nil {
		return err
	}
	err = c.unet()
	if err != nil {
		return err
	}
	return nil
}

// fws
func (c *Crdv) fws(id cid.Id) (*nod.Nod, error) {
	if c.nod == nil || c.nod.Suc == nil {
		return nil, errors.New("no node or successor found")
	}

	msg, err := drl.Rsc(c.msr, id[:], c.nod.Suc)
	if err != nil {
		return nil, err
	}
	var suc pb.Nod
	if err = proto.Unmarshal(msg.Data, &suc); err != nil {
		return nil, err
	}
	return nod.Dbf(suc)
}

func (c *Crdv) fs(id cid.Id) (*nod.Nod, error) {
	if c.nod == nil {
		return nil, errors.New("no node found")
	}

	n := c.nod.Ffe(id)
	if n != nil {
		return n, nil
	}

	return c.fws(id)
}

func (c *Crdv) bsp(bts *nod.Nod) (*nod.Nod, error) {
	msg, err := drl.Rsc(c.msr, c.nod.ID[:], bts)
	if err != nil {
		return nil, err
	}
	var suc pb.Nod
	if err = proto.Unmarshal(msg.Data, &suc); err != nil {
		return nil, err
	}
	return nod.Dbf(suc)
}

func (c *Crdv) ift() error {
	for i := 0; i < cid.M; i++ {
		str := c.nod.ID.Strt(i)
		suc, err := c.fs(str)
		if err != nil {
			return fmt.Errorf("error initializing finger table entry %d: %v", i, err)
		}
		c.nod.Fgs[i] = *nod.NewFge(str, suc)
	}
	return nil
}

func (c *Crdv) unet() error {
	// TODO:
	return nil
}
