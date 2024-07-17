package cvt

import (
	"bytes"
	"encoding/binary"
)

func Bti(d []byte) (int, error) {
	var i uint32
	err := binary.Read(bytes.NewReader(d), binary.BigEndian, &i)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

func Itb(i int) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(i))
	return b
}
