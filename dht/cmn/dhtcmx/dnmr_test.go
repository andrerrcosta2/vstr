package dhtcmx

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/andrerrcosta2/vstr/dht/cmn/cmcod"
	"github.com/andrerrcosta2/vstr/dht/cmn/dhtms"
	"github.com/andrerrcosta2/vstr/dht/crd/cid"
	"github.com/andrerrcosta2/vstr/msg/tpl"
	"github.com/andrerrcosta2/vstr/nmm/hchk"
	"github.com/andrerrcosta2/vstr/nwk/nwu"
	"github.com/andrerrcosta2/vstr/srd/cvt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net"
	"sort"

	"github.com/andrerrcosta2/vstr/nwk"
	"testing"
)

func TestHchk(t1 *testing.T) {
	thckr := newMhckr()
	dnmr := NewDnmr(newMcnmx(), thckr, 0)
	tcs := []struct {
		nm   string
		addr nwk.NwkAddr
		exp  error
		out  error
	}{
		{
			nm:   "Should return nil when host is reachable",
			addr: nwk.NewSaddr("89.0.142.86"),
			exp:  nil,
			out:  nil,
		},
		{
			nm:   "Should return error when addr is invld",
			addr: nwk.NewSaddr("89.0.142.86.19b"),
			exp:  errors.New("invld addr\n"),
			out:  errors.New("invld addr\n"),
		},
		{
			nm:   "Should return error when host is nt reachable",
			addr: nwk.NewSaddr("89.0.142.87"),
			exp:  errors.New("host unreachable\n"),
			out:  errors.New("host unreachable\n"),
		},
	}
	for _, tc := range tcs {
		t1.Run(tc.nm, func(t1 *testing.T) {
			thckr.On("Chk", tc.addr).Return(tc.out)
			err := dnmr.Hchk(tc.addr)
			if err != nil && tc.exp != nil {
				if err.Error() != tc.exp.Error() {
					t1.Errorf("\n{\n\tHchk() error:\n\n\texp:%v\tret:%v\n}", tc.exp, err)
				}
			} else if err != tc.exp {
				t1.Errorf("\n{\n\tHchk() error:\n\n\texp:%v\tret:%v\n}", tc.exp, err)
			}
		})
	}
}

func TestQm(t1 *testing.T) {
	tcnmx := newMcnmx()
	dnmr := NewDnmr(tcnmx, newMhckr(), 0)
	tcs := []struct {
		nm   string
		msg  *dhtms.Nms
		addr []nwk.NwkAddr
		exp  []*dhtms.Nms
		out  []tplout
	}{
		{
			nm: "Should ret a msg slc when host is reachable",
			addr: []nwk.NwkAddr{
				nwk.NewSaddr("89.0.142.86:8080"),
				nwk.NewSaddr("89.0.142.87:8080"),
				nwk.NewSaddr("89.0.142.88:8080"),
			},
			msg: dhtms.NewNms(cid.New(net.IP{127, 0, 0, 1}, 8080).Byt(), nwu.DNM_SV, nil),
			exp: []*dhtms.Nms{
				dhtms.NewNms(cid.New(net.ParseIP("89.0.142.86"), 8080).Byt(), nwu.DNM_SVR, nil),
				dhtms.NewNms(cid.New(net.ParseIP("89.0.142.87"), 8080).Byt(), nwu.DNM_SVR, nil),
				dhtms.NewNms(cid.New(net.ParseIP("89.0.142.88"), 8080).Byt(), nwu.DNM_SVR, nil),
			},
			out: []tplout{
				{dhtms.NewNms(cid.New(net.ParseIP("89.0.142.86"), 8080).Byt(), nwu.DNM_SVR, nil), nil},
				{dhtms.NewNms(cid.New(net.ParseIP("89.0.142.87"), 8080).Byt(), nwu.DNM_SVR, nil), nil},
				{dhtms.NewNms(cid.New(net.ParseIP("89.0.142.88"), 8080).Byt(), nwu.DNM_SVR, nil), nil},
			},
		},
		{
			nm: "Should retrn a msg slc containingn invld addr error",
			addr: []nwk.NwkAddr{
				nwk.NewSaddr("89.0.142.89.19p:8080"),
				nwk.NewSaddr("89.0.142.90:8080"),
				nwk.NewSaddr("89.0.142.91:8080"),
			},
			msg: dhtms.NewNms(cid.New(net.IP{127, 0, 0, 1}, 8080).Byt(), nwu.DNM_SV, nil),
			exp: []*dhtms.Nms{
				dhtms.NewNms(nil, cmcod.RES_GCE, []byte(errors.New("invld addr\n").Error())),
				dhtms.NewNms(cid.New(net.ParseIP("89.0.142.90"), 8080).Byt(), nwu.DNM_SVR, nil),
				dhtms.NewNms(cid.New(net.ParseIP("89.0.142.91"), 8080).Byt(), nwu.DNM_SVR, nil),
			},
			out: []tplout{
				{nil, errors.New("invld addr\n")},
				{dhtms.NewNms(cid.New(net.ParseIP("89.0.142.90"), 8080).Byt(), nwu.DNM_SVR, nil), nil},
				{dhtms.NewNms(cid.New(net.ParseIP("89.0.142.91"), 8080).Byt(), nwu.DNM_SVR, nil), nil},
			},
		},
		{
			nm:  "Should return a msg slc containing no gtw addr error",
			msg: dhtms.NewNms(cid.New(net.IP{127, 0, 0, 1}, 8080).Byt(), nwu.DNM_SV, nil),
			addr: []nwk.NwkAddr{
				nwk.NewSaddr("89.0.142.92:8080"),
				nil,
				nwk.NewSaddr("89.0.142.93:8080"),
			},
			exp: []*dhtms.Nms{
				dhtms.NewNms(cid.New(net.ParseIP("89.0.142.92"), 8080).Byt(), nwu.DNM_SS, nil),
				dhtms.NewNms(nil, cmcod.RES_GCE, []byte(errors.New("no gtw addr\n").Error())),
				dhtms.NewNms(cid.New(net.ParseIP("89.0.142.93"), 8080).Byt(), nwu.DNM_SS, nil),
			},
			out: []tplout{
				{dhtms.NewNms(cid.New(net.ParseIP("89.0.142.92"), 8080).Byt(), nwu.DNM_SS, nil), nil},
				{nil, nil},
				{dhtms.NewNms(cid.New(net.ParseIP("89.0.142.93"), 8080).Byt(), nwu.DNM_SS, nil), nil},
			},
		},
	}
	for _, tc := range tcs {
		t1.Run(tc.nm, func(t1 *testing.T) {
			for i := 0; i < len(tc.addr); i++ {
				if tc.addr[i] != nil {
					tcnmx.On("Sm", tc.msg, tc.addr[i].Gfaddr()).Return(tc.out[i].res, tc.out[i].err)
				}
			}
			msgs := dnmr.Qm(tc.msg, tc.addr...)

			srtNms(msgs)
			srtNms(tc.exp)
			assert.Equal(t1, tc.exp, msgs)
		})
	}
}

func TestQmSmph(t1 *testing.T) {
	const lps = 1000
	tcnmx := newMcnmx()
	dnmr := NewDnmr(tcnmx, newMhckr(), 0)
	addrs := make([]nwk.NwkAddr, lps)

	for i := 0; i < lps; i++ {
		addrs[i] = nwk.NewSaddr(fmt.Sprintf("192.168.0.%d:8080", i))
	}

	for i, addr := range addrs {
		if i%3 != 0 {
			tcnmx.On("Sm", dhtms.NewNms(cid.New(net.IP{127, 0, 0, 1}, 8080).Byt(), cmcod.REQ_NP, nil),
				addr.Gfaddr()).Return(dhtms.NewNms(cid.New(net.ParseIP(
				fmt.Sprintf("192.168.0.%d:8080", i)),
				8080).Byt(), cmcod.RES_SS, cvt.Itb(i)),
				nil)
		} else {
			tcnmx.On("Sm", dhtms.NewNms(cid.New(net.IP{127, 0, 0, 1}, 8080).Byt(), cmcod.REQ_NP, nil), addr.Gfaddr()).Return(nil, errors.New("some error"))
		}
	}

	nms := dnmr.Qm(dhtms.NewNms(cid.New(net.IP{127, 0, 0, 1}, 8080).Byt(), cmcod.REQ_NP, nil), addrs...)
	assert.Equal(t1, lps, len(nms))

	for _, nm := range nms {
		if len(nm.Id) == 0 {
			assert.Equal(t1, dhtms.NewNms(nil, cmcod.RES_GCE, []byte(errors.New("some error").Error())), nm)
		} else {
			assert.Equal(t1, len(nm.Data) > 0, true)
			v, err := cvt.Bti(nm.Data)
			if err != nil {
				t1.Fatalf("cannot read nm.Data: %s", err)
			}
			assert.Equal(t1, v%3 != 0, true)
		}
	}
}

func TestQms(t1 *testing.T) {
	tcnmx := newMcnmx()
	dnmr := NewDnmr(tcnmx, newMhckr(), 0)
	tcs := []struct {
		nm   string
		msg  []*dhtms.Nms
		addr nwk.NwkAddr
		exp  []*dhtms.Nms
		out  []tplout
	}{
		{
			nm:   "Should ret a msg slc when host is reachable",
			addr: nwk.NewSaddr("89.0.142.86"),
			msg: []*dhtms.Nms{
				dhtms.NewNms(cid.New(net.ParseIP("127.0.0.1"), 8080).Byt(), nwu.DNM_SV, nil),
				dhtms.NewNms(cid.New(net.ParseIP("127.0.0.1"), 8080).Byt(), nwu.DNM_SV, nil),
				dhtms.NewNms(cid.New(net.ParseIP("127.0.0.1"), 8080).Byt(), nwu.DNM_SV, nil),
			},
			exp: []*dhtms.Nms{
				dhtms.NewNms(cid.New(net.ParseIP("89.0.142.86"), 8080).Byt(), nwu.DNM_SVR, nil),
				dhtms.NewNms(cid.New(net.ParseIP("89.0.142.86"), 8080).Byt(), nwu.DNM_SVR, nil),
				dhtms.NewNms(cid.New(net.ParseIP("89.0.142.86"), 8080).Byt(), nwu.DNM_SVR, nil),
			},
			out: []tplout{
				{dhtms.NewNms(cid.New(net.ParseIP("89.0.142.86"), 8080).Byt(), nwu.DNM_SVR, nil), nil},
				{dhtms.NewNms(cid.New(net.ParseIP("89.0.142.86"), 8080).Byt(), nwu.DNM_SVR, nil), nil},
				{dhtms.NewNms(cid.New(net.ParseIP("89.0.142.86"), 8080).Byt(), nwu.DNM_SVR, nil), nil},
			},
		},
		{
			nm:   "Should retrn a msg slc containingn invld addr errors",
			addr: nwk.NewSaddr("89.0.142.86.19p"),
			msg: []*dhtms.Nms{
				dhtms.NewNms(cid.New(net.ParseIP("127.0.0.1"), 8080).Byt(), nwu.DNM_SV, nil),
				dhtms.NewNms(cid.New(net.ParseIP("127.0.0.1"), 8080).Byt(), nwu.DNM_SV, nil),
				dhtms.NewNms(cid.New(net.ParseIP("127.0.0.1"), 8080).Byt(), nwu.DNM_SV, nil),
			},
			exp: []*dhtms.Nms{
				dhtms.NewNms(nil, cmcod.RES_GCE, []byte(errors.New("invld addr\n").Error())),
				dhtms.NewNms(nil, cmcod.RES_GCE, []byte(errors.New("invld addr\n").Error())),
				dhtms.NewNms(nil, cmcod.RES_GCE, []byte(errors.New("invld addr\n").Error())),
			},
			out: []tplout{
				{nil, errors.New("invld addr\n")},
				{nil, errors.New("invld addr\n")},
				{nil, errors.New("invld addr\n")},
			},
		},
		{
			nm:   "Should retrn a msg slc containingn no gtw addr error",
			addr: nil,
			msg: []*dhtms.Nms{
				dhtms.NewNms(cid.New(net.ParseIP("127.0.0.1"), 8080).Byt(), nwu.DNM_SV, nil),
				dhtms.NewNms(cid.New(net.ParseIP("127.0.0.1"), 8080).Byt(), "urcgzmsg", nil),
				dhtms.NewNms(cid.New(net.ParseIP("127.0.0.1"), 8080).Byt(), nwu.DNM_SV, nil),
			},
			exp: []*dhtms.Nms{
				dhtms.NewNms(nil, cmcod.RES_GCE, []byte(errors.New("no gtw addr\n").Error())),
				dhtms.NewNms(nil, cmcod.RES_GCE, []byte(errors.New("no gtw addr\n").Error())),
				dhtms.NewNms(nil, cmcod.RES_GCE, []byte(errors.New("no gtw addr\n").Error())),
			},
			out: []tplout{},
		},
		{
			nm:   "Should retrn a msg slc containingn a unrm msg error",
			addr: nwk.NewSaddr("89.0.142.87"),
			msg: []*dhtms.Nms{
				dhtms.NewNms(cid.New(net.ParseIP("127.0.0.1"), 8080).Byt(), nwu.DNM_SV, nil),
				dhtms.NewNms(cid.New(net.ParseIP("127.0.0.1"), 8080).Byt(), "urcgzmsg", nil),
				dhtms.NewNms(cid.New(net.ParseIP("127.0.0.1"), 8080).Byt(), nwu.DNM_SV, nil),
			},
			exp: []*dhtms.Nms{
				dhtms.NewNms(cid.New(net.ParseIP("89.0.142.87"), 8080).Byt(), nwu.DNM_SVR, nil),
				dhtms.NewNms(cid.New(net.ParseIP("89.0.142.87"), 8080).Byt(), nwu.DNM_UNRM, []byte(errors.New("DNM-UNRM\n").Error())),
				dhtms.NewNms(cid.New(net.ParseIP("89.0.142.87"), 8080).Byt(), nwu.DNM_SVR, nil),
			},
			out: []tplout{
				{dhtms.NewNms(cid.New(net.ParseIP("89.0.142.87"), 8080).Byt(), nwu.DNM_SVR, nil), nil},
				{dhtms.NewNms(cid.New(net.ParseIP("89.0.142.87"), 8080).Byt(), nwu.DNM_UNRM, []byte(errors.New("DNM-UNRM\n").Error())), nil},
				{dhtms.NewNms(cid.New(net.ParseIP("89.0.142.87"), 8080).Byt(), nwu.DNM_SVR, nil), nil},
			},
		},
	}
	for _, tc := range tcs {
		t1.Run(tc.nm, func(t1 *testing.T) {
			for i := 0; i < len(tc.msg); i++ {
				if tc.addr != nil {
					tcnmx.On("Sm", tc.msg[i], tc.addr.Gfaddr()).Return(tc.out[i].res, tc.out[i].err)
				}
			}
			msgs := dnmr.Qms(tc.addr, tc.msg...)

			//for i := 0; i < len(msgs); i++ {
			//	fmt.Printf("\n{\n\tQms() error:\n\n\texp:%+v\n\tret:%+v\n}", tc.exp[i], msgs[i])
			//}

			srtNms(msgs)
			srtNms(tc.exp)
			assert.Equal(t1, tc.exp, msgs)
		})
	}
}

func TestQmsSmph(t1 *testing.T) {
	const lps = 1000
	tcnmx := newMcnmx()
	dnmr := NewDnmr(tcnmx, newMhckr(), 0)
	msgs := make([]*dhtms.Nms, lps)

	for i := 0; i < lps; i++ {
		msgs[i] = dhtms.NewNms(cid.New(net.IP{127, 0, 0, 1}, 8080).Byt(), cmcod.REQ_NP, cvt.Itb(i))
	}

	for i, msg := range msgs {
		if i%3 != 0 {
			tcnmx.On("Sm", msg,
				net.ParseIP("89.0.142.86")).Return(dhtms.NewNms(cid.New(net.ParseIP("89.0.142.86"),
				8080).Byt(), cmcod.RES_SS, cvt.Itb(i)),
				nil)
		} else {
			tcnmx.On("Sm", mock.Anything, mock.Anything).Return(nil, errors.New("some error"))
		}
	}

	nms := dnmr.Qms(nwk.NewSaddr("89.0.142.86"), msgs...)
	assert.Equal(t1, lps, len(nms))

	for _, nm := range nms {
		if len(nm.Id) == 0 {
			assert.Equal(t1, dhtms.NewNms(nil, cmcod.RES_GCE, []byte(errors.New("some error").Error())), nm)
		} else {
			assert.Equal(t1, len(nm.Data) > 0, true)
			v, err := cvt.Bti(nm.Data)
			if err != nil {
				t1.Fatalf("cannot read nm.Data: %s", err)
			}
			assert.Equal(t1, v%3 != 0, true)
		}
	}
}

func TestQs(t1 *testing.T) {
	tcnmx := newMcnmx()
	dnmr := NewDnmr(tcnmx, newMhckr(), 0)
	tcs := []struct {
		nm   string
		msg  *dhtms.Nms
		addr nwk.NwkAddr
		exp  tplout
		out  tplout
	}{
		{
			nm:   "Should ret a success msg when host is reachable",
			addr: nwk.NewSaddr("89.0.142.86:8080"),
			msg:  dhtms.NewNms(cid.New(net.ParseIP("127.0.0.1"), 8080).Byt(), nwu.DNM_SV, nil),
			exp: tplout{
				dhtms.NewNms(cid.New(net.ParseIP("89.0.142.86"), 8080).Byt(), nwu.DNM_SVR, nil),
				nil,
			},
			out: tplout{
				res: dhtms.NewNms(cid.New(net.ParseIP("89.0.142.86"), 8080).Byt(), nwu.DNM_SVR, nil),
				err: nil,
			},
		},
		{
			nm:   "Should retrn a msg slc containingn invld addr errors",
			addr: nwk.NewSaddr("89.0.142.86.19p:8080"),
			msg:  dhtms.NewNms(cid.New(net.ParseIP("127.0.0.1"), 8080).Byt(), nwu.DNM_SV, nil),
			exp:  tplout{res: nil, err: errors.New("invld addr\n")},
			out:  tplout{res: nil, err: errors.New("invld addr\n")},
		},
		{
			nm:   "Should retrn a msg slc containingn no gtw addr error",
			addr: nil,
			msg:  dhtms.NewNms(nil, nwu.DNM_SV, nil),
			exp: tplout{
				nil,
				errors.New("no gtw addr\n"),
			},
			out: tplout{},
		},
		{
			nm:   "Should retrn a msg slc containingn a unrm msg error",
			addr: nwk.NewSaddr("89.0.142.87:8080"),
			msg:  dhtms.NewNms(cid.New(net.ParseIP("127.0.0.1"), 8080).Byt(), "urcgzmsg", nil),
			exp:  tplout{res: dhtms.NewNms(cid.New(net.ParseIP("89.0.142.87"), 8080).Byt(), nwu.DNM_UNRM, []byte(errors.New("DNM-UNRM\n").Error())), err: nil},
			out:  tplout{res: dhtms.NewNms(cid.New(net.ParseIP("89.0.142.87"), 8080).Byt(), nwu.DNM_UNRM, []byte(errors.New("DNM-UNRM\n").Error())), err: nil},
		},
	}
	for _, tc := range tcs {
		t1.Run(tc.nm, func(t1 *testing.T) {
			if tc.addr != nil {
				tcnmx.On("Sm", tc.msg, tc.addr.Gfaddr()).Return(tc.out.res, tc.out.err)
			} else {
				tcnmx.On("Sm", tc.msg, nil).Return(tc.out.res, tc.out.err)
			}
			msg, err := dnmr.Qs(tc.msg, tc.addr)

			//fmt.Printf("\n{\n\tQs() error:\n\n\texp-err:%+v\tret:%+v\n\texp-res:%+v\n\tret:%+v\n}", tc.exp.err, err, tc.exp.res, msg)

			assert.Equal(t1, tc.exp.res, msg)
			assert.Equal(t1, tc.exp.err, err)
		})
	}
}

func TestNewDnmr(t *testing.T) {
	tcs := []struct {
		nm   string
		tpl  tpl.Tpl[*dhtms.Nms, *dhtms.Nms]
		hchk hchk.Hchkr[nwk.NwkAddr]
		exp  *Dnmr
	}{
		{
			nm:   "valid tpl and hchk",
			tpl:  newMcnmx(),
			hchk: newMhckr(),
			exp: &Dnmr{
				tpl:  newMcnmx(),
				hchk: newMhckr(),
			},
		},
		{
			nm:   "nil tpl and hchk",
			tpl:  nil,
			hchk: nil,
			exp: &Dnmr{
				tpl:  nil,
				hchk: nil,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.nm, func(t *testing.T) {
			result := NewDnmr(tc.tpl, tc.hchk, 0)
			assert.Equal(t, tc.exp, result)
		})
	}
}

type tplout struct {
	res *dhtms.Nms
	err error
}

func srtNms(msgs []*dhtms.Nms) {
	sort.Slice(msgs, func(i, j int) bool {
		cmp := bytes.Compare(msgs[i].Id, msgs[j].Id)
		if cmp != 0 {
			return cmp < 0
		}
		cmp = bytes.Compare([]byte(msgs[i].Type), []byte(msgs[j].Type))
		if cmp != 0 {
			return cmp < 0
		}
		return bytes.Compare(msgs[i].Data, msgs[j].Data) < 0
	})
}
