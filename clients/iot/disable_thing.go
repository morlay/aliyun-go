package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DisableThingRequest struct {
	requests.RpcRequest
	IotId      string `position:"Query" name:"IotId"`
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
}

func (req *DisableThingRequest) Invoke(client *sdk.Client) (resp *DisableThingResponse, err error) {
	req.InitWithApiInfo("Iot", "2018-01-20", "DisableThing", "", "")
	resp = &DisableThingResponse{}
	err = client.DoAction(req, resp)
	return
}

type DisableThingResponse struct {
	responses.BaseResponse
	RequestId    common.String
	Success      bool
	ErrorMessage common.String
}
