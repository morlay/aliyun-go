package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type OnsMessageGetByMsgIdRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	MsgId        string `position:"Query" name:"MsgId"`
	Topic        string `position:"Query" name:"Topic"`
}

func (req *OnsMessageGetByMsgIdRequest) Invoke(client *sdk.Client) (resp *OnsMessageGetByMsgIdResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMessageGetByMsgId", "", "")
	resp = &OnsMessageGetByMsgIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsMessageGetByMsgIdResponse struct {
	responses.BaseResponse
	RequestId common.String
	HelpUrl   common.String
	Data      OnsMessageGetByMsgIdData
}

type OnsMessageGetByMsgIdData struct {
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
	PropertyList   OnsMessageGetByMsgIdMessagePropertyList
}

type OnsMessageGetByMsgIdMessageProperty struct {
	Name  common.String
	Value common.String
}

type OnsMessageGetByMsgIdMessagePropertyList []OnsMessageGetByMsgIdMessageProperty

func (list *OnsMessageGetByMsgIdMessagePropertyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMessageGetByMsgIdMessageProperty)
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
