package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsSubscriptionGetRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	Topic        string `position:"Query" name:"Topic"`
}

func (r OnsSubscriptionGetRequest) Invoke(client *sdk.Client) (response *OnsSubscriptionGetResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsSubscriptionGetRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsSubscriptionGet", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsSubscriptionGetResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsSubscriptionGetResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsSubscriptionGetResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsSubscriptionGetSubscribeInfoDoList
}

type OnsSubscriptionGetSubscribeInfoDo struct {
	Id          int64
	ChannelId   int
	ChannelName string
	OnsRegionId string
	RegionName  string
	Owner       string
	ConsumerId  string
	Topic       string
	Status      int
	StatusName  string
	CreateTime  int64
	UpdateTime  int64
}

type OnsSubscriptionGetSubscribeInfoDoList []OnsSubscriptionGetSubscribeInfoDo

func (list *OnsSubscriptionGetSubscribeInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsSubscriptionGetSubscribeInfoDo)
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
