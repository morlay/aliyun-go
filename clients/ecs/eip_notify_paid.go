package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type EipNotifyPaidRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Data                 string `position:"Query" name:"Data"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *EipNotifyPaidRequest) Invoke(client *sdk.Client) (resp *EipNotifyPaidResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "EipNotifyPaid", "ecs", "")
	resp = &EipNotifyPaidResponse{}
	err = client.DoAction(req, resp)
	return
}

type EipNotifyPaidResponse struct {
	responses.BaseResponse
	RequestId string
	Data      string
	Code      string
	Message   string
	Success   bool
}
