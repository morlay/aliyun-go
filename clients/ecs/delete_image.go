package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteImageRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId              string `position:"Query" name:"ImageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Force                string `position:"Query" name:"Force"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteImageRequest) Invoke(client *sdk.Client) (response *DeleteImageResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteImageRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteImage", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DeleteImageResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteImageResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteImageResponse struct {
	RequestId string
}
