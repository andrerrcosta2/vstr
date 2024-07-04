// Andre R. R. Costa *** github.com/andrerrcosta2

package msg

import (
	"github.com/andrerrcosta/vstr/crd/nod"
	"github.com/andrerrcosta/vstr/pb"
)

type Nmr interface {
	Qs(msg *pb.Msg, n *nod.Nod) (pb.Msg, error)
	Qm(msg *pb.Msg, n ...*nod.Nod) (pb.Msg, error)
	Qms(n *nod.Nod, msg ...*pb.Msg) (pb.Msg, error)
}
