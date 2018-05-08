package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteImageRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId              string `position:"Query" name:"ImageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Force                string `position:"Query" name:"Force"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteImageRequest) Invoke(client *sdk.Client) (resp *DeleteImageResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteImage", "ecs", "")
	resp = &DeleteImageResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteImageResponse struct {
	responses.BaseResponse
	RequestId common.String
}
