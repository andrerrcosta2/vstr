package dhtcmx

import (
	"github.com/andrerrcosta2/vstr/vstr/pkg/dht/cmn/dhtms"
	"github.com/andrerrcosta2/vstr/vstr/pkg/dht/crd/cid"
	"github.com/andrerrcosta2/vstr/vstr/pkg/ntk/ntku"
	"github.com/andrerrcosta2/vstr/vstr/pkg/nwk"
	"github.com/andrerrcosta2/vstr/vstr/pkg/srd/ret"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"net"
	"testing"
)

func TestSm(t *testing.T) {
	type sms struct {
		msg  *dhtms.Nms
		addr string
	}
	type smr struct {
		msg *dhtms.Nms
		err error
	}
	tcs := []struct {
		nm  string
		req sms
		res smr
		exp smr
	}{
		{
			nm:  "valid msg and addr",
			req: sms{dhtms.NewNms(cid.New(net.ParseIP("127.0.0.1"), 8080).Byt(), ntku.DNM_SV, nil), "128.0.0.1:8080"},
			res: smr{dhtms.NewNms(cid.New(net.ParseIP("127.0.0.1"), 8080).Byt(), ntku.DNM_SVR, nil), nil},
			exp: smr{dhtms.NewNms(cid.New(net.ParseIP("127.0.0.1"), 8080).Byt(), ntku.DNM_SVR, nil), nil},
		},
	}

	ctr := gomock.NewController(t)
	defer ctr.Finish()
	cli := nwk.NewMockGrpcnmsClient(ctr)

	ins := NewGrpns(nwk.NewMgrpcf[Dhtpmsr[proto.Message]](), ret.Opt[any]{})

	for _, tc := range tcs {
		t.Run(tc.nm, func(t *testing.T) {
			msg, err := ins.Sm(tc.req.msg, tc.req.addr)
			cli.EXPECT().
				Snm(gomock.Any(), tc.req.msg, gomock.Any()).Return(tc.res.msg.Pbf(), tc.res.err).Times(1)

			assert.Equal(t, tc.exp.msg, msg)
			assert.Equal(t, tc.exp.err, err)
		})
	}
}
