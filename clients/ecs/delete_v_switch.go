package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteVSwitchRequest struct {
	requests.RpcRequest
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteVSwitchRequest) Invoke(client *sdk.Client) (resp *DeleteVSwitchResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteVSwitch", "ecs", "")
	resp = &DeleteVSwitchResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteVSwitchResponse struct {
	responses.BaseResponse
	RequestId string
}
