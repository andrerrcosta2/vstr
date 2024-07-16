package dhtcfg

import (
	"github.com/andrerrcosta2/vstr/nwk"
	"time"
)

type CrdCfg struct {
	UnetRt  uint
	UnetDl  time.Duration
	UnetFrc float64
	JnRt    uint
	JnDl    time.Duration
	JnFrc   float64
	Btln    []nwk.NwkAddr
}
