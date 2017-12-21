package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BindAliasRequest struct {
	AliasName string `position:"Query" name:"AliasName"`
	AppKey    int64  `position:"Query" name:"AppKey"`
	DeviceId  string `position:"Query" name:"DeviceId"`
}

func (r BindAliasRequest) Invoke(client *sdk.Client) (response *BindAliasResponse, err error) {
	req := struct {
		*requests.RpcRequest
		BindAliasRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Push", "2016-08-01", "BindAlias", "", "")

	resp := struct {
		*responses.BaseResponse
		BindAliasResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.BindAliasResponse

	err = client.DoAction(&req, &resp)
	return
}

type BindAliasResponse struct {
	RequestId string
}
