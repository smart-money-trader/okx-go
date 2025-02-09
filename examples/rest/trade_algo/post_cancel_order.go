package main

import (
	"log"

	"github.com/smart-money-trader/okx-go/examples/rest"
	"github.com/smart-money-trader/okx-go/rest/api/trade_algo"
)

func main() {
	param := []*trade_algo.CancelAlgoOrderParam{}

	param = append(param, &trade_algo.CancelAlgoOrderParam{
		InstId: "FET-USDT",
		AlgoId: "1938680425887756288",
	})
	req, resp := trade_algo.CancelAlgoOrder(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		log.Println(err)
		return
	}
	log.Println(req, resp.(*trade_algo.CancelAlgoOrderResponse))
}
