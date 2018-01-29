package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateCategoryNameRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CateId               string `position:"Query" name:"CateId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CateName             string `position:"Query" name:"CateName"`
}

func (req *UpdateCategoryNameRequest) Invoke(client *sdk.Client) (resp *UpdateCategoryNameResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdateCategoryName", "mts", "")
	resp = &UpdateCategoryNameResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateCategoryNameResponse struct {
	responses.BaseResponse
	RequestId string
}
