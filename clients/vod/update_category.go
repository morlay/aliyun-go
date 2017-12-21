package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateCategoryRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CateId               int64  `position:"Query" name:"CateId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	CateName             string `position:"Query" name:"CateName"`
}

func (r UpdateCategoryRequest) Invoke(client *sdk.Client) (response *UpdateCategoryResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateCategoryRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "UpdateCategory", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateCategoryResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateCategoryResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateCategoryResponse struct {
	RequestId string
}
