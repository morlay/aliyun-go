package ccc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RequestLoginInfoRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (req *RequestLoginInfoRequest) Invoke(client *sdk.Client) (resp *RequestLoginInfoResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "RequestLoginInfo", "ccc", "")
	resp = &RequestLoginInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type RequestLoginInfoResponse struct {
	responses.BaseResponse
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
	LoginInfo      RequestLoginInfoLoginInfo
}

type RequestLoginInfoLoginInfo struct {
	UserName       common.String
	DisplayName    common.String
	PhoneNumber    common.String
	Region         common.String
	WebRtcUrl      common.String
	AgentServerUrl common.String
	Extension      common.String
	TenantId       common.String
	Signature      common.String
	SignData       common.String
}
