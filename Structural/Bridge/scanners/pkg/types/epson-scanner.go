package types

import "fmt"

type EpsonScanner struct {
}

func (c EpsonScanner) ScanFile() {
	fmt.Println("Epson scanner scan file")
}
