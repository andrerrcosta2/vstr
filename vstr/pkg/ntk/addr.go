package nwk

type NwkAddr interface {
	Gfaddr() string
}

// TUSF
type Saddr struct {
	addr string
}

func (d *Saddr) Gfaddr() string {
	return d.addr
}

func NewSaddr(addr string) NwkAddr {
	return &Saddr{addr: addr}
}
