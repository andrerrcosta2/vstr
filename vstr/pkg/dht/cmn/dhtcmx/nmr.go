// Andre R. R. Costa *** github.com/andrerrcosta2

package dhtcmx

import (
	"github.com/andrerrcosta2/vstr/dht/cmn/dhtms"
	"github.com/andrerrcosta2/vstr/nwk"
)

type Nmr interface {
	Qs(msg *dhtms.Nms, n nwk.NwkAddr) (*dhtms.Nms, error)
	Qm(msg *dhtms.Nms, n ...nwk.NwkAddr) []*dhtms.Nms
	Qms(n nwk.NwkAddr, msg ...*dhtms.Nms) []*dhtms.Nms
	Hchk(n nwk.NwkAddr) error
}
