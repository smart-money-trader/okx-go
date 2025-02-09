package main

import (
	"log"

	"github.com/smart-money-trader/okx-go/examples/rest"
	"github.com/smart-money-trader/okx-go/rest/api/subaccount"
)

func main() {
	param := &subaccount.GetEntrustSubaccountListParam{}
	req, resp := subaccount.NewGetEntrustSubaccountList(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*subaccount.GetEntrustSubaccountListResponse))
}
