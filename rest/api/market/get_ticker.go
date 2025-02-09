package market

import "github.com/smart-money-trader/okx-go/rest/api"

// https://www.okx.com/docs-v5/zh/#order-book-trading-market-data-get-ticker

// GET 获取单个产品行情信息
// 获取产品行情信息
func NewGetTicker(param *GetTickerParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/market/ticker",
		Method: api.MethodGet,
		Param:  param,
	}, &GetTickerResponse{}
}

type GetTickerParam struct {
	InstId string `url:"instId"` // 产品ID 如 BTC-USD-SWAP
}

type GetTickerResponse struct {
	api.Response
	Data []Ticker `json:"data"`
}

type Ticker struct {
	InstType  string `json:"instType"`  // 产品类型
	InstId    string `json:"instId"`    // 产品ID
	Last      string `json:"last"`      // 最新成交价
	LastSz    string `json:"lastSz"`    // 最新成交的数量，0 代表没有成交量
	AskPx     string `json:"askPx"`     // 卖一价
	AskSz     string `json:"askSz"`     // 卖一价的挂单数数量
	BidPx     string `json:"bidPx"`     // 买一价
	BidSz     string `json:"bidSz"`     // 买一价的挂单数量
	Open24h   string `json:"open24h"`   // 24小时开盘价
	High24h   string `json:"high24h"`   // 24小时最高价
	Low24h    string `json:"low24h"`    // 24小时最低价
	VolCcy24h string `json:"volCcy24h"` // 24小时成交量，以币为单位
	Vol24h    string `json:"vol24h"`    // 24小时成交量，以张为单位 如果是币币/币币杠杆，数值为交易货币的数量
	SodUtc0   string `json:"sodUtc0"`   // UTC 0 时开盘价
	SodUtc8   string `json:"sodUtc8"`   // UTC+8 时开盘价
	Ts        int64  `json:"ts,string"` // ticker数据产生时间，Unix时间戳的毫秒数格式
}
