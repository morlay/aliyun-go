package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetHttpHeaderConfigRequest struct {
	requests.RpcRequest
	HeaderValue   string `position:"Query" name:"HeaderValue"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	HeaderKey     string `position:"Query" name:"HeaderKey"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *SetHttpHeaderConfigRequest) Invoke(client *sdk.Client) (resp *SetHttpHeaderConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetHttpHeaderConfig", "", "")
	resp = &SetHttpHeaderConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetHttpHeaderConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
