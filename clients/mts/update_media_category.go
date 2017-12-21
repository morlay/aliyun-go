package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateMediaCategoryRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CateId               int64  `position:"Query" name:"CateId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
}

func (r UpdateMediaCategoryRequest) Invoke(client *sdk.Client) (response *UpdateMediaCategoryResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateMediaCategoryRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdateMediaCategory", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateMediaCategoryResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateMediaCategoryResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateMediaCategoryResponse struct {
	RequestId string
}
