package ocs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveAuthenticIPRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	AuthenticIP          string `position:"Query" name:"AuthenticIP"`
}

func (r RemoveAuthenticIPRequest) Invoke(client *sdk.Client) (response *RemoveAuthenticIPResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RemoveAuthenticIPRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ocs", "2015-03-01", "RemoveAuthenticIP", "", "")

	resp := struct {
		*responses.BaseResponse
		RemoveAuthenticIPResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RemoveAuthenticIPResponse

	err = client.DoAction(&req, &resp)
	return
}

type RemoveAuthenticIPResponse struct {
	RequestId string
}
