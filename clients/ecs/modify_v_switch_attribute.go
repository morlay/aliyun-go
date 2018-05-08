package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyVSwitchAttributeRequest struct {
	requests.RpcRequest
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VSwitchName          string `position:"Query" name:"VSwitchName"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyVSwitchAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyVSwitchAttributeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyVSwitchAttribute", "ecs", "")
	resp = &ModifyVSwitchAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyVSwitchAttributeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
