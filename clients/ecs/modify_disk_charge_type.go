package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyDiskChargeTypeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DiskIds              string `position:"Query" name:"DiskIds"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyDiskChargeTypeRequest) Invoke(client *sdk.Client) (resp *ModifyDiskChargeTypeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyDiskChargeType", "ecs", "")
	resp = &ModifyDiskChargeTypeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDiskChargeTypeResponse struct {
	responses.BaseResponse
	RequestId common.String
	OrderId   common.String
}
