package pkg

import (
	"errors"
)

var (
	ErrItemOutOfStack         = errors.New("item out of stock")
	ErrItemAlreadyRequested   = errors.New("item already requested")
	ErrItemDispenseInProgress = errors.New("item Dispense in progress")
	ErrMoneyIsLess            = errors.New("inserted money is less")
	ErrMoneyFirst             = errors.New("please insert money first")
	ErrNoItems                = errors.New("no item present")
	ErrSelectItemFirst        = errors.New("please select item first")
)

type VendingMachine struct {
	hasItem       State
	itemRequested State
	hasMoney      State
	noItem        State
	currentState  State
	itemCount     int
	itemPrice     int
}

func NewVendingMachine(itemCount int, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}

	hasItemState := &HasItemState{
		vm: v,
	}

	itemRequestedState := &ItemRequestState{
		vm: v,
	}

	hasMoneyState := &HasMoneyState{
		vm: v,
	}

	noItemState := &NoItemState{
		vm: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState

	return v
}

func (v *VendingMachine) RequestItem() error {
	return v.currentState.RequestItem()
}

func (v *VendingMachine) AddItem(count int) error {
	return v.currentState.AddItem(count)
}

func (v *VendingMachine) InsertMoney(money int) error {
	return v.currentState.InsertMoney(money)
}

func (v *VendingMachine) DispenseItem() error {
	return v.currentState.DispenseItem()
}

func (v *VendingMachine) setState(s State) {
	v.currentState = s
}

func (v *VendingMachine) incrementItemCount(count int) {
	v.itemCount += count
}
