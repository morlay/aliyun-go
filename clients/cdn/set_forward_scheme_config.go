package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetForwardSchemeConfigRequest struct {
	requests.RpcRequest
	SchemeOrigin     string `position:"Query" name:"SchemeOrigin"`
	SchemeOriginPort string `position:"Query" name:"SchemeOriginPort"`
	SecurityToken    string `position:"Query" name:"SecurityToken"`
	Enable           string `position:"Query" name:"Enable"`
	DomainName       string `position:"Query" name:"DomainName"`
	OwnerId          int64  `position:"Query" name:"OwnerId"`
}

func (req *SetForwardSchemeConfigRequest) Invoke(client *sdk.Client) (resp *SetForwardSchemeConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetForwardSchemeConfig", "", "")
	resp = &SetForwardSchemeConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetForwardSchemeConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
