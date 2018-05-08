package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateRulesRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         int    `position:"Query" name:"ListenerPort"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RuleList             string `position:"Query" name:"RuleList"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *CreateRulesRequest) Invoke(client *sdk.Client) (resp *CreateRulesResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "CreateRules", "slb", "")
	resp = &CreateRulesResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateRulesResponse struct {
	responses.BaseResponse
	RequestId common.String
	Rules     CreateRulesRuleList
}

type CreateRulesRule struct {
	RuleId   common.String
	RuleName common.String
}

type CreateRulesRuleList []CreateRulesRule

func (list *CreateRulesRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateRulesRule)
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
