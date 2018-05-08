package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetHttpsOptionConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	Http2         string `position:"Query" name:"Http.2"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *SetHttpsOptionConfigRequest) Invoke(client *sdk.Client) (resp *SetHttpsOptionConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetHttpsOptionConfig", "", "")
	resp = &SetHttpsOptionConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetHttpsOptionConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
