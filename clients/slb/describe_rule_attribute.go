package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeRuleAttributeRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RuleId               string `position:"Query" name:"RuleId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r DescribeRuleAttributeRequest) Invoke(client *sdk.Client) (response *DescribeRuleAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeRuleAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeRuleAttribute", "slb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeRuleAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeRuleAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeRuleAttributeResponse struct {
	RequestId      string
	RuleName       string
	LoadBalancerId string
	ListenerPort   string
	Domain         string
	Url            string
	VServerGroupId string
}
