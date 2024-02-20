package types

import (
	"fmt"
	"github.com/binaryty/Design-Design-Patterns-Golang/Structural/Bridge/pkg/base"
)

type LinuxPC struct {
	scanner base.Scanner
}

func (l *LinuxPC) Scan() {
	fmt.Println("Attempt scan to file on Linux system")
	l.scanner.ScanFile()
}

func (l *LinuxPC) AddScanner(scanner base.Scanner) {
	l.scanner = scanner
}
