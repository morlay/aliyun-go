package ocs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyInstanceCapacityRequest struct {
	requests.RpcRequest
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	Capacity             int64  `position:"Query" name:"Capacity"`
}

func (req *ModifyInstanceCapacityRequest) Invoke(client *sdk.Client) (resp *ModifyInstanceCapacityResponse, err error) {
	req.InitWithApiInfo("Ocs", "2015-03-01", "ModifyInstanceCapacity", "", "")
	resp = &ModifyInstanceCapacityResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyInstanceCapacityResponse struct {
	responses.BaseResponse
	RequestId common.String
}
