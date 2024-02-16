package pkg

type Director struct {
	b Builder
}

func NewDirector(b Builder) *Director {
	return &Director{
		b: b,
	}
}

func (d *Director) SetBuilder(builder Builder) {
	d.b = builder
}

func (d *Director) Construct(template string) {
	space := true
	word := true

	for _, t := range template {
		if t == ' ' {
			d.b.BuildDelim()
			space = false
		} else {
			if space {
				d.b.BuildStart(t)
			} else if word {
				space = false
				d.b.BuildFirstChar(t)
				word = false
			} else {
				d.b.BuildNextChar(t)
			}
		}
	}
}

func (d *Director) GetResult() string {
	return d.b.GetResult()
}
