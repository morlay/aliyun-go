package ocs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ReplaceAuthenticIPRequest struct {
	requests.RpcRequest
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	OldAuthenticIP       string `position:"Query" name:"OldAuthenticIP"`
	NewAuthenticIP       string `position:"Query" name:"NewAuthenticIP"`
}

func (req *ReplaceAuthenticIPRequest) Invoke(client *sdk.Client) (resp *ReplaceAuthenticIPResponse, err error) {
	req.InitWithApiInfo("Ocs", "2015-03-01", "ReplaceAuthenticIP", "", "")
	resp = &ReplaceAuthenticIPResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReplaceAuthenticIPResponse struct {
	responses.BaseResponse
	RequestId common.String
}
