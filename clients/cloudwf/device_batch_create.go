package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeviceBatchCreateRequest struct {
	requests.RpcRequest
	Sn         string `position:"Query" name:"Sn"`
	DeviceType int    `position:"Query" name:"DeviceType"`
}

func (req *DeviceBatchCreateRequest) Invoke(client *sdk.Client) (resp *DeviceBatchCreateResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "DeviceBatchCreate", "", "")
	resp = &DeviceBatchCreateResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeviceBatchCreateResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
