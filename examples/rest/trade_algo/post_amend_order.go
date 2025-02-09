package main

import (
	"log"

	"github.com/smart-money-trader/okx-go/examples/rest"
	"github.com/smart-money-trader/okx-go/rest/api/trade_algo"
)

func main() {
	req, resp := trade_algo.AmendAlgoOrderStopLimit("FET-USDT", "1938763812979798016", "6", "1.281", "1.271")
	if err := rest.TestClient.Do(req, resp); err != nil {
		log.Println(err)
		return
	}
	log.Println(req, resp.(*trade_algo.AmendAlgoOrderResponse))
}

// 2024/10/30 18:01:30 &{/api/v5/trade/amend-algos POST 0x1400013c000} &{{0  <nil>} [{1938763812979798016   0 }]}
