package hchk

import (
	"errors"
	"github.com/andrerrcosta2/vstr/nmm/srs/ht"
	"github.com/andrerrcosta2/vstr/nmm/srs/reg"
	"github.com/andrerrcosta2/vstr/srd/obs"
	"github.com/andrerrcosta2/vstr/srd/obs/sbj"
	"sync"
	"time"
)

type Hsts int

const (
	ST_INIT Hsts = iota
	ST_UPD
	ST_DWN
	ST_ERR
	ST_HTH
	ST_RLS
)

type Hchks[T reg.Srrg] struct {
	st    Hsts
	lserr error
	lschk time.Time
	urgs  ht.Uhtl[T]
	obs   obs.Sbj[[]T]
	mt    sync.Mutex
}

func NewHchks[T reg.Srrg](urgs ht.Uhtl[T]) *Hchks[T] {
	return &Hchks[T]{
		st:    ST_RLS,
		lschk: time.Now(),
		urgs:  urgs,
		obs:   sbj.NewDsbj[[]T](),
		mt:    sync.Mutex{},
	}
}

func (h *Hchks[T]) Upd(st Hsts, err error) {
	h.mt.Lock()
	defer h.mt.Unlock()
	h.st = st
	h.lserr = err
	h.lschk = time.Now()
	h.nxt()
}

func (h *Hchks[T]) sts() (Hsts, error, time.Time) {
	h.mt.Lock()
	defer h.mt.Unlock()
	return h.st, h.lserr, h.lschk
}

func (h *Hchks[T]) Sub(obv obs.Obv[[]T]) (*obs.Sub[[]T], error) {
	if h.obs.Cld() {
		return nil, errors.New("this sjc is closed\n")
	}
	return h.obs.Sub(obv)
}

func (h *Hchks[T]) Uns() {
	h.obs.Uns()
}

func (h *Hchks[T]) Cld() bool {
	return h.obs.Cld()
}

func (h *Hchks[T]) nxt() {
	switch h.st {
	case ST_RLS:
		h.obs.Nxt(h.urgs.Xtr())
		break

	case ST_DWN:
		h.obs.Nxt(h.urgs.Xtr())
		break

	default:
		break
	}
}
