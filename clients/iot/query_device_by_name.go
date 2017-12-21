package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryDeviceByNameRequest struct {
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
}

func (r QueryDeviceByNameRequest) Invoke(client *sdk.Client) (response *QueryDeviceByNameResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryDeviceByNameRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Iot", "2017-04-20", "QueryDeviceByName", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryDeviceByNameResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryDeviceByNameResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryDeviceByNameResponse struct {
	RequestId    string
	Success      bool
	ErrorMessage string
	DeviceInfo   QueryDeviceByNameDeviceInfo
}

type QueryDeviceByNameDeviceInfo struct {
	DeviceId     string
	DeviceSecret string
	ProductKey   string
	DeviceStatus string
	DeviceName   string
	DeviceType   string
	GmtCreate    string
	GmtModified  string
}
