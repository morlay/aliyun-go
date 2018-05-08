package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListHealthRuleResultRequest struct {
	requests.RpcRequest
	Component       string `position:"Query" name:"Component"`
	HostName        string `position:"Query" name:"HostName"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Pass            int    `position:"Query" name:"Pass"`
	Service         string `position:"Query" name:"Service"`
	PageSize        int    `position:"Query" name:"PageSize"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
}

func (req *ListHealthRuleResultRequest) Invoke(client *sdk.Client) (resp *ListHealthRuleResultResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListHealthRuleResult", "", "")
	resp = &ListHealthRuleResultResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListHealthRuleResultResponse struct {
	responses.BaseResponse
	RequestId            common.String
	Total                common.Integer
	PageNumber           common.Integer
	PageSize             common.Integer
	HealthRuleResultList ListHealthRuleResultHealthRuleResultList
}

type ListHealthRuleResultHealthRuleResult struct {
	Id              common.Long
	ClusterId       common.Long
	RuleId          common.Long
	RuleName        common.String
	RuleTitle       common.String
	RuleStatus      common.String
	RuleDescription common.String
	Service         common.String
	Component       common.String
	Pass            common.String
	HostNames       common.String
}

type ListHealthRuleResultHealthRuleResultList []ListHealthRuleResultHealthRuleResult

func (list *ListHealthRuleResultHealthRuleResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListHealthRuleResultHealthRuleResult)
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
