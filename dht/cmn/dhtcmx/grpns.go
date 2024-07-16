package dhtcmx

import (
	"context"
	"fmt"
	"github.com/andrerrcosta2/vstr/dht/cmn/dhtms"
	"github.com/andrerrcosta2/vstr/pb"
	"github.com/andrerrcosta2/vstr/srd/ret"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

type grpns struct {
	dopts []grpc.DialOption
	copts []grpc.CallOption
	rtopt ret.Opt
}

func (s *grpns) Sm(n *dhtms.Nms, addr string) (*dhtms.Nms, error) {
	cli, conn, err := s.gc(addr, s.dopts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC client: %v\n", err)
	}
	defer cls(conn)

	rsc := map[codes.Code]bool{
		codes.Unavailable:       true,
		codes.ResourceExhausted: true,
	}

	res, err := ret.Rtf(s.rtopt.Att, s.rtopt.Dl, s.rtopt.Dlf, func() (*pb.Nms, error) {
		ctx, cnc := context.WithTimeout(context.Background(), 5*time.Second)

		out, err := cli.Snm(ctx, &pb.Nms{
			Type: n.Type,
			Id:   n.Id,
			Data: n.Data,
		}, s.copts...)

		cnc()

		if rsc[status.Code(err)] {
			return nil, err
		}
		return out, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to call Snm: %v", err)
	}

	if res == nil {
		return nil, fmt.Errorf("no response from Snm")
	}

	return &dhtms.Nms{
		Type: res.Type,
		Id:   res.Id,
		Data: res.Data,
	}, nil
}

func (s *grpns) gc(addr string, opts ...grpc.DialOption) (pb.GrpcnmsClient, *grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create gRPC client connection: %v", err)
	}
	return pb.NewGrpcnmsClient(conn), conn, nil
}

func cls(cnc *grpc.ClientConn) {
	err := cnc.Close()
	if err != nil {
		log.Printf("failed to close gRPC client connection: %v", err)
	}
}

func NewGrpns(dopts []grpc.DialOption, copts []grpc.CallOption, rtopt ret.Opt) *grpns {
	return &grpns{dopts: dopts, copts: copts, rtopt: rtopt}
}
