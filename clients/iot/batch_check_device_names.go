package iot

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BatchCheckDeviceNamesRequest struct {
	requests.RpcRequest
	DeviceNames *BatchCheckDeviceNamesDeviceNameList `position:"Query" type:"Repeated" name:"DeviceName"`
	ProductKey  string                               `position:"Query" name:"ProductKey"`
}

func (req *BatchCheckDeviceNamesRequest) Invoke(client *sdk.Client) (resp *BatchCheckDeviceNamesResponse, err error) {
	req.InitWithApiInfo("Iot", "2018-01-20", "BatchCheckDeviceNames", "", "")
	resp = &BatchCheckDeviceNamesResponse{}
	err = client.DoAction(req, resp)
	return
}

type BatchCheckDeviceNamesResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorMessage string
	Data         BatchCheckDeviceNamesData
}

type BatchCheckDeviceNamesData struct {
	ApplyId int64
}

type BatchCheckDeviceNamesDeviceNameList []string

func (list *BatchCheckDeviceNamesDeviceNameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
