package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UnbindAliasRequest struct {
	AliasName string `position:"Query" name:"AliasName"`
	AppKey    int64  `position:"Query" name:"AppKey"`
	DeviceId  string `position:"Query" name:"DeviceId"`
	UnbindAll string `position:"Query" name:"UnbindAll"`
}

func (r UnbindAliasRequest) Invoke(client *sdk.Client) (response *UnbindAliasResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UnbindAliasRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Push", "2016-08-01", "UnbindAlias", "", "")

	resp := struct {
		*responses.BaseResponse
		UnbindAliasResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UnbindAliasResponse

	err = client.DoAction(&req, &resp)
	return
}

type UnbindAliasResponse struct {
	RequestId string
}
