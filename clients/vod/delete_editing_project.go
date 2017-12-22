package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteEditingProjectRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ProjectIds           string `position:"Query" name:"ProjectIds"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (req *DeleteEditingProjectRequest) Invoke(client *sdk.Client) (resp *DeleteEditingProjectResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "DeleteEditingProject", "vod", "")
	resp = &DeleteEditingProjectResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteEditingProjectResponse struct {
	responses.BaseResponse
	RequestId string
}
