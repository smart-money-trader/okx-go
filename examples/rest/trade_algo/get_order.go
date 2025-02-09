package main

import (
	"log"

	"github.com/smart-money-trader/okx-go/examples/rest"
	"github.com/smart-money-trader/okx-go/rest/api/trade_algo"
)

func main() {
	// history()

	pending()
}

func info() {
	param := &trade_algo.AlgoOrderInfoParam{
		AlgoId: "1938763812979798016",
	}
	req, resp := trade_algo.AlgoOrderInfo(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		log.Println(err)
		return
	}
	log.Println(req)
	log.Printf("%+v \n", resp.(*trade_algo.AlgoOrderInfoResponse))
}

func history() {
	req, resp := trade_algo.AlgoOrderHistorySpot("FET-USDT")
	if err := rest.TestClient.Do(req, resp); err != nil {
		log.Println(err)
		return
	}
	log.Println(req)
	log.Printf("%+v \n", resp.(*trade_algo.AlgoOrderHistoryResponse))
}

func pending() {
	req, resp := trade_algo.AlgoOrderPendingSpot("FET-USDT")
	if err := rest.TestClient.Do(req, resp); err != nil {
		log.Println(err)
		return
	}
	log.Println(req)
	log.Printf("%+v \n", resp.(*trade_algo.AlgoOrderPendingResponse))
}
