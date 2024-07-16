// Andre R. R. Costa *** github.com/andrerrcosta2

package lut

import (
	"github.com/andrerrcosta2/vstr/dht/crd/cid"
	"github.com/andrerrcosta2/vstr/dht/crd/nod"
)

type LUT interface {
	Add(ip string, pt int32) cid.Id
	Get(id cid.Id) (*nod.Nod, error)
}
