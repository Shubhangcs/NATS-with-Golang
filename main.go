package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nats-io/nats.go"
)

func main() {
	http.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request) {
		natsTrigger()
		w.WriteHeader(http.StatusOK)
	})
	log.Fatal(http.ListenAndServe(":8000" , nil))
}

func natsTrigger(){
	nat, err := nats.Connect(":4222")
	if err != nil {
		fmt.Println(err.Error())
	}
	nat.Publish("k", []byte("Hello"))
}
