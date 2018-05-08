package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type OnsMqttQueryMsgByPubInfoRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	ClientId     string `position:"Query" name:"ClientId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	EndTime      int64  `position:"Query" name:"EndTime"`
	BeginTime    int64  `position:"Query" name:"BeginTime"`
}

func (req *OnsMqttQueryMsgByPubInfoRequest) Invoke(client *sdk.Client) (resp *OnsMqttQueryMsgByPubInfoResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMqttQueryMsgByPubInfo", "", "")
	resp = &OnsMqttQueryMsgByPubInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsMqttQueryMsgByPubInfoResponse struct {
	responses.BaseResponse
	RequestId      common.String
	HelpUrl        common.String
	MqttMessageDos OnsMqttQueryMsgByPubInfoMqttMessageDoList
}

type OnsMqttQueryMsgByPubInfoMqttMessageDo struct {
	TraceId common.String
	PubTime common.Long
}

type OnsMqttQueryMsgByPubInfoMqttMessageDoList []OnsMqttQueryMsgByPubInfoMqttMessageDo

func (list *OnsMqttQueryMsgByPubInfoMqttMessageDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMqttQueryMsgByPubInfoMqttMessageDo)
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
