package trade

import (
	"github.com/smart-money-trader/okx-go/rest/api"
)

func NewGetFillsHistory(param *GetFillsParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/fills-history",
		Method: api.MethodGet,
		Param:  param,
	}, &GetFillsResponse{}
}
