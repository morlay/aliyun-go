package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId    common.String
	Success      bool
	ErrorMessage common.String
	Data         BatchRegisterDeviceData
}

type BatchRegisterDeviceData struct {
	ApplyId common.Long
}
