package drl

import (
	"fmt"
	"github.com/andrerrcosta2/vstr/crd/msg"
	"github.com/andrerrcosta2/vstr/crd/nod"
	"github.com/andrerrcosta2/vstr/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestRsc(t *testing.T) {
	nms := new(mnms)

	tcs := []struct {
		msr msg.Nmr
		ref []byte
		bts *nod.Nod
		rms pb.Msg
		exp error
	}{
		{
			msr: nms,
			ref: []byte{0x00, 0xAA, 0x3F, 0x00},
			bts: &nod.Nod{Ip: "192.168.0.0", Pt: 8080},
			rms: pb.Msg{
				Type: "test-res",
				Id:   []byte{0x00, 0xAA, 0x3F, 0x00},
				Data: []byte{0x03, 0x07, 0x09, 0x2F},
			},
			exp: nil,
		},
		{
			msr: nms,
			ref: nil,
			rms: pb.Msg{
				Type: "test-res",
				Id:   []byte{0x00, 0xAA, 0x3F, 0x00},
				Data: []byte{0x03, 0x07, 0x09, 0x2F},
			},
			bts: &nod.Nod{Ip: "192.168.0.2", Pt: 8081},
			exp: fmt.Errorf("ref cannot be nil"),
		},
		{
			msr: nil,
			ref: []byte{0x00, 0xAA, 0x3F, 0x00},
			bts: &nod.Nod{Ip: "192.168.0.2", Pt: 8081},
			rms: pb.Msg{
				Type: "test-res",
				Id:   []byte{0x00, 0xAA, 0x3F, 0x00},
				Data: []byte{0x03, 0x07, 0x09, 0x2F},
			},
			exp: fmt.Errorf("msr cannot be nil"),
		},
		{
			msr: nms,
			ref: []byte{0x00, 0xAA, 0x3F, 0x00},
			rms: pb.Msg{
				Type: "test-res",
				Id:   []byte{0x00, 0xAA, 0x3F, 0x00},
				Data: []byte{0x03, 0x07, 0x09, 0x2F},
			},
			bts: nil,
			exp: fmt.Errorf("bts cannot be nil"),
		},
	}

	for _, tc := range tcs {
		nms.On("Qs", mock.Anything, mock.AnythingOfType("*nod.Nod")).Return(tc.rms, nil)
		rsc, err := Rsc(tc.msr, tc.ref, tc.bts)

		if tc.exp != nil {
			assert.Equal(t, pb.Msg{}, rsc)
			assert.EqualError(t, err, tc.exp.Error())
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tc.rms, rsc)
		}
	}
	nms.AssertExpectations(t)
}

func TestNjn(t *testing.T) {
	nms := new(mnms)

	tcs := []struct {
		msr msg.Nmr
		jn  *nod.Nod
		suc *nod.Nod
		rms pb.Msg
		exp error
	}{
		{
			msr: nms,
			jn:  &nod.Nod{Ip: "192.168.0.0", Pt: 8080},
			suc: &nod.Nod{Ip: "192.168.0.1", Pt: 8081},
			rms: pb.Msg{
				Type: "test-res",
				Id:   []byte{0x00, 0xAA, 0x3F, 0x00},
				Data: []byte{0x03, 0x07, 0x09, 0x2F},
			},
			exp: nil,
		},
		{
			msr: nms,
			jn:  nil,
			suc: &nod.Nod{Ip: "192.168.0.1", Pt: 8081},
			rms: pb.Msg{
				Type: "test-res",
				Id:   []byte{0x00, 0xAA, 0x3F, 0x00},
				Data: []byte{0x03, 0x07, 0x09, 0x2F},
			},
			exp: fmt.Errorf("jn cannot be nil\n"),
		},
		{
			msr: nms,
			jn:  &nod.Nod{Ip: "192.168.0.0", Pt: 8080},
			suc: nil,
			rms: pb.Msg{
				Type: "test-res",
				Id:   []byte{0x00, 0xAA, 0x3F, 0x00},
				Data: []byte{0x03, 0x07, 0x09, 0x2F},
			},
			exp: fmt.Errorf("suc cannot be nil\n"),
		},
		{
			msr: nil,
			jn:  &nod.Nod{Ip: "192.168.0.0", Pt: 8080},
			suc: &nod.Nod{Ip: "192.168.0.1", Pt: 8081},
			rms: pb.Msg{
				Type: "test-res",
				Id:   []byte{0x00, 0xAA, 0x3F, 0x00},
				Data: []byte{0x03, 0x07, 0x09, 0x2F},
			},
			exp: fmt.Errorf("msr cannot be nil\n"),
		},
	}

	for _, tc := range tcs {
		if tc.exp == nil {
			nms.On("Qs", mock.Anything, mock.AnythingOfType("*nod.Nod")).Return(tc.rms, nil)
		}
		njn, err := Njn(tc.msr, tc.jn, tc.suc)

		if tc.exp != nil {
			assert.Equal(t, pb.Msg{}, njn)
			assert.EqualError(t, err, tc.exp.Error())
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tc.rms, njn)
		}
	}
	nms.AssertExpectations(t)
}
