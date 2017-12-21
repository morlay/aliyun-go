package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StopCdnDomainRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *StopCdnDomainRequest) Invoke(client *sdk.Client) (resp *StopCdnDomainResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "StopCdnDomain", "", "")
	resp = &StopCdnDomainResponse{}
	err = client.DoAction(req, resp)
	return
}

type StopCdnDomainResponse struct {
	responses.BaseResponse
	RequestId string
}
