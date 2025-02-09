package trade_algo

import "github.com/smart-money-trader/okx-go/rest/api"

// https://www.okx.com/docs-v5/zh/?shell#order-book-trading-algo-trading-get-algo-order-details

// GET 获取策略委托单信息
func AlgoOrderInfo(param *AlgoOrderInfoParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/order-algo",
		Method: api.MethodGet,
		Param:  param,
	}, &AlgoOrderInfoResponse{}
}

type AlgoOrderInfoParam struct {
	AlgoId      string `url:"algoId,omitempty"`      // 策略委托单ID
	AlgoClOrdId string `url:"algoClOrdId,omitempty"` // 客户自定义策略订单ID
}

type AlgoOrderInfoResponse struct {
	api.Response
	Data []AlgoOrderInfoData `json:"data"`
}

type AlgoOrderInfoData struct {
	InstType    string `json:"instType"`    // 产品类型
	InstId      string `json:"instId"`      // 产品ID
	Ccy         string `json:"ccy"`         // 保证金币种，仅适用于现货和合约模式下的全仓杠杆订单
	TgtCcy      string `json:"tgtCcy"`      // 币币市价单委托数量sz的单位
	AlgoId      string `json:"algoId"`      // 策略委托单ID
	OrdId       string `json:"ordId"`       // 最新一笔订单ID，即将废弃
	ClOrdId     string `json:"clOrdId"`     // 客户自定义订单ID
	AlgoClOrdId string `json:"algoClOrdId"` // 客户自定义策略订单ID
	Tag         string `json:"tag"`         // 订单标签
	Sz          string `json:"sz"`          // 委托数量
	OrdType     string `json:"ordType"`     // 订单类型
	Side        string `json:"side"`        // 订单方向
	PosSide     string `json:"posSide"`     // 持仓方向
	TdMode      string `json:"tdMode"`      // 交易模式

	State           string `json:"state"`           // 订单状态
	Lever           string `json:"lever"`           // 杠杆倍数，0.01到125之间的数值，仅适用于 币币杠杆/交割/永续
	TpTriggerPx     string `json:"tpTriggerPx"`     // 止盈触发价
	TpTriggerPxType string `json:"tpTriggerPxType"` // 止盈触发价类型
	TpOrdPx         string `json:"tpOrdPx"`         // 止盈委托价
	SlTriggerPx     string `json:"slTriggerPx"`     // 止损触发价
	SlTriggerPxType string `json:"slTriggerPxType"` // 止损触发价类型
	SlOrdPx         string `json:"slOrdPx"`         // 止损委托价

	TriggerPx     string `json:"triggerPx"`     // 计划委托触发价格
	TriggerPxType string `json:"triggerPxType"` // 计划委托触发价格类型
	OrdPx         string `json:"ordPx"`         // 计划委托单的委托价格
	ActualSz      string `json:"actualSz"`      // 实际委托量
	ActualPx      string `json:"actualPx"`      // 实际委托价
	ActualSide    string `json:"actualSide"`    // 实际触发方向

	TriggerTime string `json:"triggerTime"` // 策略委托触发时间，Unix时间戳的毫秒数格式

	// 仅适用于冰山委托和时间加权委托
	PxVar    string `json:"pxVar"`    // 价格比例
	PxSpread string `json:"pxSpread"` // 价距
	SzLimit  string `json:"szLimit"`  // 单笔数量
	PxLimit  string `json:"pxLimit"`  // 挂单限制价

	UTime int64 `json:"uTime,string"` // 订单更新时间，Unix时间戳的毫秒数格式
	CTime int64 `json:"cTime,string"` // 订单创建时间，Unix时间戳的毫秒数格式

	// 仅适用于时间加权委托
	TimeInterval string `json:"timeInterval"` // 下单间隔

	// 仅适用于移动止盈止损
	CallbackRatio  string `json:"callbackRatio"`  // 回调幅度的比例
	CallbackSpread string `json:"callbackSpread"` // 回调幅度的价距
	ActivePx       string `json:"activePx"`       // 移动止盈止损激活价格
	MoveTriggerPx  string `json:"moveTriggerPx"`  // 移动止盈止损触发价格

	ReduceOnly string `json:"reduceOnly"` // 是否只减仓
}

func (od AlgoOrderInfoData) IsFilled() bool {
	return od.State == api.OrdStateFilled
}
