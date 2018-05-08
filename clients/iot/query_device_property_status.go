package iot

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryDevicePropertyStatusRequest struct {
	requests.RpcRequest
	IotId      string `position:"Query" name:"IotId"`
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
}

func (req *QueryDevicePropertyStatusRequest) Invoke(client *sdk.Client) (resp *QueryDevicePropertyStatusResponse, err error) {
	req.InitWithApiInfo("Iot", "2018-01-20", "QueryDevicePropertyStatus", "", "")
	resp = &QueryDevicePropertyStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryDevicePropertyStatusResponse struct {
	responses.BaseResponse
	RequestId    common.String
	Success      bool
	ErrorMessage common.String
	Data         QueryDevicePropertyStatusData
}

type QueryDevicePropertyStatusData struct {
	List QueryDevicePropertyStatusPropertyStatusInfoList
}

type QueryDevicePropertyStatusPropertyStatusInfo struct {
	Unit       common.String
	Identifier common.String
	DataType   common.String
	Time       common.String
	Value      common.String
	Name       common.String
}

type QueryDevicePropertyStatusPropertyStatusInfoList []QueryDevicePropertyStatusPropertyStatusInfo

func (list *QueryDevicePropertyStatusPropertyStatusInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryDevicePropertyStatusPropertyStatusInfo)
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
