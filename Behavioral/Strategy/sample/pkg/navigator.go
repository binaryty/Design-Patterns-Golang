package pkg

type Navigator struct {
	Strategy
}

func (nav *Navigator) SetStrategy(stratrgy Strategy) {
	nav.Strategy = stratrgy
}
