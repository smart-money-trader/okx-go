package main

import (
	"log"

	"github.com/smart-money-trader/okx-go/examples/rest"
	"github.com/smart-money-trader/okx-go/rest/api"
	"github.com/smart-money-trader/okx-go/rest/api/trade"
	"github.com/smart-money-trader/okx-go/rest/api/trade_algo"
)

func main() {
	// StopLimitSell()
	ProfitLimitSell()
}

func LimitBuy() {
	param := &trade.PostOrderParam{
		InstId:  "FET-USDT",
		TdMode:  api.TdModeCash,
		Side:    api.SideBuy,
		OrdType: api.OrdTypeLimit,
		Sz:      "10",
		Px:      "1.306",
	}
	req, resp := trade.NewPostOrder(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.PostOrderResponse))
}

// Stop Loss Order
func StopLimitSell() {
	req, resp := trade_algo.NewOrderAlgoStopLimitSell("FET-USDT", "6", "1.28", "1.27")
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}

	log.Printf("%+v \n", req)
	log.Printf("%+v \n", resp.(*trade_algo.AlgoOrderResponse))
}

// Take Profit
func ProfitLimitSell() {
	req, resp := trade_algo.NewOrderAlgoProfitLimitSell("FET-USDT", "8", "1.315", "1.31")
	if err := rest.TestClient.Do(req, resp); err != nil {
		log.Println(err)
		return
	}

	log.Printf("%+v \n", req)
	log.Printf("%+v \n", resp.(*trade_algo.AlgoOrderResponse))
}
