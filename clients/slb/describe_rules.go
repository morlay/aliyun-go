package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeRulesRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         int    `position:"Query" name:"ListenerPort"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *DescribeRulesRequest) Invoke(client *sdk.Client) (resp *DescribeRulesResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeRules", "slb", "")
	resp = &DescribeRulesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeRulesResponse struct {
	responses.BaseResponse
	RequestId common.String
	Rules     DescribeRulesRuleList
}

type DescribeRulesRule struct {
	RuleId         common.String
	RuleName       common.String
	Domain         common.String
	Url            common.String
	VServerGroupId common.String
}

type DescribeRulesRuleList []DescribeRulesRule

func (list *DescribeRulesRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRulesRule)
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
