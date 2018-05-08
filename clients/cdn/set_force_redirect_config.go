package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetForceRedirectConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	RedirectType  string `position:"Query" name:"RedirectType"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *SetForceRedirectConfigRequest) Invoke(client *sdk.Client) (resp *SetForceRedirectConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetForceRedirectConfig", "", "")
	resp = &SetForceRedirectConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetForceRedirectConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
