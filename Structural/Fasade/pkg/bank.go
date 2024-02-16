package pkg

import (
	"errors"
	"fmt"
	"time"
)

type Bank struct {
	Name  string
	Cards []Card
}

func (b Bank) CheckBalance(cardID string) error {
	fmt.Printf("[Bank]: cardID: %s - get card balance\n", cardID)
	time.Sleep(time.Millisecond * 300)
	for _, c := range b.Cards {
		if c.ID != cardID {
			continue
		}
		if c.Balance <= 0 {
			return errors.New("[Bank]: insufficient funds")
		}
	}

	fmt.Println("[Bank]: the balance is positive")

	return nil
}
