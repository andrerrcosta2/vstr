package srs

import "github.com/andrerrcosta2/vstr/nmm/srcfg"

type Rcv interface {
	Rcv(s srcfg.Srcfg)
}
