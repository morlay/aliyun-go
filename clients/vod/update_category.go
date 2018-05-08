package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UpdateCategoryRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CateId               int64  `position:"Query" name:"CateId"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	CateName             string `position:"Query" name:"CateName"`
}

func (req *UpdateCategoryRequest) Invoke(client *sdk.Client) (resp *UpdateCategoryResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "UpdateCategory", "vod", "")
	resp = &UpdateCategoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateCategoryResponse struct {
	responses.BaseResponse
	RequestId common.String
}
