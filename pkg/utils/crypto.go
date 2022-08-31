package utils

import (
	"crypto"
	"fmt"
	"io"
)

type Hasher struct {
	ha crypto.Hash
}

func NewHasher(ha crypto.Hash) *Hasher {
	return &Hasher{ha: ha}
}

func (h Hasher) hash(data string) (s string, err error) {
	w := h.ha.New()
	_, err = io.WriteString(w, data)
	s = fmt.Sprintf("%x", w.Sum(nil))
	return
}

func Hash(ha crypto.Hash, data string) (s string, err error) {
	h := NewHasher(ha)
	s, err = h.hash(data)
	return
}
