package main

import (
	"log"

	"github.com/smart-money-trader/okx-go/examples"
	"github.com/smart-money-trader/okx-go/ws"
	"github.com/smart-money-trader/okx-go/ws/private"
)

func main() {
	args := &ws.Args{
		InstType: "SPOT",
	}
	handler := func(c private.EventOrders) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := private.SubscribeOrders(args, examples.Auth, handler, handlerError); err != nil {
		panic(err)
	}
	select {}
}
