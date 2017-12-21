package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ClearAccountAliasRequest struct {
	requests.RpcRequest
}

func (req *ClearAccountAliasRequest) Invoke(client *sdk.Client) (resp *ClearAccountAliasResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "ClearAccountAlias", "", "")
	resp = &ClearAccountAliasResponse{}
	err = client.DoAction(req, resp)
	return
}

type ClearAccountAliasResponse struct {
	responses.BaseResponse
	RequestId string
}
