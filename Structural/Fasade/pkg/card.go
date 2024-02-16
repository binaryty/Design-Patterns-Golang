package pkg

import (
	"fmt"
	"time"
)

type Card struct {
	ID      string
	Type    string
	Balance float64
	Bank    *Bank
}

func (c Card) CheckBalance() error {
	fmt.Println("[Card]: Request bank to check balance")
	time.Sleep(time.Millisecond * 500)

	return c.Bank.CheckBalance(c.ID)
}
