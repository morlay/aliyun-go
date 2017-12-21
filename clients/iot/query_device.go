package iot

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryDeviceRequest struct {
	PageSize    int    `position:"Query" name:"PageSize"`
	CurrentPage int    `position:"Query" name:"CurrentPage"`
	ProductKey  string `position:"Query" name:"ProductKey"`
}

func (r QueryDeviceRequest) Invoke(client *sdk.Client) (response *QueryDeviceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryDeviceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Iot", "2017-04-20", "QueryDevice", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryDeviceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryDeviceResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryDeviceResponse struct {
	RequestId    string
	Success      bool
	ErrorMessage string
	Total        int
	PageSize     int
	PageCount    int
	Page         int
	Data         QueryDeviceDeviceInfoList
}

type QueryDeviceDeviceInfo struct {
	DeviceId     string
	DeviceSecret string
	ProductKey   string
	DeviceStatus string
	DeviceName   string
	DeviceType   string
	GmtCreate    string
	GmtModified  string
}

type QueryDeviceDeviceInfoList []QueryDeviceDeviceInfo

func (list *QueryDeviceDeviceInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryDeviceDeviceInfo)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
