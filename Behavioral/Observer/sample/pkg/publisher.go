package pkg

type Publisher struct {
	Name      string
	Consumers map[string]Consumer
}

func (pub *Publisher) Subscribe(consumer Consumer) {
	pub.Consumers[consumer.GetName()] = consumer
}

func (pub *Publisher) UnSubscribe(consumer Consumer) {
	delete(pub.Consumers, consumer.GetName())
}

func (pub *Publisher) Notify() {
	for _, consumer := range pub.Consumers {
		consumer.Update(pub.Name)
	}
}
