package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteRulesRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	RuleIds              string `position:"Query" name:"RuleIds"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *DeleteRulesRequest) Invoke(client *sdk.Client) (resp *DeleteRulesResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DeleteRules", "slb", "")
	resp = &DeleteRulesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteRulesResponse struct {
	responses.BaseResponse
	RequestId string
}
