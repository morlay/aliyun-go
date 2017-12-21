package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeListenerAccessControlAttributeRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         int    `position:"Query" name:"ListenerPort"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r DescribeListenerAccessControlAttributeRequest) Invoke(client *sdk.Client) (response *DescribeListenerAccessControlAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeListenerAccessControlAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeListenerAccessControlAttribute", "slb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeListenerAccessControlAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeListenerAccessControlAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeListenerAccessControlAttributeResponse struct {
	RequestId           string
	AccessControlStatus string
	SourceItems         string
}
