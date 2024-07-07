package drl

import (
	"github.com/andrerrcosta2/vstr/crd/nod"
	"github.com/andrerrcosta2/vstr/pb"
	"github.com/stretchr/testify/mock"
)

type mnms struct {
	mock.Mock
}

func (m *mnms) Qs(msg *pb.Msg, n *nod.Nod) (pb.Msg, error) {
	args := m.Called(msg, n)
	return args.Get(0).(pb.Msg), args.Error(1)
}

func (m *mnms) Qm(msg *pb.Msg, n ...*nod.Nod) []pb.Msg {
	args := m.Called(msg, n)
	return args.Get(0).([]pb.Msg)
}

func (m *mnms) Qms(n *nod.Nod, msg ...*pb.Msg) []pb.Msg {
	args := m.Called(msg, n)
	return args.Get(0).([]pb.Msg)
}
