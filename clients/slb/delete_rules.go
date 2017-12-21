package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteRulesRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	RuleIds              string `position:"Query" name:"RuleIds"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r DeleteRulesRequest) Invoke(client *sdk.Client) (response *DeleteRulesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteRulesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "DeleteRules", "slb", "")

	resp := struct {
		*responses.BaseResponse
		DeleteRulesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteRulesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteRulesResponse struct {
	RequestId string
}
