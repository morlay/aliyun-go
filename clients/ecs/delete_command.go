package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteCommandRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	CommandId            string `position:"Query" name:"CommandId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteCommandRequest) Invoke(client *sdk.Client) (response *DeleteCommandResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteCommandRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteCommand", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DeleteCommandResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteCommandResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteCommandResponse struct {
	RequestId string
}
