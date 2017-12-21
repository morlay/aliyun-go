package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddCategoryRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	ParentId             int64  `position:"Query" name:"ParentId"`
	CateName             string `position:"Query" name:"CateName"`
}

func (r AddCategoryRequest) Invoke(client *sdk.Client) (response *AddCategoryResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddCategoryRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "AddCategory", "", "")

	resp := struct {
		*responses.BaseResponse
		AddCategoryResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddCategoryResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddCategoryResponse struct {
	RequestId string
	Category  AddCategoryCategory
}

type AddCategoryCategory struct {
	CateId   int64
	CateName string
	ParentId int64
	Level    int64
}
