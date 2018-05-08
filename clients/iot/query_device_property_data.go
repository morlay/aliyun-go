package iot

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryDevicePropertyDataRequest struct {
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

func (req *QueryDevicePropertyDataRequest) Invoke(client *sdk.Client) (resp *QueryDevicePropertyDataResponse, err error) {
	req.InitWithApiInfo("Iot", "2018-01-20", "QueryDevicePropertyData", "", "")
	resp = &QueryDevicePropertyDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryDevicePropertyDataResponse struct {
	responses.BaseResponse
	RequestId    common.String
	Success      bool
	ErrorMessage common.String
	Data         QueryDevicePropertyDataData
}

type QueryDevicePropertyDataData struct {
	NextValid bool
	NextTime  common.Long
	List      QueryDevicePropertyDataPropertyInfoList
}

type QueryDevicePropertyDataPropertyInfo struct {
	Time  common.String
	Value common.String
}

type QueryDevicePropertyDataPropertyInfoList []QueryDevicePropertyDataPropertyInfo

func (list *QueryDevicePropertyDataPropertyInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryDevicePropertyDataPropertyInfo)
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
