package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	req.InitWithApiInfo("Mts", "2014-06-18", "AddCategory", "mts", "")
	resp = &AddCategoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddCategoryResponse struct {
	responses.BaseResponse
	RequestId common.String
	Category  AddCategoryCategory
}

type AddCategoryCategory struct {
	CateId   common.String
	CateName common.String
	ParentId common.String
	Level    common.String
}
