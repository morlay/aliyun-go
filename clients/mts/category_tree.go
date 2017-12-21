package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CategoryTreeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r CategoryTreeRequest) Invoke(client *sdk.Client) (response *CategoryTreeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CategoryTreeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "CategoryTree", "", "")

	resp := struct {
		*responses.BaseResponse
		CategoryTreeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CategoryTreeResponse

	err = client.DoAction(&req, &resp)
	return
}

type CategoryTreeResponse struct {
	RequestId    string
	CategoryTree string
}
