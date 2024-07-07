// Andre R. R. Costa *** github.com/andrerrcosta2 *** andrerrcosta@gmail.com

package msg

import (
	"fmt"
	"github.com/andrerrcosta2/vstr/crd/nod"
	"github.com/andrerrcosta2/vstr/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestQs(t *testing.T) {
	cpmx := new(mcpmx)
	tcpNms := &TcpNms{cpmx: cpmx}

	tcs := []struct {
		req pb.Msg
		res pb.Msg
		nod nod.Nod
		exp error
	}{
		{
			req: pb.Msg{
				Type: "test-msg",
				Id:   []byte{0xFF, 0x02, 0x03, 0x04},
				Data: []byte{0x04, 0x88, 0xAD, 0xE4},
			},
			res: pb.Msg{
				Type: "test-res",
				Id:   []byte{0x01, 0x02, 0x03, 0x04},
				Data: []byte{0xFF, 0xAA, 0xBB, 0xCC},
			},
			nod: nod.Nod{
				Ip: "127.0.0.1",
				Pt: 8080,
			},
			exp: nil,
		},
	}

	for _, tc := range tcs {
		cpmx.On("Spm", mock.Anything, mock.AnythingOfType("string")).Return(&tc.res, nil)
		res, err := tcpNms.Qs(&tc.req, &tc.nod)
		assert.Equal(t, tc.res, res)
		if tc.exp != nil {
			assert.Error(t, err, tc.exp)
		} else {
			assert.NoError(t, err)
		}
	}
	cpmx.AssertExpectations(t)
}

func TestQm(t *testing.T) {
	cpmx := new(mcpmx)
	tcpNms := &TcpNms{cpmx: cpmx}

	tcs := []struct {
		req pb.Msg
		res []pb.Msg
		nod []*nod.Nod
		exp error
	}{
		{
			req: pb.Msg{
				Type: "test-msg",
				Id:   []byte{0xFF, 0x02, 0x03, 0x04},
				Data: []byte{0x04, 0x88, 0xAD, 0xE4},
			},
			res: []pb.Msg{
				{
					Type: "test-res",
					Id:   []byte{0x01, 0x02, 0x03, 0x04},
					Data: []byte{0xFF, 0xAA, 0xBB, 0xCC},
				},
				{
					Type: "test-res",
					Id:   []byte{0x01, 0x02, 0x03, 0x04},
					Data: []byte{0xFF, 0xAA, 0xBB, 0xCC},
				},
				{
					Type: "test-res",
					Id:   []byte{0x01, 0x02, 0x03, 0x04},
					Data: []byte{0xFF, 0xAA, 0xBB, 0xCC},
				},
			},
			nod: []*nod.Nod{
				{Ip: "127.0.0.1", Pt: 8080},
				{Ip: "127.0.0.2", Pt: 8081},
				{Ip: "127.0.0.3", Pt: 8082},
			},
		},
	}

	for _, tc := range tcs {
		for i, msg := range tc.res {
			cpmx.On("Spm", mock.Anything, fmt.Sprintf("%s:%d", tc.nod[i].Ip, tc.nod[i].Pt)).Return(&msg, nil)
		}

		res := tcpNms.Qm(&tc.req, tc.nod...)
		assert.Len(t, res, len(tc.res))
		for i, msg := range res {
			if msg.Type == "RES_GCE" {
				assert.Contains(t, string(msg.Data), "error")
			} else {
				assert.Equal(t, tc.res[i].Type, msg.Type)
			}
		}
	}
	cpmx.AssertExpectations(t)
}

func TestQms(t *testing.T) {
	cpmx := new(mcpmx)
	tcpNms := &TcpNms{cpmx: cpmx}

	tcs := []struct {
		req []*pb.Msg
		res []pb.Msg
		nod nod.Nod
		exp error
	}{
		{
			req: []*pb.Msg{
				{
					Type: "test-msg",
					Id:   []byte{0xFF, 0x02, 0x03, 0x04},
					Data: []byte{0x04, 0x88, 0xAD, 0xE4},
				},
				{
					Type: "test-msg",
					Id:   []byte{0xFF, 0x02, 0x03, 0x04},
					Data: []byte{0x04, 0x88, 0xAD, 0xE4},
				},
				{
					Type: "test-msg",
					Id:   []byte{0xFF, 0x02, 0x03, 0x04},
					Data: []byte{0x04, 0x88, 0xAD, 0xE4},
				},
			},
			res: []pb.Msg{
				{
					Type: "test-res",
					Id:   []byte{0x01, 0x02, 0x03, 0x04},
					Data: []byte{0xFF, 0xAA, 0xBB, 0xCC},
				},
				{
					Type: "test-res",
					Id:   []byte{0x01, 0x02, 0x03, 0x04},
					Data: []byte{0xFF, 0xAA, 0xBB, 0xCC},
				},
				{
					Type: "test-res",
					Id:   []byte{0x01, 0x02, 0x03, 0x04},
					Data: []byte{0xFF, 0xAA, 0xBB, 0xCC},
				},
			},
			nod: nod.Nod{Ip: "127.0.0.1", Pt: 8080},
			exp: nil,
		},
	}

	for _, tc := range tcs {
		for i, msg := range tc.res {
			cpmx.On("Spm", tc.req[i], mock.AnythingOfType("string")).Return(&msg, nil)
		}
		res := tcpNms.Qms(&tc.nod, tc.req...)
		assert.Len(t, res, len(tc.res))
		for i, msg := range res {
			if msg.Type == "RES_GCE" {
				assert.Contains(t, string(msg.Data), "error")
			} else {
				assert.Equal(t, tc.res[i].Type, msg.Type)
			}
		}
	}
	cpmx.AssertExpectations(t)
}
