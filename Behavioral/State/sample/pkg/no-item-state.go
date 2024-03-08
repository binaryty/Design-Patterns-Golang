package pkg

type NoItemState struct {
	vm *VendingMachine
}

func (i *NoItemState) RequestItem() error {
	return ErrItemOutOfStack
}

func (i *NoItemState) AddItem(count int) error {
	i.vm.incrementItemCount(count)
	i.vm.setState(i.vm.hasItem)

	return nil
}

func (i *NoItemState) InsertMoney(money int) error {
	return ErrItemOutOfStack
}

func (i *NoItemState) DispenseItem() error {
	return ErrItemOutOfStack
}
