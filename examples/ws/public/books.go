package main

import (
	"log"

	"github.com/smart-money-trader/okx-go/ws"
	"github.com/smart-money-trader/okx-go/ws/public"
)

func main() {
	args := &ws.Args{
		Channel: "books5",
		InstId:  "BTC-USDT",
	}
	handler := func(c public.EventBooks) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := public.SubscribeBooks(args, handler, handlerError, false); err != nil {
		panic(err)
	}
	select {}
}
