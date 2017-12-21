package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetHttpsOptionConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	Http2         string `position:"Query" name:"Http.2"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r SetHttpsOptionConfigRequest) Invoke(client *sdk.Client) (response *SetHttpsOptionConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetHttpsOptionConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetHttpsOptionConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SetHttpsOptionConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetHttpsOptionConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetHttpsOptionConfigResponse struct {
	RequestId string
}
