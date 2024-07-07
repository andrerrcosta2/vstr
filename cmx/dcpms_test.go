// github.com/andrerrcosta2 *** andrerrcosta@gmail.com

package cmx

import (
	"fmt"
	"github.com/andrerrcosta2/vstr/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestSpm(t *testing.T) {
	var cpmx *Dcpms

	tcs := []struct {
		msg  proto.Message
		addr string
		exp  error
	}{
		{
			msg: &pb.Msg{
				Type: "test-msg",
				Id:   []byte{0xFF, 0x02, 0x03, 0x04},
				Data: []byte{0x04, 0x88, 0xAD, 0xE4},
			},
			addr: "192.168.0.1",
			exp:  nil,
		},
	}

	for _, tc := range tcs {
		cpmx = buildCpmx(tc.msg)
		spm, err := cpmx.Spm(tc.msg, tc.addr)
		if tc.exp != nil {
			assert.Error(t, err)
		} else {
			fmt.Printf("\nMessage: %v\n", spm)
			assert.NoError(t, err)
			assert.NotNil(t, spm)
		}
	}
}

func buildCpmx(res proto.Message) *Dcpms {
	rd, _ := proto.Marshal(res)
	return NewDcpmx(&mdlr{
		Mock: mock.Mock{},
		Conn: &mcon{
			ReadData:  rd,
			WriteData: nil,
		},
		Err: nil,
	})
}
