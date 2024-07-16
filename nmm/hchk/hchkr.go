package hchk

type Nhckr interface {
	Chk(tgt any) error
}
