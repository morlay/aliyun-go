package phoenixsp_inner

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CountUserAvailableResourcesByRequest struct {
	requests.RpcRequest
	BeginEndTime       int64  `position:"Query" name:"BeginEndTime"`
	EndResCreateTime   int64  `position:"Query" name:"EndResCreateTime"`
	BillingTag         string `position:"Query" name:"BillingTag"`
	CommodityCode      string `position:"Query" name:"CommodityCode"`
	AvailableDays      int    `position:"Query" name:"AvailableDays"`
	ResourceType       string `position:"Query" name:"ResourceType"`
	BeginResCreateTime int64  `position:"Query" name:"BeginResCreateTime"`
	Caller             string `position:"Query" name:"Caller"`
	AutoRenewal        string `position:"Query" name:"AutoRenewal"`
	EndEndTime         int64  `position:"Query" name:"EndEndTime"`
	InstanceIdList     string `position:"Query" name:"InstanceIdList"`
	HasExpiredRes      string `position:"Query" name:"HasExpiredRes"`
	Aliuid             int64  `position:"Query" name:"Aliuid"`
	IsPrepaid          string `position:"Query" name:"IsPrepaid"`
	Bid                string `position:"Query" name:"Bid"`
	Region             string `position:"Query" name:"Region"`
}

func (req *CountUserAvailableResourcesByRequest) Invoke(client *sdk.Client) (resp *CountUserAvailableResourcesByResponse, err error) {
	req.InitWithApiInfo("Phoenixsp-inner", "2016-08-05", "CountUserAvailableResourcesBy", "", "")
	resp = &CountUserAvailableResourcesByResponse{}
	err = client.DoAction(req, resp)
	return
}

type CountUserAvailableResourcesByResponse struct {
	responses.BaseResponse
	RequestId common.String
	Code      common.String
	Message   common.String
	Success   bool
	Data      common.Integer
}
