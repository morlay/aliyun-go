package sts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetCallerIdentityRequest struct {
}

func (r GetCallerIdentityRequest) Invoke(client *sdk.Client) (response *GetCallerIdentityResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetCallerIdentityRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Sts", "2015-04-01", "GetCallerIdentity", "", "")

	resp := struct {
		*responses.BaseResponse
		GetCallerIdentityResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetCallerIdentityResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetCallerIdentityResponse struct {
	AccountId string
	UserId    string
	Arn       string
	RequestId string
}
