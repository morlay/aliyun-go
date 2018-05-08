package ocs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RemoveAuthenticIPRequest struct {
	requests.RpcRequest
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	AuthenticIP          string `position:"Query" name:"AuthenticIP"`
}

func (req *RemoveAuthenticIPRequest) Invoke(client *sdk.Client) (resp *RemoveAuthenticIPResponse, err error) {
	req.InitWithApiInfo("Ocs", "2015-03-01", "RemoveAuthenticIP", "", "")
	resp = &RemoveAuthenticIPResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveAuthenticIPResponse struct {
	responses.BaseResponse
	RequestId common.String
}
