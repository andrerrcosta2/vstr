package nwu

import (
	"errors"
	"github.com/andrerrcosta2/vstr/nmm/srs/reg"
	"time"
)

const (
	UFT_NNFE = "UFT-NNFE"
	UFT_NNFR = "UFT-NNFR"
	STB_DSNU = "STB-DSNU"
	STB_DSNR = "STB-DSNR"
	STB_DPNU = "STB-DPNU"
	STB_DPNR = "STB-DPNR"
	JN_BNNF  = "JN-BNNF"
	UNET_ERR = "UNET-ERR"
	SE_ERR   = "SE-ERR"
	UNET     = "UNET"
)

type Ureg struct {
	err error
	cod string
	dat any
	ts  time.Time
}

func NewUreg(cod string, err error, dat any) *Ureg {
	return &Ureg{
		err: err,
		cod: cod,
		dat: dat,
		ts:  time.Now(),
	}
}

func (u *Ureg) Gts() time.Time {
	return u.ts
}

func (u *Ureg) Grc() string {
	return u.cod
}

func (u *Ureg) Ger() error {
	return u.err
}

func (u *Ureg) Gdt() any {
	return u.dat
}

func (u *Ureg) Equal(ot reg.Srrg) bool {
	o, ok := ot.(*Ureg)
	if !ok {
		return false
	}
	if u == o {
		return true
	}
	if o == nil {
		return false
	}
	return errors.Is(u.err, o.err) &&
		u.cod == o.cod &&
		u.dat == o.dat &&
		u.ts.Equal(o.ts)
}
