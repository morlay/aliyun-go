package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsMqttQueryClientByClientIdRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	ClientId     string `position:"Query" name:"ClientId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
}

func (req *OnsMqttQueryClientByClientIdRequest) Invoke(client *sdk.Client) (resp *OnsMqttQueryClientByClientIdResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMqttQueryClientByClientId", "", "")
	resp = &OnsMqttQueryClientByClientIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsMqttQueryClientByClientIdResponse struct {
	responses.BaseResponse
	RequestId        string
	HelpUrl          string
	MqttClientInfoDo OnsMqttQueryClientByClientIdMqttClientInfoDo
}

type OnsMqttQueryClientByClientIdMqttClientInfoDo struct {
	Online          bool
	ClientId        string
	SocketChannel   string
	LastTouch       int64
	SubScriptonData OnsMqttQueryClientByClientIdSubscriptionDoList
}

type OnsMqttQueryClientByClientIdSubscriptionDo struct {
	ParentTopic string
	SubTopic    string
	Qos         int
}

type OnsMqttQueryClientByClientIdSubscriptionDoList []OnsMqttQueryClientByClientIdSubscriptionDo

func (list *OnsMqttQueryClientByClientIdSubscriptionDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMqttQueryClientByClientIdSubscriptionDo)
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
