package trade_algo

import "github.com/smart-money-trader/okx-go/rest/api"

// https://www.okx.com/docs-v5/zh/?shell#order-book-trading-algo-trading-get-algo-order-details

// GET 获取历史策略委托单列表
// 获取最近3个月当前账户下所有策略委托单列表
func AlgoOrderHistory(param *AlgoOrderHistoryParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/orders-algo-history",
		Method: api.MethodGet,
		Param:  param,
	}, &AlgoOrderHistoryResponse{}
}

type AlgoOrderHistoryParam struct {
	OrdType  string `json:"ordType"`  // 订单类型
	State    string `json:"state"`    // 订单状态
	AlgoId   string `json:"algoId"`   // 策略委托单ID
	InstType string `json:"instType"` // 产品类型
	InstId   string `json:"instId"`   // 产品ID，BTC-USDT
	After    string `json:"after"`    // 请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的algoId
	Before   string `json:"before"`   // 请求此ID之后（更新的数据）的分页内容，传的值为对应接口的algoId
	Limit    string `json:"limit"`    // 返回结果的数量，最大为100，默认100条
}

type AlgoOrderHistoryResponse struct {
	api.Response
	Data []AlgoOrderInfoData `json:"data"`
}

func AlgoOrderHistorySpot(instId string) (api.IRequest, api.IResponse) {
	param := &AlgoOrderHistoryParam{
		OrdType:  "conditional,oco",
		InstType: api.InstTypeSPOT,
		InstId:   instId,
		Limit:    "100",
		// State:    api.StateEffective,
	}
	return AlgoOrderHistory(param)
}
