package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsMqttQueryMsgByPubInfoRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	ClientId     string `position:"Query" name:"ClientId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	EndTime      int64  `position:"Query" name:"EndTime"`
	BeginTime    int64  `position:"Query" name:"BeginTime"`
}

func (r OnsMqttQueryMsgByPubInfoRequest) Invoke(client *sdk.Client) (response *OnsMqttQueryMsgByPubInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsMqttQueryMsgByPubInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMqttQueryMsgByPubInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsMqttQueryMsgByPubInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsMqttQueryMsgByPubInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsMqttQueryMsgByPubInfoResponse struct {
	RequestId      string
	HelpUrl        string
	MqttMessageDos OnsMqttQueryMsgByPubInfoMqttMessageDoList
}

type OnsMqttQueryMsgByPubInfoMqttMessageDo struct {
	TraceId string
	PubTime int64
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
