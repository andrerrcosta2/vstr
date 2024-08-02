package nod

import "github.com/andrerrcosta2/vstr/dht/crd/cid"

func gft(id cid.Id) []*Fge {
	out := make([]*Fge, cid.M)
	for i := 0; i < cid.M; i++ {
		out[i] = NewFge(id.Strt(uint(i)), &Nod{})
	}
	return out
}
