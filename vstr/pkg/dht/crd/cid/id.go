// Andre R. R. Costa *** github.com/andrerrcosta2

package cid

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/big"
	"net"
)

const (
	M = 160
)

type Id [M / 8]byte

func New(ip net.IP, port uint16) Id {
	addr := append(ip, byte(port>>8), byte(port))
	return sha1.Sum(addr)
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

func (id Id) Big() *big.Int {
	return new(big.Int).SetBytes(id[:])
}

func (id Id) Byt() []byte {
	return id[:]
}

func (id Id) Strt(i uint) Id {
	idi := id.Big()
	ofs := new(big.Int).Lsh(big.NewInt(1), i) // 2^i
	rsz := new(big.Int).Lsh(big.NewInt(1), M) // 2^M
	strt := new(big.Int).Add(idi, ofs)
	strt.Mod(strt, rsz) // (i + 2^i) % 2^M
	return Fbi(strt)
}

func (id Id) Vld() bool {
	for _, b := range id {
		if b != 0 {
			return true
		}
	}
	return false
}

func (id Id) String() string {
	return hex.EncodeToString(id[:])
}

func Rdm() (Id, error) {
	var id Id
	_, err := rand.Read(id[:])
	if err != nil {
		return id, fmt.Errorf("Rdm failed: %w", err)
	}
	return id, nil
}

func BytesToId(bytes []byte) (Id, error) {
	var id Id

	if len(bytes) != M/8 {
		return id, fmt.Errorf("unexpected length of bytes")
	}

	copy(id[:], bytes)

	return id, nil
}
