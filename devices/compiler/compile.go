package main

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/Minizbot2012/minxdr"
	"github.com/OrbTools/OrbCommon/devices/structs"
)

func main() {
	files, _ := os.ReadDir("devices/json/")
	for _, file := range files {
		data, _ := os.ReadFile("devices/json/" + file.Name())
		DevDef := &structs.DeviceDef{}
		json.Unmarshal(data, DevDef)
		xdo, _ := os.Create("devices/xdr/" + strings.Split(file.Name(), ".")[0] + ".bin")
		minxdr.Marshal(xdo, DevDef)
		xdo.Close()
	}
}
