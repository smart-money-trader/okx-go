package trade_algo

import "github.com/smart-money-trader/okx-go/rest/api"

// https://www.okx.com/docs-v5/zh/#order-book-trading-algo-trading-post-place-algo-order

// POST 策略委托下单
// 提供单向止盈止损委托、双向止盈止损委托、计划委托、时间加权委托、移动止盈止损委托
func NewAlgoOrder(param *AlgoOrderParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/order-algo",
		Method: api.MethodPost,
		Param:  param,
	}, &AlgoOrderResponse{}
}

type AlgoOrderParam struct {
	InstId        string `json:"instId"`                  // 产品ID，如 BTC-USDT
	TdMode        string `json:"tdMode"`                  // 交易模式
	Ccy           string `json:"ccy,omitempty"`           // 保证金币种
	Side          string `json:"side"`                    // 订单方向
	PosSide       string `json:"posSide,omitempty"`       // 持仓方向
	OrdType       string `json:"ordType"`                 // 订单类型
	Sz            string `json:"sz"`                      // 委托数量
	Tag           string `json:"tag,omitempty"`           // 订单标签
	TgtCcy        string `json:"tgtCcy,omitempty"`        // 委托数量的类型
	AlgoClOrdId   string `json:"algoClOrdId,omitempty"`   // 客户自定义策略订单ID
	CloseFraction string `json:"closeFraction,omitempty"` // 策略委托触发时，平仓的百分比

	TpOrdKind       string `json:"tpOrdKind,omitempty"`       // 止盈止损订单类型
	TpTriggerPx     string `json:"tpTriggerPx,omitempty"`     // 止盈触发价
	TpTriggerPxType string `json:"tpTriggerPxType,omitempty"` // 止盈触发价类型
	TpOrdPx         string `json:"tpOrdPx,omitempty"`         // 止盈委托价

	SlTriggerPx     string `json:"slTriggerPx,omitempty"`     // 止损触发价
	SlTriggerPxType string `json:"slTriggerPxType,omitempty"` // 止损触发价类型
	SlOrdPx         string `json:"slOrdPx,omitempty"`         // 止损委托价

	CxlOnClosePos bool `json:"cxlOnClosePos,omitempty"` // 决定用户所下的止盈止损订单是否与该交易产品对应的仓位关联
	ReduceOnly    bool `json:"reduceOnly,omitempty"`    // 是否只减仓，true 或 false，默认false
}

type AlgoOrderResponse struct {
	api.Response
	Data []AlgoOrderData `json:"data"`
}

type AlgoOrderData struct {
	AlgoId      string `json:"algoId"`
	ClOrdId     string `json:"clOrdId"`
	AlgoClOrdId string `json:"algoClOrdId"`
	Tag         string `json:"tag"`
	SCode       string `json:"sCode"`
	SMsg        string `json:"sMsg"`
}

// 单向止盈止损
// 限价止损单
func NewOrderAlgoStopLimitSell(instId string, sz string, triggerPx, ordPx string) (api.IRequest, api.IResponse) {
	return NewAlgoOrder(&AlgoOrderParam{
		InstId:          instId,
		TdMode:          api.TdModeCash,
		Side:            api.SideSell,
		OrdType:         api.OrdTypeConditional,
		Sz:              sz,
		SlTriggerPx:     triggerPx,
		SlTriggerPxType: api.TriggerPxTypeLast,
		SlOrdPx:         ordPx,
		Tag:             "bot",
	})
}

// 单向止盈止损
// 限价止盈单
func NewOrderAlgoProfitLimitSell(instId string, sz string, triggerPx, ordPx string) (api.IRequest, api.IResponse) {
	return NewAlgoOrder(&AlgoOrderParam{
		InstId:          instId,
		TdMode:          api.TdModeCash,
		Side:            api.SideSell,
		OrdType:         api.OrdTypeConditional,
		Sz:              sz,
		TpTriggerPx:     triggerPx,
		TpTriggerPxType: api.TriggerPxTypeLast,
		TpOrdPx:         ordPx,
		Tag:             "bot",
	})
}
