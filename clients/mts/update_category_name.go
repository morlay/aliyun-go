package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateCategoryNameRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CateId               string `position:"Query" name:"CateId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CateName             string `position:"Query" name:"CateName"`
}

func (r UpdateCategoryNameRequest) Invoke(client *sdk.Client) (response *UpdateCategoryNameResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateCategoryNameRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdateCategoryName", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateCategoryNameResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateCategoryNameResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateCategoryNameResponse struct {
	RequestId string
}
