package fly

import (
	"github.com/andrerrcosta2/vstr/nwk"
	"github.com/andrerrcosta2/vstr/pb"
	"github.com/stretchr/testify/mock"
)

type mnms struct {
	mock.Mock
}

func (m *mnms) Qs(msg *pb.Msg, n nwk.NwkAddr) (pb.Msg, error) {
	args := m.Called(msg, n)
	return args.Get(0).(pb.Msg), args.Error(1)
}

func (m *mnms) Qm(msg *pb.Msg, n ...nwk.NwkAddr) []pb.Msg {
	args := m.Called(msg, n)
	return args.Get(0).([]pb.Msg)
}

func (m *mnms) Qms(n nwk.NwkAddr, msg ...*pb.Msg) []pb.Msg {
	args := m.Called(msg, n)
	return args.Get(0).([]pb.Msg)
}

func (m *mnms) Hchk(n nwk.NwkAddr) error {
	args := m.Called(n)
	return args.Error(0)
}
