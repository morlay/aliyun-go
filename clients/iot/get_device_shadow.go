package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetDeviceShadowRequest struct {
	requests.RpcRequest
	ShadowMessage string `position:"Query" name:"ShadowMessage"`
	DeviceName    string `position:"Query" name:"DeviceName"`
	ProductKey    string `position:"Query" name:"ProductKey"`
}

func (req *GetDeviceShadowRequest) Invoke(client *sdk.Client) (resp *GetDeviceShadowResponse, err error) {
	req.InitWithApiInfo("Iot", "2017-04-20", "GetDeviceShadow", "", "")
	resp = &GetDeviceShadowResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetDeviceShadowResponse struct {
	responses.BaseResponse
	RequestId     string
	Success       bool
	ErrorMessage  string
	ShadowMessage string
}
