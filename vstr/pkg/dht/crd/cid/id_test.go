package cid

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

const (
	tk1 = "127.0.0.1"
	tk2 = "128.0.0.1"
	tk3 = "129.0.0.1"
)

func TestNewId(t *testing.T) {
	exp := sha1.Sum([]byte(tk1))
	id := New(tk1)
	var idArr [20]byte
	copy(idArr[:], id[:])
	assert.Equal(t, exp, idArr, "New() should generate the correct ID")
}

func TestEq(t *testing.T) {
	id1 := New(tk1)
	id2 := New(tk1)
	id3 := New(tk2)

	assert.True(t, id1.Eq(id2), "Equal() should return true for equal IDs")
	assert.False(t, id1.Eq(id3), "Equal() should return false for different IDs")
}

func TestLt(t *testing.T) {
	id1 := New(tk1)
	id2 := New(tk2)

	assert.Equal(t, id1.Lt(id2), id1.Big().Cmp(id2.Big()) < 0, "Lt() should return correct comparison")
}

func TestGt(t *testing.T) {
	id1 := New(tk1)
	id2 := New(tk2)

	assert.Equal(t, id1.Gt(id2), id1.Big().Cmp(id2.Big()) > 0, "Gt() should return correct comparison")
}

func TestBtw(t *testing.T) {
	fmt.Printf("tk1: %s, tk2: %s, tk3: %s", tk1, tk2, tk3)

	id1 := New(tk1) // 4b84b15bff6ee5796152495a230e45e3d7e947d9
	id2 := New(tk2) // 9c678c09e0163cb9f0bdaf0363047ac5b549704e
	id3 := New(tk3) // 4cb76ef2b5226800f33753538f867f39fd13c0ab

	fmt.Printf("tk1: %s, tk2: %s, tk3: %s", id1.String(), id2.String(), id3.String())

	assert.True(t, id3.Btw(id1, id2), "Btw() should return true if id is between start and end in a wrap-around range")
	assert.False(t, id1.Btw(id2, id3), "Btw() should return false if id is not between start and end")
}

func TestFbi_Tbi(t *testing.T) {
	bi := big.NewInt(123456789)
	id := Fbi(bi)
	assert.Equal(t, bi, id.Big(), "Fbi() and Big() should be consistent")
}

func TestStrt(t *testing.T) {
	id := New(tk1)
	strt := id.Strt(5)
	exp := new(big.Int).Add(id.Big(), new(big.Int).Lsh(big.NewInt(1), 5))
	exp.Mod(exp, new(big.Int).Lsh(big.NewInt(1), M))

	assert.Equal(t, exp, strt.Big(), "Strt() should calculate the correct start ID")
}

func TestString(t *testing.T) {
	id := New(tk1)
	exp := hex.EncodeToString(id[:])
	assert.Equal(t, exp, id.String(), "String() should return the correct hex representation")
}

func TestBytesToId(t *testing.T) {
	id := New(tk1)
	bytes := id[:]
	_id, err := BytesToId(bytes)
	assert.NoError(t, err, "BytesToId() should not return an error for correct byte length")
	assert.Equal(t, id, _id, "BytesToId() should return the correct ID")

	_, err = BytesToId([]byte("short"))
	assert.Error(t, err, "BytesToId() should return an error for incorrect byte length")
}
