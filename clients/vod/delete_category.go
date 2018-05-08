package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteCategoryRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CateId               int64  `position:"Query" name:"CateId"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (req *DeleteCategoryRequest) Invoke(client *sdk.Client) (resp *DeleteCategoryResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "DeleteCategory", "vod", "")
	resp = &DeleteCategoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCategoryResponse struct {
	responses.BaseResponse
	RequestId common.String
}
