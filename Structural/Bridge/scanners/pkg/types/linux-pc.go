package types

import (
	"fmt"
	"github.com/binaryty/Design-Design-Patterns-Golang/Structural/Bridge/scanners/pkg/base"
)

type LinuxPC struct {
	scanner base.Scanner
}

func (pc *LinuxPC) Scan() {
	fmt.Println("Attempt to scan document to file on linux pc system")
	pc.scanner.ScanFile()
}

func (pc *LinuxPC) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner
}
