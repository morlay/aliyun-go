package sas_api

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetApplicationAttackListRequest struct {
	requests.RpcRequest
	Start int `position:"Query" name:"Start"`
	Limit int `position:"Query" name:"Limit"`
}

func (req *GetApplicationAttackListRequest) Invoke(client *sdk.Client) (resp *GetApplicationAttackListResponse, err error) {
	req.InitWithApiInfo("Sas-api", "2017-07-05", "GetApplicationAttackList", "", "")
	resp = &GetApplicationAttackListResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetApplicationAttackListResponse struct {
	responses.BaseResponse
	Code    string
	Message string
	Success bool
	Data    GetApplicationAttackListData
}

type GetApplicationAttackListData struct {
	Total int64
	Items GetApplicationAttackListItemList
}

type GetApplicationAttackListItem struct {
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
	Details            GetApplicationAttackListDetailList
}

type GetApplicationAttackListDetail struct {
	Value    string
	Type     int
	Label    string
	LinkText string
}

type GetApplicationAttackListItemList []GetApplicationAttackListItem

func (list *GetApplicationAttackListItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetApplicationAttackListItem)
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

type GetApplicationAttackListDetailList []GetApplicationAttackListDetail

func (list *GetApplicationAttackListDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetApplicationAttackListDetail)
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
