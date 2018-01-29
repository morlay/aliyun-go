package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CategoryTreeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CategoryTreeRequest) Invoke(client *sdk.Client) (resp *CategoryTreeResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "CategoryTree", "mts", "")
	resp = &CategoryTreeResponse{}
	err = client.DoAction(req, resp)
	return
}

type CategoryTreeResponse struct {
	responses.BaseResponse
	RequestId    string
	CategoryTree string
}
