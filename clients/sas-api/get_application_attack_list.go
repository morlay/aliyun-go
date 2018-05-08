package sas_api

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	Code    common.String
	Message common.String
	Success bool
	Data    GetApplicationAttackListData
}

type GetApplicationAttackListData struct {
	Total common.Long
	Items GetApplicationAttackListItemList
}

type GetApplicationAttackListItem struct {
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
	Details            GetApplicationAttackListDetailList
}

type GetApplicationAttackListDetail struct {
	Value    common.String
	Type     common.Integer
	Label    common.String
	LinkText common.String
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
