package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CancelCopyImageRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId              string `position:"Query" name:"ImageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r CancelCopyImageRequest) Invoke(client *sdk.Client) (response *CancelCopyImageResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CancelCopyImageRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "CancelCopyImage", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		CancelCopyImageResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CancelCopyImageResponse

	err = client.DoAction(&req, &resp)
	return
}

type CancelCopyImageResponse struct {
	RequestId string
}
