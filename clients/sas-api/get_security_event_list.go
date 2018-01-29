package sas_api

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetSecurityEventListRequest struct {
	requests.RpcRequest
	Start int `position:"Query" name:"Start"`
	Limit int `position:"Query" name:"Limit"`
}

func (req *GetSecurityEventListRequest) Invoke(client *sdk.Client) (resp *GetSecurityEventListResponse, err error) {
	req.InitWithApiInfo("Sas-api", "2017-07-05", "GetSecurityEventList", "", "")
	resp = &GetSecurityEventListResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetSecurityEventListResponse struct {
	responses.BaseResponse
	Code    string
	Message string
	Success bool
	Data    GetSecurityEventListData
}

type GetSecurityEventListData struct {
	Total int64
	Items GetSecurityEventListItemList
}

type GetSecurityEventListItem struct {
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
	Details            GetSecurityEventListDetailList
}

type GetSecurityEventListDetail struct {
	Value    string
	Type     int
	Label    string
	LinkText string
}

type GetSecurityEventListItemList []GetSecurityEventListItem

func (list *GetSecurityEventListItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetSecurityEventListItem)
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

type GetSecurityEventListDetailList []GetSecurityEventListDetail

func (list *GetSecurityEventListDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetSecurityEventListDetail)
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
