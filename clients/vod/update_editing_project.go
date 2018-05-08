package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UpdateEditingProjectRequest struct {
	requests.RpcRequest
	CoverURL             string `position:"Query" name:"CoverURL"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Timeline             string `position:"Query" name:"Timeline"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	Title                string `position:"Query" name:"Title"`
	ProjectId            string `position:"Query" name:"ProjectId"`
}

func (req *UpdateEditingProjectRequest) Invoke(client *sdk.Client) (resp *UpdateEditingProjectResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "UpdateEditingProject", "vod", "")
	resp = &UpdateEditingProjectResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateEditingProjectResponse struct {
	responses.BaseResponse
	RequestId common.String
}
