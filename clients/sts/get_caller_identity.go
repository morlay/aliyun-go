package sts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetCallerIdentityRequest struct {
	requests.RpcRequest
}

func (req *GetCallerIdentityRequest) Invoke(client *sdk.Client) (resp *GetCallerIdentityResponse, err error) {
	req.InitWithApiInfo("Sts", "2015-04-01", "GetCallerIdentity", "", "")
	resp = &GetCallerIdentityResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetCallerIdentityResponse struct {
	responses.BaseResponse
	AccountId common.String
	UserId    common.String
	Arn       common.String
	RequestId common.String
}
