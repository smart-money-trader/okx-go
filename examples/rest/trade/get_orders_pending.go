package main

import (
	"log"

	"github.com/smart-money-trader/okx-go/examples/rest"
	"github.com/smart-money-trader/okx-go/rest/api/trade"
)

func main() {
	param := &trade.GetOrdersQueryParam{}
	req, resp := trade.NewGetOrdersPending(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.GetOrderResponse))
}
