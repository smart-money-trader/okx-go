package main

import (
	"log"

	"github.com/smart-money-trader/okx-go/examples/rest"
	"github.com/smart-money-trader/okx-go/rest/api/trade"
)

func main() {
	param := &trade.GetOrderParam{
		InstId: "OKB-USDT",
		OrdId:  "501163171776954368",
	}
	req, resp := trade.NewGetOrder(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.GetOrderResponse))
}
