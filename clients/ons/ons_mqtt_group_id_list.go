package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsMqttGroupIdListRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
}

func (req *OnsMqttGroupIdListRequest) Invoke(client *sdk.Client) (resp *OnsMqttGroupIdListResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMqttGroupIdList", "", "")
	resp = &OnsMqttGroupIdListResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsMqttGroupIdListResponse struct {
	responses.BaseResponse
	RequestId string
	HelpUrl   string
	Data      OnsMqttGroupIdListMqttGroupIdDoList
}

type OnsMqttGroupIdListMqttGroupIdDo struct {
	Id          int64
	ChannelId   int
	OnsRegionId string
	RegionName  string
	Owner       string
	GroupId     string
	Topic       string
	Status      int
	CreateTime  int64
	UpdateTime  int64
}

type OnsMqttGroupIdListMqttGroupIdDoList []OnsMqttGroupIdListMqttGroupIdDo

func (list *OnsMqttGroupIdListMqttGroupIdDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMqttGroupIdListMqttGroupIdDo)
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
