package types

import "fmt"

type HPScanner struct {
}

func (c HPScanner) ScanFile() {
	fmt.Println("HP scanner scan file")
}
