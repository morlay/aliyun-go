package phoenixsp_inner

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CountRegionAvailableResourcesByRequest struct {
	requests.RpcRequest
	MarketType         string `position:"Query" name:"MarketType"`
	Filter             string `position:"Query" name:"Filter"`
	Caller             string `position:"Query" name:"Caller"`
	ResCreateTimeEnd   int64  `position:"Query" name:"ResCreateTimeEnd"`
	ResourceSource     string `position:"Query" name:"ResourceSource"`
	AliUID             int64  `position:"Query" name:"AliUID"`
	ChargeType         string `position:"Query" name:"ChargeType"`
	CommodityCode      string `position:"Query" name:"CommodityCode"`
	Bid                string `position:"Query" name:"Bid"`
	Region             string `position:"Query" name:"Region"`
	ResourceType       string `position:"Query" name:"ResourceType"`
	ResCreateTimeBegin int64  `position:"Query" name:"ResCreateTimeBegin"`
}

func (req *CountRegionAvailableResourcesByRequest) Invoke(client *sdk.Client) (resp *CountRegionAvailableResourcesByResponse, err error) {
	req.InitWithApiInfo("Phoenixsp-inner", "2016-08-05", "CountRegionAvailableResourcesBy", "", "")
	resp = &CountRegionAvailableResourcesByResponse{}
	err = client.DoAction(req, resp)
	return
}

type CountRegionAvailableResourcesByResponse struct {
	responses.BaseResponse
	RequestId string
	Code      string
	Message   string
	Success   bool
	Data      int
}
