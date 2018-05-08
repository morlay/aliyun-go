package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyIntranetBandwidthKbRequest struct {
	requests.RpcRequest
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	IntranetMaxBandwidthOut int    `position:"Query" name:"IntranetMaxBandwidthOut"`
	InstanceId              string `position:"Query" name:"InstanceId"`
	IntranetMaxBandwidthIn  int    `position:"Query" name:"IntranetMaxBandwidthIn"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount            string `position:"Query" name:"OwnerAccount"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyIntranetBandwidthKbRequest) Invoke(client *sdk.Client) (resp *ModifyIntranetBandwidthKbResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyIntranetBandwidthKb", "ecs", "")
	resp = &ModifyIntranetBandwidthKbResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyIntranetBandwidthKbResponse struct {
	responses.BaseResponse
	RequestId common.String
}
