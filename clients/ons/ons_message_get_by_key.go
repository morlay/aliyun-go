package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId string
	HelpUrl   string
	Data      OnsMessageGetByKeyOnsRestMessageDoList
}

type OnsMessageGetByKeyOnsRestMessageDo struct {
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
	PropertyList   OnsMessageGetByKeyMessagePropertyList
}

type OnsMessageGetByKeyMessageProperty struct {
	Name  string
	Value string
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
