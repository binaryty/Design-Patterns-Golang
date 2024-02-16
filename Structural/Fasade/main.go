package main

import (
	"fmt"
	"github.com/binaryty/Design-Design-Patterns-Golang/Structural/Fasade/pkg"
)

var (
	bank = pkg.Bank{
		Name:  "Alfa Bank",
		Cards: []pkg.Card{},
	}

	card1 = pkg.Card{
		ID:      "123",
		Type:    "Debit",
		Balance: 200,
		Bank:    &bank,
	}

	card2 = pkg.Card{
		ID:      "456",
		Type:    "Debit",
		Balance: 10,
		Bank:    &bank,
	}

	user1 = pkg.User{
		Name: "Ivan",
		Card: &card1,
	}

	user2 = pkg.User{
		Name: "Petr",
		Card: &card2,
	}

	product1 = pkg.Product{
		Name:  "Cheese",
		Price: 120,
	}

	product2 = pkg.Product{
		Name:  "Bread",
		Price: 50,
	}

	shop = pkg.Shop{
		Name: "TopStore",
		Products: []pkg.Product{
			product1,
			product2,
		},
	}
)

func main() {
	fmt.Println("[Bank]: Card issuance")
	bank.Cards = append(bank.Cards, card1, card2)

	fmt.Printf("[%s]\n", user1.Name)
	if err := shop.Sell(user1, product1.Name); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("[%s]\n", user2.Name)
	if err := shop.Sell(user2, product2.Name); err != nil {
		fmt.Println(err.Error())
	}
}
