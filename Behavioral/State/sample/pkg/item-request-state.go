package pkg

import (
	"fmt"
)

type ItemRequestState struct {
	vm *VendingMachine
}

func (i *ItemRequestState) AddItem(count int) error {
	return ErrItemDispenseInProgress
}

func (i *ItemRequestState) RequestItem() error {
	return ErrItemAlreadyRequested
}

func (i *ItemRequestState) InsertMoney(money int) error {
	if money < i.vm.itemPrice {
		return ErrMoneyIsLess
	}

	fmt.Println("Money entered is ok")
	i.vm.setState(i.vm.hasMoney)

	return nil
}

func (i *ItemRequestState) DispenseItem() error {
	return ErrMoneyFirst
}
