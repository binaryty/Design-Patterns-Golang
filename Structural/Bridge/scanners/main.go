package main

import "github.com/binaryty/Design-Design-Patterns-Golang/Structural/Bridge/scanners/pkg/types"

var (
	cannon = types.CannonScanner{}
	epson  = types.EpsonScanner{}
	hp     = types.HPScanner{}

	windowsPC = types.WinPC{}
	linuxPC   = types.LinuxPC{}
	macPC     = types.MacPC{}
)

func main() {
	windowsPC.AddScanner(cannon)
	windowsPC.Scan()

	linuxPC.AddScanner(epson)
	linuxPC.Scan()

	macPC.AddScanner(hp)
	macPC.Scan()
}
