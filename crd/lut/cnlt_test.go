// Andre R. R. Costa *** github.com/andrerrcosta2

package lut

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/andrerrcosta/vstr/crd/cid"
	"github.com/andrerrcosta/vstr/crd/nod"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

const (
	testIP1   = "127.0.0.1"
	testPort1 = 7000
	testIP2   = "192.168.1.1"
	testPort2 = 8000
)

func TestGid(t *testing.T) {
	expectedID := sha1.Sum([]byte(fmt.Sprintf("%s:%d", testIP1, testPort1)))
	id := Gid(testIP1, testPort1)
	var idArr [20]byte
	copy(idArr[:], id[:])
	assert.Equal(t, expectedID, idArr, "Gid() should generate the correct ID")
}

func TestCnlt_AddAndGet(t *testing.T) {
	tb := make(map[string]*nod.Nod)
	tbl := &Cnlt{
		mt: sync.Mutex{},
		tb: tb,
	}

	id1 := tbl.Add(testIP1, testPort1)
	id2 := tbl.Add(testIP2, testPort2)

	node1, err := tbl.Get(id1)
	assert.NoError(t, err, "Get() should not return an error for existing node")
	assert.NotNil(t, node1, "Get() should return a non-nil node")
	assert.Equal(t, testIP1, node1.Ip, "Node IP should match the added IP")
	assert.Equal(t, int32(testPort1), node1.Pt, "Node port should match the added port")

	node2, err := tbl.Get(id2)
	assert.NoError(t, err, "Get() should not return an error for existing node")
	assert.NotNil(t, node2, "Get() should return a non-nil node")
	assert.Equal(t, testIP2, node2.Ip, "Node IP should match the added IP")
	assert.Equal(t, int32(testPort2), node2.Pt, "Node port should match the added port")

	invalidID := cid.Id(sha1.Sum([]byte("invalid:9999")))
	_, err = tbl.Get(invalidID)
	assert.Error(t, err, "Get() should return an error for non-existing node")
	assert.Equal(t, errors.New("node not found"), err, "Error message should be 'node not found'")
}
