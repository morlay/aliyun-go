package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeRulesRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         int    `position:"Query" name:"ListenerPort"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r DescribeRulesRequest) Invoke(client *sdk.Client) (response *DescribeRulesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeRulesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeRules", "slb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeRulesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeRulesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeRulesResponse struct {
	RequestId string
	Rules     DescribeRulesRuleList
}

type DescribeRulesRule struct {
	RuleId         string
	RuleName       string
	Domain         string
	Url            string
	VServerGroupId string
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
