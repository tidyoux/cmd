package plugin

import (
	"encoding/base64"
	"fmt"

	"github.com/decred/base58"
)

func init() {
	Register("base64", &Encode{base64.StdEncoding.EncodeToString})
	Register("base58", &Encode{base58.Encode})
	Register("hex2str", &Encode{
		func(d []byte) string {
			return string(d)
		}})

	Register("base64decode", &Decode{base64.StdEncoding.DecodeString})
	Register("base58decode", &Decode{
		func(s string) ([]byte, error) {
			return base58.Decode(s), nil
		}})
}

type Encode struct {
	encoder func([]byte) string
}

func (e *Encode) Exe(cmd string, args []string) ([]string, error) {
	result := make([]string, 0, len(args))
	for _, s := range args {
		result = append(result, e.encoder(toBytes(s)))
	}
	return result, nil
}

type Decode struct {
	decoder func(string) ([]byte, error)
}

func (d *Decode) Exe(cmd string, args []string) ([]string, error) {
	result := make([]string, 0, len(args))
	for _, s := range args {
		data, err := d.decoder(s)
		if err != nil {
			return nil, fmt.Errorf("decode, %v", err)
		}

		result = append(result, toHex(data))
	}
	return result, nil
}
