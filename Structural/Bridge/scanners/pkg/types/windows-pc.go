package types

import (
	"fmt"
	"github.com/binaryty/Design-Design-Patterns-Golang/Structural/Bridge/scanners/pkg/base"
)

type WinPC struct {
	scanner base.Scanner
}

func (pc *WinPC) Scan() {
	fmt.Println("Attempt to scan document to file on windows pc system")
	pc.scanner.ScanFile()
}

func (pc *WinPC) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner
}
