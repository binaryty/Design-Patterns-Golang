package types

import "fmt"

type Epson struct {
}

func (e *Epson) ScanFile() {
	fmt.Println("Document already scanned to file on Epson Scanner")
}
