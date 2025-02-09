package trade

import (
	"github.com/smart-money-trader/okx-go/rest/api"
)

func NewGetOrdersHistory(param *GetOrdersQueryParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/orders-history",
		Method: api.MethodGet,
		Param:  param,
	}, &GetOrderResponse{}
}

func GetOrdersHistory(instId string) (api.IRequest, api.IResponse) {
	return NewGetOrdersHistory(&GetOrdersQueryParam{
		InstId:   instId,
		InstType: api.InstTypeSPOT,
		OrdType:  "market,limit",
	})
}
