// Andre R. R. Costa *** github.com/andrerrcosta2

package msg

import (
	"fmt"
	"github.com/andrerrcosta2/vstr/cmx"
	"github.com/andrerrcosta2/vstr/crd/nod"
	"github.com/andrerrcosta2/vstr/pb"
	"sync"
)

type TcpNms struct {
	cpmx cmx.Cpmx
}

func (t *TcpNms) Qs(msg *pb.Msg, n *nod.Nod) (pb.Msg, error) {
	res, err := t.cpmx.Spm(msg, fmt.Sprintf("%s:%d", n.Ip, n.Pt))
	if err != nil {
		return pb.Msg{}, err
	}

	out, ok := res.(*pb.Msg)
	if !ok {
		return pb.Msg{}, fmt.Errorf("unexpected response type")
	}

	return *out, nil
}

func (t *TcpNms) Qm(msg *pb.Msg, nds ...*nod.Nod) []pb.Msg {
	var wg sync.WaitGroup
	rch := make(chan pb.Msg, len(nds))

	for _, nd := range nds {
		wg.Add(1)
		go func(nd *nod.Nod) {
			defer wg.Done()
			res, err := t.Qs(msg, nd)
			if err != nil {
				rch <- pb.Msg{
					Type: RES_GCE,
					Data: []byte(err.Error()),
				}
			}
			rch <- res
		}(nd)
	}
	wg.Wait()
	close(rch)

	opt := make([]pb.Msg, 0, len(nds))
	for r := range rch {
		opt = append(opt, r)
	}
	return opt
}

func (t *TcpNms) Qms(n *nod.Nod, msgs ...*pb.Msg) []pb.Msg {
	var wg sync.WaitGroup
	rch := make(chan pb.Msg, len(msgs))

	for _, m := range msgs {
		wg.Add(1)
		go func(m *pb.Msg) {
			defer wg.Done()
			res, err := t.Qs(m, n)
			if err != nil {
				rch <- pb.Msg{
					Type: RES_GCE,
					Data: []byte(err.Error()),
				}
				return
			}
			rch <- res
		}(m)
	}
	wg.Wait()
	close(rch)

	opt := make([]pb.Msg, 0, len(msgs))
	for r := range rch {
		opt = append(opt, r)
	}
	return opt
}
