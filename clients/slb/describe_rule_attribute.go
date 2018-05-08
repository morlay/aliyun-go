package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeRuleAttributeRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RuleId               string `position:"Query" name:"RuleId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *DescribeRuleAttributeRequest) Invoke(client *sdk.Client) (resp *DescribeRuleAttributeResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeRuleAttribute", "slb", "")
	resp = &DescribeRuleAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeRuleAttributeResponse struct {
	responses.BaseResponse
	RequestId      common.String
	RuleName       common.String
	LoadBalancerId common.String
	ListenerPort   common.String
	Domain         common.String
	Url            common.String
	VServerGroupId common.String
}
