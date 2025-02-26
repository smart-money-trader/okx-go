package rubik

import (
	"log"
	"testing"

	"github.com/smart-money-trader/okx-go/examples/rest"
	"github.com/smart-money-trader/okx-go/rest/api/rubik"
)

func TestAbc(t *testing.T) {
	// nowBegin := time.Now().UnixMilli()
	// nowEnd := nowBegin + 30*60*1000
	// fmt.Println(nowBegin)
	// fmt.Println(nowEnd)
	param := &rubik.TakerVolumeContractParam{
		InstId: "BTC-USDT-SWAP",
		Period: "30m",
		Unit:   "1",
		// Begin:  fmt.Sprintf("%d", nowBegin),
		// End:    fmt.Sprintf("%d", nowEnd),
		Limit: "10",
	}
	req, resp := rubik.GetTakerVolumeContract(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*rubik.TakerVolumeContractResp))
}
