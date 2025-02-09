package main

import (
	"log"

	"github.com/smart-money-trader/okx-go/examples/rest"
	"github.com/smart-money-trader/okx-go/rest/api/convert"
)

func main() {
	param := &convert.GetHistoryParam{}
	req, resp := convert.NewGetHistory(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*convert.GetHistoryResponse))
}
