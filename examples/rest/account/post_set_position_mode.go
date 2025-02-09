package main

import (
	"log"

	"github.com/smart-money-trader/okx-go/examples/rest"
	"github.com/smart-money-trader/okx-go/rest/api"
	"github.com/smart-money-trader/okx-go/rest/api/account"
)

func main() {
	param := &account.PostSetPositionModeParam{
		PosMode: api.PosModeLongShort,
	}
	req, resp := account.NewPostSetPositionMode(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.PostSetPositionModeResponse))
}
