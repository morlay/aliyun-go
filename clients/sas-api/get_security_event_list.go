package sas_api

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	Code    common.String
	Message common.String
	Success bool
	Data    GetSecurityEventListData
}

type GetSecurityEventListData struct {
	Total common.Long
	Items GetSecurityEventListItemList
}

type GetSecurityEventListItem struct {
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
	Details            GetSecurityEventListDetailList
}

type GetSecurityEventListDetail struct {
	Value    common.String
	Type     common.Integer
	Label    common.String
	LinkText common.String
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
