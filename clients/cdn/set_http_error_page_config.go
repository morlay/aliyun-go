package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetHttpErrorPageConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageUrl       string `position:"Query" name:"PageUrl"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	ErrorCode     string `position:"Query" name:"ErrorCode"`
}

func (r SetHttpErrorPageConfigRequest) Invoke(client *sdk.Client) (response *SetHttpErrorPageConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetHttpErrorPageConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetHttpErrorPageConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SetHttpErrorPageConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetHttpErrorPageConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetHttpErrorPageConfigResponse struct {
	RequestId string
}
