package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetAccountAliasRequest struct {
}

func (r GetAccountAliasRequest) Invoke(client *sdk.Client) (response *GetAccountAliasResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetAccountAliasRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "GetAccountAlias", "", "")

	resp := struct {
		*responses.BaseResponse
		GetAccountAliasResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetAccountAliasResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetAccountAliasResponse struct {
	RequestId    string
	AccountAlias string
}
