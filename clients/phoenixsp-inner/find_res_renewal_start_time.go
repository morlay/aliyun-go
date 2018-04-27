package phoenixsp_inner

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FindResRenewalStartTimeRequest struct {
	requests.RpcRequest
	Caller         string `position:"Query" name:"Caller"`
	InstanceIdList string `position:"Query" name:"InstanceIdList"`
	Aliuid         int64  `position:"Query" name:"Aliuid"`
	Bid            string `position:"Query" name:"Bid"`
	ResourceType   string `position:"Query" name:"ResourceType"`
}

func (req *FindResRenewalStartTimeRequest) Invoke(client *sdk.Client) (resp *FindResRenewalStartTimeResponse, err error) {
	req.InitWithApiInfo("Phoenixsp-inner", "2016-08-05", "FindResRenewalStartTime", "", "")
	resp = &FindResRenewalStartTimeResponse{}
	err = client.DoAction(req, resp)
	return
}

type FindResRenewalStartTimeResponse struct {
	responses.BaseResponse
	RequestId string
	Code      string
	Message   string
	Success   bool
	Data      int64
}
