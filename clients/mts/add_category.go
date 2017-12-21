package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddCategoryRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ParentId             int64  `position:"Query" name:"ParentId"`
	CateName             string `position:"Query" name:"CateName"`
}

func (req *AddCategoryRequest) Invoke(client *sdk.Client) (resp *AddCategoryResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "AddCategory", "", "")
	resp = &AddCategoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddCategoryResponse struct {
	responses.BaseResponse
	RequestId string
	Category  AddCategoryCategory
}

type AddCategoryCategory struct {
	CateId   string
	CateName string
	ParentId string
	Level    string
}
