package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetAccountAliasRequest struct {
	requests.RpcRequest
	AccountAlias string `position:"Query" name:"AccountAlias"`
}

func (req *SetAccountAliasRequest) Invoke(client *sdk.Client) (resp *SetAccountAliasResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "SetAccountAlias", "", "")
	resp = &SetAccountAliasResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetAccountAliasResponse struct {
	responses.BaseResponse
	RequestId common.String
}
