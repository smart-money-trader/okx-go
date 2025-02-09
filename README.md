# okx-go

golang sdk for https://www.okx.com/docs-v5

## Installation

go get：
```shell
go get github.com/smart-money-trader/okx-go
```

## Example
all examples are in the folder examples

## Rest api example code
```go
package main

import (
	"log"

	"github.com/smart-money-trader/okx-go/rest"
	"github.com/smart-money-trader/okx-go/rest/api/account"
)

func main() {
	auth := common.NewAuth("your apikey", "your key", "your passphrase", false)
	client := rest.New("", auth, nil)
	param := &account.GetBalanceParam{}
	req, resp := account.NewGetBalance(param)
	if err := client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.GetBalanceResponse))
}
```

## Public websocket example code
```go
package main

import (
	"log"

	"github.com/smart-money-trader/okx-go/ws/public"
)

func main() {
	handler := func(c public.EventTickers) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := public.SubscribeTickers("BTC-USDT", handler, handlerError, false); err != nil {
		panic(err)
	}
	select {}
}
```

## Private websocket example code
```go
package main

import (
	"log"

	"github.com/smart-money-trader/okx-go/common"
	"github.com/smart-money-trader/okx-go/ws"
	"github.com/smart-money-trader/okx-go/ws/private"
)

func main() {
	auth := common.NewAuth("your apikey", "your key", "your passphrase", false)
	args := &ws.Args{
		InstType: "SPOT",
	}
	handler := func(c private.EventOrders) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := private.SubscribeOrders(args, auth, handler, handlerError); err != nil {
		panic(err)
	}
	select {}
}
```
