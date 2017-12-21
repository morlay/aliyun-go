package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyPrepayInstanceSpecRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyPrepayInstanceSpecRequest) Invoke(client *sdk.Client) (resp *ModifyPrepayInstanceSpecResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyPrepayInstanceSpec", "ecs", "")
	resp = &ModifyPrepayInstanceSpecResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyPrepayInstanceSpecResponse struct {
	responses.BaseResponse
	RequestId string
	OrderId   string
}
