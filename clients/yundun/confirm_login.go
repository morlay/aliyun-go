package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ConfirmLoginRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	SourceIp   string `position:"Query" name:"SourceIp"`
	Time       string `position:"Query" name:"Time"`
}

func (r ConfirmLoginRequest) Invoke(client *sdk.Client) (response *ConfirmLoginResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ConfirmLoginRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Yundun", "2015-04-16", "ConfirmLogin", "", "")

	resp := struct {
		*responses.BaseResponse
		ConfirmLoginResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ConfirmLoginResponse

	err = client.DoAction(&req, &resp)
	return
}

type ConfirmLoginResponse struct {
	RequestId string
}
