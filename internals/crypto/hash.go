package crypto

import (
	"crypto"
	"github.com/bytemare/crypto/hash"
)

type KDF struct {
	h *hash.Hash
}

func NewKDF(id crypto.Hash) *KDF {
	return &KDF{h: hash.FromCrypto(id).Get()}
}

func (k *KDF) Expand(key, info []byte, length int) []byte {
	return k.h.HKDFExpand(key, info, length)
}
