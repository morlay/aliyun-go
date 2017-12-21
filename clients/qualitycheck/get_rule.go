package qualitycheck

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      GetRuleData
}

type GetRuleData struct {
	Rules GetRuleRuleInfoList
}

type GetRuleRuleInfo struct {
	Rid             string
	RuleLambda      string
	Name            string
	Type            int
	Status          int
	IsDelete        int
	StartTime       string
	EndTime         string
	Weight          string
	IsOnline        int
	CreateEmpid     string
	CreateTime      string
	LastUpdateTime  string
	LastUpdateEmpid string
	Comments        string
	AutoReview      int
	RuleScoreType   int
	ScoreName       string
	ScoreSubName    string
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
