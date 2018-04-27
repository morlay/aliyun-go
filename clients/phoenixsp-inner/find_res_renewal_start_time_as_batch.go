package phoenixsp_inner

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FindResRenewalStartTimeAsBatchRequest struct {
	requests.RpcRequest
	Caller         string `position:"Query" name:"Caller"`
	InstanceIdList string `position:"Query" name:"InstanceIdList"`
	Aliuid         int64  `position:"Query" name:"Aliuid"`
	Bid            string `position:"Query" name:"Bid"`
	ResourceType   string `position:"Query" name:"ResourceType"`
}

func (req *FindResRenewalStartTimeAsBatchRequest) Invoke(client *sdk.Client) (resp *FindResRenewalStartTimeAsBatchResponse, err error) {
	req.InitWithApiInfo("Phoenixsp-inner", "2016-08-05", "FindResRenewalStartTimeAsBatch", "", "")
	resp = &FindResRenewalStartTimeAsBatchResponse{}
	err = client.DoAction(req, resp)
	return
}

type FindResRenewalStartTimeAsBatchResponse struct {
	responses.BaseResponse
	RequestId string
	Code      string
	Message   string
	Success   bool
	Data      int64
}
