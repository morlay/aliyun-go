package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteMediaRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MediaIds             string `position:"Query" name:"MediaIds"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteMediaRequest) Invoke(client *sdk.Client) (resp *DeleteMediaResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "DeleteMedia", "", "")
	resp = &DeleteMediaResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteMediaResponse struct {
	responses.BaseResponse
	RequestId string
}
