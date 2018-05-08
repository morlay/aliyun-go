package ocs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeAuthenticIPRequest struct {
	requests.RpcRequest
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
}

func (req *DescribeAuthenticIPRequest) Invoke(client *sdk.Client) (resp *DescribeAuthenticIPResponse, err error) {
	req.InitWithApiInfo("Ocs", "2015-03-01", "DescribeAuthenticIP", "", "")
	resp = &DescribeAuthenticIPResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAuthenticIPResponse struct {
	responses.BaseResponse
	RequestId              common.String
	GetAuthenticIpResponse DescribeAuthenticIPGetAuthenticIpResponse
}

type DescribeAuthenticIPGetAuthenticIpResponse struct {
	AuthenticIPs common.String
}
