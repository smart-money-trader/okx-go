package account

import "github.com/smart-money-trader/okx-go/rest/api"

func NewGetBillsArchive(param *GetBillsParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/account/bills-archive",
		Method: api.MethodGet,
		Param:  param,
	}, &GetBillsResponse{}
}
