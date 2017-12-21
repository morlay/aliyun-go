package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteEditingProjectRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ProjectIds           string `position:"Query" name:"ProjectIds"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (r DeleteEditingProjectRequest) Invoke(client *sdk.Client) (response *DeleteEditingProjectResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteEditingProjectRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "DeleteEditingProject", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteEditingProjectResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteEditingProjectResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteEditingProjectResponse struct {
	RequestId string
}
