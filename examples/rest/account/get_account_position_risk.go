package main

import (
	"log"

	"github.com/smart-money-trader/okx-go/examples/rest"
	"github.com/smart-money-trader/okx-go/rest/api/account"
)

func main() {
	param := &account.GetAccountPositionRiskParam{}
	req, resp := account.NewGetAccountPositionRisk(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.GetAccountPositionRiskResponse))
}
