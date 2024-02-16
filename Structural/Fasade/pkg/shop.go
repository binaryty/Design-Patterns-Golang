package pkg

import (
	"errors"
	"fmt"
	"time"
)

type Shop struct {
	Name     string
	Products []Product
}

func (s Shop) Sell(user User, product string) error {
	fmt.Println("[Shop]: Request balance User card ")
	time.Sleep(time.Millisecond * 500)

	if err := user.Card.CheckBalance(); err != nil {
		return err
	}

	fmt.Printf("[Shop]: Checking whether the [%s] can buy the product\n", user.Name)
	time.Sleep(time.Millisecond * 500)

	for _, p := range s.Products {
		if p.Name != product {
			continue
		}

		if p.Price > user.GetBalance() {
			return errors.New("[Shop]: insufficient funds")
		}

		user.Card.Balance -= p.Price
		fmt.Printf("[Shop]: The product [%s] was purchased by [%s]\n", product, user.Name)
	}

	return nil
}
