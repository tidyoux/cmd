package plugin

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"strings"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

func init() {
	Register("md5", &Hash{md5.New()})
	Register("sha1", &Hash{sha1.New()})
	Register("sha256", &Hash{sha256.New()})
	Register("sha512", &Hash{sha512.New()})
	Register("ripemd160", &Hash{ripemd160.New()})
	Register("sha3-256", &Hash{sha3.New256()})
	Register("sha3-512", &Hash{sha3.New512()})
	Register("keccak-256", &Hash{sha3.NewLegacyKeccak256()})

	h, _ := blake2b.New256(nil)
	Register("blake2b-256", &Hash{h})

	h, _ = blake2b.New512(nil)
	Register("blake2b-512", &Hash{h})
}

type Hash struct {
	hasher hash.Hash
}

func (h *Hash) Exe(cmd string, args []string) ([]string, error) {
	s := strings.Join(args, " ")
	h.hasher.Reset()
	h.hasher.Write(toBytes(s))
	d := h.hasher.Sum(nil)
	return []string{toHex(d)}, nil
}
