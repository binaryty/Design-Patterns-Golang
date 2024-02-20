package types

import (
	"fmt"
	"github.com/binaryty/Design-Design-Patterns-Golang/Structural/Bridge/pkg/base"
)

type MacPC struct {
	scanner base.Scanner
}

func (m *MacPC) Scan() {
	fmt.Println("Attempt scan to file on Mac system")
	m.scanner.ScanFile()
}

func (m *MacPC) AddScanner(scanner base.Scanner) {
	m.scanner = scanner
}
