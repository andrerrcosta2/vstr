package dhtcmx

import (
	"fmt"
	"github.com/andrerrcosta2/vstr/dht/cmn/dhtms"
	"github.com/andrerrcosta2/vstr/pb"
	"github.com/andrerrcosta2/vstr/srd/ret"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
)

func TestNewGrpns(t *testing.T) {
	type args struct {
		dopts []grpc.DialOption
		copts []grpc.CallOption
		rtopt ret.Opt
	}
	tests := []struct {
		name string
		args args
		want *grpns
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewGrpns(tt.args.dopts, tt.args.copts, tt.args.rtopt), "NewGrpns(%v, %v, %v)", tt.args.dopts, tt.args.copts, tt.args.rtopt)
		})
	}
}

func Test_cls(t *testing.T) {
	type args struct {
		cnc *grpc.ClientConn
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls(tt.args.cnc)
		})
	}
}

func Test_grpns_Sm(t *testing.T) {
	type fields struct {
		dopts []grpc.DialOption
		copts []grpc.CallOption
		rtopt ret.Opt
	}
	type args struct {
		n    *dhtms.Nms
		addr string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dhtms.Nms
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &grpns{
				dopts: tt.fields.dopts,
				copts: tt.fields.copts,
				rtopt: tt.fields.rtopt,
			}
			got, err := s.Sm(tt.args.n, tt.args.addr)
			if !tt.wantErr(t, err, fmt.Sprintf("Sm(%v, %v)", tt.args.n, tt.args.addr)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Sm(%v, %v)", tt.args.n, tt.args.addr)
		})
	}
}

func Test_grpns_gc(t *testing.T) {
	type fields struct {
		dopts []grpc.DialOption
		copts []grpc.CallOption
		rtopt ret.Opt
	}
	type args struct {
		addr string
		opts []grpc.DialOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    pb.GrpcnmsClient
		want1   *grpc.ClientConn
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &grpns{
				dopts: tt.fields.dopts,
				copts: tt.fields.copts,
				rtopt: tt.fields.rtopt,
			}
			got, got1, err := s.gc(tt.args.addr, tt.args.opts...)
			if !tt.wantErr(t, err, fmt.Sprintf("gc(%v, %v)", tt.args.addr, tt.args.opts...)) {
				return
			}
			assert.Equalf(t, tt.want, got, "gc(%v, %v)", tt.args.addr, tt.args.opts...)
			assert.Equalf(t, tt.want1, got1, "gc(%v, %v)", tt.args.addr, tt.args.opts...)
		})
	}
}
