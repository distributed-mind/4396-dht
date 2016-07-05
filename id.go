package dht

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

// ZeroID "0000000000000000000000000000000000000000"
var ZeroID ID

// ID consists of 160 bits
type ID [5]uint32

// NewID returns a id
func NewID(hash []byte) (id ID, err error) {
	h := make([]byte, 20)
	n, err := hex.Decode(h, hash)
	if err != nil {
		return
	}
	if n != 20 {
		err = fmt.Errorf("invalid hash")
		return
	}

	for i := 0; i < 5; i++ {
		for j, k := i*4, uint32(24); j < (i+1)*4; j, k = j+1, k-8 {
			id[i] |= uint32(h[j]) << k
		}
	}
	return
}

// NewRandomID returns a random id
func NewRandomID() (id ID) {
	for i := 0; i < 5; i++ {
		id[i] = rand.Uint32()
	}
	return
}

// Compare two id
func (id ID) Compare(o ID) int {
	for i := 0; i < 5; i++ {
		if id[i] > o[i] {
			return 1
		} else if id[i] < o[i] {
			return -1
		}
	}
	return 0
}

func (id ID) String() string {
	return fmt.Sprintf("%08x%08x%08x%08x%08x", id[0], id[1], id[2], id[3], id[4])
}

func init() {
	rand.Seed(time.Now().UnixNano())
}