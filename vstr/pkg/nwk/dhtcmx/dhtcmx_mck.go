// Andre R. R. Costa *** github.com/andrerrcosta2 *** andrerrcosta@gmail.com

package nwk

import (
	"github.com/andrerrcosta2/vstr/vstr/pkg/dht/cmn/dhtms"
	"github.com/andrerrcosta2/vstr/vstr/pkg/msg/msgc"
	"github.com/andrerrcosta2/vstr/vstr/pkg/msg/msrf"
	"github.com/andrerrcosta2/vstr/vstr/pkg/ntk"
	"github.com/andrerrcosta2/vstr/vstr/pkg/ntk/dlr"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

// Nmrmck: mock for tpl.Tpl
type Nmrmck interface {
	Sm(msg *dhtms.Nms, addr string) (*dhtms.Nms, error)
}

type mcnmx struct {
	mock.Mock
}

func (mck *mcnmx) Sm(msg *dhtms.Nms, addr string) (*dhtms.Nms, error) {
	args := mck.Called(msg, addr)
	var result *dhtms.Nms
	if args.Get(0) != nil {
		result = args.Get(0).(*dhtms.Nms)
	}
	return result, args.Error(1)
}

func NewMcnmx() *mcnmx {
	return &mcnmx{}
}

// mhckr: mock for hchk.Hchkr
type mhckr struct {
	mock.Mock
}

func (mck *mhckr) Chk(addr ntk.NwkAddr) error {
	args := mck.Called(addr)
	return args.Error(0)
}

func NewMhckr() *mhckr {
	return &mhckr{}
}

type Mgrpcf[T msgc.Rpcc] struct {
	mock.Mock
}

func (mck *Mgrpcf[T]) Gin(addr string, fct func(conn dlr.RpcConn) (T, error)) (msgc.Rpcci[T, grpc.CallOption], error) {
	a := mck.Called(addr, fct)
	return a.Get(0).(msgc.Rpcci[T, grpc.CallOption]), a.Error(1)
}

func NewMgrpcf[T any]() msrf.Rpccf[T, grpc.CallOption] {
	return &Mgrpcf[T]{}
}

type Mrpci struct {
	mock.Mock
	ctr *gomock.Controller
}

func (mck *Mrpci) End() {
	mck.Called()
}

func (mck *Mrpci) Copts() []grpc.CallOption {
	args := mck.Called()
	return args.Get(0).([]grpc.CallOption)
}

func (mck *Mrpci) Cli() *MockGrpcnmsClient {
	return NewMockGrpcnmsClient(mck.ctr)
}

type Mgrpcnms struct {
	mock.Mock
}

func (mck *Mgrpcnms) Snm(msg *dhtms.Nms) (*dhtms.Nms, error) {
	args := mck.Called(msg)
	return args.Get(0).(*dhtms.Nms), args.Error(1)
}
