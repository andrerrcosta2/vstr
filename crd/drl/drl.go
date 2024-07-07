// Andre R. R. Costa *** github.com/andrerrcosta2

package drl

import (
	"fmt"
	"github.com/andrerrcosta2/vstr/crd/msg"
	"github.com/andrerrcosta2/vstr/crd/nod"
	"github.com/andrerrcosta2/vstr/pb"
	"google.golang.org/protobuf/proto"
)

func Rsc(msr msg.Nmr, ref []byte, bts *nod.Nod) (pb.Msg, error) {
	if msr == nil {
		return pb.Msg{}, fmt.Errorf("msr cannot be nil")
	}
	if ref == nil {
		return pb.Msg{}, fmt.Errorf("ref cannot be nil")
	}
	if bts == nil {
		return pb.Msg{}, fmt.Errorf("bts cannot be nil")
	}
	return msr.Qs(&pb.Msg{
		Id:   ref,
		Type: msg.REQ_NS,
	}, bts)
}

func Njn(msr msg.Nmr, jn *nod.Nod, suc *nod.Nod) (pb.Msg, error) {
	if msr == nil {
		return pb.Msg{}, fmt.Errorf("msr cannot be nil\n")
	}
	if jn == nil {
		return pb.Msg{}, fmt.Errorf("jn cannot be nil\n")
	}
	if suc == nil {
		return pb.Msg{}, fmt.Errorf("suc cannot be nil\n")
	}

	jnb, err := proto.Marshal(jn.Pbf())
	if err != nil {
		return pb.Msg{}, fmt.Errorf("failed to marshal join node: %v\n", err)
	}

	return msr.Qs(&pb.Msg{
		Id:   jn.ID[:],
		Data: jnb,
		Type: msg.REQ_UNET,
	}, suc)
}
