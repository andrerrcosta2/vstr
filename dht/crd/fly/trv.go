// Andre R. R. Costa *** github.com/andrerrcosta2

package fly

import (
	"errors"
	"fmt"
	"github.com/andrerrcosta2/vstr/dht/cmn/cmcod"
	"github.com/andrerrcosta2/vstr/dht/cmn/dhtcmx"
	"github.com/andrerrcosta2/vstr/dht/cmn/dhtms"
	"github.com/andrerrcosta2/vstr/dht/crd/nod"
	"github.com/andrerrcosta2/vstr/nwk"
	"github.com/andrerrcosta2/vstr/pb"
	"google.golang.org/protobuf/proto"
)

func Rsc(msr dhtcmx.Nmr, ref []byte, bts nwk.NwkAddr) (*nod.Nod, error) {
	if msr == nil {
		return nil, fmt.Errorf("msr cannot be nil")
	}
	if ref == nil {
		return nil, fmt.Errorf("ref cannot be nil")
	}
	if bts == nil {
		return nil, fmt.Errorf("bts cannot be nil")
	}

	res, err := msr.Qs(&dhtms.Nms{
		Id:   ref,
		Type: cmcod.REQ_NS,
	}, bts)

	if err != nil {
		return nil, err
	}

	switch res.Type {
	case cmcod.RES_NSNF:
		return nil, errors.New(cmcod.RES_NSNF)
	case cmcod.RES_NSS:
		msg := &pb.Nod{}
		err = proto.Unmarshal(res.Data, msg)
		if err != nil {
			return nil, err
		}
		dbf, err := nod.Dbf(msg)
		if err != nil {
			return nil, err
		}
		return dbf, nil
	default:
		return nil, errors.New("Unknown Response Type")
	}
}

func Rpc(msr dhtcmx.Nmr, ref []byte, bts nwk.NwkAddr) (*nod.Nod, error) {
	if msr == nil {
		return nil, fmt.Errorf("msr cannot be nil")
	}
	if ref == nil {
		return nil, fmt.Errorf("ref cannot be nil")
	}
	if bts == nil {
		return nil, fmt.Errorf("bts cannot be nil")
	}

	res, err := msr.Qs(&dhtms.Nms{
		Id:   ref,
		Type: cmcod.REQ_NP,
	}, bts)

	if err != nil {
		return nil, err
	}

	switch res.Type {
	case cmcod.RES_NPNF:
		return nil, errors.New(cmcod.RES_NSNF)
	case cmcod.RES_NPS:
		msg := &pb.Nod{}
		err = proto.Unmarshal(res.Data, msg)
		if err != nil {
			return nil, err
		}
		dbf, err := nod.Dbf(msg)
		if err != nil {
			return nil, err
		}
		return dbf, nil
	default:
		return nil, errors.New("Unknown Response Type")
	}
}

func Njn(msr dhtcmx.Nmr, jn *nod.Nod, ref *nod.Nod) error {
	if msr == nil {
		return fmt.Errorf("msr cannot be nil\n")
	}
	if jn == nil {
		return fmt.Errorf("jn cannot be nil\n")
	}
	if ref == nil {
		return fmt.Errorf("ref cannot be nil\n")
	}

	_, err := proto.Marshal(jn.Pbf())
	if err != nil {
		return fmt.Errorf("failed marshaling nod: %v\n", err)
	}

	return nil
}
