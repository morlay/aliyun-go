package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RenewInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int    `position:"Query" name:"Period"`
	PeriodUnit           string `position:"Query" name:"PeriodUnit"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *RenewInstanceRequest) Invoke(client *sdk.Client) (resp *RenewInstanceResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "RenewInstance", "ecs", "")
	resp = &RenewInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type RenewInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
