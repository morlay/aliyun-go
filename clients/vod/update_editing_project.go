package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateEditingProjectRequest struct {
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

func (r UpdateEditingProjectRequest) Invoke(client *sdk.Client) (response *UpdateEditingProjectResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateEditingProjectRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "UpdateEditingProject", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateEditingProjectResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateEditingProjectResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateEditingProjectResponse struct {
	RequestId string
}
