package main

import (
	"log"

	"github.com/smart-money-trader/okx-go/examples/rest"
	"github.com/smart-money-trader/okx-go/rest/api/market"
)

func main() {
	param := &market.GetBooksParam{
		InstId: "BTC-USDT",
	}
	req, resp := market.NewGetBooks(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*market.GetBooksResponse))
}
