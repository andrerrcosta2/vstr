// Andre R. R. Costa *** github.com/andrerrcosta2 *** andrerrcosta@gmail.com

package dhtcmx

import (
	"github.com/andrerrcosta2/vstr/dht/cmn/dhtms"
	"github.com/andrerrcosta2/vstr/nwk"
	"github.com/stretchr/testify/mock"
)

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

func newMcnmx() *mcnmx {
	return &mcnmx{}
}

type mhckr struct {
	mock.Mock
}

func (mck *mhckr) Chk(addr nwk.NwkAddr) error {
	args := mck.Called(addr)
	return args.Error(0)
}

func newMhckr() *mhckr {
	return &mhckr{}
}
