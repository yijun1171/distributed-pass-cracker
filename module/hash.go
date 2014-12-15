package module

import (
	"crypto/md5"
	"encoding/hex"
	"hash"
)

type Hash struct {
	hashStr string
	hashCtx hash.Hash
}

func NewHash() *Hash {
	return &Hash{"", md5.New()}
}

func (h *Hash) Init(src []byte, salt []byte) string {
	h.hashCtx.Write(src)
	h.hashStr = hex.EncodeToString(h.hashCtx.Sum(salt))
	return h.GetHash()
}

func (h *Hash) GetHash() string {
	return h.hashStr
}

func (h *Hash) Check(other []byte, salt []byte) bool {
	if other == nil {
		return false
	}
	h.hashCtx.Reset() //reset
	h.hashCtx.Write(other)
	otherStr := hex.EncodeToString(h.hashCtx.Sum(salt))
	if otherStr == h.hashStr {
		return true
	} else {
		return false
	}
}
