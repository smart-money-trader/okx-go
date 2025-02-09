package main

import (
	"log"

	"github.com/smart-money-trader/okx-go/examples/rest"
	"github.com/smart-money-trader/okx-go/rest/api"
	"github.com/smart-money-trader/okx-go/rest/api/market"
)

func main() {
	param := &market.GetTickersParam{
		InstType: api.InstTypeSPOT,
	}
	req, resp := market.NewGetTickers(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*market.GetTickersResponse))
	log.Println(len(resp.(*market.GetTickersResponse).Data))
}
