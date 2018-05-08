package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId    common.String
	Success      bool
	ErrorMessage common.String
	Data         QueryDeviceDetailData
}

type QueryDeviceDetailData struct {
	IotId           common.String
	ProductKey      common.String
	ProductName     common.String
	DeviceName      common.String
	DeviceSecret    common.String
	FirmwareVersion common.String
	GmtCreate       common.String
	GmtActive       common.String
	GmtOnline       common.String
	Status          common.String
	IpAddress       common.String
	NodeType        common.Integer
	Region          common.String
}
