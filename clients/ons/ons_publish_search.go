package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsPublishSearchRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	Search       string `position:"Query" name:"Search"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
}

func (req *OnsPublishSearchRequest) Invoke(client *sdk.Client) (resp *OnsPublishSearchResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsPublishSearch", "", "")
	resp = &OnsPublishSearchResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsPublishSearchResponse struct {
	responses.BaseResponse
	RequestId string
	HelpUrl   string
	Data      OnsPublishSearchPublishInfoDoList
}

type OnsPublishSearchPublishInfoDo struct {
	Id          int64
	ChannelId   int
	ChannelName string
	OnsRegionId string
	RegionName  string
	Owner       string
	ProducerId  string
	Topic       string
	Status      int
	StatusName  string
	CreateTime  int64
	UpdateTime  int64
}

type OnsPublishSearchPublishInfoDoList []OnsPublishSearchPublishInfoDo

func (list *OnsPublishSearchPublishInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsPublishSearchPublishInfoDo)
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
