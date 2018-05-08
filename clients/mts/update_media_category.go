package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UpdateMediaCategoryRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CateId               int64  `position:"Query" name:"CateId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
}

func (req *UpdateMediaCategoryRequest) Invoke(client *sdk.Client) (resp *UpdateMediaCategoryResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdateMediaCategory", "mts", "")
	resp = &UpdateMediaCategoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateMediaCategoryResponse struct {
	responses.BaseResponse
	RequestId common.String
}
