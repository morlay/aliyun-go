package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type OnsTopicListRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
}

func (req *OnsTopicListRequest) Invoke(client *sdk.Client) (resp *OnsTopicListResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsTopicList", "", "")
	resp = &OnsTopicListResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsTopicListResponse struct {
	responses.BaseResponse
	RequestId common.String
	HelpUrl   common.String
	Data      OnsTopicListPublishInfoDoList
}

type OnsTopicListPublishInfoDo struct {
	Id           common.Long
	ChannelId    common.Integer
	ChannelName  common.String
	OnsRegionId  common.String
	RegionName   common.String
	Topic        common.String
	Owner        common.String
	Relation     common.Integer
	RelationName common.String
	Status       common.Integer
	StatusName   common.String
	Appkey       common.String
	CreateTime   common.Long
	UpdateTime   common.Long
	Remark       common.String
}

type OnsTopicListPublishInfoDoList []OnsTopicListPublishInfoDo

func (list *OnsTopicListPublishInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsTopicListPublishInfoDo)
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
