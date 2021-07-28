// +build !xdr

package hid

import (
	_ "embed"
	"encoding/json"
)

//go:embed generated.json
var file []byte

func init() {
	json.Unmarshal(file, &Mappings)
}
