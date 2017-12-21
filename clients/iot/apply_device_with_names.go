package iot

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ApplyDeviceWithNamesRequest struct {
	requests.RpcRequest
	DeviceNames *ApplyDeviceWithNamesDeviceNameList `position:"Query" type:"Repeated" name:"DeviceName"`
	ProductKey  string                              `position:"Query" name:"ProductKey"`
}

func (req *ApplyDeviceWithNamesRequest) Invoke(client *sdk.Client) (resp *ApplyDeviceWithNamesResponse, err error) {
	req.InitWithApiInfo("Iot", "2017-04-20", "ApplyDeviceWithNames", "", "")
	resp = &ApplyDeviceWithNamesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ApplyDeviceWithNamesResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorMessage string
	ApplyId      int64
}

type ApplyDeviceWithNamesDeviceNameList []string

func (list *ApplyDeviceWithNamesDeviceNameList) UnmarshalJSON(data []byte) error {
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
