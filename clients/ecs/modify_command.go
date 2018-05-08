package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyCommandRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	WorkingDir           string `position:"Query" name:"WorkingDir"`
	Description          string `position:"Query" name:"Description"`
	CommandId            string `position:"Query" name:"CommandId"`
	CommandContent       string `position:"Query" name:"CommandContent"`
	Timeout              int64  `position:"Query" name:"Timeout"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Name                 string `position:"Query" name:"Name"`
}

func (req *ModifyCommandRequest) Invoke(client *sdk.Client) (resp *ModifyCommandResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyCommand", "ecs", "")
	resp = &ModifyCommandResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyCommandResponse struct {
	responses.BaseResponse
	RequestId common.String
}
