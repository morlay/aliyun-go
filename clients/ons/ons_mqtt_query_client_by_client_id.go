package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsMqttQueryClientByClientIdRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	ClientId     string `position:"Query" name:"ClientId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
}

func (r OnsMqttQueryClientByClientIdRequest) Invoke(client *sdk.Client) (response *OnsMqttQueryClientByClientIdResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsMqttQueryClientByClientIdRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMqttQueryClientByClientId", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsMqttQueryClientByClientIdResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsMqttQueryClientByClientIdResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsMqttQueryClientByClientIdResponse struct {
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
