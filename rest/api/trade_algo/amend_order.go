package trade_algo

import "github.com/smart-money-trader/okx-go/rest/api"

// https://www.okx.com/docs-v5/zh/?shell#order-book-trading-algo-trading-post-amend-algo-order

// POST 修改策略委托订单
// 修改策略委托订单（仅支持止盈止损和计划委托订单，不包含、冰山委托、时间加权、移动止盈止损等订单）
func AmendAlgoOrder(param *AmendAlgoOrderParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/amend-algos",
		Method: api.MethodPost,
		Param:  param,
	}, &AmendAlgoOrderResponse{}
}

type AmendAlgoOrderParam struct {
	InstId      string `json:"instId"`                // 产品ID
	AlgoId      string `json:"algoId,omitempty"`      // 策略委托单ID
	AlgoClOrdId string `json:"algoClOrdId,omitempty"` // 客户自定义策略订单ID
	CxlOnFail   bool   `json:"cxlOnFail,omitempty"`   // 当订单修改失败时，该订单是否需要自动撤销。默认为false
	ReqId       string `json:"reqId,omitempty"`       // 用户自定义修改事件ID
	NewSz       string `json:"newSz"`                 // 修改的新数量，必须大于0

	// 止盈止损
	NewTpTriggerPx     string `json:"newTpTriggerPx,omitempty"`     // 止盈触发价
	NewTpOrdPx         string `json:"newTpOrdPx,omitempty"`         // 止盈委托价
	NewTpTriggerPxType string `json:"newTpTriggerPxType,omitempty"` // 止盈触发价类型
	NewSlTriggerPx     string `json:"newSlTriggerPx,omitempty"`     // 止损触发价
	NewSlOrdPx         string `json:"newSlOrdPx,omitempty"`         // 止损委托价
	NewSlTriggerPxType string `json:"newSlTriggerPxType,omitempty"` // 止损触发价类型
}

type AmendAlgoOrderResponse struct {
	api.Response
	Data []AmendAlgoOrderData `json:"data"`
}

type AmendAlgoOrderData struct {
	AlgoId      string `json:"algoId"`      // 订单ID
	AlgoClOrdId string `json:"algoClOrdId"` // 客户自定义策略订单ID
	ReqId       string `json:"reqId"`       // 用户自定义修改事件ID
	SCode       string `json:"sCode"`       // 事件执行结果的code，0代表成功
	SMsg        string `json:"sMsg"`        // 事件执行失败时的msg
}

// 修改策略止损限价单
func AmendAlgoOrderStopLimit(instId, algoId string, sz string, triggerPx, ordPx string) (api.IRequest, api.IResponse) {
	return AmendAlgoOrder(&AmendAlgoOrderParam{
		InstId:             instId,
		AlgoId:             algoId,
		CxlOnFail:          false,
		NewSz:              sz,
		NewSlTriggerPx:     triggerPx,
		NewSlOrdPx:         ordPx,
		NewSlTriggerPxType: api.TriggerPxTypeLast,
	})
}

// 修改策略止盈限价单
func AmendAlgoOrderProfitLimit(instId, algoId string, sz string, triggerPx, ordPx string) (api.IRequest, api.IResponse) {
	return AmendAlgoOrder(&AmendAlgoOrderParam{
		InstId:             instId,
		AlgoId:             algoId,
		CxlOnFail:          false,
		NewSz:              sz,
		NewTpTriggerPx:     triggerPx,
		NewTpOrdPx:         ordPx,
		NewTpTriggerPxType: api.TriggerPxTypeLast,
	})
}
