package pkg

import "fmt"

type Subscriber struct {
	Name string
}

func (cons *Subscriber) Update(pubName string) {
	fmt.Printf("Sending to subscribe %s from %s\n", cons.GetName(), pubName)
}

func (cons *Subscriber) GetName() string {
	return cons.Name
}
