package trade_algo

import "github.com/smart-money-trader/okx-go/rest/api"

// https://www.okx.com/docs-v5/zh/?shell#order-book-trading-algo-trading-get-algo-order-list

// GET 获取未完成策略委托单列表
// 获取当前账户下未触发的策略委托单列表
func AlgoOrderPending(param *AlgoOrderPendingParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/orders-algo-pending",
		Method: api.MethodGet,
		Param:  param,
	}, &AlgoOrderPendingResponse{}
}

type AlgoOrderPendingParam struct {
	AlgoId      string `json:"algoId"`      // 策略委托单ID
	OrdType     string `json:"ordType"`     // 订单类型
	InstType    string `json:"instType"`    // 产品类型
	InstId      string `json:"instId"`      // 产品ID，BTC-USDT
	After       string `json:"after"`       // 请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的algoId
	Before      string `json:"before"`      // 请求此ID之后（更新的数据）的分页内容，传的值为对应接口的algoId
	Limit       string `json:"limit"`       // 返回结果的数量，最大为100，默认100条
	AlgoClOrdId string `json:"algoClOrdId"` // 客户自定义策略订单ID
}

type AlgoOrderPendingResponse struct {
	api.Response
	Data []AlgoOrderInfoData `json:"data"`
}

func AlgoOrderPendingSpot(instId string) (api.IRequest, api.IResponse) {
	param := &AlgoOrderPendingParam{
		OrdType:  "conditional,oco",
		InstType: api.InstTypeSPOT,
		InstId:   instId,
		Limit:    "100",
	}
	return AlgoOrderPending(param)
}
