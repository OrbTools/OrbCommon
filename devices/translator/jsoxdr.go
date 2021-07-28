package main

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/OrbTools/OrbCommon/devices"
	xdr "github.com/davecgh/go-xdr/xdr2"
)

func main() {
	files, _ := os.ReadDir("devices/json/")
	for _, file := range files {
		data, _ := os.ReadFile("devices/json/" + file.Name())
		DevDef := &devices.DeviceDef{}
		json.Unmarshal(data, DevDef)
		xdo, _ := os.Create("devices/xdr/" + strings.Split(file.Name(), ".")[0] + ".bin")
		xdr.Marshal(xdo, DevDef)
		xdo.Close()
	}
}
