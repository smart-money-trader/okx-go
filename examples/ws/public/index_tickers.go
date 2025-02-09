package main

import (
	"log"

	"github.com/smart-money-trader/okx-go/ws/public"
)

func main() {
	handler := func(c public.EventIndexTickers) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := public.SubscribeIndexTickers("BTC-USDT", handler, handlerError, false); err != nil {
		panic(err)
	}
	select {}
}
