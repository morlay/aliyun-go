package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetRuleRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	VServerGroupId       string `position:"Query" name:"VServerGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RuleId               string `position:"Query" name:"RuleId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r SetRuleRequest) Invoke(client *sdk.Client) (response *SetRuleResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetRuleRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "SetRule", "slb", "")

	resp := struct {
		*responses.BaseResponse
		SetRuleResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetRuleResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetRuleResponse struct {
	RequestId string
}
