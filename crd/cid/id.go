// Andre R. R. Costa *** github.com/andrerrcosta2

package cid

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/big"
)

const (
	M = 160
)

type Id [M / 8]byte

func NewId(key string) Id {
	return sha1.Sum([]byte(key))
}

func (id Id) Eq(other Id) bool {
	for i := range id {
		if id[i] != other[i] {
			return false
		}
	}
	return true
}

func (id Id) Lt(other Id) bool {
	for i := range id {
		if id[i] < other[i] {
			return true
		} else if id[i] > other[i] {
			return false
		}
	}
	return false
}

func (id Id) Gt(other Id) bool {
	for i := range id {
		if id[i] > other[i] {
			return true
		} else if id[i] < other[i] {
			return false
		}
	}
	return false
}

func (id Id) Btw(start, end Id) bool {
	if start.Lt(id) {
		return id.Lt(end)
	}
	return false
}

func Fbi(bi *big.Int) Id {
	var id Id
	b := bi.Bytes()
	copy(id[M/8-len(b):], b)
	return id
}

func (id Id) Tbi() *big.Int {
	return new(big.Int).SetBytes(id[:])
}

func (id Id) Strt(i int) Id {
	idi := id.Tbi()
	ofs := new(big.Int).Lsh(big.NewInt(1), uint(i)) // 2^i
	rsz := new(big.Int).Lsh(big.NewInt(1), M)       // 2^M
	strt := new(big.Int).Add(idi, ofs)
	strt.Mod(strt, rsz) // (i + 2^i) % 2^M
	return Fbi(strt)
}

func (id Id) String() string {
	return hex.EncodeToString(id[:])
}

func BytesToId(bytes []byte) (Id, error) {
	var id Id

	if len(bytes) != M/8 {
		return id, fmt.Errorf("unexpected length of bytes")
	}

	copy(id[:], bytes)

	return id, nil
}
