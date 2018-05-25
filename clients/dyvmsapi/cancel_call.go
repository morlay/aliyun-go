package dyvmsapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CancelCallRequest struct {
	requests.RpcRequest
	CallId               string `position:"Query" name:"CallId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CancelCallRequest) Invoke(client *sdk.Client) (resp *CancelCallResponse, err error) {
	req.InitWithApiInfo("Dyvmsapi", "2017-05-25", "CancelCall", "", "")
	resp = &CancelCallResponse{}
	err = client.DoAction(req, resp)
	return
}

type CancelCallResponse struct {
	responses.BaseResponse
	RequestId common.String
	Status    bool
	Code      common.String
	Message   common.String
}
