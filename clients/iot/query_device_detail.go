package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryDeviceDetailRequest struct {
	requests.RpcRequest
	IotId      string `position:"Query" name:"IotId"`
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
}

func (req *QueryDeviceDetailRequest) Invoke(client *sdk.Client) (resp *QueryDeviceDetailResponse, err error) {
	req.InitWithApiInfo("Iot", "2018-01-20", "QueryDeviceDetail", "", "")
	resp = &QueryDeviceDetailResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryDeviceDetailResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorMessage string
	Data         QueryDeviceDetailData
}

type QueryDeviceDetailData struct {
	IotId           string
	ProductKey      string
	ProductName     string
	DeviceName      string
	DeviceSecret    string
	FirmwareVersion string
	GmtCreate       string
	GmtActive       string
	GmtOnline       string
	Status          string
	IpAddress       string
	NodeType        int
	Region          string
}
