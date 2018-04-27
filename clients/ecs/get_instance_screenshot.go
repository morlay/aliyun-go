package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetInstanceScreenshotRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Wakeup               string `position:"Query" name:"Wakeup"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
}

func (req *GetInstanceScreenshotRequest) Invoke(client *sdk.Client) (resp *GetInstanceScreenshotResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "GetInstanceScreenshot", "ecs", "")
	resp = &GetInstanceScreenshotResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetInstanceScreenshotResponse struct {
	responses.BaseResponse
	RequestId  string
	InstanceId string
	Screenshot string
}
