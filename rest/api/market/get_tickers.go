package market

import "github.com/smart-money-trader/okx-go/rest/api"

// https://www.okx.com/docs-v5/zh/#order-book-trading-market-data-get-tickers

// GET 获取所有产品行情信息
// 获取产品行情信息
func NewGetTickers(param *GetTickersParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/market/tickers",
		Method: api.MethodGet,
		Param:  param,
	}, &GetTickersResponse{}
}

type GetTickersParam struct {
	InstType   string `url:"instType"`             // 产品类型
	Uly        string `url:"uly,omitempty"`        // 标的指数
	InstFamily string `url:"instFamily,omitempty"` // 交易品种
}

type GetTickersResponse struct {
	api.Response
	Data []Ticker `json:"data"`
}
