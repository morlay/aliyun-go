package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	HelpUrl   common.String
	Data      OnsMqttGroupIdListMqttGroupIdDoList
}

type OnsMqttGroupIdListMqttGroupIdDo struct {
	Id          common.Long
	ChannelId   common.Integer
	OnsRegionId common.String
	RegionName  common.String
	Owner       common.String
	GroupId     common.String
	Topic       common.String
	Status      common.Integer
	CreateTime  common.Long
	UpdateTime  common.Long
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
