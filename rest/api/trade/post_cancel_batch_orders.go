package trade

import "github.com/smart-money-trader/okx-go/rest/api"

func NewPostCancelBatchOrders(param []*PostCancelOrderParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/cancel-batch-orders",
		Method: api.MethodPost,
		Param:  param,
	}, &PostCancelOrderResponse{}
}
