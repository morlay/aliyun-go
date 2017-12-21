package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetAccountAliasRequest struct {
	AccountAlias string `position:"Query" name:"AccountAlias"`
}

func (r SetAccountAliasRequest) Invoke(client *sdk.Client) (response *SetAccountAliasResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetAccountAliasRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "SetAccountAlias", "", "")

	resp := struct {
		*responses.BaseResponse
		SetAccountAliasResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetAccountAliasResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetAccountAliasResponse struct {
	RequestId string
}
