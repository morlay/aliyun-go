package iot

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ApplyDeviceWithNamesRequest struct {
	DeviceNames *ApplyDeviceWithNamesDeviceNameList `position:"Query" type:"Repeated" name:"DeviceName"`
	ProductKey  string                              `position:"Query" name:"ProductKey"`
}

func (r ApplyDeviceWithNamesRequest) Invoke(client *sdk.Client) (response *ApplyDeviceWithNamesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ApplyDeviceWithNamesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Iot", "2017-04-20", "ApplyDeviceWithNames", "", "")

	resp := struct {
		*responses.BaseResponse
		ApplyDeviceWithNamesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ApplyDeviceWithNamesResponse

	err = client.DoAction(&req, &resp)
	return
}

type ApplyDeviceWithNamesResponse struct {
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
