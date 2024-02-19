package types

import "fmt"

type HP struct {
}

func (H *HP) ScanFile() {
	fmt.Println("Document already scanned to file on HP Scanner")
}
