package sas_api

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetThreatAnalyseListRequest struct {
	requests.RpcRequest
	Start int `position:"Query" name:"Start"`
	Limit int `position:"Query" name:"Limit"`
}

func (req *GetThreatAnalyseListRequest) Invoke(client *sdk.Client) (resp *GetThreatAnalyseListResponse, err error) {
	req.InitWithApiInfo("Sas-api", "2017-07-05", "GetThreatAnalyseList", "", "")
	resp = &GetThreatAnalyseListResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetThreatAnalyseListResponse struct {
	responses.BaseResponse
	Code    string
	Message string
	Success bool
	Data    GetThreatAnalyseListData
}

type GetThreatAnalyseListData struct {
	Total int64
	Items GetThreatAnalyseListItemList
}

type GetThreatAnalyseListItem struct {
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
	Details            GetThreatAnalyseListDetailList
}

type GetThreatAnalyseListDetail struct {
	Value    string
	Type     int
	Label    string
	LinkText string
}

type GetThreatAnalyseListItemList []GetThreatAnalyseListItem

func (list *GetThreatAnalyseListItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetThreatAnalyseListItem)
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

type GetThreatAnalyseListDetailList []GetThreatAnalyseListDetail

func (list *GetThreatAnalyseListDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetThreatAnalyseListDetail)
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
