package ocs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeAuthenticIPRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
}

func (r DescribeAuthenticIPRequest) Invoke(client *sdk.Client) (response *DescribeAuthenticIPResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeAuthenticIPRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ocs", "2015-03-01", "DescribeAuthenticIP", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeAuthenticIPResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeAuthenticIPResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeAuthenticIPResponse struct {
	RequestId              string
	GetAuthenticIpResponse DescribeAuthenticIPGetAuthenticIpResponse
}

type DescribeAuthenticIPGetAuthenticIpResponse struct {
	AuthenticIPs string
}
