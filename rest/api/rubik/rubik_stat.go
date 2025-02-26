package rubik

import "github.com/smart-money-trader/okx-go/rest/api"

// GET 获取合约主动买入/卖出情况
func GetTakerVolumeContract(param *TakerVolumeContractParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/rubik/stat/taker-volume-contract",
		Method: api.MethodGet,
		Param:  param,
	}, &TakerVolumeContractResp{}
}

type TakerVolumeContractParam struct {
	InstId string `json:"instId"` // 产品ID
	Period string `json:"period"` // 时间粒度，默认值5m, 如 [5m/15m/30m/1H/2H/4H]
	Unit   string `json:"unit"`   // 买入、卖出的单位，默认值是1
	End    string `json:"end"`    // 筛选的结束时间戳 ts，Unix 时间戳为毫秒数格式，如 1597027383085
	Begin  string `json:"begin"`  // 筛选的开始时间戳 ts，Unix 时间戳为毫秒数格式，如 1597026383085
	Limit  string `json:"limit"`  // 分页返回的结果集数量，最大为100，不填默认返回100条
}

type TakerVolumeContractResp struct {
	api.Response
	Data [][]string `json:"data"`
}

// TakerVolumeContractResp.Data
// Ts      string `json:"ts"`      // 时间戳，Unix时间戳的毫秒数格式，如 1597026383085
// SellVol string `json:"sellVol"` // 卖出量
// BuyVol  string `json:"buyVol"`  // 买入量
