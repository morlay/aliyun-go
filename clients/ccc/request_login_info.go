package ccc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RequestLoginInfoRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (r RequestLoginInfoRequest) Invoke(client *sdk.Client) (response *RequestLoginInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RequestLoginInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CCC", "2017-07-05", "RequestLoginInfo", "ccc", "")

	resp := struct {
		*responses.BaseResponse
		RequestLoginInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RequestLoginInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type RequestLoginInfoResponse struct {
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
	Region         string
	WebRtcUrl      string
	AgentServerUrl string
	Extension      string
	TenantId       string
	Signature      string
	SignData       string
}
