package nwu

import (
	"sort"
	"sync"
)

type Urstk struct {
	mtx sync.Mutex
	reg map[string][]*Ureg
}

func NewUrstk(s int) *Urstk {
	return &Urstk{
		mtx: sync.Mutex{},
		reg: make(map[string][]*Ureg, s),
	}
}

func (u *Urstk) Add(code string, err error, data any) {
	u.mtx.Lock()
	defer u.mtx.Unlock()

	reg := NewUreg(code, err, data)

	if _, ok := u.reg[code]; !ok {
		u.reg[code] = []*Ureg{}
	}
	u.reg[code] = append(u.reg[code], reg)
}

func (u *Urstk) Get(code string) []*Ureg {
	regs := u.reg[code]
	srrgs := make([]*Ureg, len(regs))
	for i, r := range regs {
		srrgs[i] = r
	}
	return srrgs
}

func (u *Urstk) Mpt() bool {
	u.mtx.Lock()
	defer u.mtx.Unlock()

	return len(u.reg) == 0
}

func (u *Urstk) Ctn(cod string) bool {
	u.mtx.Lock()
	defer u.mtx.Unlock()

	_, ok := u.reg[cod]
	return ok
}

func (u *Urstk) RK(cod string) {
	u.mtx.Lock()
	defer u.mtx.Unlock()

	delete(u.reg, cod)
}

func (u *Urstk) RR(reg *Ureg) {
	u.mtx.Lock()
	defer u.mtx.Unlock()

	for i := range u.reg[reg.cod] {
		if u.reg[reg.cod][i].Equal(reg) {
			u.reg[reg.cod] = append(u.reg[reg.cod][:i], u.reg[reg.cod][i+1:]...)
			break
		}
	}
}

func (u *Urstk) Sz() int {
	u.mtx.Lock()
	defer u.mtx.Unlock()

	sz := 0
	for _, logs := range u.reg {
		sz += len(logs)
	}
	return sz
}

func (u *Urstk) Xtr() []*Ureg {
	u.mtx.Lock()
	defer u.mtx.Unlock()

	var wrgs []*Ureg
	for _, regs := range u.reg {
		wrgs = append(wrgs, regs...)
	}
	sort.SliceStable(wrgs, func(i, j int) bool {
		return wrgs[i].ts.Before(wrgs[j].ts)
	})

	out := make([]*Ureg, len(wrgs))
	copy(out, wrgs)
	return out
}
