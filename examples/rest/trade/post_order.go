package main

import (
	"log"

	"github.com/smart-money-trader/okx-go/examples/rest"
	"github.com/smart-money-trader/okx-go/rest/api"
	"github.com/smart-money-trader/okx-go/rest/api/trade"
)

func main() {
	param := &trade.PostOrderParam{
		InstId:  "OKB-USDT",
		TdMode:  api.TdModeCash,
		Side:    api.SideBuy,
		OrdType: api.OrdTypeLimit,
		Sz:      "9",
		Px:      "5",
	}
	req, resp := trade.NewPostOrder(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.PostOrderResponse))
}
