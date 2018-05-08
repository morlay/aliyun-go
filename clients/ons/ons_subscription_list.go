package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type OnsSubscriptionListRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
}

func (req *OnsSubscriptionListRequest) Invoke(client *sdk.Client) (resp *OnsSubscriptionListResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsSubscriptionList", "", "")
	resp = &OnsSubscriptionListResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsSubscriptionListResponse struct {
	responses.BaseResponse
	RequestId common.String
	HelpUrl   common.String
	Data      OnsSubscriptionListSubscribeInfoDoList
}

type OnsSubscriptionListSubscribeInfoDo struct {
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

type OnsSubscriptionListSubscribeInfoDoList []OnsSubscriptionListSubscribeInfoDo

func (list *OnsSubscriptionListSubscribeInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsSubscriptionListSubscribeInfoDo)
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
