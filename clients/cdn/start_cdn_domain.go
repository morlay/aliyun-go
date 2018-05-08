package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type StartCdnDomainRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *StartCdnDomainRequest) Invoke(client *sdk.Client) (resp *StartCdnDomainResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "StartCdnDomain", "", "")
	resp = &StartCdnDomainResponse{}
	err = client.DoAction(req, resp)
	return
}

type StartCdnDomainResponse struct {
	responses.BaseResponse
	RequestId common.String
}
