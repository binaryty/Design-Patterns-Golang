package pkg

type Light interface {
	GreenLight() error
	YellowLight() error
	RedLight() error
}
