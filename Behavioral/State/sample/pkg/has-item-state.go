package pkg

import "fmt"

type HasItemState struct {
	vm *VendingMachine
}

func (i *HasItemState) AddItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vm.incrementItemCount(count)

	return nil
}

func (i *HasItemState) RequestItem() error {
	if i.vm.itemCount == 0 {
		i.vm.setState(i.vm.noItem)

		return ErrNoItems
	}

	fmt.Printf("Item requested\n")

	i.vm.setState(i.vm.itemRequested)

	return nil
}

func (i *HasItemState) InsertMoney(money int) error {
	return ErrSelectItemFirst
}

func (i *HasItemState) DispenseItem() error {
	return ErrSelectItemFirst
}
