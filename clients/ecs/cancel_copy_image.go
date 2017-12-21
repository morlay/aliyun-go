package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CancelCopyImageRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId              string `position:"Query" name:"ImageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CancelCopyImageRequest) Invoke(client *sdk.Client) (resp *CancelCopyImageResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "CancelCopyImage", "ecs", "")
	resp = &CancelCopyImageResponse{}
	err = client.DoAction(req, resp)
	return
}

type CancelCopyImageResponse struct {
	responses.BaseResponse
	RequestId string
}
