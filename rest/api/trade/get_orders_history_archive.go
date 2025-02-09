package trade

import (
	"github.com/smart-money-trader/okx-go/rest/api"
)

func NewGetOrdersHistoryArchive(param *GetOrdersQueryParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/orders-history-archive",
		Method: api.MethodGet,
		Param:  param,
	}, &GetOrderResponse{}
}

func OrdersHistoryArchiveSpot(instId string) (api.IRequest, api.IResponse) {
	return NewGetOrdersHistoryArchive(&GetOrdersQueryParam{
		InstType: api.InstTypeSPOT,
		// InstId:   instId,
		// OrdType:  "market,limit",
		// State:    api.OrdStateFilled,
	})
}
