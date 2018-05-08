package qualitycheck

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetRuleRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *GetRuleRequest) Invoke(client *sdk.Client) (resp *GetRuleResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "GetRule", "", "")
	resp = &GetRuleResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetRuleResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
	Data      GetRuleData
}

type GetRuleData struct {
	Rules GetRuleRuleInfoList
}

type GetRuleRuleInfo struct {
	Rid                      common.String
	RuleLambda               common.String
	Name                     common.String
	Type                     common.Integer
	Status                   common.Integer
	IsDelete                 common.Integer
	StartTime                common.String
	EndTime                  common.String
	Weight                   common.String
	IsOnline                 common.Integer
	CreateEmpid              common.String
	CreateTime               common.String
	LastUpdateTime           common.String
	LastUpdateEmpid          common.String
	Comments                 common.String
	AutoReview               common.Integer
	RuleScoreType            common.Integer
	ScoreName                common.String
	ScoreSubName             common.String
	BusinessCategoryNameList GetRuleBusinessCategoryNameListList
}

type GetRuleRuleInfoList []GetRuleRuleInfo

func (list *GetRuleRuleInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRuleRuleInfo)
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

type GetRuleBusinessCategoryNameListList []common.String

func (list *GetRuleBusinessCategoryNameListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
