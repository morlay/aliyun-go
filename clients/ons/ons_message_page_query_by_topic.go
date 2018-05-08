package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type OnsMessagePageQueryByTopicRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	PageSize     int    `position:"Query" name:"PageSize"`
	Topic        string `position:"Query" name:"Topic"`
	EndTime      int64  `position:"Query" name:"EndTime"`
	BeginTime    int64  `position:"Query" name:"BeginTime"`
	CurrentPage  int    `position:"Query" name:"CurrentPage"`
	TaskId       string `position:"Query" name:"TaskId"`
}

func (req *OnsMessagePageQueryByTopicRequest) Invoke(client *sdk.Client) (resp *OnsMessagePageQueryByTopicResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMessagePageQueryByTopic", "", "")
	resp = &OnsMessagePageQueryByTopicResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsMessagePageQueryByTopicResponse struct {
	responses.BaseResponse
	RequestId  common.String
	HelpUrl    common.String
	MsgFoundDo OnsMessagePageQueryByTopicMsgFoundDo
}

type OnsMessagePageQueryByTopicMsgFoundDo struct {
	TaskId       common.String
	MaxPageCount common.Long
	CurrentPage  common.Long
	MsgFoundList OnsMessagePageQueryByTopicOnsRestMessageDoList
}

type OnsMessagePageQueryByTopicOnsRestMessageDo struct {
	Topic          common.String
	Flag           common.Integer
	Body           common.String
	StoreSize      common.Integer
	BornTimestamp  common.Long
	BornHost       common.String
	StoreTimestamp common.Long
	StoreHost      common.String
	MsgId          common.String
	OffsetId       common.String
	BodyCRC        common.Integer
	ReconsumeTimes common.Integer
	PropertyList   OnsMessagePageQueryByTopicMessagePropertyList
}

type OnsMessagePageQueryByTopicMessageProperty struct {
	Name  common.String
	Value common.String
}

type OnsMessagePageQueryByTopicOnsRestMessageDoList []OnsMessagePageQueryByTopicOnsRestMessageDo

func (list *OnsMessagePageQueryByTopicOnsRestMessageDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMessagePageQueryByTopicOnsRestMessageDo)
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

type OnsMessagePageQueryByTopicMessagePropertyList []OnsMessagePageQueryByTopicMessageProperty

func (list *OnsMessagePageQueryByTopicMessagePropertyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMessagePageQueryByTopicMessageProperty)
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
