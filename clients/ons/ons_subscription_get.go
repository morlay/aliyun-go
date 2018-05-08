package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type OnsSubscriptionGetRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	Topic        string `position:"Query" name:"Topic"`
}

func (req *OnsSubscriptionGetRequest) Invoke(client *sdk.Client) (resp *OnsSubscriptionGetResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsSubscriptionGet", "", "")
	resp = &OnsSubscriptionGetResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsSubscriptionGetResponse struct {
	responses.BaseResponse
	RequestId common.String
	HelpUrl   common.String
	Data      OnsSubscriptionGetSubscribeInfoDoList
}

type OnsSubscriptionGetSubscribeInfoDo struct {
	Id          common.Long
	ChannelId   common.Integer
	ChannelName common.String
	OnsRegionId common.String
	RegionName  common.String
	Owner       common.String
	ConsumerId  common.String
	Topic       common.String
	Status      common.Integer
	StatusName  common.String
	CreateTime  common.Long
	UpdateTime  common.Long
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
