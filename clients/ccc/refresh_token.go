package ccc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RefreshTokenRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (req *RefreshTokenRequest) Invoke(client *sdk.Client) (resp *RefreshTokenResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "RefreshToken", "ccc", "")
	resp = &RefreshTokenResponse{}
	err = client.DoAction(req, resp)
	return
}

type RefreshTokenResponse struct {
	responses.BaseResponse
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
	Token          RefreshTokenToken
}

type RefreshTokenToken struct {
	Signature common.String
	SignData  common.String
}
