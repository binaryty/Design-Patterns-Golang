package types

import (
	"fmt"
	"github.com/binaryty/Design-Design-Patterns-Golang/Structural/Bridge/scanners/pkg/base"
)

type MacPC struct {
	scanner base.Scanner
}

func (pc *MacPC) Scan() {
	fmt.Println("Attempt to scan document to file on mac pc system")
	pc.scanner.ScanFile()
}

func (pc *MacPC) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner
}
