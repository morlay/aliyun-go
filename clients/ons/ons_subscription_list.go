package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsSubscriptionListRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
}

func (r OnsSubscriptionListRequest) Invoke(client *sdk.Client) (response *OnsSubscriptionListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsSubscriptionListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsSubscriptionList", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsSubscriptionListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsSubscriptionListResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsSubscriptionListResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsSubscriptionListSubscribeInfoDoList
}

type OnsSubscriptionListSubscribeInfoDo struct {
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
