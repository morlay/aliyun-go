package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddEditingProjectRequest struct {
	requests.RpcRequest
	CoverURL             string `position:"Query" name:"CoverURL"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	Timeline             string `position:"Query" name:"Timeline"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	Title                string `position:"Query" name:"Title"`
}

func (req *AddEditingProjectRequest) Invoke(client *sdk.Client) (resp *AddEditingProjectResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "AddEditingProject", "vod", "")
	resp = &AddEditingProjectResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddEditingProjectResponse struct {
	responses.BaseResponse
	RequestId common.String
	Project   AddEditingProjectProject
}

type AddEditingProjectProject struct {
	ProjectId    common.String
	CreationTime common.String
	ModifiedTime common.String
	Status       common.String
	Description  common.String
	Title        common.String
}
