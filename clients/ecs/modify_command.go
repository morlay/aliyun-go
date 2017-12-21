package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyCommandRequest struct {
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

func (r ModifyCommandRequest) Invoke(client *sdk.Client) (response *ModifyCommandResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyCommandRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyCommand", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ModifyCommandResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyCommandResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyCommandResponse struct {
	RequestId string
}
