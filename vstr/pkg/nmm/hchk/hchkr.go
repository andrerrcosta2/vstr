package hchk

import "github.com/andrerrcosta2/vstr/nwk"

type Hchkr[T nwk.NwkAddr] interface {
	Chk(tgt T) error
}
