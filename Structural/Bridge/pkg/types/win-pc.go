package types

import (
	"fmt"
	"github.com/binaryty/Design-Design-Patterns-Golang/Structural/Bridge/pkg/base"
)

type WinPC struct {
	scanner base.Scanner
}

func (w *WinPC) Scan() {
	fmt.Println("Attempt scan to file on Windows system")
	w.scanner.ScanFile()
}

func (w *WinPC) AddScanner(scanner base.Scanner) {
	w.scanner = scanner
}
