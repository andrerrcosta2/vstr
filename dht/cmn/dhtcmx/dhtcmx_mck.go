// Andre R. R. Costa *** github.com/andrerrcosta2 *** andrerrcosta@gmail.com

package dhtcmx

import (
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
)

type mcpmx struct {
	mock.Mock
}

func (m *mcpmx) Spm(msg proto.Message, addr string) (proto.Message, error) {
	args := m.Called(msg, addr)
	return args.Get(0).(proto.Message), args.Error(1)
}
