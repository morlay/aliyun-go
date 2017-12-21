package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddEditingProjectRequest struct {
	CoverURL             string `position:"Query" name:"CoverURL"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	Timeline             string `position:"Query" name:"Timeline"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	Title                string `position:"Query" name:"Title"`
}

func (r AddEditingProjectRequest) Invoke(client *sdk.Client) (response *AddEditingProjectResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddEditingProjectRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "AddEditingProject", "", "")

	resp := struct {
		*responses.BaseResponse
		AddEditingProjectResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddEditingProjectResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddEditingProjectResponse struct {
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
