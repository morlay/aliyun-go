package sas_api

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetCrackListRequest struct {
	requests.RpcRequest
	Start int `position:"Query" name:"Start"`
	Limit int `position:"Query" name:"Limit"`
}

func (req *GetCrackListRequest) Invoke(client *sdk.Client) (resp *GetCrackListResponse, err error) {
	req.InitWithApiInfo("Sas-api", "2017-07-05", "GetCrackList", "", "")
	resp = &GetCrackListResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetCrackListResponse struct {
	responses.BaseResponse
	Code    string
	Message string
	Success bool
	Data    GetCrackListData
}

type GetCrackListData struct {
	Total int64
	Items GetCrackListItemList
}

type GetCrackListItem struct {
	Id                 int64
	Level              string
	Uuid               string
	Product            string
	VmIp               string
	Url                string
	Method             string
	SourceIp           string
	SourceUuid         string
	SourceDomain       string
	SourcePort         int
	SourceLocal        string
	DstIp              string
	DstUuid            string
	DstDomain          string
	DstPort            int
	DstLocal           string
	AttackCount        int64
	ThreatRate         string
	AttackStartTime    int64
	AttackEndTime      int64
	AttackCategory     int
	AttackCategoryName string
	AttackType         string
	AttackTypeName     string
	AttackStatus       int
	AttackSource       string
	Details            GetCrackListDetailList
}

type GetCrackListDetail struct {
	Value    string
	Type     int
	Label    string
	LinkText string
}

type GetCrackListItemList []GetCrackListItem

func (list *GetCrackListItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetCrackListItem)
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

type GetCrackListDetailList []GetCrackListDetail

func (list *GetCrackListDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetCrackListDetail)
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
