package pkg

import "fmt"

type HasMoneyState struct {
	vm *VendingMachine
}

func (i *HasMoneyState) AddItem(count int) error {
	return ErrItemDispenseInProgress
}

func (i *HasMoneyState) RequestItem() error {
	return ErrItemDispenseInProgress
}

func (i *HasMoneyState) InsertMoney(money int) error {
	return ErrItemOutOfStack
}

func (i *HasMoneyState) DispenseItem() error {
	fmt.Println("Dispensing item")
	i.vm.itemCount = i.vm.itemCount - 1

	if i.vm.itemCount == 0 {
		i.vm.setState(i.vm.noItem)
	} else {
		i.vm.setState(i.vm.hasItem)
	}

	return nil
}
