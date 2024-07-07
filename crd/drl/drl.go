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
	return msr.Qs(&pb.Msg{
		Id:   ref,
		Type: msg.REQ_NS,
	}, bts)
}

func Njn(msr msg.Nmr, jn *nod.Nod, suc *nod.Nod) (pb.Msg, error) {
	jnb, err := proto.Marshal(jn.Pbf())
	if err != nil {
		return pb.Msg{}, fmt.Errorf("failed to marshal join node: %v", err)
	}

	return msr.Qs(&pb.Msg{
		Id:   jn.ID[:],
		Data: jnb,
		Type: msg.REQ_UNET,
	}, suc)
}
