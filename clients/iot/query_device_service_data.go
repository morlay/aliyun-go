package iot

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryDeviceServiceDataRequest struct {
	requests.RpcRequest
	Asc        int    `position:"Query" name:"Asc"`
	Identifier string `position:"Query" name:"Identifier"`
	IotId      string `position:"Query" name:"IotId"`
	PageSize   int    `position:"Query" name:"PageSize"`
	EndTime    int64  `position:"Query" name:"EndTime"`
	DeviceName string `position:"Query" name:"DeviceName"`
	StartTime  int64  `position:"Query" name:"StartTime"`
	ProductKey string `position:"Query" name:"ProductKey"`
}

func (req *QueryDeviceServiceDataRequest) Invoke(client *sdk.Client) (resp *QueryDeviceServiceDataResponse, err error) {
	req.InitWithApiInfo("Iot", "2018-01-20", "QueryDeviceServiceData", "", "")
	resp = &QueryDeviceServiceDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryDeviceServiceDataResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorMessage string
	Data         QueryDeviceServiceDataData
}

type QueryDeviceServiceDataData struct {
	NextTime  int64
	NextValid bool
	List      QueryDeviceServiceDataServiceInfoList
}

type QueryDeviceServiceDataServiceInfo struct {
	Time       string
	Identifier string
	Name       string
	InputData  string
	OutputData string
}

type QueryDeviceServiceDataServiceInfoList []QueryDeviceServiceDataServiceInfo

func (list *QueryDeviceServiceDataServiceInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryDeviceServiceDataServiceInfo)
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
