package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetForceRedirectConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	RedirectType  string `position:"Query" name:"RedirectType"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r SetForceRedirectConfigRequest) Invoke(client *sdk.Client) (response *SetForceRedirectConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetForceRedirectConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetForceRedirectConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SetForceRedirectConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetForceRedirectConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetForceRedirectConfigResponse struct {
	RequestId string
}
