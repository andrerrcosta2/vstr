package hchk

import (
	"fmt"
	"github.com/andrerrcosta2/vstr/nmm/srcfg"
	"github.com/andrerrcosta2/vstr/nwk"
	"github.com/andrerrcosta2/vstr/nwk/dlr"
	"github.com/andrerrcosta2/vstr/srd/ret"
)

type Chckr[T nwk.NwkAddr] struct {
	cfg srcfg.Cnmm
	dlr dlr.Dlr
}

func (c *Chckr[T]) Chk(addr T) error {
	conn, err := ret.Rt(c.cfg.Att, c.cfg.Dl, func() (dlr.Conn, error) {
		return c.dlr.Dlt(addr.Gfaddr(), c.cfg.Intv)
	})
	if err != nil {
		return fmt.Errorf("addr '%s' is unreachable: %w\n", addr.Gfaddr(), err)
	}
	defer dlr.Cls(conn)
	return nil
}
