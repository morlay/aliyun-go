package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	req.InitWithApiInfo("vod", "2017-03-21", "AddEditingProject", "", "")
	resp = &AddEditingProjectResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddEditingProjectResponse struct {
	responses.BaseResponse
	RequestId string
	Project   AddEditingProjectProject
}

type AddEditingProjectProject struct {
	ProjectId    string
	CreationTime string
	ModifiedTime string
	Status       string
	Description  string
	Title        string
}
