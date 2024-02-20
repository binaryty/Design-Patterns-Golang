package types

import "fmt"

type CannonScanner struct {
}

func (c CannonScanner) ScanFile() {
	fmt.Println("Cannon scanner scan file")
}
