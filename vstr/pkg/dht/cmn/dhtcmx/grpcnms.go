// Andre R. R. Costa *** github.com/andrerrcosta2 *** andrerrcosta@gmail.com

package dhtcmx

import (
	"context"
	"fmt"
	"github.com/andrerrcosta2/vstr/vstr/pkg/dht/cmn/dhtms"
	"github.com/andrerrcosta2/vstr/vstr/pkg/msg/msrf"
	"github.com/andrerrcosta2/vstr/vstr/pkg/ntk/dlr"
	"github.com/andrerrcosta2/vstr/vstr/pkg/pb"
	"github.com/andrerrcosta2/vstr/vstr/pkg/srd/ret"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"time"
)

type Grpcns struct {
	ftr   msrf.Grpcf[pb.GrpcnmsClient]
	rtopt ret.Opt[any]
}

func (s *Grpcns) Sm(n *dhtms.Nms, addr string) (*dhtms.Nms, error) {
	ins, err := s.ftr.Gin(addr, func(conn dlr.RpcConn) pb.GrpcnmsClient {
		return pb.NewGrpcnmsClient(conn.(grpc.ClientConnInterface))
	})
	if err != nil {
		return nil, fmt.Errorf("sm failed: %v\n", err)
	}
	defer ins.End()

	rsc := map[codes.Code]bool{
		codes.Unavailable:       true,
		codes.ResourceExhausted: true,
	}

	res, err := ret.Rt(&s.rtopt, func() (*pb.Nms, error) {
		tmo := s.rtopt.Tmo
		if tmo == 0 {
			tmo = 5 * time.Second
		}

		ctx, cnc := context.WithTimeout(s.rtopt.Ctx, tmo)

		out, err := ins.Cli().Snm(ctx, &pb.Nms{
			Type: n.Type,
			Id:   n.Id,
			Data: n.Data,
		}, ins.Copts()...)

		cnc()

		if rsc[status.Code(err)] {
			return nil, err
		}
		return out, nil
	})

	if err != nil {
		return nil, fmt.Errorf("sm failed: %v", err)
	}

	return &dhtms.Nms{
		Type: res.Type,
		Id:   res.Id,
		Data: res.Data,
	}, nil
}

func NewGrpns(ftr msrf.Rpccf[Dhtpmsr[proto.Message]], rtopt ret.Opt[any]) *Grpcns {
	return &Grpcns{
		ftr:   ftr,
		rtopt: rtopt,
	}
}

type Dhtrcppcf[T proto.Message] struct {
	fct func(conn dlr.RpcConn) Dhtpmsr[T]
}

type Dhtpmsr[T proto.Message] interface {
	Snm(ctx context.Context, in T, opts ...grpc.CallOption) (T, error)
}
