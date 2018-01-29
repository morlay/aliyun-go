package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteMediaTagRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Tag                  string `position:"Query" name:"Tag"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
}

func (req *DeleteMediaTagRequest) Invoke(client *sdk.Client) (resp *DeleteMediaTagResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "DeleteMediaTag", "mts", "")
	resp = &DeleteMediaTagResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteMediaTagResponse struct {
	responses.BaseResponse
	RequestId string
}
