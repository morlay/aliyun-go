package ccc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	LoginInfo      RequestLoginInfoLoginInfo
}

type RequestLoginInfoLoginInfo struct {
	UserName       string
	DisplayName    string
	PhoneNumber    string
	Region         string
	WebRtcUrl      string
	AgentServerUrl string
	Extension      string
	TenantId       string
	Signature      string
	SignData       string
}
