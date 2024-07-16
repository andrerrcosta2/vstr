// github.com/andrerrcosta2 *** andrerrcosta@gmail.com

package msg

import (
	"fmt"
	"github.com/andrerrcosta2/vstr/nwk/dlr"
	"github.com/andrerrcosta2/vstr/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestSpm(t *testing.T) {
	tcs := []struct {
		msg  proto.Message
		addr string
		dlr  dlr.Dlr
		exp  error
	}{
		{
			msg: &pb.Msg{
				Type: "test-crdnmr",
				Id:   []byte{0xFF, 0x02, 0x03, 0x04},
				Data: []byte{0x04, 0x88, 0xAD, 0xE4},
			},
			addr: "192.168.0.1",
			dlr:  buildMockDialer(nil, nil),
			exp:  nil,
		},
		{
			msg:  nil,
			addr: "192.168.0.1",
			dlr:  buildMockDialer(nil, nil),
			exp:  fmt.Errorf("message is nil\n"),
		},
		{
			msg: &pb.Msg{
				Type: "test-crdnmr",
				Id:   []byte{0xFF, 0x02, 0x03, 0x04},
				Data: []byte{0x04, 0x88, 0xAD, 0xE4},
			},
			addr: "",
			dlr:  buildMockDialer(nil, nil),
			exp:  fmt.Errorf("address is empty\n"),
		},
		{
			msg: &pb.Msg{
				Type: "test-crdnmr",
				Id:   []byte{0xFF, 0x02, 0x03, 0x04},
				Data: []byte{0x04, 0x88, 0xAD, 0xE4},
			},
			addr: "192.168.0.1",
			dlr:  nil,
			exp:  fmt.Errorf("dialer is nil\n"),
		},
	}

	for _, tc := range tcs {
		cpmx := &dcpmx{dlr: tc.dlr}
		spm, err := cpmx.Spm(tc.msg, tc.addr)

		if tc.exp != nil {
			assert.Error(t, err)
			assert.EqualError(t, err, tc.exp.Error())
		} else {
			assert.NoError(t, err)
			assert.NotNil(t, spm)
		}
	}
}

func buildMockDialer(res proto.Message, err error) dlr.Dlr {
	return &mdlr{
		Mock: mock.Mock{},
		Conn: &mcon{
			rd: func() []byte {
				if res != nil {
					rd, _ := proto.Marshal(res)
					return rd
				}
				return nil
			}(),
			wd: nil,
		},
		Err: err,
	}
}
