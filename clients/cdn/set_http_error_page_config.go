package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetHttpErrorPageConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageUrl       string `position:"Query" name:"PageUrl"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	ErrorCode     string `position:"Query" name:"ErrorCode"`
}

func (req *SetHttpErrorPageConfigRequest) Invoke(client *sdk.Client) (resp *SetHttpErrorPageConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetHttpErrorPageConfig", "", "")
	resp = &SetHttpErrorPageConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetHttpErrorPageConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
