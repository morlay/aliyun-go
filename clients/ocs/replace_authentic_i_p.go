package ocs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReplaceAuthenticIPRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	OldAuthenticIP       string `position:"Query" name:"OldAuthenticIP"`
	NewAuthenticIP       string `position:"Query" name:"NewAuthenticIP"`
}

func (r ReplaceAuthenticIPRequest) Invoke(client *sdk.Client) (response *ReplaceAuthenticIPResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReplaceAuthenticIPRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ocs", "2015-03-01", "ReplaceAuthenticIP", "", "")

	resp := struct {
		*responses.BaseResponse
		ReplaceAuthenticIPResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReplaceAuthenticIPResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReplaceAuthenticIPResponse struct {
	RequestId string
}
