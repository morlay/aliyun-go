package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type OnsMessageGetByKeyRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	Key          string `position:"Query" name:"Key"`
}

func (req *OnsMessageGetByKeyRequest) Invoke(client *sdk.Client) (resp *OnsMessageGetByKeyResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMessageGetByKey", "", "")
	resp = &OnsMessageGetByKeyResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsMessageGetByKeyResponse struct {
	responses.BaseResponse
	RequestId common.String
	HelpUrl   common.String
	Data      OnsMessageGetByKeyOnsRestMessageDoList
}

type OnsMessageGetByKeyOnsRestMessageDo struct {
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
	PropertyList   OnsMessageGetByKeyMessagePropertyList
}

type OnsMessageGetByKeyMessageProperty struct {
	Name  common.String
	Value common.String
}

type OnsMessageGetByKeyOnsRestMessageDoList []OnsMessageGetByKeyOnsRestMessageDo

func (list *OnsMessageGetByKeyOnsRestMessageDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMessageGetByKeyOnsRestMessageDo)
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

type OnsMessageGetByKeyMessagePropertyList []OnsMessageGetByKeyMessageProperty

func (list *OnsMessageGetByKeyMessagePropertyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMessageGetByKeyMessageProperty)
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
