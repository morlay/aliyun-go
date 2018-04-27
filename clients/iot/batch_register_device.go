package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BatchRegisterDeviceRequest struct {
	requests.RpcRequest
	Count      int    `position:"Query" name:"Count"`
	ProductKey string `position:"Query" name:"ProductKey"`
}

func (req *BatchRegisterDeviceRequest) Invoke(client *sdk.Client) (resp *BatchRegisterDeviceResponse, err error) {
	req.InitWithApiInfo("Iot", "2018-01-20", "BatchRegisterDevice", "", "")
	resp = &BatchRegisterDeviceResponse{}
	err = client.DoAction(req, resp)
	return
}

type BatchRegisterDeviceResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorMessage string
	Data         BatchRegisterDeviceData
}

type BatchRegisterDeviceData struct {
	ApplyId int64
}
