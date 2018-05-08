package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetEditingProjectRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	ProjectId            string `position:"Query" name:"ProjectId"`
}

func (req *GetEditingProjectRequest) Invoke(client *sdk.Client) (resp *GetEditingProjectResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "GetEditingProject", "vod", "")
	resp = &GetEditingProjectResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetEditingProjectResponse struct {
	responses.BaseResponse
	RequestId common.String
	Project   GetEditingProjectProject
}

type GetEditingProjectProject struct {
	ProjectId    common.String
	CreationTime common.String
	ModifiedTime common.String
	Status       common.String
	Description  common.String
	Title        common.String
	Timeline     common.String
	CoverURL     common.String
}
