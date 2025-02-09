package main

import (
	"log"

	"github.com/smart-money-trader/okx-go/examples/rest"
	"github.com/smart-money-trader/okx-go/rest/api"
	"github.com/smart-money-trader/okx-go/rest/api/trade"
)

func main() {
	param := []*trade.PostOrderParam{
		&trade.PostOrderParam{
			InstId:  "OKB-USDT",
			TdMode:  api.TdModeCash,
			Side:    api.SideBuy,
			OrdType: api.OrdTypeLimit,
			Sz:      "9",
			Px:      "5",
		},
		&trade.PostOrderParam{
			InstId:  "XRP-USDT",
			TdMode:  api.TdModeCash,
			Side:    api.SideBuy,
			OrdType: api.OrdTypeLimit,
			Sz:      "1",
			Px:      "0.03",
		},
	}
	req, resp := trade.NewPostBatchOrders(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.PostOrderResponse))
}
