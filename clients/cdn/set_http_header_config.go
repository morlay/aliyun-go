package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetHttpHeaderConfigRequest struct {
	HeaderValue   string `position:"Query" name:"HeaderValue"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	HeaderKey     string `position:"Query" name:"HeaderKey"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r SetHttpHeaderConfigRequest) Invoke(client *sdk.Client) (response *SetHttpHeaderConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetHttpHeaderConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetHttpHeaderConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SetHttpHeaderConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetHttpHeaderConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetHttpHeaderConfigResponse struct {
	RequestId string
}
