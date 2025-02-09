package main

import (
	"log"

	"github.com/smart-money-trader/okx-go/examples/rest"
	"github.com/smart-money-trader/okx-go/rest/api/asset"
)

func main() {
	param := &asset.PostTransferParam{
		Ccy:  "USDT",
		Amt:  "1",
		From: "18",
		To:   "6",
	}
	req, resp := asset.NewPostTransfer(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*asset.PostTransferResponse))
}
