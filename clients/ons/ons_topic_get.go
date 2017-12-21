package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsTopicGetRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
}

func (r OnsTopicGetRequest) Invoke(client *sdk.Client) (response *OnsTopicGetResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsTopicGetRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsTopicGet", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsTopicGetResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsTopicGetResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsTopicGetResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsTopicGetPublishInfoDoList
}

type OnsTopicGetPublishInfoDo struct {
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

type OnsTopicGetPublishInfoDoList []OnsTopicGetPublishInfoDo

func (list *OnsTopicGetPublishInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsTopicGetPublishInfoDo)
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
