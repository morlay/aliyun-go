package iot

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryDeviceEventDataRequest struct {
	requests.RpcRequest
	Asc        int    `position:"Query" name:"Asc"`
	Identifier string `position:"Query" name:"Identifier"`
	IotId      string `position:"Query" name:"IotId"`
	PageSize   int    `position:"Query" name:"PageSize"`
	EndTime    int64  `position:"Query" name:"EndTime"`
	EventType  string `position:"Query" name:"EventType"`
	DeviceName string `position:"Query" name:"DeviceName"`
	StartTime  int64  `position:"Query" name:"StartTime"`
	ProductKey string `position:"Query" name:"ProductKey"`
}

func (req *QueryDeviceEventDataRequest) Invoke(client *sdk.Client) (resp *QueryDeviceEventDataResponse, err error) {
	req.InitWithApiInfo("Iot", "2018-01-20", "QueryDeviceEventData", "", "")
	resp = &QueryDeviceEventDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryDeviceEventDataResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorMessage string
	Data         QueryDeviceEventDataData
}

type QueryDeviceEventDataData struct {
	NextTime  int64
	NextValid bool
	List      QueryDeviceEventDataEventInfoList
}

type QueryDeviceEventDataEventInfo struct {
	Time       string
	Identifier string
	Name       string
	EventType  string
	OutputData string
}

type QueryDeviceEventDataEventInfoList []QueryDeviceEventDataEventInfo

func (list *QueryDeviceEventDataEventInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryDeviceEventDataEventInfo)
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
