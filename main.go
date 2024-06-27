package main

import (

	"github.com/nats-io/nats.go"
)

func main() {
	nat , err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	nat.Publish("MyNet" , []byte("Hello"))
}

