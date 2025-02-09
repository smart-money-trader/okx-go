package trade_algo

import "github.com/smart-money-trader/okx-go/rest/api"

// https://www.okx.com/docs-v5/zh/#order-book-trading-algo-trading-post-cancel-algo-order

// POST 撤销策略委托订单
// 撤销策略委托订单，每次最多可以撤销10个策略委托单
func CancelAlgoOrder(param []*CancelAlgoOrderParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/cancel-algos",
		Method: api.MethodPost,
		Param:  param,
	}, &CancelAlgoOrderResponse{}
}

type CancelAlgoOrderParam struct {
	InstId string `json:"instId"` // 产品ID 如 BTC-USDT
	AlgoId string `json:"algoId"` // 策略委托单ID
}

type CancelAlgoOrderResponse struct {
	api.Response
	Data []CancelOrder `json:"data"`
}

type CancelOrder struct {
	AlgoId string `json:"algoId"` // 策略委托单ID
	SCode  string `json:"sCode"`  // 事件执行结果的code，0代表成功
	SMsg   string `json:"sMsg"`   // 事件执行失败时的msg
}
