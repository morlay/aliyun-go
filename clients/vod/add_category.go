package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddCategoryRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	ParentId             int64  `position:"Query" name:"ParentId"`
	CateName             string `position:"Query" name:"CateName"`
}

func (req *AddCategoryRequest) Invoke(client *sdk.Client) (resp *AddCategoryResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "AddCategory", "vod", "")
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
	CateId   common.Long
	CateName common.String
	ParentId common.Long
	Level    common.Long
}
