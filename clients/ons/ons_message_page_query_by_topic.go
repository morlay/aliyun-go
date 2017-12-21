package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsMessagePageQueryByTopicRequest struct {
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

func (r OnsMessagePageQueryByTopicRequest) Invoke(client *sdk.Client) (response *OnsMessagePageQueryByTopicResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsMessagePageQueryByTopicRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMessagePageQueryByTopic", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsMessagePageQueryByTopicResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsMessagePageQueryByTopicResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsMessagePageQueryByTopicResponse struct {
	RequestId  string
	HelpUrl    string
	MsgFoundDo OnsMessagePageQueryByTopicMsgFoundDo
}

type OnsMessagePageQueryByTopicMsgFoundDo struct {
	TaskId       string
	MaxPageCount int64
	CurrentPage  int64
	MsgFoundList OnsMessagePageQueryByTopicOnsRestMessageDoList
}

type OnsMessagePageQueryByTopicOnsRestMessageDo struct {
	Topic          string
	Flag           int
	Body           string
	StoreSize      int
	BornTimestamp  int64
	BornHost       string
	StoreTimestamp int64
	StoreHost      string
	MsgId          string
	OffsetId       string
	BodyCRC        int
	ReconsumeTimes int
	PropertyList   OnsMessagePageQueryByTopicMessagePropertyList
}

type OnsMessagePageQueryByTopicMessageProperty struct {
	Name  string
	Value string
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
