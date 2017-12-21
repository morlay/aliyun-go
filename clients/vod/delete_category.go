package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteCategoryRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CateId               int64  `position:"Query" name:"CateId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (r DeleteCategoryRequest) Invoke(client *sdk.Client) (response *DeleteCategoryResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteCategoryRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "DeleteCategory", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteCategoryResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteCategoryResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteCategoryResponse struct {
	RequestId string
}
