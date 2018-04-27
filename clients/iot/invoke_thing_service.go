package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type InvokeThingServiceRequest struct {
	requests.RpcRequest
	Args       string `position:"Query" name:"Args"`
	Identifier string `position:"Query" name:"Identifier"`
	IotId      string `position:"Query" name:"IotId"`
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
}

func (req *InvokeThingServiceRequest) Invoke(client *sdk.Client) (resp *InvokeThingServiceResponse, err error) {
	req.InitWithApiInfo("Iot", "2018-01-20", "InvokeThingService", "", "")
	resp = &InvokeThingServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type InvokeThingServiceResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorMessage string
	Data         InvokeThingServiceData
}

type InvokeThingServiceData struct {
	Result string
}
