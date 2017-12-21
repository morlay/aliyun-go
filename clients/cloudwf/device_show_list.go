package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeviceShowListRequest struct {
	requests.RpcRequest
	Dirc       string `position:"Query" name:"Dirc"`
	Page       int    `position:"Query" name:"Page"`
	Per        int    `position:"Query" name:"Per"`
	DeviceType int    `position:"Query" name:"DeviceType"`
	Sid        int64  `position:"Query" name:"Sid"`
}

func (req *DeviceShowListRequest) Invoke(client *sdk.Client) (resp *DeviceShowListResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "DeviceShowList", "", "")
	resp = &DeviceShowListResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeviceShowListResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
