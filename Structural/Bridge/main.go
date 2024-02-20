package main

import (
	"fmt"
	"github.com/binaryty/Design-Design-Patterns-Golang/Structural/Bridge/scanners/pkg/types"
	"github.com/binaryty/Design-Design-Patterns-Golang/Structural/Bridge/shapes"
)

var (
	hpScanner    = types.HPScanner{}
	epsonScanner = types.EpsonScanner{}

	winPC   = types.WinPC{}
	linuxPC = types.LinuxPC{}
	macPC   = types.MacPC{}
)

func main() {
	winPC.AddScanner(&hpScanner)
	winPC.Scan()

	linuxPC.AddScanner(&epsonScanner)
	linuxPC.Scan()

	macPC.AddScanner(&hpScanner)
	macPC.Scan()

	fmt.Println()

	shapes.RenderShapes()
}
