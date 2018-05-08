package sas_api

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	Code    common.String
	Message common.String
	Success bool
	Data    GetThreatAnalyseListData
}

type GetThreatAnalyseListData struct {
	Total common.Long
	Items GetThreatAnalyseListItemList
}

type GetThreatAnalyseListItem struct {
	Id                 common.Long
	Level              common.String
	Uuid               common.String
	Product            common.String
	VmIp               common.String
	Url                common.String
	Method             common.String
	SourceIp           common.String
	SourceUuid         common.String
	SourceDomain       common.String
	SourcePort         common.Integer
	SourceLocal        common.String
	DstIp              common.String
	DstUuid            common.String
	DstDomain          common.String
	DstPort            common.Integer
	DstLocal           common.String
	AttackCount        common.Long
	ThreatRate         common.String
	AttackStartTime    common.Long
	AttackEndTime      common.Long
	AttackCategory     common.Integer
	AttackCategoryName common.String
	AttackType         common.String
	AttackTypeName     common.String
	AttackStatus       common.Integer
	AttackSource       common.String
	Details            GetThreatAnalyseListDetailList
}

type GetThreatAnalyseListDetail struct {
	Value    common.String
	Type     common.Integer
	Label    common.String
	LinkText common.String
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
