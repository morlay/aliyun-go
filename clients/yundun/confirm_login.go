package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ConfirmLoginRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	SourceIp   string `position:"Query" name:"SourceIp"`
	Time       string `position:"Query" name:"Time"`
}

func (req *ConfirmLoginRequest) Invoke(client *sdk.Client) (resp *ConfirmLoginResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "ConfirmLogin", "", "")
	resp = &ConfirmLoginResponse{}
	err = client.DoAction(req, resp)
	return
}

type ConfirmLoginResponse struct {
	responses.BaseResponse
	RequestId string
}
