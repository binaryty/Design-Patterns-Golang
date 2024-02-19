package main

import (
	"fmt"
	"github.com/binaryty/Design-Design-Patterns-Golang/Structural/Bridge/pkg/types"
)

var (
	hpScanner    = types.HP{}
	epsonScanner = types.Epson{}

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

	RenderShapes()
}
