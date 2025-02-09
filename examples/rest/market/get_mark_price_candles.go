package main

import (
	"log"

	"github.com/smart-money-trader/okx-go/examples/rest"
	"github.com/smart-money-trader/okx-go/rest/api/market"
)

func main() {
	param := &market.GetCandlesParam{
		InstId: "ETH-USDT",
	}
	req, resp := market.NewGetMarkPriceCandles(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*market.GetCandlesResponse))
}
