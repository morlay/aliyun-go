package ocs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddAuthenticIPRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	AuthenticIP          string `position:"Query" name:"AuthenticIP"`
}

func (r AddAuthenticIPRequest) Invoke(client *sdk.Client) (response *AddAuthenticIPResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddAuthenticIPRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ocs", "2015-03-01", "AddAuthenticIP", "", "")

	resp := struct {
		*responses.BaseResponse
		AddAuthenticIPResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddAuthenticIPResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddAuthenticIPResponse struct {
	RequestId string
}
