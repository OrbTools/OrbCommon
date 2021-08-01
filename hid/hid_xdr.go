// +build xdr

package hid

import (
	"bytes"
	_ "embed"
	xdr "github.com/minizbot2012/minxdr"
)

//go:embed generated.bin
var file []byte

func init() {
	xdr.Unmarshal(bytes.NewReader(file), &Mappings)
}
