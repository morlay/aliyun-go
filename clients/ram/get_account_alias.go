package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetAccountAliasRequest struct {
	requests.RpcRequest
}

func (req *GetAccountAliasRequest) Invoke(client *sdk.Client) (resp *GetAccountAliasResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "GetAccountAlias", "", "")
	resp = &GetAccountAliasResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetAccountAliasResponse struct {
	responses.BaseResponse
	RequestId    string
	AccountAlias string
}
