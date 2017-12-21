package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateDeviceShadowRequest struct {
	requests.RpcRequest
	ShadowMessage string `position:"Query" name:"ShadowMessage"`
	DeviceName    string `position:"Query" name:"DeviceName"`
	ProductKey    string `position:"Query" name:"ProductKey"`
}

func (req *UpdateDeviceShadowRequest) Invoke(client *sdk.Client) (resp *UpdateDeviceShadowResponse, err error) {
	req.InitWithApiInfo("Iot", "2017-04-20", "UpdateDeviceShadow", "", "")
	resp = &UpdateDeviceShadowResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateDeviceShadowResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorMessage string
}
