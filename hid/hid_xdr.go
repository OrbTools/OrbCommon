// +build xdr

package hid

import (
	"bytes"
	_ "embed"
	xdr "github.com/davecgh/go-xdr/xdr2"
)

//go:embed generated.bin
var file []byte

func init() {
	xdr.Unmarshal(bytes.NewReader(file), Mappings)
}
