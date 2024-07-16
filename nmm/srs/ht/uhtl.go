package ht

import (
	"github.com/andrerrcosta2/vstr/nmm/srs/reg"
)

type Uhtl[T reg.Srrg] interface {
	Ctn(code string) bool
	Mpt() bool
	Get(code string) []T
	RK(code string)
	RR(reg T)
	Xtr() []T
}
