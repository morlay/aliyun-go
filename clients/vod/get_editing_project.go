package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetEditingProjectRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	ProjectId            string `position:"Query" name:"ProjectId"`
}

func (r GetEditingProjectRequest) Invoke(client *sdk.Client) (response *GetEditingProjectResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetEditingProjectRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "GetEditingProject", "", "")

	resp := struct {
		*responses.BaseResponse
		GetEditingProjectResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetEditingProjectResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetEditingProjectResponse struct {
	RequestId string
	Project   GetEditingProjectProject
}

type GetEditingProjectProject struct {
	ProjectId    string
	CreationTime string
	ModifiedTime string
	Status       string
	Description  string
	Title        string
	Timeline     string
	CoverURL     string
}
