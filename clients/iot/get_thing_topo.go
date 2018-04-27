package iot

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetThingTopoRequest struct {
	requests.RpcRequest
	IotId      string `position:"Query" name:"IotId"`
	PageNo     int    `position:"Query" name:"PageNo"`
	PageSize   int    `position:"Query" name:"PageSize"`
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
}

func (req *GetThingTopoRequest) Invoke(client *sdk.Client) (resp *GetThingTopoResponse, err error) {
	req.InitWithApiInfo("Iot", "2018-01-20", "GetThingTopo", "", "")
	resp = &GetThingTopoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetThingTopoResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorMessage string
	Data         GetThingTopoData
}

type GetThingTopoData struct {
	Total       int64
	CurrentPage int
	PageSize    int
	PageCount   int64
	List        GetThingTopoDeviceInfoList
}

type GetThingTopoDeviceInfo struct {
	IotId      string
	ProductKey string
	DeviceName string
}

type GetThingTopoDeviceInfoList []GetThingTopoDeviceInfo

func (list *GetThingTopoDeviceInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetThingTopoDeviceInfo)
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
