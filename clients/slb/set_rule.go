package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetRuleRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	VServerGroupId       string `position:"Query" name:"VServerGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RuleId               string `position:"Query" name:"RuleId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *SetRuleRequest) Invoke(client *sdk.Client) (resp *SetRuleResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "SetRule", "slb", "")
	resp = &SetRuleResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetRuleResponse struct {
	responses.BaseResponse
	RequestId common.String
}
