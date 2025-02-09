package main

import (
	"log"

	"github.com/smart-money-trader/okx-go/ws/public"
)

func main() {
	handler := func(c public.EventFundingRate) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := public.SubscribeFundingRate("BTC-USDT-SWAP", handler, handlerError, false); err != nil {
		panic(err)
	}
	select {}
}
