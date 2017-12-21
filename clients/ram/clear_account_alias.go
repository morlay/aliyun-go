package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ClearAccountAliasRequest struct {
}

func (r ClearAccountAliasRequest) Invoke(client *sdk.Client) (response *ClearAccountAliasResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ClearAccountAliasRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "ClearAccountAlias", "", "")

	resp := struct {
		*responses.BaseResponse
		ClearAccountAliasResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ClearAccountAliasResponse

	err = client.DoAction(&req, &resp)
	return
}

type ClearAccountAliasResponse struct {
	RequestId string
}
