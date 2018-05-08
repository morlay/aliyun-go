package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyPrepayInstanceSpecRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	OperatorType         string `position:"Query" name:"OperatorType"`
	SystemDiskCategory   string `position:"Query" name:"SystemDiskCategory"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	MigrateAcrossZone    string `position:"Query" name:"MigrateAcrossZone"`
	InstanceType         string `position:"Query" name:"InstanceType"`
}

func (req *ModifyPrepayInstanceSpecRequest) Invoke(client *sdk.Client) (resp *ModifyPrepayInstanceSpecResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyPrepayInstanceSpec", "ecs", "")
	resp = &ModifyPrepayInstanceSpecResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyPrepayInstanceSpecResponse struct {
	responses.BaseResponse
	RequestId common.String
	OrderId   common.String
}
