package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsTopicSearchRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	Search       string `position:"Query" name:"Search"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
}

func (req *OnsTopicSearchRequest) Invoke(client *sdk.Client) (resp *OnsTopicSearchResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsTopicSearch", "", "")
	resp = &OnsTopicSearchResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsTopicSearchResponse struct {
	responses.BaseResponse
	RequestId string
	HelpUrl   string
	Data      OnsTopicSearchPublishInfoDoList
}

type OnsTopicSearchPublishInfoDo struct {
	Id           int64
	ChannelId    int
	ChannelName  string
	OnsRegionId  string
	RegionName   string
	Topic        string
	Owner        string
	Relation     int
	RelationName string
	Status       int
	StatusName   string
	Appkey       string
	CreateTime   int64
	UpdateTime   int64
	Remark       string
}

type OnsTopicSearchPublishInfoDoList []OnsTopicSearchPublishInfoDo

func (list *OnsTopicSearchPublishInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsTopicSearchPublishInfoDo)
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
