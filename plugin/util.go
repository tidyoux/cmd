package plugin

import (
	"encoding/hex"
	"strings"
)

func toBytes(s string) []byte {
	if strings.HasPrefix(s, "0x") {
		if d, err := hex.DecodeString(s[2:]); err == nil {
			return d
		}
	}
	return []byte(s)
}

func toHex(d []byte) string {
	return "0x" + hex.EncodeToString(d)
}
