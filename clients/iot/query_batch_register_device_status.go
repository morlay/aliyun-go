package iot

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryBatchRegisterDeviceStatusRequest struct {
	requests.RpcRequest
	ApplyId    int64  `position:"Query" name:"ApplyId"`
	ProductKey string `position:"Query" name:"ProductKey"`
}

func (req *QueryBatchRegisterDeviceStatusRequest) Invoke(client *sdk.Client) (resp *QueryBatchRegisterDeviceStatusResponse, err error) {
	req.InitWithApiInfo("Iot", "2018-01-20", "QueryBatchRegisterDeviceStatus", "", "")
	resp = &QueryBatchRegisterDeviceStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryBatchRegisterDeviceStatusResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorMessage string
	Data         QueryBatchRegisterDeviceStatusData
}

type QueryBatchRegisterDeviceStatusData struct {
	Status      string
	ValidList   QueryBatchRegisterDeviceStatusValidListList
	InvalidList QueryBatchRegisterDeviceStatusInvalidListList
}

type QueryBatchRegisterDeviceStatusValidListList []string

func (list *QueryBatchRegisterDeviceStatusValidListList) UnmarshalJSON(data []byte) error {
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

type QueryBatchRegisterDeviceStatusInvalidListList []string

func (list *QueryBatchRegisterDeviceStatusInvalidListList) UnmarshalJSON(data []byte) error {
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
