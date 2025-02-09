package rest

import (
	"github.com/smart-money-trader/okx-go/examples"
	rc "github.com/smart-money-trader/okx-go/rest"
)

// 敏感信息申请的模拟盘KEY，不确定何时会删除
var TestClient = rc.New("", examples.Auth, nil)
