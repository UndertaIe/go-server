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

func (h Hasher) Hash(data string) string {
	w := h.ha.New()
	io.WriteString(w, data)
	return fmt.Sprintf("%x", w.Sum(nil))
}

func Hash(ha crypto.Hash, data string) string {
	h := NewHasher(ha)
	return h.Hash(data)
}
